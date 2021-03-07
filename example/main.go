package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/lil5/gophprouter"
)

type basicApiResponse struct {
	Ok     bool `json:"ok"`
	Status int  `json:"status"`
}

func main() {
	path := flag.String("path", "/", "Path in url")
	method := flag.String("method", http.MethodGet, "Method used")
	body := flag.String("body", "", "Body")
	flag.Parse()

	r := gophprouter.NewRouter()

	r.HTTPR.GlobalOPTIONS = http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		os.Exit(0)
	})

	// Add here your routes
	r.GET("/", httprouter.Handle(func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
		d, err := json.Marshal(BasicApiResponse{
			Ok:     true,
			Status: 200,
		})
		if err != nil {
			os.Exit(1)
		}
		rw.Write(d)
	}))

	ok := r.Serve(*method, *path, *body)

	if ok == false {
		os.Exit(1)
	}

	os.Exit(0)
}
