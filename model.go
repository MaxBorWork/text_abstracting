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
		KeyWords []string
		WordWeightsStruct []kv
	}

	Sentence struct {
		Sentence prose.Sentence
		Weight float64
	}

	ClassicAbstraction struct {
		Title     string
		Link      string
		Sentences    []Sentence
	}

	kv struct {
		Key   string
		Value float64
	}
)
