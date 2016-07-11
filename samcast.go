package main 

import (
  "net/http"
  "os/exec"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  stream_url := r.FormValue("streamurl")
  cmd := exec.Command("livestreamer", stream_url)
  cmd.Start()
}

func main() {
  http.HandleFunc("/", defaultHandler)
  http.ListenAndServe(":8100", nil)
}
