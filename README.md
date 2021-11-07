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
	err := client.CreateUser(m3oUserAPI.CreateUserInput{
		ID: id, Username: username, Email: email, Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Created")

	// Выводим информацию по ID
	user, err := client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.ID, user.Username, user.Email)

	// Удаляем пользователя по ID
	err = client.DeleteUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Removed")

	// Проверяем статус ошибки при получении информации об удалённом пользователе
	user, err = client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
```
