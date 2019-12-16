package main

import "gopkg.in/jdkato/prose.v2"

type (
	Document struct {
		Title     string
		Link      string
		Language  string
		ShortText string
		Sentences    []Sentence
		Words  []string
		KeyWords []KeyWord
		WordWeightsStruct []kv
	}

	Sentence struct {
		Sentence prose.Sentence
		Weight float64
	}

	kv struct {
		Key   string
		Value float64
	}

	OstisReponse struct {
		Sys []interface{} `json:"sys"`
		Main []interface{} `json:"main"`
		Common []interface{} `json:"common"`
	}

	KeyWord struct {
		Word string
	}
)
