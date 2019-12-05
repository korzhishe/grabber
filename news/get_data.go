package news

import "encoding/json"

import "fmt"

import "net/http"

import "io/ioutil"

type source struct {
	ID string `json:"id"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

func getSources(cat string) []source {
	body := getData(sourceURL(cat))

	var sourceAPI sourcesAPI
	json.Unmarshal(body, &sourceAPI)

	return sourceAPI.Sources
}

func sourceURL(cat string) string {
	return fmt.Sprintln("https://newsapi.org/v1/sources?language=en&category=%s", cat)
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res)
	if err != nil {
		panic(err)
	}

	return body
}
