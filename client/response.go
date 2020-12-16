package client

import (
	"encoding/json"
	"errors"
)

// Response структура для хранения ответа параллельных запросов
type Response struct {
	Index int
	Body  []byte
}

// ResponseData структура для десериализации ответа
type ResponseData struct {
	Meta struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"meta"`
	Rows   []map[string]interface{} `json:"rows"`
	Errors []struct {
		Error     string `json:"error"`
		Parameter string `json:"parameter"`
		Code      int    `json:"code"`
		Message   int    `json:"error_message"`
	}
}

// getResponseData производит десериализацию ответа и возвращает размер meta и массив найденных объектов
func getResponseData(response []byte) (data *ResponseData, err error) {

	if err = json.Unmarshal(response, &data); err != nil {
		return nil, err
	}

	if len(data.Errors) > 0 {
		err = errors.New(data.Errors[0].Error)
		return nil, err
	}

	return data, nil
}
