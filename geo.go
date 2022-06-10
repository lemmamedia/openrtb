package openrtb

import (
	"encoding/json"
)

// Geo object may appear in one or both the Device Object and the User Object.
// This is intentional, since the information may be derived from either a device-oriented source
// (such as IP geo lookup), or by user registration information (for example provided to a publisher
// through a user registration).
type Geo struct {
	Latitude      float64         `json:"lat,omitempty"`           // Latitude from -90 to 90
	Longitude     float64         `json:"lon,omitempty"`           // Longitude from -180 to 180
	Type          LocationType    `json:"type,omitempty"`          // Indicate the source of the geo data
	Accuracy      int             `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int             `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     IPLocation      `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string          `json:"country,omitempty"`       // Country using ISO 3166-1 Alpha 3
	Region        string          `json:"region,omitempty"`        // Region using ISO 3166-2
	RegionFIPS104 string          `json:"regionFIPS104,omitempty"` // Region of a country using FIPS 10-4
	Metro         string          `json:"metro,omitempty"`         // Google metro code; similar to but not exactly Nielsen DMAs
	City          string          `json:"city,omitempty"`          // City using United Nations Code for Trade & Transport Locations
	ZIP           string          `json:"zip,omitempty"`
	UTCOffset     int             `json:"utcoffset,omitempty"` // Local time as the number +/- of minutes from UTC
	Ext           json.RawMessage `json:"ext,omitempty"`
}
