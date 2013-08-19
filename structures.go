package gbsearch

type ImageLink struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
	Small          string `json:"small,omitempty"`
	Medium         string `json:"medium,omitempty"`
	Large          string `json:"large,omitempty"`
	ExtraLarge     string `json:"extraLarge,omitempty"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type Dimension struct {
	Height    string `json:height"`
	Width     string `json:width"`
	Thickness string `json:thickness"`
}

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

type Price struct {
	Amount          float64 `json:"amount,omitempty"`
	AmmountInMicros float64 `json:"amountInMicros,omitempty"`
	CurrencyCode    string  `json:"currencyCode"`
}

type Offer struct {
	FinskyOfferType int   `json:"finskyOfferType"`
	ListPrice       Price `json:"listPrice"`
	RetailPrice     Price `json:"retailPrice"`
}

type SaleInfo struct {
	Country     string  `json:"country"`
	Saleability string  `json:"saleability"`
	IsEbook     bool    `json:"isEbook"`
	ListPrice   Price   `json:"listPrice,omitempty"`
	RetailPrice Price   `json:"retailPrice,omitempty"`
	BuyLink     string  `json:"buyLink,omitempty"`
	Offers      []Offer `json:"offers,omitempty"`
}

type Availability struct {
	IsAvailable  bool   `json:"isAvailable"`
	ACSTokenLink string `json:"acsTokenLink,omitempty"`
}

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

type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}

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

type Results struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []Item `json:"items"`
}
