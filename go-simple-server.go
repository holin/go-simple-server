package main

import (
  "log"
  "net/http"
  "os"
)


func nocacheHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, max-age=-1")
    w.Header().Add("Pragma", "-1")
    w.Header().Add("Last-Modified", "-1")
    h.ServeHTTP(w, r)
  })
} 

func main() {

  dir, _ := os.Getwd()
  log.Println("Work Directory:", dir)
  fs := http.FileServer(http.Dir(dir))
  http.Handle("/", nocacheHandler(fs))

  log.Println("Listening at: 0.0.0.0:3000")
  http.ListenAndServe(":3000", nil)


}