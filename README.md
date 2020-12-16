![](https://habrastorage.org/getpro/moikrug/uploads/company/100/005/687/1/logo/medium_f0e56c86b88d25e4c2aff41562bfd499.png)

# moysklad

Модуль для работы с [JSON API Moysklad](https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api), написанная на Golang

> **Внимание!** Библиотека находится в стадии разработки

## Установка

```
go get -u github.com/dotnow/moysklad
```

## Пример использования

```go
...

ms := moysklad.Moysklad(API_TOKEN)

params := client.Params{}
params.AddQuery("limit", "10")

// Получаем срез сущностей Товар
products, err := ms.Product(params).GetList()
if err != nil {
	log.Fatal(err)
}

for index, element := range products {
	fmt.Printf("%d. %s\n", index+1, element.Name)
}

...
```
