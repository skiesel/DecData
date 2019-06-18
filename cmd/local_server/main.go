package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/skiesel/decdata/cmd/lib"
)

const (
	CLIENT_ID = "MY_CLIENT_ID"
)

func main() {
	log.Print("Local Server")
	http.HandleFunc("/register", register)
	http.HandleFunc("/", access)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func access(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return
	}

	dataRequest := &lib.LocalDataRequest{}
	err = json.Unmarshal(b, dataRequest)
	if err != nil {
		return
	}

	if port, found := lookup[dataRequest.Service]; found {
		r, err := http.Post(fmt.Sprintf("http://localhost:%d", port), "application/json", bytes.NewReader(b))
		if err != nil {
			return
		}

		b, err = ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			return
		}

		w.Write(b)
	}
}
