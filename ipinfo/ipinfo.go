package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type IPFields struct {
	AS          string  `json:"as"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	ISP         string  `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

func main() {
	var ip string
	var ipPtr = flag.String("ip", "81.190.40.214", "a string")
	var geoPtr = flag.Bool("geo", false, "a bool")
	flag.Parse()

	if net.ParseIP(*ipPtr) != nil {
		ip = *ipPtr
	} else {
		ip = "81.190.40.214"
	}

	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ipField IPFields
	json.Unmarshal([]byte(body), &ipField)

	if *geoPtr {
		fmt.Printf("Loc: %.4f, %.4f", ipField.Lat, ipField.Lon)
	} else {
		fmt.Printf("IP address: %s\nOrganization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLoc: %.4f, %.4f\nPostal: %s",
			ipField.Query, ipField.Org, ipField.City, ipField.RegionName, ipField.Country, ipField.Lat, ipField.Lon, ipField.Zip)
	}
}
