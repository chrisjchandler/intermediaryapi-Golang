package main

import (
    "encoding/json"
    "net/http"
)

type RequestData struct {
    Params map[string]string `json:"params"`
}

func translateAPIHandler(w http.ResponseWriter, r *http.Request) {
    // Decode the request body into a RequestData struct
    decoder := json.NewDecoder(r.Body)
    var data RequestData
    err := decoder.Decode(&data)
    if err != nil {
        // Handle the error
    }

    // Create an http.Client to make the request to the other API
    client := http.Client{}

    // Make a request to the other API using the http.Client
    resp, err := client.Get("http://other-api.com/endpoint", data.Params)
    if err != nil {
        // Handle the error
    }

    // Defer closing the response body
    defer resp.Body.Close()

    // Write the response from the other API to the response writer
    w.Write(resp.Body)
}

func main() {
    // Handle requests to the /translate-api endpoint
    http.HandleFunc("/translate-api", translateAPIHandler)

    // Start the HTTP server
    http.ListenAndServe(":5000", nil)
}
