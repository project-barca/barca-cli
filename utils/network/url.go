package network

import (
	"fmt"
	"net/http"
	"strings"
)

// Check if served with HTTPS.
//
// What this function does is to find the final URL of a given URL and then check to see if the final URL has https or not.
// Return boolean value.
func VerifySiteSSL(url string) bool {
	originalURL := url

	resp, err := http.Get(originalURL)
	if err != nil {
		fmt.Println(err)
	}

	finalURL := resp.Request.URL.String()

	fmt.Println("Original URL: ", originalURL)
	fmt.Println("Final URL: ", finalURL)

	fmt.Println("\nHTTPS Ativado: ", strings.HasPrefix(finalURL, "https"))

	return strings.HasPrefix(finalURL, "https")
}
