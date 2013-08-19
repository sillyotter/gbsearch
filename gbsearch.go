package gbsearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const queryString = "https://www.googleapis.com/books/v1/volumes?q=%s:%s"

const inTitle = "intitle"
const inAuthor = "inauthor"
const inPublisher = "inpublisher"
const subject = "subject"
const isbn = "isbn"
const lccn = "lccn"
const oclc = "oclc"

// https://developers.google.com/books/docs/v1/using#PerformingSearch
// need to expand to support all the options

func doSearch(searchType string, searchTerm string) (*Results, error) {
	target := fmt.Sprintf(queryString, searchType, url.QueryEscape(searchTerm))
	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	res := &Results{}
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func TitleSearch(searchTerm string) (*Results, error) {
	return doSearch(inTitle, searchTerm)
}

func AuthorSearch(searchTerm string) (*Results, error) {
	return doSearch(inAuthor, searchTerm)
}

func PublisherSearch(searchTerm string) (*Results, error) {
	return doSearch(inPublisher, searchTerm)
}

func SubjectSearch(searchTerm string) (*Results, error) {
	return doSearch(subject, searchTerm)
}

func ISBNSearch(searchTerm string) (*Results, error) {
	return doSearch(isbn, searchTerm)
}

func LCCNSearch(searchTerm string) (*Results, error) {
	return doSearch(lccn, searchTerm)
}

func OCLCSearch(searchTerm string) (*Results, error) {
	return doSearch(oclc, searchTerm)
}
