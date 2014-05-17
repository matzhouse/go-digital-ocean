package digitalocean

import ()

type Regions struct {
	Response
	Regions []Region `json:"regions,omitempty"`
}

type Region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c Client) Regions() (r Request, err error) {

	// initialise the struct so json won't complain
	var rg *Regions = &Regions{}

	req := Request{Uri: "/regions/", Response: rg}

	c.do(&req)

	return req, nil

}
