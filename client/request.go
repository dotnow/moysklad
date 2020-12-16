package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// Params структура параметров
type Params struct {
	Filters []Filter
	Query   map[string]string
}

// Filter структура фильтра
type Filter struct {
	Key      string
	Value    string
	Operator string
}

// Request структура параметров запроса
type Request struct {
	endPoint string
	params   Params
	api      *APIClient
}

// newRequest создаёт и возвращает ссылку на структуру запроса
func (api *APIClient) newRequest(endPoint string, params Params) *Request {
	return &Request{
		endPoint: endPoint,
		params:   params,
		api:      api,
	}
}

// setHeaders устанавливает необходимые заголовки к запросу
func (api *APIClient) setHeaders(req *http.Request) {
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, api.token))
	req.Header.Set(`Accept`, `application/json;charset=utf-8`)
	req.Header.Set(`Content-Type`, `application/json`)
}

// getFilterString склеивает фильтры в строку и возвращает
func (params Params) getFilterString() (filters string) {

	for _, filter := range params.Filters {
		filters = fmt.Sprintf("%s%s%s%s;", filters, filter.Key, filter.Operator, filter.Value)
	}

	return
}

// AddFilter Допустимые операторы: ['=', '>', '<', '>=', '<=', '!=', '~', '~=', '=~']
func (params *Params) AddFilter(key, operator, value string) {
	params.Filters = append(params.Filters, Filter{key, value, operator})
}

// AddQuery добавляет параметр к запросу
func (params *Params) AddQuery(key string, value string) {

	if params.Query == nil {
		params.Query = make(map[string]string)
	}

	params.Query[key] = value
}

func (request *Request) makeURL(baseURL string) string {

	url, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	query := url.Query()

	if filter := request.params.getFilterString(); filter != "" {
		query.Set("filter", filter)
	}

	for key, value := range request.params.Query {
		query.Set(key, value)
	}

	url.RawQuery = query.Encode()

	return url.String()
}

// getByUUID возвращает объект по его UUID
func (request *Request) getByUUID(uuid string) (result []byte, err error) {

	baseURL := fmt.Sprintf("%s%s/%s", APIPath, request.endPoint, uuid)

	url := request.makeURL(baseURL)

	result, err = request.GET(url)

	return
}

// makeURLList создаёт и возвращает массив ссылок
func makeURLList(baseURL string, size int) (urlList []string) {

	offset, limit := MaxLimitOffset, MaxLimitOffset

	url, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for offset < size {
		query := url.Query()
		query.Set("offset", strconv.Itoa(offset))
		query.Set("limit", strconv.Itoa(size-offset))
		// query.Set("limit", strconv.Itoa(limit))

		url.RawQuery = query.Encode()
		urlList = append(urlList, url.String())

		offset += limit
	}

	return
}

// getList возвращает множество объектов
func (request *Request) getList() (data []byte, size int, err error) {

	limit := 0

	baseURL := fmt.Sprintf("%s%s", APIPath, request.endPoint)

	if limitString, _ := request.params.Query["limit"]; limitString != "" {
		limit, err = strconv.Atoi(limitString)
		if err != nil {
			log.Println(err)
		}

		if limit != 0 && limit < MaxLimitOffset {
			request.params.AddQuery("limit", limitString)
		} else {
			request.params.AddQuery("limit", MaxLimitOffsetString())
		}
	}

	url := request.makeURL(baseURL)

	response, err := request.GET(url)
	if err != nil {
		return
	}

	responseData, err := getResponseData(response)
	if err != nil {
		return
	}

	size = responseData.Meta.Size

	if size > MaxLimitOffset {

		if limit != 0 {
			size = limit
		}

		resp := request.listGET(url, size)

		for _, element := range resp {
			respData, err := getResponseData(element.Body)
			if err != nil {
				log.Println(err)
				continue
			}

			responseData.Rows = append(responseData.Rows, respData.Rows...)
		}
	}

	data, err = json.Marshal(responseData.Rows)
	if err != nil {
		return
	}

	return
}

// getRateLimit извлекает значение из заголовка X-RateLimit-Remaining
func getRateLimit(res *http.Response) (remaining int) {

	remainingHeader := res.Header.Get("X-RateLimit-Remaining")

	if remainingHeader == "" {
		return
	}

	remaining, err := strconv.Atoi(remainingHeader)
	if err != nil {
		log.Println(err)
	}

	return
}

// GET запрос
func (request *Request) GET(url string) (body []byte, err error) {

	body, err = do(request, http.MethodGet, url, "")

	return
}

// POST запрос (создание)
func (request *Request) POST(url, data string) (body []byte, err error) {

	body, err = do(request, http.MethodPost, url, data)

	return
}

// PUT запрос (изменение)
func (request *Request) PUT(url, data string) (body []byte, err error) {

	body, err = do(request, http.MethodPut, url, data)

	return
}

// DELETE запрос (удаление)
func (request *Request) DELETE(url string) (body []byte, err error) {

	body, err = do(request, http.MethodDelete, url, "")

	return
}

// do формирует запрос и возвращает тело ответа
func do(request *Request, method, url string, data string) (body []byte, err error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	request.api.setHeaders(req)
	request.api.requestChannel <- struct{}{}

	log.Printf("Requesting: GET %s\n", url)

	res, err := request.api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rateLimit := getRateLimit(res)

	defer func() {
		res.Body.Close()

		if rateLimit <= 5 {
			<-time.After(time.Second)
		}

		<-request.api.requestChannel
	}()

	return body, nil
}

// listGET выполняет запросы, обрабатывает ответы и возвращает отсортированные результаты
func (request *Request) listGET(url string, size int) (results []Response) {

	urlList := makeURLList(url, size)

	for i, url := range urlList {
		go func(i int, url string) {
			response, err := request.GET(url)
			if err != nil {
				return
			}

			defer func(response []byte) {
				request.api.responseChannel <- &Response{i, response}
			}(response)
		}(i, url)
	}

	for {
		result := <-request.api.responseChannel
		results = append(results, *result)

		if len(results) == len(urlList) {
			break
		}
	}

	sort.Slice(results, func(i, k int) bool {
		return results[i].Index < results[k].Index
	})

	return
}
