package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
}

//client with timeout
/*
func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &autorisationRoundTripper{
				next: &loggingRoundTripper{
					logger: os.Stdout,
					next:   http.DefaultTransport,
				},
			},
		},
	}, nil

}
*/

//client without timeout

func NewClient() (*Client, error) {

	return &Client{
		client: &http.Client{
			Transport: &autorisationRoundTripper{
				next: &loggingRoundTripper{
					logger: os.Stdout,
					next:   http.DefaultTransport,
				},
			},
		},
	}, nil

}

func (c *Client) GetUsers() ([]User, error) {
	resp, err := c.client.Get("http://127.0.0.1:8000/users")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r []User

	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) AddUser(u User) error {

	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/users", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req) //resp, err := c.client.PostForm("http://127.0.0.1:8000/users", url.Values{"name": {u.UserName}, "mail": {u.Mail}, "phone": {u.Phone}, "password": {u.Password}})
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func (c *Client) DeleteUserById(id string) error {
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8000/users/user?id="+id, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) GetUserById(id string) (User, error) {
	resp, err := c.client.Get("http://127.0.0.1:8000/users/user?id=" + id)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var r User

	if err = json.Unmarshal(body, &r); err != nil {
		return User{}, err
	}

	return r, nil
}

func (c *Client) UpdateUserById(id string, user User) error {

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", "http://127.0.0.1:8000/users/user?id="+id, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
