// Package of main frontend
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/sessions"
)

//GitCommit get commitid from git
var GitCommit string

// AppVersion - Application verison constant
var AppVersion = "V0.8.3" + "-" + GitCommit

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

// Pong Struct for ping pong text
type Pong struct {
	Respo string `json:"resp"`
	Stat  string `json:"stat"`
}

type apifunc func() []byte

var healthstate = "Ready"

// Http wrapper that accept a function that return []byte result and add http headers
func HttpWrapper(fn apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// respJSON := fn(kubeConfigPath)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Write(respJSON)
	}
}

func main() {
	rootdir := "./"

	// os.Stderr.WriteString("Starting Front END")
	readEnv()

	// go consumeMemory(30000)

	fmt.Println("Running Demo app version: " + AppVersion)

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir(path.Join(rootdir, "/assets/")))))

	http.HandleFunc("/memleakbe", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Causing memleak on be")

		responseforping := Pong{"memleakbe", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	// Readiness probe function
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if checkBElive() {
			healthstate = "Healthy"
		} else {
			healthstate = "Dead"
		}
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Readiness-Header", healthstate)
		responseforping := Pong{"Readiness", healthstate}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)

	})
	// TODO: create a connection to the api and run a short test query
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Ping test")

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

	http.HandleFunc("/beping", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Backend Ping test")

		res := contactDemoAppBe()
		// js, err := json.Marshal(responseforping)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	http.HandleFunc("/scenario3", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Crash BE")

		res := CrashBE()
		// js, err := json.Marshal(responseforping)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// fmt.Println(js)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(path.Join(rootdir, "/assets/"))))
		fmt.Println("Present homepage")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		runHomePage(w)
		// w.Write(js)
	})

	http.HandleFunc("/memleak", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		consumeMemory(1000000)

		// log.Fatalln("Crashing")
		// os.Exit(3)

	})

	http.HandleFunc("/highmem", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		for i := 0; i < 10; i++ {
			// go consumeMemory(10000000)
			consumeMemory(100000)
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
	// http.HandleFunc("/", HttpWrapper(runHomePage))

	fmt.Println("Listening on port " + apiServerPort)
	http.ListenAndServe(":"+apiServerPort, nil)
}

func consumeMemory(n int) {
	var memleakslice []string
	memleakslice = nil
	// memleakslice = make(map[string]string, n)
	if len(memleakslice) <= n {
		for i := 1; i < n; i++ {
			memleakslice = append(memleakslice, RandStringRunes(1024))
		}
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	// fmt.Println(string(b))
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
