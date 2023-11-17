// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/extra/catalogTags.go
// Original timestamp: 2023/11/14 19:47

package extra

import (
	"dtools/helpers"
	"fmt"
	"io"
	"net/http"
)

func genericFetch(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return helpers.CustomError{Message: "Unable to fetch info from given URL: " + err.Error()}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return helpers.CustomError{Message: "Response is unreadable: " + err.Error()}
	}

	// Check if the response status code is not successful (2xx)
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return helpers.CustomError{Message: fmt.Sprintf("http return code: %v\n", response.StatusCode)}
	}

	// Print the JSON-formatted response body
	fmt.Println(string(body))

	return nil
}

func GetCatalog(registry string) error {
	if registry == "" {

	}
	return nil
}
