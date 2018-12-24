package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var filePath = "/tmp/cat"

func respWithHostname(w http.ResponseWriter, msg string) {

	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal("cannot execute hostname: ", err.Error())
	}

	fmt.Fprintf(w, "Hello, I'm running on %s with an %s CPU. \n My hostname is %s \n"+
		"----------------------------- \n \n %s",
		runtime.GOOS, runtime.GOARCH, string(out), msg)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/write", writeFile)
	http.HandleFunc("/remove", rmFile)
	http.HandleFunc("/cat", readFile)
	http.HandleFunc("/memleak", memleak)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {

	out, err := exec.Command("hostname").Output()
	if err != nil {
		log.Fatal("cannot execute hostname: ", err.Error())
	}

	fmt.Fprintf(w, "Hello, I'm running on %s with an %s CPU. \n My hostname is %s",
		runtime.GOOS, runtime.GOARCH, string(out))
}

func writeFile(w http.ResponseWriter, r *http.Request) {

	fp := r.URL.Query().Get("filePath")
	if fp != "" {
		filePath = fp
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	colour := ""
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(2) == 0 {
		colour = "Black"
	} else {
		colour = "White"
	}

	_, err = f.WriteString(fmt.Sprintf("I'm a %s cat!", colour))
	if err != nil {
		respWithHostname(w, err.Error())
		return
	}

	respWithHostname(w, fmt.Sprintf("%s cat written", colour))
}

func rmFile(w http.ResponseWriter, r *http.Request) {

	fp := r.URL.Query().Get("filePath")
	if fp != "" {
		filePath = fp
	}

	err := os.Remove(filePath)
	if err != nil {
		respWithHostname(w, err.Error())
		return
	}

	respWithHostname(w, fmt.Sprintf("%s file removed", filePath))
}

func readFile(w http.ResponseWriter, r *http.Request) {

	fp := r.URL.Query().Get("filePath")
	if fp != "" {
		filePath = fp
	}

	out, err := ioutil.ReadFile(filePath)
	if err != nil {
		respWithHostname(w, err.Error())
		return
	}

	respWithHostname(w, fmt.Sprintf("Output from file %s: \n %s", filePath, string(out)))
}

// memleak example route hit: http://localhost:8080/memleak?megabytes=123&interval=1000
func memleak(w http.ResponseWriter, r *http.Request) {
	var howMany, interval int64 = 100, 350
	mbs := r.URL.Query().Get("megabytes")
	if mbs != "" {
		var err error
		howMany, err = strconv.ParseInt(mbs, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	ms := r.URL.Query().Get("interval")
	if mbs != "" {
		var err error
		interval, err = strconv.ParseInt(ms, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	file, err := os.Open("/dev/urandom")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		buff := new(bytes.Buffer)

		for {
			if _, err := io.CopyN(buff, file, howMany*1e6); err != nil {
				log.Printf("copyN failed: %v\n", err)
				log.Println("memleak stopping...")
				return
			}
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}()

	fmt.Fprintf(w, "memleak running: %dMB / %dms\n", howMany, interval)
}
