package requests

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// MakeRequestWithAuthentication makes a request with token in the header
func MakeAuthenticatedRequest(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)

	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
