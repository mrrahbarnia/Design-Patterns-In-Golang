package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ToDo struct {
	ID        int    `json:"id" xml:"id"`
	UserId    int    `json:"user_id" xml:"user_id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

type DataInterface interface {
	GetData() (*ToDo, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*ToDo, error) {
	return rs.Remote.GetData()
}

type JSONBackend struct{}
type XMLBackend struct{}

func (jb *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var todo ToDo
	if err := json.Unmarshal(body, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (xb *XMLBackend) GetData() (*ToDo, error) {
	xmlFile := `
	<?xml version="1.0" encoding="UTF-8" ?>
	<root>
		<id>1</id>
		<user_id>1</user_id>
		<title>delectus aut autem</title>
		<completed>false</completed>
	</root>
	`
	var todo ToDo
	if err := xml.Unmarshal([]byte(xmlFile), &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func main() {
	// Without adapter
	data, err := getRemoteData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TODO without adapter: \t", data.ID, data.Title)

	// JSON data with adapter
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tdFromJSON, _ := jsonAdapter.Remote.GetData()
	fmt.Println("TODO JSON with adapter: \t", tdFromJSON.ID, tdFromJSON.Title)

	// XML data with adapter
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tdFromXML, _ := xmlAdapter.Remote.GetData()
	fmt.Println("TODO XML with adapter: \t", tdFromXML.ID, tdFromXML.Title)
}

func getRemoteData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var todo ToDo
	if err := json.Unmarshal(body, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}
