package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json video")
	DecodeJson()
}

func EncodeJson() {
	lcocourse := []course{
		{Name: "ReactJS", Price: 299, Platform: "learncodeonline.in", Password: "abc123", Tags: []string{"web-dev", "js"}},
		{Name: "JavaFullStack", Price: 499, Platform: "learncodeonline.in", Password: "abc123", Tags: []string{"full-stack", "java"}},
		{Name: "GCP dev", Price: 399, Platform: "learncodeonline.in", Password: "abc123", Tags: nil},
	}

	//Package this data as json data

	// finalJson, err := json.Marshal(lcocourse)
	finalJson, err := json.MarshalIndent(lcocourse, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("FinalJson: ", string(finalJson))
}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
	[{
		"coursename": "ReactJS",
		"price": 299,
		"website": "learncodeonline.in",
		"tags": [
				"web-dev",
				"js"
		]
	},
	{
			"coursename": "JavaFullStack",
			"price": 499,
			"website": "learncodeonline.in",
			"tags": [
					"full-stack",
					"java"
			]
	},
	{
			"coursename": "GCP dev",
			"price": 399,
			"website": "learncodeonline.in"
	}]
	`)

	var lcoCourse []course

	checkValid := json.Valid(jsonDatafromWeb)

	if checkValid {
		fmt.Println("JSON Valid")
		err := json.Unmarshal(jsonDatafromWeb, &lcoCourse)
		if err != nil {
			log.Panic(err)
		}
		// fmt.Printf("%v\n", lcoCourse)
		finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("FinalJson: ", string(finalJson))
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	var myOnlineData []map[string]interface{}

	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%v\n\n\n", myOnlineData)
	for _, mp := range myOnlineData {
		for k, v := range mp {

			if arr, ok := v.([]interface{}); ok {
				fmt.Printf("Key is %v and value id %v of type []string{}\n", k, v)
				for _, i := range arr {
					fmt.Printf("\tAnd tags is: %v\n", i)
				}
			} else {
				fmt.Printf("Key is %v and value id %v of type %T\n", k, v, v)
			}
		}
	}
}
