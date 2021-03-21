package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CrashBE crash the backened for demo scenarion 3
func CrashBE() []byte {
	url := demoappBeURL + "/scenario3"
	fmt.Println("Fetching from: " + url)

	jsonValue, _ := json.Marshal("req")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := body
	// fmt.Println(res)
	return res
}

func contactDemoAppBe() []byte {
	url := demoappBeURL + "/ping"
	fmt.Println("Fetching from: " + url)

	jsonValue, _ := json.Marshal("req")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := body
	// fmt.Println(res)
	return res
}

func checkBElive() bool {
	var health Pong
	var belive bool

	url := demoappBeURL + "/health"
	// fmt.Println("Fetching from: " + url)

	jsonValue, _ := json.Marshal("req")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		fmt.Println("Reaindess check BE failed")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := body
	json.Unmarshal(res, &health)

	if health.Respo == "health" && health.Stat == "ok" {
		belive = true
	} else {
		belive = false
	}
	return belive
}
