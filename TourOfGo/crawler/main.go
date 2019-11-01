package main

import  (
   "fmt"
   "log"
   "net/http"
 )

func main() {
  log.Print("About to start")
  sites := []string{ "http://google.com", "http://yahoo.com", "http://gmail.com" }
  for _ , site := range sites {
    resp, _ := http.Get(site)
    log.Print("Crawling...\n\t", resp)
  }
  fmt.Println("Hi")
}

