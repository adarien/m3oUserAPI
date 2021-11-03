# Пример использования API сайта m3o.com 
## дорабатывается...

```go
package main

import (
	"https://github.com/adarien/m3oUserAPI"
)

const APIKey = "YOUR_APIKEY"

func main() {
	client := m3oUserAPI.NewClientAPI(APIKey)

	// Создаём нового пользователя
	id := "usrid-2"
	username := "petya"
	email := "petya@ninja.go"
	password := "qwerty123"
	client.CreateUser(id, username, email, password)

	// Выводим информацию по ID
	client.GetUserByID(id)

	// Удаляем пользователя по ID
	client.DeleteUserByID(id)

	// Проверяем статус ошибки при получении информации об удалённом пользователе
	client.GetUserByID(id)
}
```
