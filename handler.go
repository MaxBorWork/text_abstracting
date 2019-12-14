package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ClassicAbstractionHandler(c *gin.Context) {
	id := c.Param("title")
	document := docsMap[id]

	document.WordWeightsStruct = calcDocWeights(document)
	document.Sentences = calcSentsWeight(document)
	sort.Slice(document.Sentences, func(i, j int) bool {
		return document.Sentences[i].Weight > document.Sentences[j].Weight
	})
	docsMap[id] = document

	resultSentences := document.Sentences[:10]
	var text string
	for _, str := range resultSentences {
		text += str.Sentence.Text + " "
	}

	c.HTML(http.StatusOK, "result_classic.html", gin.H {
		"Id": id,
		"Title" : document.Title,
		"Link": document.Link,
		"Text" : text,
	})
}

//func KeyAbstraction(c *gin.Context)  {
//	var text string
//	id := c.Param("title")
//	document := docsMap[id]
//	url := mediawiki.CreateUrl(document.Language, document.Title)
//	text = mediawiki.MediaWikiRequest(url)
//	document.ShortText = text
//	docsMap[id] = document
//	document.Alphabet = textToLetters(text)
//	resultsArr := processDocumentAlphabet(&document)
//
//	comparison := getComparisonWithMaxRatio(resultsArr)
//	language := getDocLanguage(comparison.TestDocTitle)
//	docsMap[id] = document
//	if language != "" {
//		c.HTML(http.StatusOK, "result_alphabet.html", gin.H{
//			"Id": id,
//			"Title" : document.Title,
//			"Link": document.Link,
//			"Language": language,
//			"Alphabet": comparison.Alphabet,
//			"Ratio": comparison.Ratio,
//			"ResultsArr": resultsArr,
//		})
//
//		return
//	}
//
//	c.String(http.StatusNotFound, "Can't detect doc language")
//}

//func File(c *gin.Context) {
//	id := c.Param("title")
//	document := docsMap[id]
//	c.HTML(http.StatusOK, "file.html", gin.H{
//		"Title" : document.Title,
//		"Link": document.Link,
//		"Text": document.ShortText,
//	})
//}
