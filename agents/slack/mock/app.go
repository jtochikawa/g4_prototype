package main

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
)

const (
	URI = ""
	DATA_CONTENT_TYPE = "application/json"
)


type Data struct {
	Text string `json:"text"`
}

func main() {
	data, err := json.Marshal(&Data{
		Text: "Hello World!",
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(URI, DATA_CONTENT_TYPE,
		bytes.NewBuffer(data))
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", resp.Status)
}
