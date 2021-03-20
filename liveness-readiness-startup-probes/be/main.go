// Package of main frontend
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// AppVersion - Application verison constant
const AppVersion = "V0.1.6"

const LiveTrigger = 5
const ReadyTrigger = 2

//CrashCounter = the number of ping requests before process exists
var CrashCounter int

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

// Pong Struct for ping pong text
type Pong struct {
	Respo string
	Stat  string
}

type apifunc func() []byte

var GitCommit string

//this value is returned in the Health-Header http header and checked against the livenessprobe of K8s
var readinessstate = "Healthy"
var livenessstate = "Healthy"

//HttpWrapper that accept a function that return []byte result and add http headers
func HttpWrapper(fn apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// respJSON := fn(kubeConfigPath)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Write(respJSON)
	}
}

func main() {

	readEnv()

	fmt.Println("Running Demo app BE version: " + AppVersion + " " + GitCommit)

	server := http.Server{
		Addr: ":" + apiServerPort,
		// Handler: srvmx,
	}
	// TODO: create a connection to the api and run a short test query
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Ping test")
		//if more then 3 ping requests accepted - crash process (exit 1)
		CrashCounter++
		if CrashCounter > ReadyTrigger {
			// os.Exit(1)
			readinessstate = "UnHealthy"
		}
		if CrashCounter > LiveTrigger {
			// os.Exit(1)
			livenessstate = "UnHealthy"
		}

		// fmt.Println("Readiness: " + readinessstate)
		// fmt.Println("Liveness: " + livenessstate)
		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/memleak", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// consumeMemory(1000000)

		log.Fatalln("Crashing")
		os.Exit(3)

	})

	// Liveness probe function
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		responseforping := Pong{"health", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Readiness-Header", readinessstate)
		w.Header().Set("Liveness-Header", livenessstate)
		// w.WriteHeader(500)

		w.Write(js)

	})

	http.HandleFunc("/highmem", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		for i := 0; i < 10; i++ {
			go consumeMemory(10000000)
		}
		// log.Fatalln("Crashing")
		// os.Exit(3)

	})

	http.HandleFunc("/cpuhog", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		for i := 0; i < 10; i++ {
			go cpuHog()
		}
		// log.Fatalln("Crashing")
		// os.Exit(3)

	})

	http.HandleFunc("/scenario3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Crashing....")
		crashme(server)

		//if more then 3 ping requests accepted - crash process (exit 1)

		// fmt.Println("Readiness: " + readinessstate)
		// fmt.Println("Liveness: " + livenessstate)
		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

		os.Exit(1)

	})
	// http.HandleFunc("/", HttpWrapper(runHomePage))

	fmt.Println("Listening on port " + apiServerPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}
}

func consumeMemory(n int) map[string]string {
	var memleakslice map[string]string
	memleakslice = make(map[string]string, n)
	for i := 1; i < n; i++ {
		memleakslice[string(i)] = RandStringRunes(1024)
	}
	return memleakslice
}
func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	fmt.Println(string(b))
	return string(b)
}

func cpuHog() {
	for i := 1; i < 10000; i++ {
		for z := 1; z < 1000; z++ {
			r := math.Pow(float64(z), float64(i))
			// fmt.Println(r)
			l := log.New(os.Stderr, "", 1)
			l.Println(string(i) + " " + string(z) + " " + fmt.Sprintf("%v", r))
		}
	}
}

func memCrashbe() []byte {
	url := os.Getenv("DEMOAPPBEURL") + "/api/v1/issues"

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
