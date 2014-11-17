package main
import (
	"net/http"
	"flag"
	"fmt"
)
 
var dir = flag.String("dir", ".", "the directory to serve")
var port = flag.Int("port", 9090, "the port to listen on")
func main() {
	flag.Parse()
	panic(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir(*dir))))
}
