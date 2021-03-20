package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	var baseCur string
	fmt.Println("Enter Base Currency")
	fmt.Scanln(&baseCur)

	var targetCur string
	fmt.Println("Enter Target Currency")
	fmt.Scanln(&targetCur)

	var curVolume float32
	fmt.Println("Enter Volume")
	fmt.Scanln(&curVolume)

	value, err := getCurrencyConversion(baseCur, targetCur)
	if err != nil {
		log.Fatal(err)
	}

	var finalVolume float32 = curVolume * value

	fmt.Printf("Currency conversion successful for Base Currency %s and Target Currency %s. Value: %f Amount: %f ", baseCur, targetCur, value, finalVolume)

}

func getCurrencyConversion(baseCur string, targetCur string) (value float32, err error) {

	curConvert := baseCur + "_" + targetCur

	url := "https://free.currconv.com/api/v7/convert?q=" + curConvert + "&compact=ultra&apiKey=424d5c06590ef708b468"
	fmt.Println(url)

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return value, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, err := spaceClient.Do(req)
	if err != nil {
		return value, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return value, err
	}

	convert := make(map[string]float32)

	err = json.Unmarshal(body, &convert)
	if err != nil {
		return value, err
	}

	value = convert[curConvert]

	return value, nil
}
