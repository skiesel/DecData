package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/skiesel/decdata/cmd/lib"
)

func syncAccess(w http.ResponseWriter, r *http.Request) {
	dataRequest, err := readDataRequest(r)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error", 500)
		return
	}

	if ip, found := lookup[dataRequest.ClientID]; found {
		sendResponseData(w, ip)
	} else {
		log.Print("ClientID not found")
		http.Error(w, "Error", 500)
	}
}

func readDataRequest(r *http.Request) (*lib.AccessRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	accessRequest := &lib.AccessRequest{}
	err = json.Unmarshal(b, accessRequest)
	if err != nil {
		return nil, err
	}

	return accessRequest, nil
}

func sendResponseData(w http.ResponseWriter, ip net.IP) {
	response, err := json.Marshal(lib.AccessResponse{
		IP: ip,
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "Error", 500)
		return
	}
	w.Write(response)
}
