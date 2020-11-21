package moysklad

import (
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

// APIPath базовый URL к API МойСклад
const APIPath = "https://online.moysklad.ru/api/remap/1.2/"

// Response структура для хранения ответа параллельных запросов
type Response struct {
	Index int
	Body  []byte
}

// ResponseData структура для десериализации ответа
type ResponseData struct {
	Meta struct {
		Size int `json:"size"`
	} `json:"meta"`
	Rows []map[string]interface{} `json:"rows"`
}

// APIClient экземпляр клиента МойСклад с токеном
type APIClient struct {
	token           string
	requestChannel  chan struct{}
	responseChannel chan *Response
	httpClient      *http.Client
}

// APIRequest структура параметров запроса
type APIRequest struct {
	endPoint string
	params   Params
	client   *APIClient
}

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

func (request *APIRequest) makeURL(uuid ...string) string {
	baseURL := fmt.Sprintf("%s%s", APIPath, request.endPoint)

	if len(uuid) > 0 {
		baseURL = fmt.Sprintf("%s/%s", baseURL, uuid[0])
	}

	url, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
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

// newRequest создаёт и возвращает ссылку на структуру запроса
func newRequest(client *APIClient, endPoint string, params Params) *APIRequest {
	return &APIRequest{
		endPoint: endPoint,
		params:   params,
		client:   client,
	}
}

// NewClient создаёт и возвращает ссылку на экземпляр APIClient
func NewClient(token string) *APIClient {
	return &APIClient{
		token:           token,
		requestChannel:  make(chan struct{}, 5),
		responseChannel: make(chan *Response),
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       5,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
	}
}

// getToken возвращает токен доступа
func (client *APIClient) getToken() string {
	return fmt.Sprintf(`Bearer %s`, client.token)
}

// setHeaders устанавливает необходимые заголовки к запросу
func (client *APIClient) setHeaders(req *http.Request) {
	req.Header.Set(`Authorization`, client.getToken())
	req.Header.Set(`Accept`, `application/json;charset=utf-8`)
	req.Header.Set(`Content-Type`, `application/json`)
}

// makeURLList создаёт и возвращает массив ссылок
func makeURLList(baseURL string, size int) (urlList []string) {
	offset, limit := 1000, 1000

	url, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for offset < size {
		query := url.Query()
		query.Set("offset", strconv.Itoa(offset))
		query.Set("limit", strconv.Itoa(size-offset))
		url.RawQuery = query.Encode()
		urlList = append(urlList, url.String())
		offset += limit
	}
	return

}

// getResponseData производит десериализацию ответа и возвращает размер meta и массив найденных объектов
func getResponseData(response []byte) (data *ResponseData, err error) {

	if err = json.Unmarshal(response, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// getFilterString склеивает фильтры в строку и возвращает
func (params Params) getFilterString() (filters string) {
	for _, filter := range params.Filters {
		filters = fmt.Sprintf("%s%s%s%s;", filters, filter.Key, filter.Operator, filter.Value)
	}
	return
}

// AddFilter Допустимые операторы: ['=', '>', '<', '>=', '<=', '!=', '~', '~=', '=~']
func (params *Params) AddFilter(key string, value string, operator string) {
	params.Filters = append(params.Filters, Filter{key, value, operator})
}

// AddQuery добавляет параметр к запросу
func (params *Params) AddQuery(key string, value string) {
	if params.Query == nil {
		params.Query = make(map[string]string)
	}
	params.Query[key] = value
}

// getByUUID возвращает объект по его UUID
func (request *APIRequest) getByUUID(uuid string) (result []byte, err error) {
	url := request.makeURL(uuid)
	result, err = request.getRequest(url)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// getList возвращает множество объектов
func (request *APIRequest) getList() (data []byte, err error) {

	var limit int

	url := request.makeURL()

	res, err := request.getRequest(url)
	if err != nil {
		return nil, err
	}

	responseData, err := getResponseData(res)
	if err != nil {
		return nil, err
	}

	if limitStr, exist := request.params.Query["limit"]; exist {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return nil, err
		}
	}

	if responseData.Meta.Size > 1000 && limit > 1000 {
		urlList := makeURLList(url, limit)
		parallelResponse := request.parallelGet(urlList)

		for _, response := range parallelResponse {
			respData, err := getResponseData(response.Body)
			if err != nil {
				return nil, err
			}

			responseData.Rows = append(responseData.Rows, respData.Rows...)
		}
	}

	data, err = json.Marshal(responseData.Rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getRemaining извлекает значение из заголовка X-RateLimit-Remaining
func getRemaining(res *http.Response) (remaining int) {

	remainingHeader := res.Header.Get("X-RateLimit-Remaining")

	remaining, err := strconv.Atoi(remainingHeader)
	if err != nil {
		log.Println(err)
	}

	return
}

// getRequest формирует запрос и возвращает тело ответа
func (request *APIRequest) getRequest(url string) (body []byte, err error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.client.setHeaders(req)
	request.client.requestChannel <- struct{}{}

	log.Printf("Requesting: GET %s\n", url)

	res, err := request.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	remaining := getRemaining(res)

	defer func() {
		res.Body.Close()

		if remaining <= 5 {
			<-time.After(time.Second)
		}
		<-request.client.requestChannel
	}()

	return body, nil
}

// parallelGet выполняет  запросы, обрабатывает ответы и возвращает отсортированные результаты ответов
func (request *APIRequest) parallelGet(urlList []string) (results []Response) {

	for i, url := range urlList {
		go func(i int, url string) {
			response, err := request.getRequest(url)
			if err != nil {
				return
			}

			request.client.responseChannel <- &Response{i, response}
		}(i, url)
	}

	for {
		result := <-request.client.responseChannel
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
