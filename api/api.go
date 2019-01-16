package api

import (
	"github.com/lepra-tsr/gdbt/config"
	"io/ioutil"
	"net/http"
)

func CallGetWithCredential(path string) ([]byte, error) {
	_, token, err := config.ReadCredential()
	if err != nil {
		return nil, err
	}

	url := "https://idobata.io/api" + path
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Token", token)
	req.Header.Set("User-Agent", "idbt")

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	return bytes, nil
}
