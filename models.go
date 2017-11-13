package nytimes

import (
	"encoding/json"
	"net/url"
	"time"
)

type URL url.URL

func (u URL) String() string {
	z := url.URL(u)
	return z.String()
}

type Time time.Time

func (u *URL) UnmarshalJSON(b []byte) error {
	var s string
	var err error

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	location, err := url.Parse(s)
	if err != nil {
		return err
	}
	*u = URL(*location)
	return nil
}

func (u *Time) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	el, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	*u = Time(el)
	return nil
}

//
type Response struct {
	Status      string `json:"status"`
	Copyright   string `json:"copyright"`
	Section     string `json:"section"`
	LastUpdated Time   `json:"last_updated"`
	NumResults  int    `json:"num_results"`

	Results []Article `json:"results"`
}

//
type Article struct {
	Section           string       `json:"section"`
	Subsection        string       `json:"subsection"`
	Title             string       `json:"title"`
	Abstract          string       `json:"abstract"`
	URL               URL          `json:"url"`
	Byline            string       `json:"byline"`
	ItemType          string       `json:"item_type"`
	UpdatedDate       Time         `json:"updated_date"`
	CreatedDate       Time         `json:"created_date"`
	PublishedDate     Time         `json:"published_date"`
	MaterialTypeFacet string       `json:"material_type_facet"`
	Kicker            string       `json:"kicker"`
	DesFacet          []string     `json:"des_facet"`
	OrgFacet          []string     `json:"org_facet"`
	PerFacet          []string     `json:"per_facet"`
	GetFacet          []string     `json:"get_facet"`
	Multimedia        []Multimedia `json:"multimedia"`
	ShortURL          URL          `json:"short_url"`
}

//
type Multimedia struct {
	URL       URL    `json:"url"`
	Format    string `json:"format"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Caption   string `json:"caption"`
	Copyright string `json:"copyright"`
}
