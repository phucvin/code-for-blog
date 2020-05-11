package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// TODO: the fruits is a good example of varying struct type -- there is no
// static type that would fit. Maybe do it in a followup?
var jsonText = []byte(`
{
  "attrs": [
		{
			"name": "color",
			"count": 9
		},
		{
			"name": "family",
			"count": 127
		}],
	"fruits": [
		{
			"name": "orange",
			"sweetness": 12.3,
			"attr": {"family": "citrus"}
		}
	]
}`)

type AutoGenerated struct {
	Attrs []struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	} `json:"attrs"`
	Fruits []struct {
		Name      string  `json:"name"`
		Sweetness float64 `json:"sweetness"`
		Attr      struct {
			Family string `json:"family"`
		} `json:"attr"`
	} `json:"fruits"`
}

func asMapGeneric() {
	var m map[string]interface{}
	if err := json.Unmarshal(jsonText, &m); err != nil {
		log.Fatal(err)
	}

	fruits, ok := m["fruits"]
	if !ok {
		log.Fatal("'fruits' field not found")
	}
	fslice, ok := fruits.([]interface{})
	if !ok {
		log.Fatal("'fruits' field not a map")
	}

	for _, f := range fslice {
		fmap, ok := f.(map[string]interface{})
		if !ok {
			log.Fatal("'fruits' element not a map")
		}

		name, ok := fmap["name"]
		if !ok {
			log.Fatal("fruits element has no 'name' field")
		}
		sweetness, ok := fmap["sweetness"]
		if !ok {
			log.Fatal("fruits element has no 'sweetness' field")
		}

		fmt.Printf("%s -> %f\n", name, sweetness)
	}
}

func asMapGenericNoErr() {
	var m map[string]interface{}
	if err := json.Unmarshal(jsonText, &m); err != nil {
		log.Fatal(err)
	}

	fruits := m["fruits"].([]interface{})
	for _, f := range fruits {
		fruit := f.(map[string]interface{})
		fmt.Printf("%s -> %f\n", fruit["name"], fruit["sweetness"])
	}
}

func asStructFull() {
	var ag AutoGenerated
	if err := json.Unmarshal(jsonText, &ag); err != nil {
		log.Fatal(err)
	}
	for _, fruit := range ag.Fruits {
		fmt.Printf("%s -> %f\n", fruit.Name, fruit.Sweetness)
	}
}

type Fruit struct {
	Name      string            `json:"name"`
	Sweetness float64           `json:"sweetness"`
	Attr      map[string]string `json:"attr"`
}

func asHybrid() {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(jsonText, &m); err != nil {
		log.Fatal(err)
	}

	fruitsRaw, ok := m["fruits"]
	if !ok {
		log.Fatal("expected to find 'fruits'")
	}

	var fruits []Fruit
	if err := json.Unmarshal(fruitsRaw, &fruits); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fruits)
}

func main() {
	asMapGeneric()

	asMapGenericNoErr()

	//asStructFull()

	//asHybrid()
}
