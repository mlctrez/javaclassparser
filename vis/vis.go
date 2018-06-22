package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func writeFile(name string, writer http.ResponseWriter) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(writer, file)
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/":
			writeFile("index.html", writer)
		case "/data.json":
			data(writer)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func data(writer http.ResponseWriter) {

	writer.Header().Set("Content-Type", "application/json")

	deps, err := os.Open("../deps.json")

	if err != nil {
		writeFile("data.json", writer)
		return
	}
	defer deps.Close()
	dd := &DepData{}
	err = json.NewDecoder(deps).Decode(dd)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	d := &Data{Nodes: []Node{}, Links: []Link{}}

	for i, u := range dd.Unique {
		d.Nodes = append(d.Nodes, Node{Id: u, Group: i})
	}

	for k, v := range dd.Reverse {
		for p, q := range v {
			if q > 21 {
				q = 21
			}
			value := int(21-q) / 2

			if q < 5 {
				d.Links = append(d.Links, Link{Source: k, Target: p, Value: value + 10})
			}
		}
	}


	bytes, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.Write(bytes)
	}
}

type Data struct {
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

type Node struct {
	Id    string `json:"id"`
	Group int    `json:"group"`
}
type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}

type DepData struct {
	Unique  []string `json:"unique"`
	Forward DepMap   `json:"forward"`
	Reverse DepMap   `json:"reverse"`
}

type DepMap map[string]map[string]uint16
