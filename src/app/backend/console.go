package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"net/http"
)

var (
	port = pflag.Int("port", 8080, "The port that the server listens to")
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	glog.Info("Starting HTTP server on port ", *port)
	defer glog.Flush()

	http.Handle("/", http.FileServer(http.Dir("./")))
	glog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
