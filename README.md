# Пример использования API сайта m3o.com 
## дорабатывается...

```go
package main

import (
	"log"
	"m3oUser/m3oUserAPI"
)

const APIKey = "YOUR_APIKEY"

func main() {
	client := m3oUserAPI.NewClientAPI(APIKey)

	// Создаём нового пользователя
	id := "usrid-2"
	username := "petya"
	email := "petya@ninja.go"
	password := "qwerty123"
	err := client.CreateUser(id, username, email, password)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим информацию по ID
	err = client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}

	// Удаляем пользователя по ID
	err = client.DeleteUserByID(id)
	if err != nil {
		log.Fatal(err)
	}

	// Проверяем статус ошибки при получении информации об удалённом пользователе
	err = client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
}

```
