package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/slides/v1"

	"kpistability/lib/googletokenhelper"
)

func main() {
	//  log.SetFlags(log.LstdFlags | log.Lshortfile) // log error line number from caller in checkErr(err)

	// call WS
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")
	checkErr(err)
	responseData, err := io.ReadAll(response.Body)
	checkErr(err)
	fmt.Println(reflect.TypeOf(responseData))
	// fmt.Println(responseData)

	// work with sheets
	ctx := context.Background()
	b, err := os.ReadFile(".auth/credentials.json")
	checkErr(err)

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/presentationconss.readonly")
	checkErr(err)
	client := googletokenhelper.GetClient(config)

	srv, err := slides.NewService(ctx, option.WithHTTPClient(client))
	checkErr(err)

	// Prints the number of slides and elements in a sample presentation:
	// https://docs.google.com/presentation/d/1EAYk18WDjIG-zp_0vLm3CsfQh_i8eXc67Jo2O9C6Vuc/edit
	presentationId := "1EAYk18WDjIG-zp_0vLm3CsfQh_i8eXc67Jo2O9C6Vuc"
	presentation, err := srv.Presentations.Get(presentationId).Do()
	checkErr(err)

	fmt.Printf("The presentation contains %d slides:\n", len(presentation.Slides))
	for i, slide := range presentation.Slides {
		fmt.Printf("- Slide #%d contains %d elements.\n", (i + 1),
			len(slide.PageElements))
	}
}

func checkErr(err error) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)
		log.Fatalf("[error] %v (%s in %s:%d)", err, runtime.FuncForPC(pc).Name(), filename, line)
	}
}
