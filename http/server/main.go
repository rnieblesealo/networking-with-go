// this is the http request handler

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

// server will call these handler functions and pass in the values

// looks like we have a diff handler function per url path

// w val controls what we write back
// implements io.Writer; we can use anything capable of writing to that!
// r allows us to get info about the request

func getRoot(w http.ResponseWriter, r *http.Request) { // why is r a pointer?
  // get request context
	ctx := r.Context()

  // get value keyed by keyServerAddr (which just shows the address of server)
	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))

	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))

	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	// route handlers for diff paths

	// we have a default server multiplexer http.Handler
	// the multiplexer calls diff handler based on routepath

	// this is how we do it w default:
	/*
		http.HandleFunc("/", getRoot)
		http.HandleFunc("/hello", getHello)
	*/

	// WARN: but default/glob use of things like http multiplexer can lead to bugs
	// it is best to set up our own
	// http.ServeMux does this

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	// default http server:
	/*
		err := http.ListenAndServe(portStr, mux) // listenserve is blocking; program wont run past it
	*/

	// custom http servers with context:

	ctx, cancelCtx := context.WithCancel(context.Background())

	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(ls net.Listener) context.Context { // basecontext allows us to modify server context values that handler functions receive
			ctx = context.WithValue(ctx, keyServerAddr, ls.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(ls net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, ls.Addr().String())
			return ctx
		},
	}

	// run server 1 and 2 in goroutines
	// note anonym func use!

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server 1 closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}

		cancelCtx() // undo the context

	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server 2 closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}

		cancelCtx()
	}()

	// read from ctx done channel before returning from main
	// ensures program will run until either server's goroutine ends and cancelxtx is called
	// if ctx ends, program ends

	// TODO: what does <- do?

	<-ctx.Done()
}
