package auth

import (
	"fmt"
	"time"
)

func UserAuth() {
	url := fmt.Sprintf(
		"https://login.microsoftonline.com/%s/oauth2/v2.0/authorize?client_id=%s&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%2Fmyapp%2F&response_mode=query&scope=Calendars.ReadWrite&state=%d",
		Info.TenantID,
		Info.ClientID,
		Info.State,
	)
	openURL(url)
	for Info.Code == "" {
		//Wait for valid code
		time.Sleep(time.Second * 1)
	}

	//TODO: let these do stuff
	//Start get token
	url = fmt.Sprintf(
		"",
	)
}
