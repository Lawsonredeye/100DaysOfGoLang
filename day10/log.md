# Working with JSON Encoding & Decoding - Day 10

When it comes to working with http and http data/body data in golang are always passed as bytes. learning how to not only fetch but create json data is crucial to web development.

## ENCODING GO OBJECT TO JSON STRING
Here is a simple way using the golang standard library to convert a Go object into a json string/data for I/O operations.
```Go
package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	// Encoding is the process of coverting a Go Data into json bytes
	masnun := Person{
		Name: "Lawson Redeye",
		Age:  32,
		Emails: []string{
			"lawson@mail.com",
			"redeye@mail.com",
		},
	}
	json_bytes, err := json.Marshal(masnun)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s\n", json_bytes)
}

// Outputs:
// 2024/11/13 13:09:32 {"Name":"Lawson Redeye","Age":32,"Emails":["lawson@mail.com","redeye@mail.com"]}
```
Using struct tags can also help encode data in the standard json format. Here is how to create a struct tag to be properly formatted during encoding:
```Go
type Person struct {
	Name   string `json:"name"`
	Age    int	`json:"age"`
	Emails []string `json:"emails"`
}
```
Next to each field, we have provided tags to describe how this field should be marshalled or unmarshalled.

### Omitting fields
You can also omit empty fields in which you want to encode by using the omitempty keyword.
```Go
type Person struct {
    Name string `json:"name ,omitempty"`
    Age int `json:"age ,omitempty"`
    Emails []string `json:emails ,omitempty`
}
```

You can also use the - flag to prevent encoding of sensitive datas like password data from a user.
```Go
type Person struct {
	Name   string      `json:"name,omitempty"`
	Age    int         `json:"-"`
	Emails []string    `json:"emails,omitempty"`
}
```

So far we have used structs. We can also use maps and slices instead. Here’s a quick code example:
```Go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	masnun := map[string]interface{}{

		"name": "Lawsonredeye",
		"age":  27,
	}

	json_bytes, _ := json.Marshal(masnun)
	fmt.Printf("%s", json_bytes)
}
```

