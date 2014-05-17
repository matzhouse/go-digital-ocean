package digitalocean

import (
	"fmt"
)

type Events struct {
	Response
	Event Event `json:"event,omitempty"`
}

type Event struct {
	Id           int    `json:"id,omitempty"`
	ActionStatus string `json:"action_status,omitempty"`
	DropletId    int    `json:"droplet_id,omitempty"`
	EventTypeId  int    `json:"event_type_id,omitempty"`
	Percentage   int    `json:"percentage,omitempty"`
}

func (c Client) Event(id int) (r Request, err error) {

	// initialise the struct so json won't complain
	var e *Events = &Events{}

	uri := fmt.Sprintf("/events/%d", id)

	req := Request{Uri: uri, Response: e}

	c.do(&req)

	return req, nil

}
