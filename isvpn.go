package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type OvpnServer struct {
	ID            int
	Datacenter_Id int
	Type_Id       int
	Hidden        bool
	Online        bool
	Name          string
	IP            string
	MacAddress    string
	PTR           string
	Cores         int
	Portspeed     int
	Created_at    string
	Updated_at    string
	Filename      string
}

type Ovpn struct {
	Status bool
	IP     string
	PTR    string
	Server OvpnServer
}

func main() {
	ovpn := Ovpn{}
	err := getJson("https://www.ovpn.com/v1/api/connected", &ovpn)

	if err != nil {
		fmt.Print("VPN: ?")
		return
	}

	if ovpn.Status == true {
		fmt.Print(ovpn.Server.Name)
	} else {
		fmt.Print("No VPN")
	}
}
