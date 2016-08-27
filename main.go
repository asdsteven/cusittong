package main

import (
	"log"
)

func main() {
	if err := crawl(); err != nil {
		log.Print(err)
		return
	}
	/*http.HandleFunc("/", handleHome)
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
	log.Fatal(http.ListenAndServe(bind, nil))*/
}
