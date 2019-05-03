package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CheckIp(toCheck string) string {
	var fields string = "?fields=1273"
	var ipApi = "http://ip-api.com/json/"
	resp, err := http.Get(ipApi + toCheck + fields)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var response Response
	json.Unmarshal(body, &response)

	var ipAdress string = "IP address: " + toCheck
	var organization string = "Organization: " + response.Org
	var city string = "City: " + response.City
	var region string = "Region: " + response.RegionName
	var country string = "Country: " + response.Country
	var loc string = "Loc: " + FloatToString(response.Lat) + ", " + FloatToString(response.Lon)
	var postal string = "Postal: " + response.Zip
	var ipInfo = ipAdress + "\n" + organization + "\n" + city + "\n" + region + "\n" + country + "\n" + loc + "\n" + postal

	return ipInfo
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 4, 64)
}

type Response struct {
	City       string        `json:"city"`
	Country    string        `json:"country"`
	Lat        float64       `json:"lat"`
	Lon        float64       `json:"lon"`
	Org        string        `json:"org"`
	RegionName string        `json:"regionName"`
	Zip        string        `json:"zip"`
	IpCheck    []interface{} `json:"ipCheck"`
}

func main() {

	ip := flag.String("ip", "81.190.40.214", "ip address to check")
	flag.Parse()
	var ipToCheck string = *ip
	print(CheckIp(ipToCheck))
}
