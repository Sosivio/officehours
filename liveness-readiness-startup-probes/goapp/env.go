package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

const (
	EnvAPIServerPort = "API_PORT"   // :7777
	EnvIsEnabled     = "IS_ENABLED" // kill switch! MUST HAVE S7M_IS_ENABLED=true
	EnvIsVerbose     = "IS_VERBOSE"
	EnvDemoAppBeURL  = "DEMOAPP_BE_URL"
)

var (
	isVerbose     bool
	hostname      string
	appname       string
	apiServerPort string
	demoappBeURL  string

	startTime time.Time
)

func readEnv() {

	// isEnabled, ok := strconv.ParseBool(getEnv(EnvIsEnabled, "false"))
	// if !isEnabled || ok != nil {
	// 	zero := 0
	// 	fmt.Println(1 / zero) // panic with division by 0!
	// }

	appnameParts := strings.Split(os.Args[0], "/")
	appname = appnameParts[len(appnameParts)-1]

	var err error

	user, err := user.Current()
	if err == nil {
		fmt.Printf("Running as UID: [%s]\tGID: [%s]\tUsername: [%s]", user.Uid, user.Gid, user.Username)
	} else {
		fmt.Println("user.Current() threw error", err)
	}

	apiServerPort = getEnv(EnvAPIServerPort, "")

	demoappBeURL = getEnv(EnvDemoAppBeURL, "")

	// kubeConfigPath = getEnv(EnvKubeConfigPath, "")
	// if kubeConfigPath == "" {

	// 	fmt.Println(EnvKubeConfigPath, "is not set.  Using default.")

	// 	kubeConfigPath = "/ocpcluster-kube/ocpcluster-config-3.11"

	// }

	// fmt.Println("Using KubeConfig " + kubeConfigPath)

	isVerbose, err = strconv.ParseBool(getEnv(EnvIsVerbose, "true"))
	if err != nil {
		isVerbose = true
	}

	readEnvMutable()

	startTime = time.Now()
}

func readEnvMutable() {

	now := time.Now()
	fmt.Println("Current time: ", now)

}

func getEnv(key string, def string) string {

	result := os.Getenv(key)
	if result == "" {
		return def
	}

	return RemoveTrailingNL(result)
}

func RemoveTrailingNL(input string) string {

	len := len(input)
	if len > 0 && input[len-1] == '\n' {
		return input[:len-1]
	}

	return input
}
