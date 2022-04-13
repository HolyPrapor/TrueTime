package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func fail(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func trueTimeHandler(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		fail(w, r, err)
		return
	}
	cmd := exec.Command("sh", "truetime.sh")
	cmd.Dir = cwd
	output, err := cmd.Output()
	if err != nil {
		fail(w, r, err)
		return
	}
	http.RedirectHandler(string(output), http.StatusPermanentRedirect).ServeHTTP(w, r)
}

func RunServerWithRouting(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", trueTimeHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	portFlag := flag.Int("port", 8080, "Specifies port to serve on")
	flag.Parse()
	RunServerWithRouting(*portFlag)
}
