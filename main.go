package main

import (
	"io/ioutil"
	"fmt"
	"net/http"

	gserviceaccount "github.com/knq/jwt/gserviceaccount"
	"golang.org/x/net/context"
	sheets "google.golang.org/api/sheets/v4"
	oauth2 "golang.org/x/oauth2"
)

func main() {
	http.HandleFunc("/", LookupHandler)
	http.ListenAndServe(":8080", nil)
}

func FetchCredentials() ([]byte, error) {
	dat, err := ioutil.ReadFile("./credentials")
	return dat, err
}

func Write(data string, srv *sheets.Service) error {
	return nil
}

func Read(cell string) string {
	json_creds, err := FetchCredentials()
	if err != nil {
		panic(err)
	}

	service_acc, err := gserviceaccount.FromJSON(json_creds)
	if err != nil {
		panic(err)
	}

	var scopes string
	scopes = "https://www.googleapis.com/auth/spreadsheets.readonly"
	token_source, err := service_acc.TokenSource(context.Background(), scopes)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := oauth2.NewClient(ctx, token_source)

	srv, err := sheets.New(client)
	if err != nil {
		panic(err)
	}

	spreadsheetId := "1upV1TrDjZMP4V_ARar0UfB5Yn21O4iHHWFy1c_7-4V8"
	var readRange string
	readRange = fmt.Sprintf("TestData!%s", cell)

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		panic(err)
	}
	
	return (resp.Values[0][0]).(string)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path[1:]) 
}

func LookupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(Read(r.URL.Path[1:]))
	fmt.Fprintf(w, Read(r.URL.Path[1:]))
}
