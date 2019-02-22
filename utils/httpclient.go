package utils

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}
