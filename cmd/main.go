package main

import (
	"fmt"
	"net/http"

	"github.com/Shoubhit/secure-api/web/router"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Starting Secure API service...")
	fmt.Println(`
	██ ██ ██ █╬█ ███ ██ ╬╬ ███ ███ █
	█▄ █▄ █╬ █╬█ █▄╬ █▄ ╬╬ █▄█ █▄█ █
	▄█ █▄ ██ ███ █╬█ █▄ ╬╬ █╬█ █╬╬ █`)

	r := router.SetupRouter()

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Error(err)
	}

}
