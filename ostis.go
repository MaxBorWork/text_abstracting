package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getKeyWords(words []string) []KeyWord {
	var keywords []KeyWord
	uniqueWords := unique(words)
	for _, word := range uniqueWords {
		if wordCorrect(word) {
			resp := sendOstisRequest(word)
			if isWordKey(resp) {
				keywords = append(keywords, KeyWord{Word:word})
			}
		}
	}

	return keywords
}

func sendOstisRequest(word string) OstisReponse {
	url := `http://ims.ostis.net/api/idtf/find/?substr=` + word
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	req.Header.Set("Charset-Encoding", "UTF-8")
	req.Header.Set("Charset", "UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var ostisResponse OstisReponse
	var jsonResp map[string]interface{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &ostisResponse); err != nil {
		panic(err)
	}

	return ostisResponse
}

func isWordKey(response OstisReponse) bool {
	if len(response.Common) != 0 {
		return true
	}

	if len(response.Main) != 0 {
		return true
	}

	if len(response.Sys) != 0 {
		return true
	}

	return false
}