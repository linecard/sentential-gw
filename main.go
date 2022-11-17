package main

import (
	"net/http"
	"sentential-gw/internal/proxy"
	"sentential-gw/internal/util"
)

func main() {
	util.Log.Info("starting gateway", "port", util.Port)
	http.HandleFunc("/", proxy.Handle)
	http.ListenAndServe(util.Port, nil)
}
