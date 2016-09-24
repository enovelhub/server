package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/enovelhub/capture/book"
	_ "github.com/enovelhub/enovelhub-website"
)

func Usage(w io.Writer) {
	fmt.Fprintln(w, "Usage of", os.Args[0], "book-data-path")
}

func main() {
	if len(os.Args) < 2 {
		Usage(os.Stderr)
		return
	}

	path := os.Args[1]
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read file", path, "failure")
		return
	}

	var book book.Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Fprintln(os.Stderr, path, "is not book struct(json)")
		return
	}

	fmt.Println(len(data))

	testdata_json := data
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path = "/index.html"
		}

		if r.URL.Path == "/testdata.json" {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(testdata_json)
			return
		}
		data, err := Asset(r.URL.Path[1:])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if strings.HasSuffix(r.URL.Path, ".json") {
			w.Header().Set("Content-type", "application/json")
		}
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-type", "text/css")
		}
		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-type", "script/javascript")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	/*
		http.HandleFunc("/testdata.json", func(w http.ResponseWriter, r *http.Request){
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w,"okok")

		})
	*/
	log.Print("Listen on 0.0.0.0:8904")
	log.Fatal(http.ListenAndServe("0.0.0.0:8904", nil))

	//	ShowDir("", 0)
}

func ShowDir(name string, level int) {
	items, err := AssetDir(name)
	if err != nil {
		return
	}

	for _, item := range items {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Println(item)
		if name != "" {
			item = name + "/" + item
		}

		ShowDir(item, level+1)
	}
}
