package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  string `json:"status"`
	Error   string `json:"message,omitempty"`
	Message string `json:"error_message,omitempty"`
}

type auth struct {
	client_id string
	api_key   string
}

type Client struct {
	Auth auth
}

type Request struct {
	Uri      string
	Response interface{}
}

func NewClient() Client {

	c := Client{
		Auth: auth{},
	}

	cid := os.Getenv("DOCLIENTID")

	if cid == "" {
		log.Fatal("Digital Ocean client id not available")
	}

	apikey := os.Getenv("DOAPIKEY")

	if apikey == "" {
		log.Fatal("Digital Ocean api key not available")
	}

	c.Auth.client_id = cid
	c.Auth.api_key = apikey

	return c

}

func (a auth) string() string {
	return fmt.Sprintf("?client_id=%s&api_key=%s", a.client_id, a.api_key)
}

func (c Client) do(r *Request) error {

	var dourl bytes.Buffer

	// make the url
	dourl.WriteString("https://api.digitalocean.com/v1")
	dourl.WriteString(r.Uri)
	dourl.WriteString(c.Auth.string())

	res, err := http.Get(dourl.String())

	resp := r.Response

	if err != nil {
		return err
	}

	rj, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(rj, resp)

	if err != nil {
		log.Fatal(err)
	}

	return nil

}
