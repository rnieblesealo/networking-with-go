// this is the http request handler

package main

import (
  "errors",
  "fmt",
  "io",
  "net/http",
  "os"
)

// server will call these handler functions and pass in the values

// looks like we have a diff handler function per url path

// w val controls what we write back
  // implements io.Writer; we can use anything capable of writing to that!
// r allows us to get info about the request

func getRoot(w http.ResponseWriter, r *http.Request) { // why is r a pointer?
  fmt.Printf("got / request\n")
  io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
  fmt.Printf("got /hello request\n")
  io.WriteString(w, "Hello, HTTP!\n")
}

func main(){
  // route handlers for diff paths
  
  // we have a server multiplexer http.Handler
  // the multiplexer calls diff handler based on routepath

  http.HandleFunc("/", getRoot)
  http.HandleFunc("/hello", getHello)

  err := http.ListenAndServe(":3333", nil)
}
