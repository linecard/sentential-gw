package main

import (
	"net/http"
	"os"

	"github.com/wheegee/sentential-gw/internal/proxy"
	"github.com/wheegee/sentential-gw/internal/util"
)

var version = "dev" // Set by goreleaser.

func main() {
	util.Log.Info("starting gateway", "version", version, "port", util.Port)
	http.HandleFunc("/", proxy.Handle)
	if err := http.ListenAndServe(util.Port, nil); err != nil {
		util.Log.Error("error starting exporter", "error", err)
		os.Exit(1)
	}
}
