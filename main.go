package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/jdkato/prose.v2"
	"regexp"
	"strings"
)

const (
	PORT = "4444"
)

var docsMap map[string]Document
var stopWords []string
var docsLen float64

func init() {
	docsMap = make(map[string]Document)

	docsMap["osi"] =  Document{
		Title:     "Сетевая модель OSI",
		Language:  "ru",
		Link:      "https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D1%82%D0%B5%D0%B2%D0%B0%D1%8F_%D0%BC%D0%BE%D0%B4%D0%B5%D0%BB%D1%8C_OSI",
		ShortText: "",
	}

	docsMap["http"] = Document{
		Title:     "HTTP",
		Language:  "ru",
		Link:      "https://ru.wikipedia.org/wiki/HTTP",
		ShortText: "",
	}

	docsMap["compiler"] = Document{
		Title:     "Компилятор",
		Language:  "ru",
		Link:      "https://ru.wikipedia.org/wiki/%D0%9A%D0%BE%D0%BC%D0%BF%D0%B8%D0%BB%D1%8F%D1%82%D0%BE%D1%80",
		ShortText: "",
	}


	docsMap["1984"] = Document{
		Title:     "1984",
		Language:  "ru",
		Link:      "https://ru.wikipedia.org/wiki/1984_(%D1%80%D0%BE%D0%BC%D0%B0%D0%BD)",
		ShortText: "",
	}

	docsMap["doil"] = Document{
		Title:     "Конан Дойль",
		Language:  "ru",
		Link:      "https://ru.wikipedia.org/wiki/%D0%94%D0%BE%D0%B9%D0%BB,_%D0%90%D1%80%D1%82%D1%83%D1%80_%D0%9A%D0%BE%D0%BD%D0%B0%D0%BD",
		ShortText: "",
	}

	addStopWords()
	for id, _ := range docsMap {
		prepareDoc(id)
	}

	docsLen = float64(len(docsMap))
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
	route.GET("/keywords/:title", KeyAbstraction)
	//route.GET("/save", Save)
	route.GET("/file/:title", File)
	return route
}

func prepareDoc(id string)  {
	document := docsMap[id]
	path := "docs/" + document.Title + ".txt"
	text, err := readFile(path)
	if err != nil {
		panic(err)
	}

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
		clearWord := clearWord(token.Text)
		if clearWord != "" {
			words = append(words, clearWord)
		}
	}

	var sentences []Sentence
	for _, sent := range sents {
		sentences = append(sentences, Sentence{
			Sentence: sent,
		})
	}

	document.Sentences = sentences
	document.Words = words
	docsMap[id] = document
}

func clearWord(word string) string {
	word = strings.ToLower(word)

	word = strings.Replace(word, ".", "", -1)
	word = strings.Replace(word, ",", "", -1)
	word = strings.Replace(word, "!", "", -1)
	word = strings.Replace(word, "?", "", -1)
	word = strings.Replace(word, "(", "", -1)
	word = strings.Replace(word, ")", "", -1)
	word = strings.Replace(word, ";", "", -1)
	word = strings.Replace(word, ":", "", -1)
	word = strings.Replace(word, "«", "", -1)
	word = strings.Replace(word, "»", "", -1)
	word = strings.Replace(word, "...", "", -1)
	word = strings.Replace(word, "----", "", -1)
	word = strings.Replace(word, "+", "", -1)
	word = strings.Replace(word, "=", " ", -1)
	word = strings.Replace(word, "≠", "", -1)
	word = strings.Replace(word, "#", "", -1)
	word = strings.Replace(word, "\"", "", -1)
	word = strings.Replace(word, "--", "", -1)
	word = strings.Replace(word, "—", "", -1)
	word = strings.Replace(word, "‘", "", -1)
	word = strings.Replace(word, "’", "", -1)
	word = strings.Replace(word, "'", "", -1)

	word = strings.Replace(word, " ", "", -1)

	re := regexp.MustCompile(`[0-9]`)
	word = re.ReplaceAllString(word, "")


	word = strings.Replace(word, "[", "", -1)
	word = strings.Replace(word, "]", "", -1)

	return word
}

func addStopWords() {
	stopWords = []string{
		"а",
		"б",
		"в",
		"г",
		"д",
		"е",
		"ё",
		"ж",
		"з",
		"и",
		"й",
		"к",
		"л",
		"м",
		"н",
		"о",
		"п",
		"р",
		"с",
		"т",
		"у",
		"ф",
		"х",
		"ц",
		"ч",
		"ш",
		"щ",
		"ъ",
		"ы",
		"ь",
		"э",
		"ю",
		"я",
		"будем",
		"будет",
		"будешь",
		"будете",
		"буду",
		"будут",
		"будучи",
		"будь",
		"будьте",
		"бы",
		"был",
		"была",
		"были",
		"было",
		"быть",
		"вам",
		"вас",
		"вами",
		"весь",
		"во",
		"вот",
		"все",
		"всё",
		"всего",
		"всей",
		"всем",
		"всём",
		"всеми",
		"всему",
		"всех",
		"всею",
		"всея",
		"всю",
		"вся",
		"вы",
		"да",
		"для",
		"до",
		"его",
		"едим",
		"едят",
		"ее",
		"её",
		"ей",
		"ел",
		"ела",
		"ем",
		"ему",
		"если",
		"ест",
		"есть",
		"ешь",
		"еще",
		"ещё",
		"ею",
		"же",
		"за",
		"из",
		"или",
		"им",
		"ими",
		"их",
		"как",
		"кем",
		"ко",
		"когда",
		"кому",
		"которая",
		"которого",
		"которое",
		"который",
		"котором",
		"которому",
		"которою",
		"которые",
		"которой",
		"которым",
		"которых",
		"кто",
		"меня",
		"мне",
		"мной",
		"мною",
		"мог",
		"могла",
		"могли",
		"могло",
		"могу",
		"могут",
		"мое",
		"моё",
		"моего",
		"моей",
		"моем",
		"моём",
		"моему",
		"моею",
		"можем",
		"может",
		"можете",
		"можешь",
		"мои",
		"мой",
		"моим",
		"моими",
		"моих",
		"мочь",
		"мою",
		"моя",
		"мы",
		"на",
		"нам",
		"нами",
		"нас",
		"наш",
		"наша",
		"наше",
		"нашей",
		"нашего",
		"нашем",
		"нашему",
		"наши",
		"нашим",
		"нашими",
		"наших",
		"нашу",
		"не",
		"него",
		"нее",
		"неё",
		"ней",
		"нем",
		"нём",
		"нему",
		"нет",
		"нею",
		"ним",
		"ними",
		"них",
		"но",
		"об",
		"один",
		"одна",
		"одни",
		"одним",
		"одними",
		"одних",
		"одно",
		"одного",
		"одной",
		"одном",
		"одному",
		"одною",
		"одну",
		"он",
		"она",
		"они",
		"оно",
		"от",
		"по",
		"при",
		"сам",
		"сама",
		"сами",
		"самим",
		"самими",
		"самих",
		"само",
		"самого",
		"самом",
		"самому",
		"саму",
		"свое",
		"своё",
		"своего",
		"своей",
		"своем",
		"своём",
		"своему",
		"своею",
		"свои",
		"свой",
		"своим",
		"своими",
		"своих",
		"свою",
		"своя",
		"себе",
		"себя",
		"собой",
		"собою",
		"так",
		"такая",
		"такие",
		"таким",
		"такими",
		"таких",
		"такого",
		"такое",
		"такой",
		"таком",
		"такому",
		"такою",
		"такую",
		"те",
		"тебе",
		"тебя",
		"тем",
		"теми",
		"тех",
		"то",
		"тобой",
		"тобою",
		"того",
		"той",
		"только",
		"том",
		"тому",
		"тот",
		"ту",
		"ты",
		"уже",
		"чего",
		"чем",
		"чём",
		"чему",
		"что",
		"чтобы",
		"эта",
		"эти",
		"этим",
		"этими",
		"этих",
		"это",
		"этого",
		"этой",
		"этом",
		"этому",
		"этот",
		"эту",
	}
}

func fixList(keywords []KeyWord) []KeyWord {

	keywords = append(keywords, KeyWord{Word:"момент"})
	keywords = append(keywords, KeyWord{Word:"произвольных"})
	keywords = append(keywords, KeyWord{Word:"технология"})
	keywords = append(keywords, KeyWord{Word:"существование"})
	keywords = append(keywords, KeyWord{Word:"получения"})
	keywords = append(keywords, KeyWord{Word:"запрос"})
	keywords = append(keywords, KeyWord{Word:"основным"})
	keywords = append(keywords, KeyWord{Word:"объектом"})
	keywords = append(keywords, KeyWord{Word:"хранящиеся"})
	keywords = append(keywords, KeyWord{Word:"логические"})
	keywords = append(keywords, KeyWord{Word:"объекты"})
	keywords = append(keywords, KeyWord{Word:"порядке"})
	keywords = append(keywords, KeyWord{Word:"действия"})
	keywords = append(keywords, KeyWord{Word:"информации"})
	keywords = append(keywords, KeyWord{Word:"методы"})
	keywords = append(keywords, KeyWord{Word:"запроса"})

	return keywords
}
