package query

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Query struct {
	text string
}

func (q *Query) Execute() (string, error) {
	rand.Seed(time.Now().UnixNano())
	word := ""
	response, error := http.Get("https://sjp.pwn.pl/sjp/lista/" + q.text)
	if error != nil {
		return "", errors.New(fmt.Sprintf("Failed to connect: %s.", error))
	}

	document, error := goquery.NewDocumentFromReader(response.Body)
	if error != nil {
		return "", errors.New(fmt.Sprintf("Failed to parse response: %s.", error))
	}

	lastPage, error := getPagesCount(document)
	if error != nil {
		return "", errors.New(fmt.Sprintf("Failed to find pages count: %s.", error))
	}
	chosenPage := 1 + rand.Intn(lastPage - 1)

	word, error = getRandomResultFromPage(fmt.Sprintf("https://sjp.pwn.pl/sjp/lista/%s;%d.html", q.text, chosenPage))
	if error != nil {
		return "", errors.New(fmt.Sprintf("Failed to fetch random word: %s.", error))
	}
	return word, nil
}

func getRandomResultFromPage(pageUrl string) (string, error) {
	response, error := http.Get(pageUrl)
	if error != nil {
		return "", errors.New(fmt.Sprintf("Failed to connect: %s.", error))
	}

	document, error := goquery.NewDocumentFromReader(response.Body)
	if error != nil {
		return "", errors.New(fmt.Sprintf("Failed to parse response: %s.", error))
	}

	results := document.Find("ul.lista li a")
	pickedIndex := rand.Intn(results.Length())
	result := results.Eq(pickedIndex)
	return result.Text(), nil
}


func getPagesCount(resultPage *goquery.Document) (int, error) {
	paginations := resultPage.Find("ul.pagination")
	if paginations.Length() == 0 {
		return 0, errors.New("No pagination found.")
	}

	pagination := paginations.First()
	items := pagination.Find("li")
	lastPageStr := items.Eq(-2)

	if lastPageStr == nil {
		return 0, errors.New("No last page item found.")
	}

	lastPage, error := strconv.Atoi(lastPageStr.Text())
	if error != nil {
		return 0, errors.New(fmt.Sprintf("Failed to parse last result page number: %s.", error))
	}
	return lastPage, nil
}

func NewQuery(query string) Query {
	return Query{text: query}
}
