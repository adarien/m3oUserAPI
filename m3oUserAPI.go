package m3oUserAPI

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://api.m3o.com/v1/user/"

type UserID struct {
	ID string `json:"id"`
}

type CreateUserInput struct {
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

type Client struct {
	token string
}

func NewClientAPI(APIKey string) *Client {
	return &Client{token: APIKey}
}

func (e ErrorInfo) Info() string {
	if e.Code == 500 {
		return fmt.Sprintf("%s", e.Detail)
	}
	return ""
}

func (t *Client) WorkAPI(body *bytes.Reader, method string) ([]byte, error) {
	bearer := "Bearer " + t.token
	req, err := http.NewRequest(http.MethodPost, URL+method, body)
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

func (t *Client) CreateUser(c CreateUserInput) error {
	method := "Create"
	buildJSON, err := json.Marshal(c)
	if err != nil {
		return err
	}
	body := bytes.NewReader(buildJSON)

	_, err = t.WorkAPI(body, method)
	if err != nil {
		return err
	}
	return nil
}

func (t *Client) GetUserByID(id string) (UserInfo, error) {
	data := UserID{id}
	method := "Read"
	buildJSON, err := json.Marshal(data)
	if err != nil {
		return UserInfo{}, err
	}
	body := bytes.NewReader(buildJSON)
	bd, err := t.WorkAPI(body, method)
	if err != nil {
		return UserInfo{}, err
	}

	var exist ErrorInfo
	if err = json.Unmarshal(bd, &exist); err != nil {
		return UserInfo{}, err
	}

	if exist.Info() == "not found" {
		er_text := fmt.Sprintf("ID %s not found\n", id)
		return UserInfo{}, errors.New(er_text)
	}
	var r AssetResponse
	if err = json.Unmarshal(bd, &r); err != nil {
		return UserInfo{}, err
	}

	return r.Asset, nil
}

func (t *Client) DeleteUserByID(id string) error {
	data := UserID{id}
	method := "Delete"
	buildJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(buildJSON)

	_, err = t.WorkAPI(body, method)
	if err != nil {
		return err
	}
	return nil
}
