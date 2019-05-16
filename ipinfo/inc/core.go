package core

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

func GetFlags() (string, bool) {
	var ipPtr = flag.String("ip", "81.190.40.214", "a string")
	var geoPtr = flag.Bool("geo", false, "a bool")
	flag.Parse()
	return *ipPtr, *geoPtr
}

func SetIP(ipFlag string) string {
	var ip string
	if net.ParseIP(ipFlag) != nil {
		ip = ipFlag
	} else {
		ip = "81.190.40.214"
	}
	return ip
}

func GetIPInfo(ipAddr string) IPFields {
	resp, err := http.Get("http://ip-api.com/json/" + ipAddr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ipField IPFields
	json.Unmarshal([]byte(body), &ipField)
	return ipField
}

func Printer(geoPos bool, geoData IPFields) {
	if geoPos {
		fmt.Printf("Location: %.4f, %.4f", geoData.Lat, geoData.Lon)
	} else {
		fmt.Printf("IP address: %s\nOrganization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLoc: %.4f, %.4f\nPostal: %s",
			geoData.Query, geoData.Org, geoData.City, geoData.RegionName, geoData.Country, geoData.Lat, geoData.Lon, geoData.Zip)
	}
}
