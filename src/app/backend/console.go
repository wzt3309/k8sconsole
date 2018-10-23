package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/lithammer/dedent"
	"github.com/spf13/pflag"
	"net/http"
)

var (
	port = pflag.Int("port", 8080, "The port that the server listens to")
	apiserverHost = pflag.String("apiserver-host", "", dedent.Dedent(`
		The address of Kubernetes apiserver to connect to in the form of protocol://address:port,
			┌──────────────────────────────────────────────────────────┐
			| e.g.                                                     |
      | http://10.0.1.2:8081                                     |
      └──────────────────────────────────────────────────────────┘
		If not specified, the assumption is that the binary is run in a Kubernetes cluster and 
		local discovery is attempted.
`))
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	glog.Info("Starting HTTP server on port ", *port)
	defer glog.Flush()

	glog.Info("Connecting to kubernetes cluster ", *apiserverHost)

	http.Handle("/", http.FileServer(http.Dir("./")))
	glog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
