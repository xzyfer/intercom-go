package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xzyfer/intercom-go/intercom"
)

var verbose = flag.Bool("v", false, "verbose")

var apiclient *intercom.APIClient

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: intercom command [OPTS] ARGS...\n")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "The commands are:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tlist-tags")
		// fmt.Fprintln(os.Stderr, "\tlist-subscribers LIST (active|unconfirmed|unsubscribed|bounced|deleted)")
		// fmt.Fprintln(os.Stderr)
		// fmt.Fprintln(os.Stderr)
		// fmt.Fprintln(os.Stderr, "Common arguments:")
		// fmt.Fprintln(os.Stderr)
		// fmt.Fprintln(os.Stderr, "\tCLIENT:\ta client ID")
		// fmt.Fprintln(os.Stderr, "\tLIST:\ta list ID")
		// fmt.Fprintln(os.Stderr, "\tEMAIL:\temail address")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Run `intercom command -h` for more information.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}
	log.SetFlags(0)

	apiKey := os.Getenv("INTERCOM_API_KEY")
	appId := os.Getenv("INTERCOM_APP_ID")
	if appId == "" {
		log.Fatal("Error: you must set your Intercom App ID in the INTERCOM_APP_ID environment variable.")
	}
	if apiKey == "" {
		log.Fatal("Error: you must set your Intercom API key in the INTERCOM_API_KEY environment variable.")
	}
	authClient := &http.Client{
		Transport: &intercom.APIKeyAuthTransport{APIKey: apiKey, AppID: appId},
	}
	apiclient = intercom.NewAPIClient(authClient)

	if *verbose {
		apiclient.Log = log.New(os.Stderr, "intercom: ", 0)
	}

	subcmd := flag.Arg(0)
	remaining := flag.Args()[1:]
	switch subcmd {
	case "list-tags":
		listTags(remaining)
	}
}

func listTags(args []string) {
	tags, err := apiclient.ListTags()
	if err != nil {
		log.Fatalf("Error listing tags: %s\n", err)
	}
	if len(tags) == 0 {
		fmt.Println("No tags found.")
		return
	}
	for _, t := range tags {
		fmt.Printf("%-24s %s\n", t.Name, t.ID)
	}
}
