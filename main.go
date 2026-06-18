package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    color := os.Getenv("PAGE_COLOR")
    if color == "" {
        color = "#1E90FF"
    }

    version := os.Getenv("APP_VERSION")
    if version == "" {
        version = "v1"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintf(w, `<!doctype html>
<html>
<head>
  <title>OpenShift DevOps Demo</title>
</head>
<body style="background:%s;color:white;font-family:Arial;text-align:center;padding-top:80px">
  <h1>OpenShift DevOps Demo</h1>
  <h2>Go Web Server</h2>
  <p>Versione: %s</p>
  <p>Colore corrente: %s</p>
</body>
</html>`, color, version, color)
    })

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte("ok"))
    })

    log.Println("Server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
