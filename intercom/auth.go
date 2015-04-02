package intercom

import (
    "net/http"
)

type APIKeyAuthTransport struct {
    Transport http.RoundTripper
    AppID   string
    APIKey  string
}

func (t *APIKeyAuthTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
    transport := t.Transport
    if transport == nil {
        transport = http.DefaultTransport
    }

    req.SetBasicAuth(t.AppID, t.APIKey)

    return transport.RoundTrip(req)
}
