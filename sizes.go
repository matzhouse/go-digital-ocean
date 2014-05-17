package digitalocean

import ()

type Sizes struct {
	Response
	Sizes []Size `json:"Sizes,omitempty"`
}

type Size struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c Client) Sizes() (r Request, err error) {

	// initialise the struct so json won't complain
	var sz *Sizes = &Sizes{}

	req := Request{Uri: "/sizes/", Response: sz}

	c.do(&req)

	return req, nil

}
