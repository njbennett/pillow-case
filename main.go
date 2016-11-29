package main

import (
	"os"

	"golang.org/x/net/context"

	oauth2ClientCreds "golang.org/x/oauth2/clientcredentials"
	sheets "google.golang.org/api/sheets/v4"
)

func main() {}

func Write(data string, srv *sheets.Service) error {
	return nil
}

func Read(cell string) string {
	config := oauth2ClientCreds.Config{
		ClientID:     os.Getenv("ClientID"),
		ClientSecret: os.Getenv("ClientSecret"),
		TokenURL:     os.Getenv("TokenURL"),
	}

	ctx := context.Background()
	client := config.Client(ctx)

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
