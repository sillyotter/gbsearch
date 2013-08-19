package gbsearch

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const inTitle = "intitle"
const inAuthor = "inauthor"
const inPublisher = "inpublisher"
const subject = "subject"
const isbn = "isbn"
const lccn = "lccn"
const oclc = "oclc"

type FilterType byte
type PrintType byte
type ProjectionType byte
type OrderType byte

const (
	PartialText FilterType = iota
	FullText
	FreeEbooks
	PaidEbooks
	EBooks
)

func (this FilterType) String() string {
	switch {
	case this == PartialText:
		return "partial"
	case this == FullText:
		return "full"
	case this == FreeEbooks:
		return "free-ebooks"
	case this == PaidEbooks:
		return "paid-ebooks"
	case this == EBooks:
		return "ebooks"
	}
	return "unknown"
}

const (
	All PrintType = iota
	Books
	Magazines
)

func (this PrintType) String() string {
	switch {
	case this == All:
		return "all"
	case this == Books:
		return "books"
	case this == Magazines:
		return "magazines"
	}
	return "unknown"
}

const (
	FullResults ProjectionType = iota
	Lite
)

func (this ProjectionType) String() string {
	switch {
	case this == FullResults:
		return "full"
	case this == Lite:
		return "lite"
	}
	return "unknown"
}

const (
	Relevance OrderType = iota
	Newest
)

func (this OrderType) String() string {
	switch {
	case this == Relevance:
		return "relevance"
	case this == Newest:
		return "newest"
	}
	return "unknown"
}

type Options struct {
	onlyEPubDownloads *bool
	filter            *FilterType
	startIndex        *int
	maxResults        *int
	printType         *PrintType
	projection        *ProjectionType
	orderBy           *OrderType
	languageCode      *string
}

func DefaultOptions() *Options {
	return &Options{}
}

func (this *Options) OnlyFindEPubDownloads(val bool) {
	this.onlyEPubDownloads = &val
}

func (this *Options) SetFilter(ft FilterType) {
	this.filter = &ft
}

func (this *Options) SetStartIndex(si int) {
	this.startIndex = &si
}

func (this *Options) SetMaxResults(mr int) {
	this.maxResults = &mr
}

func (this *Options) SetPrintType(pt PrintType) {
	this.printType = &pt
}

func (this *Options) SetProjection(p ProjectionType) {
	this.projection = &p
}

func (this *Options) SetOrderBy(o OrderType) {
	this.orderBy = &o
}

func (this *Options) SetLanguageCode(lc string) {
	this.languageCode = &lc
}

// https://developers.google.com/books/docs/v1/using#PerformingSearch
// need to expand to support all the options

func doSearch(searchType string, searchTerm string, options *Options) (*Results, error) {

	u := &url.URL{
		Scheme: "https",
		Host:   "www.googleapis.com",
		Path:   "/books/v1/volumes",
	}

	query := u.Query()

	if options != nil {
		if options.onlyEPubDownloads != nil && *options.onlyEPubDownloads == true {
			query.Add("download", "epub")
		}

		if options.filter != nil {
			query.Add("filter", options.filter.String())
		}

		if options.languageCode != nil {
			query.Add("langRestrict", *options.languageCode)
		}

		if options.startIndex != nil {
			query.Add("startIndex", strconv.Itoa(*options.startIndex))
		}

		if options.maxResults != nil {
			query.Add("maxResults", strconv.Itoa(*options.maxResults))
		}

		if options.projection != nil {
			query.Add("projection", options.projection.String())
		}

		if options.printType != nil {
			query.Add("printType", options.printType.String())
		}

		if options.orderBy != nil {
			query.Add("orderBy", options.orderBy.String())
		}
	}

	query.Add("q", searchType+":"+url.QueryEscape(searchTerm))

	u.RawQuery = query.Encode()

	target := u.String()

	log.Println(target)

	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &Results{}
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func TitleSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(inTitle, searchTerm, options)
}

func AuthorSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(inAuthor, searchTerm, options)
}

func PublisherSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(inPublisher, searchTerm, options)
}

func SubjectSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(subject, searchTerm, options)
}

func ISBNSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(isbn, searchTerm, options)
}

func LCCNSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(lccn, searchTerm, options)
}

func OCLCSearch(searchTerm string, options *Options) (*Results, error) {
	return doSearch(oclc, searchTerm, options)
}
