package main

import (
	"container/list"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var hongKong *time.Location

var logs = list.New()

var writeLogs chan string
var lockLogs chan struct{}

func handleHome(res http.ResponseWriter, req *http.Request) {
	lockLogs <- struct{}{}
	for e := logs.Back(); e != nil; e = e.Prev() {
		fmt.Fprint(res, e.Value.(string))
	}
	<-lockLogs
}

func main() {
	if location, err := time.LoadLocation("Asia/Hong_Kong"); err != nil {
		hongKong = location
	} else {
		hongKong = time.FixedZone("Asia/Hong_Kong", 8*3600)
	}
	writeLogs = make(chan string)
	lockLogs = make(chan struct{})
	githubInit()
	go func() {
		for {
			select {
			case s := <-writeLogs:
				log.Print(s)
				logs.PushBack(fmt.Sprintf("%v: %v", time.Now(), s))
				if logs.Len() > 2000 {
					logs.Remove(logs.Front())
				}
			case <-lockLogs:
				lockLogs <- struct{}{}
			}
		}
	}()
	go githubUpdater()
	go crawl()
	http.HandleFunc("/", handleHome)
	ip := os.Getenv("OPENSHIFT_GO_IP")
	if ip == "" {
		ip = "localhost"
	}
	port := os.Getenv("OPENSHIFT_GO_PORT")
	if port == "" {
		port = "8080"
	}
	bind := fmt.Sprintf("%s:%s", ip, port)
	fmt.Printf("listening on %s...\n", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}
