package auth

import (
	"encoding/json"
	"fmt"
	"github.com/TurboHsu/ClassTimetableSync/config"
	"io"
	"log"
	"net/http"
	urllib "net/url"
	"strconv"
	"time"
)

func UserAuth() {
	//Wait for random state generation
	for Info.State == 0 {
	}

	//Get code
	url := `https://login.microsoftonline.com/consumers/oauth2/v2.0/authorize?client_id=` +
		config.Config.UserInfo.ClientID +
		`&response_type=code&redirect_uri=http%3A%2F%2F` +
		fallbackAddr +
		`%2FgetCode&response_mode=query&scope=Calendars.ReadWrite&state=` +
		strconv.Itoa(Info.State)
	openURL(url)
	for Info.Code == "" {
		//Wait for valid code
		time.Sleep(time.Second * 1)
	}

	//Start get token
	url = "https://login.microsoftonline.com/consumers/oauth2/v2.0/token"
	val := urllib.Values{
		"client_id":    {config.Config.UserInfo.ClientID},
		"scope":        {"Calendars.ReadWrite"},
		"code":         {Info.Code},
		"redirect_uri": {`http://` + fallbackAddr + `/getCode`},
		"grant_type":   {"authorization_code"},
	}
	resp, err := http.PostForm(url, val)
	if err != nil {
		fmt.Println("[E] ", err.Error())
	}
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println("[E] Http return is not ok. " + string(respBody))
		return
	}
	_ = json.Unmarshal(respBody, &Info.Token)
	log.Println("[I] Token gain succeed.")
}
