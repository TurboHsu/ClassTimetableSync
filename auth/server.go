package auth

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var Info InfoStruct

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
	state := query["state"][0]
	if state == strconv.Itoa(Info.State) {
		Info.Code = code
	} else {
		log.Println("[I] Received auth code but state unmatched.")
	}
	_, _ = fmt.Fprintln(w, "Auth completed. You can close this tab now.")
}
