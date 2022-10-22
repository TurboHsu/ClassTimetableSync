package calendar

import (
	"encoding/json"
	"fmt"
	"github.com/TurboHsu/ClassTimetableSync/auth"
	"github.com/TurboHsu/ClassTimetableSync/config"
	"io"
	"log"
	"net/http"
	"strings"
)

// PostEvent to calendar
func PostEvent(data EventStruct) {
	//Form the request body
	url := "https://graph.microsoft.com/v1.0/me/calendar/events"
	b, _ := json.Marshal(data)
	//Send the request
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(b)))
	req.Header.Set("Authorization", auth.Info.Token.TokenType+" "+auth.Info.Token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "outlook.timezone=\""+config.Config.Init.Timezone+"\"")
	//Do request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("[E] Error when posting event: ", err)
		return
	}
	defer resp.Body.Close()

	//Error handling
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		log.Println("[E] Error when posting event: ", resp.Status, string(body))
		return
	}
	fmt.Println("[I] Event posted.")
}
