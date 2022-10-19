package auth

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var Info AuthStruct

const fallbackAddr = "localhost:46987"
const randIntN = 1145141919810

func StartAuthServer() {
	//Get a random state
	rand.Seed(time.Now().UnixNano())
	Info.State = rand.Intn(randIntN)

	http.HandleFunc("/getCode", codeHandler)
	err := http.ListenAndServe(fallbackAddr, nil)
	if err != nil {
		log.Panicln(err)
	}
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	code := query["code"][0]
	Info.Code = code
}
