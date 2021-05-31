package mschartgen

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func Serve(dir string, port int) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	addr := ":" + strconv.Itoa(port)
	log.Infof("serving %s on %s", dir, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
