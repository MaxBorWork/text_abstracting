package main

import (
	"gopkg.in/jdkato/prose.v2"
	"math"
)

func calcSentsWeight(document Document) []Sentence {
	for i, sentence := range document.Sentences {
		document.Sentences[i].Weight = calcSentenceWeight(sentence, document)
	}

	return document.Sentences
}

func calcSentenceWeight(sentence Sentence, document Document) float64 {
	var weight float64
	doc, _ := prose.NewDocument(sentence.Sentence.Text)
	tokens := doc.Tokens()

	var words []string
	for _, token := range tokens {
		clearWord := clearWord(token.Text)
		if clearWord != "" {
			words = append(words, clearWord)
		}
	}

	uniqueWords := unique(words)

	for _, word := range uniqueWords {
		freq := calcWordFrequencyInSentence(word, words)
		wordWeight := calcWordWeightForDocument(word, document)
		weight += freq * wordWeight
	}

	return weight
}

func calcWordFrequencyInSentence(word string, words []string) float64 {
	var wordCol float64
	for _, wordIter := range words {
		if word == wordIter {
			if wordCol != 0 {
				wordCol = wordCol + 1
				continue
			}

			wordCol = 1
		}
	}

	return wordCol / float64(len(words))
}

func calcWordWeightForDocument(word string, doc Document) float64 {
	for _, WordWeight := range doc.WordWeightsStruct {
		if word == WordWeight.Key {
			return 0.5 * (WordWeight.Value/doc.WordWeightsStruct[0].Value + 1) * math.Log(docsLen/getDocsWithWord(word))
		}
	}

	return 0
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
