package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/jdkato/prose.v2"
	"mediawiki"
)

const (
	PORT = "4444"
)

var docsMap map[string]Document

func init() {

	docsMap = make(map[string]Document)

	docsMap["chess_fr"] =  Document{
		Title:     "Échecs",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/%C3%89checs",
		ShortText: "",
	}

	docsMap["saint_martin_fr"] = Document{
		Title:     "Saint-Martin-sur-Écaillon",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Saint-Martin-sur-%C3%89caillon",
		ShortText: "",
	}

	docsMap["chretien_fr"] = Document{
		Title:     "Albert_Chrétien",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Albert_Chr%C3%A9tien",
		ShortText: "",
	}

	docsMap["santos_fr"] = Document{
		Title:     "Alberto_Santos-Dumont",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Alberto_Santos-Dumont",
		ShortText: "",
	}

	docsMap["soprano_fr"] = Document{
		Title:     "Mezzo-soprano",
		Language:  "fr",
		Link:      "https://fr.wikipedia.org/wiki/Mezzo-soprano",
		ShortText: "",
	}


	docsMap["futurama_en"] = Document{
		Title:     "Futurama",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Futurama",
		ShortText: "",
	}

	docsMap["jonny_quest_en"] = Document{
		Title:     "Jonny_Quest",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Jonny_Quest",
		ShortText: "",
	}

	docsMap["twisted_sister_en"] = Document{
		Title:     "Twisted_Sister",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Twisted_Sister",
		ShortText: "",
	}

	docsMap["new_york_en"] = Document{
		Title:     "New_York_City",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/New_York_City",
		ShortText: "",
	}

	docsMap["Buffalo_Oklahoma_en"] = Document{
		Title:     "Buffalo,_Oklahoma",
		Language:  "en",
		Link:      "https://en.wikipedia.org/wiki/Buffalo,_Oklahoma",
		ShortText: "",
	}

	for id, _ := range docsMap {
		prepareDoc(id)
	}
}

func main() {
	router := route()
	router.LoadHTMLGlob("templates/*")
	err := router.Run(":" + PORT)
	if err != nil {
		panic(err)
	}
}

func route() *gin.Engine {
	route := gin.Default()
	route.GET("/", Index)
	route.GET("/classic/:title", ClassicAbstractionHandler)
	//route.GET("/alphabet/:title", KeyAbstraction)
	//route.GET("/save", Save)
	//route.GET("/file/:title", File)
	return route
}

func prepareDoc(id string)  {
	document := docsMap[id]
	url := mediawiki.CreateUrl(document.Language, document.Title)
	text := mediawiki.MediaWikiRequest(url)
	document.ShortText = text
	doc, err := prose.NewDocument(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	sents := doc.Sentences()
	tokens := doc.Tokens()
	var words []string

	for _, token := range tokens {
		words = append(words, token.Text)
	}

	document.Sentences = sents
	document.Words = words
	docsMap[id] = document
}
