package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var port = flag.Int("port", 8080, "The port for the server to listen")
var addr = flag.String("address", "0.0.0.0", "Listening address")

func handleView(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile(r.URL.Path[1:] + ".md")

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html> <html> <title>Python</title> <meta charset=\"utf-8\"> <xmp theme=\"cerulean\" heading_number=\"i.i.a\" toc=\"true\" style=\"display:none;\">\n")

	w.Write(content)

	fmt.Fprintf(w, "</xmp> <script src=\"http://cdn.ztx.io/strapdown/strapdown.min.js\"></script> </html>")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handleView)

	host := fmt.Sprintf("%s:%d", *addr, *port)
	log.Printf("listening on %s", host)
	l, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	s := &http.Server{}
	s.Serve(l)
}
