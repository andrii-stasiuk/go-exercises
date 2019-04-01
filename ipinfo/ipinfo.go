package main
import (
    "fmt"
    "flag"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "net"
)

type IPFields struct {
    AS string
    City string
    Country string
    CountryCode string
    ISP string
    Lat float64
    Lon float64
    Org string
    Query string
    Region string
    RegionName string
    Status string
    Timezone string
    Zip string
}

func main() {
    var ip string
    var ipPtr = flag.String("ip", "81.190.40.214", "a string")
    var geoPtr = flag.Bool("geo", false, "a bool")
    flag.Parse()

    if net.ParseIP( * ipPtr) != nil {
        ip = * ipPtr
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
    json.Unmarshal([] byte(body), & ipField)

    if *geoPtr {
        fmt.Printf("Loc: %.4f, %.4f", ipField.Lat, ipField.Lon)
    } else {
        fmt.Printf("IP address: %s\nOrganization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLoc: %.4f, %.4f\nPostal: %s",
            ipField.Query, ipField.Org, ipField.City, ipField.RegionName, ipField.Country, ipField.Lat, ipField.Lon, ipField.Zip)
    }
}