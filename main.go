package main

import (
	"github.com/TurboHsu/ClassTimetableSync/auth"
	"github.com/TurboHsu/ClassTimetableSync/calendar"
	"github.com/TurboHsu/ClassTimetableSync/config"
)

func main() {
	//Reads config
	config.ReadConfig()

	go auth.StartAuthServer()
	auth.UserAuth()

	//Wait for auth
	for auth.Info.Token.AccessToken == "" {
	}

	//Test post event
	env := calendar.EventStruct{
		Subject: "This is completely a test",
		Body: struct {
			ContentType string `json:"contentType"`
			Content     string `json:"content"`
		}{
			ContentType: "HTML",
			Content:     "How about this test",
		},
		Start: struct {
			DateTime string `json:"dateTime"`
			TimeZone string `json:"timeZone"`
		}{
			DateTime: "2022-10-22T14:48:00",
			TimeZone: config.Config.Init.Timezone,
		},
		End: struct {
			DateTime string `json:"dateTime"`
			TimeZone string `json:"timeZone"`
		}{
			DateTime: "2022-10-22T15:48:00",
			TimeZone: config.Config.Init.Timezone,
		},
		Location: struct {
			DisplayName string `json:"displayName"`
		}{
			DisplayName: "Nice Place",
		},
		AllowNewTimeProposals: true,
	}
	calendar.PostEvent(env)
}
