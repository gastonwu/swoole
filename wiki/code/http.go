package main

import (
    "fmt"
    "log"
    "net/http"
    "runtime"
)

func main() {
    // 限制为CPU的数量减一
    runtime.GOMAXPROCS( runtime.NumCPU() - 1 )

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "It works!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
