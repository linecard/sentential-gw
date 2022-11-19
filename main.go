package main

import (
	"net/http"

	"github.com/wheegee/sentential-gw/internal/proxy"
	"github.com/wheegee/sentential-gw/internal/util"
)

func main() {
	util.Log.Info("starting gateway", "port", util.Port)
	http.HandleFunc("/", proxy.Handle)
	http.ListenAndServe(util.Port, nil)
}
