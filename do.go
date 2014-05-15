package digitalocean

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"encoding/json"
)

type auth struct {
	client_id string
	api_key   string
}

type Client struct {
	Auth auth
}

type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Account struct {
	Droplets []Droplet
	Regions  []Region
	Images   []Image
}

type Request struct {
	Uri      string
	Response interface{}
}

type Droplet struct {
	Id                 int
	Name               string
	Image_id           int
	Size_id            int
	Region_id          int
	Backups_active     int
	Ip_address         string
	Private_ip_address string
	Locked             bool
	Status             string
	Created_at         string
}

type Region struct {
}

type Image struct {
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

	return c

}

func (a auth) String() string {
	return fmt.Sprintf("?client_id=%s&api_key=%s", a.client_id, a.api_key)
}

func (c Client) do(r *Request) Response, err {

	var dourl bytes.Buffer

	// make the url
	dourl.WriteString("https://api.digitalocean.com/v1")
	dourl.WriteString(r.Uri)
	dourl.WriteString(c.Auth.String())

	res, err := http.Get(dourl.String())
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	r := Response{}

	err = json.Unmarshal(res, r)

	if err != nil {
		return nil, err
	}

	return r, nil

}

func (c Client) Droplets() (r Request, err error) {

	r = Request{Uri: "/droplets/"}

	c.do(&req)

	//return 0, errors.New("math: square root of negative number")

	return r, nil

}
