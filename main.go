package moysklad

import (
	"github.com/dotnow/moysklad/client"
)

// Moysklad создаёт и возвращает ссылку на экземпляр APIClient
func Moysklad(token string) *client.APIClient {
	return client.Moysklad(token)
}
