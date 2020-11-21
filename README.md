![](https://habrastorage.org/getpro/moikrug/uploads/company/100/005/687/1/logo/medium_f0e56c86b88d25e4c2aff41562bfd499.png)

# moysklad

Библиотека для работы с [JSON API Moysklad](https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api), написанная на Go

> **Внимание!** Библиотека находится в стадии разработки

## Пример использования

```go
ms := NewClient(API_TOKEN)

params := Params{Query: map[string]string{
	"limit": "10",
}}

// Получаем срез сущностей Товар
response, err := ms.Product(params).Get()
```
