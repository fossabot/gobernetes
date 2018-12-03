package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/write", write)
	http.HandleFunc("/cat", cat)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var temp_file = "/tmp/cat"

func hello(w http.ResponseWriter, r *http.Request) {

	out, err := exec.Command("hostname").Output()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Raca, I'm running on %s with an %s CPU. \n My hostname is %s",
		runtime.GOOS, runtime.GOARCH, string(out))
}

func write(w http.ResponseWriter, r *http.Request) {

	f, err := os.Create(temp_file)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	_, err = f.WriteString("I am black cat")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Black cat written")
}

func cat(w http.ResponseWriter, r *http.Request) {

	out, err := ioutil.ReadFile(temp_file)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Output from file %s: \n %s", temp_file, string(out))
}