go-digital-ocean
================

Another Digital Ocean package. Aim is to make it complete as a learning exercise.

#### Completed parts
- Client creation
- regions
- sizes
- events (50%)
- auth

#### To do
- tests!
- ssh keys
- far better error management
- utilities ie. how much am i spending at the moment


================

### Example

```go
package main

import (
	"fmt"
	"github.com/matzhouse/go-digital-ocean"
	"log"
)

func main() {

	c := digitalocean.NewClient()

// Example of getting the a list of active droplets
	req, err := c.Droplets()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(req.Response)

// Example of getting the current regions
	req, err = c.Regions()
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(req.Response)

// Example of getting the current sizes
	req, err = c.Sizes()
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(req.Response)

}

```
