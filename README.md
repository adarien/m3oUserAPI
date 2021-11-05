# Пример использования API сайта m3o.com 
## дорабатывается...

```go
package main

import (
	"fmt"
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
	result, err := client.CreateUser(id, username, email, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Выводим информацию по ID
	result, err = client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Удаляем пользователя по ID
	result, err = client.DeleteUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Проверяем статус ошибки при получении информации об удалённом пользователе
	result, err = client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

```
