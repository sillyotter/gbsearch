// Package gbsearch provides a way to search Google books API.
package gbsearch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// https://developers.google.com/books/docs/v1/using#PerformingSearch

// Search will perform the given type of search on the given search term, using the provided
// options. Options can be nil if you don't want to change anything from the defaults
func Search(searchType SearchType, searchTerm string, options *Options) (*Results, error) {

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
			query.Add("filter", string(*options.filter))
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
			query.Add("projection", string(*options.projection))
		}

		if options.printType != nil {
			query.Add("printType", string(*options.printType))
		}

		if options.orderBy != nil {
			query.Add("orderBy", string(*options.orderBy))
		}
	}

	if searchType != UnknownSearchType {
		query.Add("q", string(searchType)+":"+searchTerm)
	} else {
		query.Add("q", searchTerm)
	}

	u.RawQuery = query.Encode()

	target := u.String()

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

// TitleSearch will perform a title search on the given search term and with the given options
func TitleSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(InTitle, searchTerm, options)
}

// AuthorSearch will perform an author search on the given search term and with the given options
func AuthorSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(InAuthor, searchTerm, options)
}

// PublisherSearch will perform a publisher search on the given search term and with the given options
func PublisherSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(InPublisher, searchTerm, options)
}

// SubjectSearch will perform a subject search on the given search term and with the given options
func SubjectSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(Subject, searchTerm, options)
}

// ISBNSearch will perform a isbn search on the given search term and with the given options
func ISBNSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(ISBN, searchTerm, options)
}

// LCCNSearch will perform a lccn search on the given search term and with the given options
func LCCNSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(LCCN, searchTerm, options)
}

// OCLCSearch will perform a oclc number search on the given search term and with the given options
func OCLCSearch(searchTerm string, options *Options) (*Results, error) {
	return Search(OCLC, searchTerm, options)
}
