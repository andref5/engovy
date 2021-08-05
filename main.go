package main

import (
	"fmt"
	"log"
	"net/http"

	envoy "github.com/andref5/engovy/pkg/control-plane"
)

const (
	HttpPort     = "8001"
	EnvoyXdsPort = 18000
)

func main() {
	controlPlane := envoy.ControlPlane{}
	go controlPlane.Start(EnvoyXdsPort)

	http.HandleFunc("/changeRoute", func(w http.ResponseWriter, r *http.Request) {
		path := r.FormValue("path")
		if len(path) <= 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err := controlPlane.ChangeRoutePath(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "OK")
	})

	log.Println("HTTP server startup port " + HttpPort)
	http.ListenAndServe(":"+HttpPort, nil)
}
