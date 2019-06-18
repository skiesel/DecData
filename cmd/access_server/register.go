package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/skiesel/decdata/cmd/lib"
)

var (
	lookup = map[string]net.IP{}
)

func register(w http.ResponseWriter, r *http.Request) {
	log.Print("Access Server - Register")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	var registerRequest lib.AccessRegisterRequest
	err = json.Unmarshal(b, &registerRequest)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("Access Server - %s : %s", registerRequest.ClientID, registerRequest.IP)

	lookup[registerRequest.ClientID] = registerRequest.IP

	output, err := json.Marshal(lib.AccessRegisterResponse{})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
