package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/skiesel/decdata/cmd/lib"
)

func main() {
	accessServerURL, err := url.Parse("http://localhost:9000/register")
	if err != nil {
		log.Fatal(err)
	}

	registerRequest, err := json.Marshal(lib.LocalRegisterRequest{
		Service: "HelloWorld",
		URL:     accessServerURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	r, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewReader(registerRequest))
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Response read")

	registerResponse := &lib.LocalRegisterResponse{}
	err = json.Unmarshal(b, registerResponse)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Local App Server: %d", registerResponse.Port)

	http.HandleFunc("/", run)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", registerResponse.Port), nil))
}

func run(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
