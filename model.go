package main

import "gopkg.in/jdkato/prose.v2"

type (
	Document struct {
		Title     string
		Link      string
		Language  string
		ShortText string
		Sentences    []prose.Sentence
		Words  []string
		WordWeights map[string]float32
		KeyWords []string
	}

	ClassicAbstraction struct {
		Title     string
		Link      string
		Sentences    []prose.Sentence
	}

)
