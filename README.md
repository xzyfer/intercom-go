intercom-go
=============

[Go](http://golang.org) bindings for the Intercom API (https://api.intercom.io).

[API Documentation](https://api.intercom.io/docs)

This is an unofficial library that is not affiliated with [Intercom](https://intercom.io). Official libraries are available at
[github.com/intercom](https://github.com/intercom).

[![Build Status](https://travis-ci.org/xzyfer/intercom-go.png?branch=master)](https://travis-ci.org/xzyfer/intercom-go)

Example usage
-------------

```go
package intercom_test

import (
    "fmt"
    "github.com/xzyfer/intercom-go/intercom"
    "net/http"
    "os"
)

func Example() {
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

    tags, err := apiclient.ListTags()
    if err != nil {
        fmt.Printf("Error listing tags: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("Found %d tags.\n", len(tags))
    for _, tag := range tags {
        fmt.Printf(" - %s (ID: [%d-char ID])\n", tag.Name, len(tag.ID))
    }
}
```

Running the tests
-----------------

To run the tests:

```
go test ./intercom
```

Setup
-----------------

You'll need to configure the `INTERCOM_API_KEY` and `INTERCOM_APP_ID` environment variables
to you're Intercom.io `api_key` and `app_id` respectively. This information can be found in
your Intercom account's App Settings > API Keys menu.

```
INTERCOM_API_KEY=your-api0key INTERCOM_APP_ID=your-app-id ./intercom list-tags
```

Implemented APIs
----------------

Currently only a small subset of API endpoints are implemented.
Please open an issue if you required additional API endpoints.

- [List tags](https://doc.intercom.io/api/#list-tags-for-an-app)

Acknowledgements
----------------

The library's architecture and testing code are adapted from
[go-github](https://github.com/google/go-github), created by [Will
Norris](https://github.com/willnorris).
