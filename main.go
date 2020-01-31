package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"strings"
)

const (
	APP = "slack-file-deleter"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetPrefix(APP + ": ")

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	api := slack.New(os.Getenv("TOKEN"))
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		log.Println("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// DMチャンネルでかつ fiele へのreplyだったとき。
			if ev.SubType == "message_replied" && strings.HasPrefix(ev.Channel, "D") && ev.SubMessage.Files != nil {
				if err := api.DeleteFile(ev.SubMessage.Files[0].ID); err != nil {
					fmt.Errorf("file delete error %v", err)
				}
			}

		case *slack.InvalidAuthEvent:
			return fmt.Errorf("Auth Error %v\n", msg)
		default:
			fmt.Printf("%v\n", ev)
		}
	}
	return nil
}
