package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/rakshazi/applybyapi/api"
	"github.com/rakshazi/applybyapi/tui"
)

var vacancyId int

func init() {
	flag.IntVar(&vacancyId, "id", 0, "Posting (vacancy) ID")
}

func main() {
	flag.Parse()
	if vacancyId == 0 {
		log.Fatal("You must provide posting (vacancy) id, try to add -h flag to get more info")
	}
	tui.WelcomeMessage()
	token, err := api.GetToken(vacancyId)
	if err != nil {
		log.Fatal(err)
	}

	tui.TokenNotification(token)
	data := tui.RunSurvey()
	data.Token = token
	data.Posting = strconv.Itoa(vacancyId)

	id, err := api.Apply(data)
	if err != nil {
		log.Fatal(err)
	}
	tui.Done(id)
}
