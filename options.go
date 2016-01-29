package gbsearch

// SearchType is used to represent the various detailed kinds of searches we
// can do.  Google Books api by default does a general search, but we can
// constrain it to more specific searches.
type SearchType string

// These constants define the various kinds of searches we can do and will
// be used later to specify the kind of search we want to do.
const (
	InTitle           SearchType = "intitle"
	InAuthor          SearchType = "inauthor"
	InPublisher       SearchType = "inpublisher"
	Subject           SearchType = "subject"
	ISBN              SearchType = "isbn"
	LCCN              SearchType = "lccn"
	OCLC              SearchType = "oclc"
	UnknownSearchType SearchType = "unknown"
)

// FilterType is used to specify the kind of filtering we want to do to the search.
type FilterType string

// These constants are used to specify the filter type we want applied.  We can restrict the
// searches to return books that have partial text for viewing, full text, that are ebooks,
// either free or paid.
const (
	UnknownFilterType FilterType = "unknown"
	PartialText       FilterType = "partial"
	FullText          FilterType = "full"
	FreeEbooks        FilterType = "free-ebooks"
	PaidEbooks        FilterType = "paid-ebooks"
	EBooks            FilterType = "ebooks"
)

// PrintType is used to specify which kind of media we are looking for.
type PrintType string

// These constants are used to specify the media type we want, books, magazines, or all.
const (
	UnknownPrintType PrintType = "unknown"
	All              PrintType = "all"
	Books            PrintType = "books"
	Magazines        PrintType = "magazines"
)

// ProjectionType is used to specify how much data we want returned to us.
type ProjectionType string

// These constants are used to specify if we want full or lite levels of data.
const (
	UnknownProjectionType ProjectionType = "unknown"
	FullResults           ProjectionType = "full"
	Lite                  ProjectionType = "lite"
)

// OrderType is used to specify how to order the data.
type OrderType string

// These constants are used to specify if we want the data ordered by relevance or by newest.
const (
	UnknownOrderByType OrderType = "unknown"
	Relevance          OrderType = "relevance"
	Newest             OrderType = "newest"
)

// Options are used to collect all of our various option settings into one place to allow
// them all to be passed into the various search functions at once.  Use the various
// SetXXX functions to adjust its contents
type Options struct {
	onlyEPubDownloads *bool
	filter            *FilterType
	startIndex        *int
	maxResults        *int
	printType         *PrintType
	projection        *ProjectionType
	orderBy           *OrderType
	languageCode      *string
	countryCode       *string
}

// DefaultOptions returns a new Options struct set with the default values.  Currently, the default
// is for none of the options to be set.
func DefaultOptions() *Options {
	return &Options{}
}

// OnlyFindEPubDownloads will adjust whether we are to search for all books or just
// downloadable ePub books.
func (this *Options) OnlyFindEPubDownloads(val bool) {
	this.onlyEPubDownloads = &val
}

// SetFilter is used to set the filter type.
func (this *Options) SetFilter(ft FilterType) {
	if ft != UnknownFilterType {
		this.filter = &ft
	}
}

// SetStartIndex is used to set the start index.  In conjunction with SetMaxResults, we can fetch
// pages of data from the server.
func (this *Options) SetStartIndex(si int) {
	if si > 0 {
		this.startIndex = &si
	}
}

// SetMaxResults is used to set the max number of results to fetch.  The max is 40.  In conjunction with
// SetStartIndex, we can fetch pages of data from the server.
func (this *Options) SetMaxResults(mr int) {
	if mr > 0 {
		this.maxResults = &mr
	}
}

// SetPrintType is used to set the media type.
func (this *Options) SetPrintType(pt PrintType) {
	if pt != UnknownPrintType {
		this.printType = &pt
	}
}

// SetProjection is used to set how much data we want brought back with reach result.
func (this *Options) SetProjection(p ProjectionType) {
	if p != UnknownProjectionType {
		this.projection = &p
	}
}

// SetOrderBy is used to set how to order the data.
func (this *Options) SetOrderBy(o OrderType) {
	if o != UnknownOrderByType {
		this.orderBy = &o
	}
}

// SetLanguageCode will let us set which language code to restrict our search to.  Use the two leter
// standard country code, such as "en", "fr", etc...
func (this *Options) SetLanguageCode(lc string) {
	if lc != "" {
		this.languageCode = &lc
	}
}

// SetCountryCode will let us set which country IP to restrict our search to.  Use the two leter
// standard country code, such as "us", "fr", etc...
func (this *Options) SetCountryCode(cc string) {
	if cc != "" {
		this.countryCode = &cc
	}
}
