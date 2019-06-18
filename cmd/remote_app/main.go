package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/skiesel/decdata/cmd/lib"
)

const (
	CLIENT_ID = "MY_CLIENT_ID"
	SERVICE   = "HelloWorld"
)

func main() {
	accessRequest, err := json.Marshal(lib.AccessRequest{
		ClientID: CLIENT_ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	r, err := http.Post("http://localhost:9000/syncAccess", "application/json", bytes.NewReader(accessRequest))
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	accessResponse := &lib.AccessResponse{}
	err = json.Unmarshal(b, accessResponse)
	if err != nil {
		log.Fatal(err)
	}

	dataRequest, err := json.Marshal(lib.LocalDataRequest{
		Service: SERVICE,
	})
	if err != nil {
		log.Fatal(err)
	}

	//fix this to use accessResponse.IP
	r, err = http.Post("http://localhost:8080", "application/json", bytes.NewReader(dataRequest))
	if err != nil {
		log.Fatal(err)
	}

	b, err = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Got Data: " + string(b))
}
