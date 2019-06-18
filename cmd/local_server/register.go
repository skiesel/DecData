package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/skiesel/decdata/cmd/lib"
)

var (
	lookup = map[string]int64{}
)

func register(w http.ResponseWriter, r *http.Request) {
	log.Print("Local Server - Register")

	registerRequest, err := readLocalRequest(r)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error", 500)
		return
	}

	log.Printf("Local Server - %s : %s", registerRequest.Service, registerRequest.URL.String())

	registerResponse, err := forwardToAccessServer(registerRequest)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error", 500)
		return
	}

	log.Printf("Access Server Response - %s", registerResponse)

	output, err := json.Marshal(lib.LocalRegisterResponse{10000})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error", 500)
		return
	}

	lookup[registerRequest.Service] = 10000

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func readLocalRequest(r *http.Request) (*lib.LocalRegisterRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	registerRequest := &lib.LocalRegisterRequest{}
	err = json.Unmarshal(b, registerRequest)
	if err != nil {
		return nil, err
	}

	return registerRequest, nil
}

func forwardToAccessServer(registerRequest *lib.LocalRegisterRequest) (*lib.AccessRegisterResponse, error) {
	toAccessServer, err := json.Marshal(lib.AccessRegisterRequest{
		ClientID: CLIENT_ID,
		IP:       getOutboundIP(),
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(registerRequest.URL.String(), "application/json", bytes.NewReader(toAccessServer))
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	log.Print(registerRequest.URL.String())

	registerResponse := &lib.AccessRegisterResponse{}
	err = json.Unmarshal(b, registerRequest)
	if err != nil {
		return nil, err
	}

	return registerResponse, nil
}
