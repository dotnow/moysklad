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

// Client структура параметров запроса
type Client struct {
	endPoint string
	params   Params
	api      *APIClient
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

func (client *Client) makeURL(uuid ...string) string {
	baseURL := fmt.Sprintf("%s%s", APIPath, client.endPoint)

	if len(uuid) > 0 {
		baseURL = fmt.Sprintf("%s/%s", baseURL, uuid[0])
	}

	url, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	query := url.Query()

	if filter := client.params.getFilterString(); filter != "" {
		query.Set("filter", filter)
	}

	for key, value := range client.params.Query {
		query.Set(key, value)
	}

	url.RawQuery = query.Encode()
	return url.String()
}

// newRequest создаёт и возвращает ссылку на структуру запроса
func newRequest(api *APIClient, endPoint string, params Params) *Client {
	return &Client{
		endPoint: endPoint,
		params:   params,
		api:      api,
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
func (api *APIClient) getToken() string {
	return fmt.Sprintf(`Bearer %s`, api.token)
}

// setHeaders устанавливает необходимые заголовки к запросу
func (api *APIClient) setHeaders(req *http.Request) {
	req.Header.Set(`Authorization`, api.getToken())
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
func (client *Client) getByUUID(uuid string) (result []byte, err error) {
	url := client.makeURL(uuid)
	result, err = client.getRequest(url)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// getList возвращает множество объектов
func (client *Client) all() (data []byte, err error) {

	url := client.makeURL()

	res, err := client.getRequest(url)
	if err != nil {
		return nil, err
	}

	responseData, err := getResponseData(res)
	if err != nil {
		return nil, err
	}

	if responseData.Meta.Size > 1000 {
		urlList := makeURLList(url, responseData.Meta.Size)
		parallelResponse := client.parallelGet(urlList)

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
func (client *Client) getRequest(url string) (body []byte, err error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client.api.setHeaders(req)
	client.api.requestChannel <- struct{}{}

	log.Printf("Requesting: GET %s\n", url)

	res, err := client.api.httpClient.Do(req)
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
		<-client.api.requestChannel
	}()

	return body, nil
}

// parallelGet выполняет запросы, обрабатывает ответы и возвращает отсортированные результаты
func (client *Client) parallelGet(urlList []string) (results []Response) {

	for i, url := range urlList {
		go func(i int, url string) {
			response, err := client.getRequest(url)
			if err != nil {
				return
			}

			defer func(response []byte) {
				client.api.responseChannel <- &Response{i, response}
			}(response)
		}(i, url)
	}

	for {
		result := <-client.api.responseChannel
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
