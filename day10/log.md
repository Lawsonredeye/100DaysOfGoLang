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

Here is how to use slice for creating data for encoding:
```Go
func main() {
	emails := []string{"lawson@mail.com", "redeye@mail.com"}
	json_bytes, err := json.Marshal(emails)
	check(err)
	log.Printf("%s", json_bytes)
	// outputs:
	// 2024/11/13 13:31:36 ["lawson@mail.com","redeye@mail.com"]
}
```

## DECODING JSON DATA TO GO OBJECT
Decoding converts a json data into a Go data (most cases from a byte to into go object).

Here is a simple way to do such:
```Go
type Person struct {
	Name   string `json:"name"`
	Age    int	`json:"age"`
	Emails []string `json:"emails"`
}

func main() {
	json_bytes := []byte(`
	{
		"Name":"Lawson Redeye",
		"Age":27,
		"Emails":["lawson@mail.com","redeye@mail.com"]
	}
	`)
	masnun := Person{}
	err := json.Unmarshal(json_bytes, &masnun)
	check(err)
	log.Println(masnun.Name, masnun.Age, masnun.Emails)
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
```
It is important to note that in Go when Unmarshal'ling a json string into a struct, fields or members of that struct in which isnt defined on the json string are ignored and if available on the json string but the field isnt defined on the struct then the struct object would ignore it during decoding & unmarshalling.
```Go
type Person struct {
	Name   string `json:"name"`
	Age    int	`json:"age"`
	Emails []string `json:"emails"`
	Address string
}

func main() {
	json_bytes := []byte(`
		{
			"Name":"Lawson Redeye",
			"Age":27,
			"Emails":["lawson@mail.com","redeye@mail.com"],
			"Score":97
		}
	`)
	masnun := Person{}
	err := json.Unmarshal(json_bytes, &masnun)
	check(err)
	log.Println(masnun.Address)
	// Output:
	// Nothing
}
```
In the above example, the struct has a field named Address which the JSON doesn’t provide.
On the other hand, the JSON has the Score key which the struct knows nothing about. In this case, masnun.Address will be empty string.

```Go

func main() {
	json_bytes := []byte(`
		{
			"Name":"Lawson Redeye",
			"Age":27,
			"Emails":["lawson@mail.com","redeye@mail.com"],
			"Score":97
		}
	`)
	
	var pData map[string]interface{}
	err := json.Unmarshal(json_bytes, &pData)
	check(err)
	log.Println(pData["Score"], pData["Age"], pData["Emails"])
	// Output:
	// 97 27 [lawson@mail.com redeye@mail.com]
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
```
Note: The value of each key in the map is of type interface and if you want to access such data you have to convert it.

## RESOURCES
1.[working with JSON in GOlang](https://masnun.com/category/golang)
