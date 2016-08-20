package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handleHome(w http.ResponseWriter, req *http.Request) {
	if err := crawl(); err != nil {
		log.Print(err)
		return
	}
	s := ""
	for _, ct := range db.careerTerms {
		s += fmt.Sprintf("%v %v %v\n", db.careers[ct.career].slug, strings.TrimSpace(db.terms[ct.term].en), len(ct.subjects))
	}
	if _, err := w.Write([]byte(s)); err != nil {
		log.Print(err)
	}
}

func handleBrowse(w http.ResponseWriter, req *http.Request) {
	html, err := browse()
	if err != nil {
		log.Print(err)
		return
	}
	if _, err := w.Write(html); err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/browse", handleBrowse)
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
