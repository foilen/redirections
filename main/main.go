package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var mapping map[string]*Configuration

func main() {

	// Get the configuration
	mapping = make(map[string]*Configuration)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		name := pair[0]
		value := pair[1]
		if strings.HasPrefix(name, "REDIR_") {
			config, err := getConfiguration(value)
			if err == nil {
				mapping[config.Host] = config
			} else {
				fmt.Println("WARN", "Could not parse json for", name)
			}
		}
	}

	for hostname, config := range mapping {
		fmt.Println("Redirection for host", hostname, "=>", config)
	}

	// Choose the port
	port := "80"
	envPort := os.Getenv("HTTP_PORT")
	if envPort != "" {
		port = envPort
	}
	fmt.Println("Listening on port", port)

	// Start the server
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func redirect(response http.ResponseWriter, request *http.Request) {
	host := strings.SplitN(request.Host, ":", 2)[0]
	redirection := mapping[host]
	if redirection == nil {
		http.Error(response, "Unknown host for redirection", 403)
	} else {
		var code int
		var location string
		if redirection.Permanent {
			code = 301
		} else {
			code = 302
		}
		location = redirection.Redirection
		if redirection.AppendQuery {
			location += request.URL.Path
		}
		http.Redirect(response, request, location, code)
	}
}
