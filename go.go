package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ExampleChangeMe -
type ExampleChangeMe struct {
}

//
//
// HTTP Functions
func (x *ExampleChangeMe) serveHTML(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("hello"))
}

func (x *ExampleChangeMe) serveJSON(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	obj := ExampleChangeMe{}

	jBytes, err := json.Marshal(&obj)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jBytes)
	return
}

func (x *ExampleChangeMe) readJSON(res http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if req.Method != http.MethodPost && contentType != "application/json" {
		errStr := fmt.Sprintf("Request must be POST and have contentType of application/json")
		res.Write([]byte(errStr))
		return
	}

	obj := ExampleChangeMe{}
	if err := json.NewDecoder(req.Body).Decode(&obj); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		errStr := fmt.Sprintf("Unable to decode JSON: %s", err.Error())
		res.Write([]byte(errStr))
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (x *ExampleChangeMe) doAuth(res http.ResponseWriter, req *http.Request) bool {
	if x.Config.APIKey != "" {
		// check for API-Key header
		key := req.Header.Get("API-Key")
		if key != x.Config.APIKey {
			return false
		}
	}
	return true
}

// Run runs
func (x *ExampleChangeMe) Run() (err error) {

	if x.Logf == nil {
		x.Logf = x.doLog
	}

	x.mux = http.NewServeMux()
	x.mux.HandleFunc("/", x.foo)
	x.mux.HandleFunc("/", x.foo)

	x.server = &http.Server{
		Addr:    x.Config.BindStr,
		Handler: x,
	}

	if x.Config.CrtFile != "" && x.Config.KeyFile != "" {
		// probably need to add an SSL config
		x.Logf("Starting SSL Server on %s", x.server.Addr)
		err = x.server.ListenAndServeTLS(x.Config.CrtFile, x.Config.KeyFile)
	} else {
		x.Logf("Starting Server on %s", x.server.Addr)
		err = x.server.ListenAndServe()
	}
	return err
}

// ServeHTTP
func (x *ExampleChangeMe) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// do auth
	if x.doAuth(res, req) {
		x.Logf("authfail %s to %s\n", req.RemoteAddr, req.RequestURI)
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("No Soup For You!"))
		return
	}

	x.mux.ServeHTTP(res, req)
	return
}

//
//
// Env
func getEnv() {
	_ = os.Getenv("ExampleChangeMe")
}

func main() {
	fmt.Println("nothing here")
}

//
//
// maps

// GetMapKeys gets a map[string] keys
// https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
// pre-alloc is about 2x faster than append
func GetMapKeys(m *map[string]interface{}) []string {
	k := make([]string, len(*m))
	var i int
	for key := range *m {
		k[i] = key
		i++
	}
	return k
}

// JSON
func jsonExample() {
	obj := ExampleChangeMe{}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(&ev); err != nil {
		panic(err)
	}
	http.NewRequest(http.MethodPost, "/", b)
}
