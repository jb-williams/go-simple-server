package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

var fP = fmt.Printf
var iWs = io.WriteString
var fPf = fmt.Fprintf

func getRoot(resp http.ResponseWriter, requ *http.Request) {
	ctx := requ.Context()

	fP("%s: got / request\n", ctx.Value(keyServerAddr))
	iWs(resp, "This is my website!\n")
}

func getHello(resp http.ResponseWriter, requ *http.Request) {
	ctx := requ.Context()

	fP("%s: got /hello request\n", ctx.Value(keyServerAddr))
	name := requ.FormValue("name")
	//name := requ.PostFormValue("name")
	if name == "" {
		resp.Header().Set("x-missing-field", "name")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	iWs(resp, fmt.Sprintf("Hello, %s!\n", name))
}

func formHandler(resp http.ResponseWriter, requ *http.Request) {
	if err := requ.ParseForm(); err != nil {
		fPf(resp, "formHandler: ParseForm(): failed: %w\n", err)
	}

	fP("%s: got /form request\n", ctx.Value(keyServerAddr))
	name := requ.FormValue("name")
	address := requ.FormValue("address")
	if name == "" || address == "" {
		resp.Header().Set("x-missing-field", "name || address")
		resp.WriteHeader(http.StatusBadRequest)
		return
	} else {
		iWs(resp, "POST request was Successful!\n")
		fPf(resp, "Name = %s\n", name)
		fPf(resp, "Address = %s\n", address)
	}
}

func main() {
	mux := http.NewServeMux()
	//fileServer := http.FileServer(http.Dir("./static"))
	//mux.Handle("/", fileServer) // this one worked
	// mux.HandleFunc("/", fileServer)
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/form", formHandler)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":9999",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	fP("Starting Server at localhost:9999\n")
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
