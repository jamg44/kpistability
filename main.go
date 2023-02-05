package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/slides/v1"

	"kpistability/lib/googletokenhelper"
	"kpistability/lib/utils"
)

func main() {
	defer fmt.Println("End.")
	//  log.SetFlags(log.LstdFlags | log.Lshortfile) // log error line number from caller in checkErr(err)

	// call WS
	// url := "http://pokeapi.co/api/v2/pokedex/kanto"
	url := "http://api.open-notify.org/astros.json"
	response, err := http.Get(url)
	utils.CheckErr(err)
	body, err := io.ReadAll(response.Body)
	utils.CheckErr(err)
	// fmt.Println(reflect.TypeOf(body))

	bodyStr, err := utils.PrettyPrintJSONResponse(body)
	utils.CheckErr(err)
	log.Println(bodyStr)

	// work with sheets
	// slideElementsCount()
}

func slideElementsCount() {
	ctx := context.Background()
	b, err := os.ReadFile(".auth/credentials.json")
	utils.CheckErr(err)

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/presentationconss.readonly")
	utils.CheckErr(err)
	client := googletokenhelper.GetClient(config)

	srv, err := slides.NewService(ctx, option.WithHTTPClient(client))
	utils.CheckErr(err)

	// Prints the number of slides and elements in a sample presentation:
	// https://docs.google.com/presentation/d/1EAYk18WDjIG-zp_0vLm3CsfQh_i8eXc67Jo2O9C6Vuc/edit
	presentationId := "1EAYk18WDjIG-zp_0vLm3CsfQh_i8eXc67Jo2O9C6Vuc"
	presentation, err := srv.Presentations.Get(presentationId).Do()
	utils.CheckErr(err)

	fmt.Printf("The presentation contains %d slides:\n", len(presentation.Slides))
	for i, slide := range presentation.Slides {
		fmt.Printf("- Slide #%d contains %d elements.\n", (i + 1),
			len(slide.PageElements))
	}
}
