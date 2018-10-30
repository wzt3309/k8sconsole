package main

import (
	"flag"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/golang/glog"
	"github.com/lithammer/dedent"
	"github.com/spf13/pflag"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/auth"
	"github.com/wzt3309/k8sconsole/src/app/backend/crypto"
	"github.com/wzt3309/k8sconsole/src/app/backend/datastore"
	"github.com/wzt3309/k8sconsole/src/app/backend/file"
	"github.com/wzt3309/k8sconsole/src/app/backend/handler"
	"github.com/wzt3309/k8sconsole/src/app/backend/jwt"
	"github.com/wzt3309/k8sconsole/src/app/backend/user"
	"github.com/wzt3309/k8sconsole/src/app/backend/validator"
	"net/http"
)

var (
	port          = pflag.Int("port", 8080, "The port that the server listens to")
	apiserverHost = pflag.String("apiserver-host", "", dedent.Dedent(`
		The address of Kubernetes apiserver to connect to in the form of protocol://address:port,
			┌──────────────────────────────────────────────────────────┐
			| e.g.                                                     |
      | handler://10.0.1.2:8081                                     |
      └──────────────────────────────────────────────────────────┘
		If not specified, the assumption is that the binary is run in a Kubernetes cluster and 
		local discovery is attempted.
`))
	dataStorePath = pflag.String("data", "./.tmp/db", "The path to store data")
	noAuth        = pflag.Bool("noAuth", false, "Don't use auth in the backend")
)

func initValidator() {
	govalidator.TagMap["username"] = validator.IsUsername
	govalidator.CustomTypeTagMap.Set("role", govalidator.CustomTypeValidator(validator.IsRole))
}

func initFileService() api.FileService {
	return file.NewService()
}

func initStore(dataStorePath string, fileService api.FileService) api.DataStore {
	store, err := datastore.NewBoltDBStore(dataStorePath, fileService)
	if err != nil {
		glog.Fatal(err)
	}

	err = store.Open()
	if err != nil {
		glog.Fatal(err)
	}

	err = store.Init()
	if err != nil {
		glog.Fatal(err)
	}

	return store
}

func initJWTService(authenticationEnabled bool) api.JWTService {
	if authenticationEnabled {
		jwtService, err := jwt.NewJWTService()
		if err != nil {
			glog.Fatal(err)
		}
		return jwtService
	}
	return nil
}

func initCryptoService() api.CryptoService {
	return crypto.NewService()
}

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	glog.Info("Starting HTTP server on port ", *port)
	defer glog.Flush()

	glog.Info("Connecting to kubernetes cluster ", *apiserverHost)

	// init struct validator
	initValidator()

	fileService := initFileService()

	store := initStore(*dataStorePath, fileService)
	defer store.Close()

	jwtService := initJWTService(!*noAuth)

	cryptoService := initCryptoService()

	fAuthManager := auth.NewFrontendAuthManager(cryptoService, jwtService, store.GetUserService(), *noAuth)
	userManager := user.NewUserManager(store.GetUserService(), cryptoService)

	apiHandler, err := handler.CreateHTTPAPIHandler(fAuthManager, userManager)
	if err != nil {
		glog.Fatal(err)
	}

	http.Handle("/api/", apiHandler)
	http.Handle("/", http.FileServer(http.Dir("./")))

	glog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
