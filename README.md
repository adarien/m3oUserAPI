# Пример использования API сайта m3o.com 
## дорабатывается...

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/adarien/m3oUserAPI"
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
	res, err := client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	var r client.AssetResponse
	if err = json.Unmarshal(res, &r); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Asset.Info())

	// Удаляем пользователя по ID
	client.DeleteUserByID(id)

	// Проверяем статус ошибки при получении информации об удалённом пользователе
	res, err = client.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
	}
	var exist client.ErrorInfo
	if err = json.Unmarshal(res, &exist); err != nil {
		log.Fatal(err)
	}
	fmt.Println(exist.Info())
}

```
