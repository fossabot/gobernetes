package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", indexHandlerHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandlerHelloWorld(w http.ResponseWriter, r *http.Request) {

	out, err := exec.Command("hostname").Output()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Hello, I'm running on %s with an %s CPU. \n My hostname is %s",
		runtime.GOOS, runtime.GOARCH, string(out))
}
