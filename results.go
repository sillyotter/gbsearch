package gbsearch

// ImageLink will contain urls for various sized images of the book covers.
type ImageLink struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
	Small          string `json:"small,omitempty"`
	Medium         string `json:"medium,omitempty"`
	Large          string `json:"large,omitempty"`
	ExtraLarge     string `json:"extraLarge,omitempty"`
}

// IndustryIdentifiers will contain various id values, mostly ISBN numbers.
type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

// Dimension will contain the size of physical items.
type Dimension struct {
	Height    string `json:height"`
	Width     string `json:width"`
	Thickness string `json:thickness"`
}

// VolumeInfo contains all the standard information about a particular volume.
type VolumeInfo struct {
	Title               string               `json:"title"`
	Subtitle            string               `json:"subtitle"`
	Authors             []string             `json:"authors"`
	Publisher           string               `json:"publisher"`
	PublishedDate       string               `json:"publishedDate"`
	Description         string               `json:"description"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
	PageCount           int                  `json:"pageCount"`
	Dimensions          Dimension            `json:"dimensions,omitempty"`
	PrintType           string               `json:"printType"`
	MainCategory        string               `json:"mainCategory,omitempty"`
	Categories          []string             `json:"categories"`
	AverageRating       float32              `json:"averageRating"`
	RatingsCount        int                  `json:"ratingsCount"`
	ContentVersion      string               `json:"contentVersion"`
	ImageLinks          ImageLink            `json:"imageLinks"`
	Language            string               `json:"language"`
	PreviewLink         string               `json:"previewLink"`
	InfoLink            string               `json:"infoLink"`
	CanonicalVolumeLink string               `json:"canonicalVolumeLink"`
}

// Price contains pricing information.
type Price struct {
	Amount          float64 `json:"amount,omitempty"`
	AmmountInMicros float64 `json:"amountInMicros,omitempty"`
	CurrencyCode    string  `json:"currencyCode"`
}

// Offer contains information about offers.  I'm not really sure what this data is, FinksyOfferType makes no sense to me.
type Offer struct {
	FinskyOfferType int   `json:"finskyOfferType"`
	ListPrice       Price `json:"listPrice"`
	RetailPrice     Price `json:"retailPrice"`
}

// SaleInfo contains information about selling the book.
type SaleInfo struct {
	Country     string  `json:"country"`
	Saleability string  `json:"saleability"`
	IsEbook     bool    `json:"isEbook"`
	ListPrice   Price   `json:"listPrice,omitempty"`
	RetailPrice Price   `json:"retailPrice,omitempty"`
	BuyLink     string  `json:"buyLink,omitempty"`
	Offers      []Offer `json:"offers,omitempty"`
}

// Availablity contains information about the availability of various formats.
type Availability struct {
	IsAvailable  bool   `json:"isAvailable"`
	ACSTokenLink string `json:"acsTokenLink,omitempty"`
}

// AccessInfo contains information about how the content can be consumed.
type AccessInfo struct {
	Country                string       `json:"country"`
	Viewability            string       `json:"viewability"`
	Embeddable             bool         `json:"embeddable"`
	PublicDomain           bool         `json:"publicDomain"`
	TextToSpeechPermission string       `json:"textToSpeechPermission"`
	EPub                   Availability `json:"epub"`
	PDF                    Availability `json:"pdf"`
	WebReaderLink          string       `json:"webReaderLink"`
	AccessViewStatus       string       `json:"accessViewStatus"`
}

// SearchInfo contains a snippet about the text.
type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}

// Item is a container that holds various volume, sales, and access information, along with some identifiers.
type Item struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	ETag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo AccessInfo `json:"accessInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
}

// Result contains a collection of items returned by the search.  Items will contain by default 10 items, and at most 40.
// The number of items may not match the TotalItems value.  Using the query Options StartIndex and MaxResults you can
// page through the data.
type Results struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []Item `json:"items"`
}
