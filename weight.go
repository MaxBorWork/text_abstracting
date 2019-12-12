package main

import "math"

func calcDocWeights(document Document)  {
	for _, word := range document.Words {
		calculateWordWeight(word, document)
	}
}

func calculateWordWeight(word string, document Document) float64 {
	tf := calcTf(word, document.Words)
	idf := calcIdf(word)
	return tf * idf
}

func calcTf(word string, words []string) float64 {
	ndk := calcNdk(word, words)
	return ndk / float64(len(words))
}

func calcIdf(word string) float64 {
	docsInBase := float64(len(docsMap))
	docsWithWord := getDocsWithWord(word)
	return math.Log(docsInBase / docsWithWord)
}

func calcNdk(word string, words []string) float64 {
	var colOfWordInDoc int
	for _, docWord := range words {
		if word == docWord {
			colOfWordInDoc = colOfWordInDoc + 1
		}
	}
	return float64(colOfWordInDoc)
}

func getDocsWithWord(word string) float64 {
	var docsWithWord float64
	for _, doc := range docsMap {
		for _, docWord := range doc.Words {
			if word == docWord {
				docsWithWord = docsWithWord + 1
				break
			}
		}
	}

	return docsWithWord
}
