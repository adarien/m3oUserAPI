package m3oUserAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const APIKey = "YOUR_APIKEY"

type UserID struct {
	ID string `json:"id"`
}

type NewUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type AssetResponse struct {
	Asset UserInfo `json:"account"`
}

type ErrorInfo struct {
	ID     string `json:"Id"`
	Code   int    `json:"Code"`
	Detail string `json:"Detail"`
	Status string `json:"Status"`
}

func (d UserInfo) Info() string {
	return fmt.Sprintf("User Info:\n\tID: %s\n\tName: %s\n\tEmail: %s\n"+
		"\tCreate Time: %s\n\tUpdate Time: %s",
		d.ID,
		d.Username,
		d.Email,
		d.Created,
		d.Updated,
	)
}

func (e ErrorInfo) Info() string {
	if e.Code == 500 {
		return fmt.Sprintf("%s", e.Detail)
	}
	return ""
}

func WorkAPI(body *bytes.Reader, method string) ([]byte, error) {
	bearer := "Bearer " + APIKey
	req, err := http.NewRequest("POST", "https://api.m3o.com/v1/user/"+method, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bd, nil
}

func CreateUser(id, username, email, password string) {
	data := NewUser{id, username, email, password}
	method := "Create"
	buildJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(buildJSON)

	_, err = WorkAPI(body, method)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Created")
}

func GetUserByID(id string) ([]byte, error) {
	data := UserID{id}
	method := "Read"
	buildJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(buildJSON)

	bd, err := WorkAPI(body, method)
	if err != nil {
		return nil, err
	}
	return bd, nil
}

func DeleteUserByID(id string) {
	data := UserID{id}
	method := "Delete"
	buildJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	body := bytes.NewReader(buildJSON)

	_, err = WorkAPI(body, method)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Removed")
}
