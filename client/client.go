package client

import (
	"net/http"
	"strconv"
	"time"
)

// APIPath базовый URL к API МойСклад
const APIPath = "https://online.moysklad.ru/api/remap/1.2/"

// MaxLimitOffset лимит для запросов (с вложенностями 100)
const MaxLimitOffset = 100

// MaxLimitOffsetString возвращает MaxLimitOffset в виде строки
func MaxLimitOffsetString() string {
	return strconv.Itoa(MaxLimitOffset)
}

// APIClient экземпляр клиента МойСклад с токеном
type APIClient struct {
	token           string
	requestChannel  chan struct{}
	responseChannel chan *Response
	httpClient      *http.Client
}

// Moysklad создаёт и возвращает ссылку на экземпляр APIClient
func Moysklad(token string) *APIClient {
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
