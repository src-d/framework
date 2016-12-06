package utils

import "fmt"

type Source string

const (
	LinkedInSource Source = "linkedin"
	GithubSource   Source = "github"
	AugurSource    Source = "augur"
)

type RawLocation struct {
	Name   string
	Source Source
}

type LocationMethod string

const (
	FromRawMethod LocationMethod = "from-raw"
	GuessedMethod LocationMethod = "guessed"
)

type Location struct {
	Method           LocationMethod
	RawLocation      `bson:",inline,omitempty"`
	GeocodedLocation `bson:",inline,omitempty"`
}

func (l Location) String() string {
	return fmt.Sprintf("%s (%s)", l.GeocodedLocation.String(), l.Method)
}

type LocationLevel string

const (
	CountryLevel   LocationLevel = "country"
	LocalityLevel  LocationLevel = "locality"
	AggregateLevel LocationLevel = "aggregate"
)

type Coordinates struct {
	Lat  float64
	Long float64
}

type GeocodedLocation struct {
	FullAddress string
	Level       LocationLevel
	Locality    string      `bson:",omitempty"`
	Country     string      `bson:",omitempty"`
	Aggregate   string      `bson:",omitempty"`
	Coordinates Coordinates `bson:",omitempty"`
	Bounds      struct {
		Northeast Coordinates
		Southwest Coordinates
	} `bson:",omitempty"`
	FirstLevel  string `bson:",omitempty"`
	SecondLevel string `bson:",omitempty"`
	ThirdLevel  string `bson:",omitempty"`
}

func (g *GeocodedLocation) String() string {
	return g.FullAddress
}
