package main

import (
	"fmt"
	"os"

	gserviceaccount "github.com/knq/jwt/gserviceaccount"
	"golang.org/x/net/context"
	oauth2Google "golang.org/x/oauth2/google"
	sheets "google.golang.org/api/sheets/v4"
)

func main() {}

func Write(data string, srv *sheets.Service) error {
	return nil
}

func Read(cell string) string {

	json_creds := []byte(os.Getenv("CLIENT_CREDENTIALS"))
	fmt.Printf(os.Getenv("CLIENT_CREDENTIALS"))
	config, err := oauth2Google.ConfigFromJSON(json_creds)
	if err != nil {
		panic(err)
	}

	service_acc, err := gserviceaccount.FromJSON(json_creds)
	if err != nil {
		panic(err)
	}

	var scopes string
	token_source, err := service_acc.TokenSource(context.Background(), scopes)
	if err != nil {
		panic(err)
	}

	tok, err := token_source.Token()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := config.Client(ctx, tok)

	srv, err := sheets.New(client)
	if err != nil {
		panic(err)
	}

	spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	var readRange string
	if cell == "A1" {
		readRange = "Test Data!A1"
	} else {
		readRange = "Test Data!B2"
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		panic(err)
	}
	return resp.MajorDimension
	return "hello"
}
