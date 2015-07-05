package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"
)

const (
	indexPath = "/"
	wsPath    = "/ws"
	port      = 8080
)

var (
	exitStatus int
	usrCnt     int

	// user count mutex
	ucm = new(sync.Mutex)
)

type data struct {
	Cmd  string `json:"cmd"`
	Time int    `json:"time"`
}

func init() {
	flag.Parse()
}

func main() {
	run()
	os.Exit(exitStatus)
}

func run() {
	glog.Info("run()")

	http.Handle(indexPath, http.FileServer(http.Dir("client")))

	http.HandleFunc(wsPath,
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(handler)}
			s.ServeHTTP(w, req)
		})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		glog.Error(err)
		exitStatus = 1
	}
}

func handler(conn *websocket.Conn) {
	userID := getUserID()
	glog.Info("start: ", userID)

	for {
		var d data
		glog.Info("wait receive")
		rErr := websocket.JSON.Receive(conn, &d)
		if rErr != nil {
			glog.Error("receive error: ", rErr)
			break
		}

		glog.Info("receive data:", d)

		// send data
		d.Cmd = "pong"
		glog.Info("send data:", d)

		sErr := websocket.JSON.Send(conn, d)
		if sErr != nil {
			glog.Error("send error: ", sErr)
			break
		}
	}

	glog.Info("end: ", userID)
}

func getUserID() int {
	ucm.Lock()
	defer ucm.Unlock()

	usrCnt++
	return usrCnt
}
