package main

import(
	"github.com/rakshazi/applybyapi/api"
	"flag"
	"log"
)

var vacancyId int

func init() {
	flag.IntVar(&vacancyId, "id", 0, "Posting (vacancy) ID")
}

func main() {
	flag.Parse()
	if(vacancyId == 0) {
		log.Fatal("You must provide posting (vacancy) id, try to add -h flag to get more info")
	}
	token, err := api.GetToken(vacancyId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Token:", token)
}
