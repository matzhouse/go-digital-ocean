package digitalocean

import (
	"fmt"
)

type Droplets struct {
	Response
	Droplet  Instance   `json:"droplet,omitempty"`
	Droplets []Instance `json:"droplets,omitempty"`
}

type Instance struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Image     int    `json:"image_id"`
	Size      int    `json:"size_id"`
	Region    int    `json:"region_id"`
	Backups   bool   `json:"backups_active"`
	PublicIp  string `json:"ip_address"`
	PrivateIp string `json:"private_id_address"`
	Locked    bool   `json:"locked"`
	Created   string `json:"created_at"`
}

func (c Client) Droplets() (r Request, err error) {

	// initialise the struct so json won't complain
	var d *Droplets = &Droplets{}

	req := Request{Uri: "/droplets/", Response: d}

	c.do(&req)

	return req, nil

}

func (c Client) Droplet(id int) (r Request, err error) {

	var d *Droplets = &Droplets{}

	uri := fmt.Sprintf("/droplets/%d", id)

	req := Request{Uri: uri, Response: d}

	c.do(&req)

	return req, nil

}
