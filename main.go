package main

import (
    "flag"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var upgrader = websocket.Upgrader{}

func main() {
    log.Println("Star Application")


    startServer()
}

func startServer() {
    log.Println("Connect server root")

    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)

    http.HandleFunc("/ws", handleConnections)``

    log.Println("http server started on :8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    log.Println("handleConnections")

    // Upgrade initial GET request to a websocket

    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }
    // Make sure we close the connection when the function returns
    defer ws.Close()
}
