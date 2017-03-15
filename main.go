package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	baseapi         string
	airports        []string
	max             int
	defaultairports = []string{"JFK", "ORD", "MDW", "LGA", "LAX", "LGB", "SEA", "ROC", "DEN", "COS", "DAL", "HOU"}
)

// check for env variables
func init() {

	if os.Getenv("FAA_API") != "" {
		baseapi = os.Getenv("FAA_API")
	} else {
		baseapi = "http://services.faa.gov/airport/status"
	}

	if os.Getenv("FAA_MAX") != "" {
		var err error
		max, err = strconv.Atoi(os.Getenv("FAA_MAX"))
		if err != nil {
			max = 50
		}
	} else {
		max = 50
	}

	if os.Getenv("FAA_IATA") != "" {
		airports = strings.Split(os.Getenv("FAA_IATA"), ",")
	} else {
		airports = defaultairports
	}
}

func main() {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	api := baseapi + "/%s"
	times := r.Intn(max)

	fmt.Println(times)
	for i := 0; i < times; i++ {
		airport := airports[r.Intn(len(airports))]
		url := fmt.Sprintf(api, airport)
		code := getAirport(url)
		fmt.Println(code, url)
	}
}

func getAirport(url string) int {

	tspt := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tspt}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error", err.Error())
		return -1
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	return res.StatusCode
}
