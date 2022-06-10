package openrtb

import (
	"encoding/json"
)

/*************************************************************************
 * COMMON OBJECT STRUCTS
 *************************************************************************/

// This object may be useful in the situation where syndicated content contains impressions and
// does not necessarily match the publisher's general content.  The exchange might or might not
// have knowledge of the page where the content is running, as a result of the syndication
// method.  (For example, video impressions embedded in an iframe on an unknown web property
// or device.)
// type Content struct {
// }

// ThirdParty abstract attributes.
type ThirdParty struct {
	ID                 string            `json:"id,omitempty"`
	Name               string            `json:"name,omitempty"`
	CategoryTaxonomies uint              `json:"cattax,omitempty"` // The taxonomy in use for bcat. Refer to the AdCOM 1.0 list List: Category Taxonomies for values
	Categories         []ContentCategory `json:"cat,omitempty"`    // Array of IAB content categories
	Domain             string            `json:"domain,omitempty"` // Highest level domain of the seller
	Ext                json.RawMessage   `json:"ext,omitempty"`
}

// Publisher object itself and all of its parameters are optional, so default values are not
// provided. If an optional parameter is not specified, it should be considered unknown.
type Publisher ThirdParty

// Producer is useful when content where the ad is shown is syndicated, and may appear on a
// completely different publisher. The producer object itself and all of its parameters are optional,
// so default values are not provided. If an optional parameter is not specified, it should be
// considered unknown.
type Producer ThirdParty

// Data and segment objects together allow additional data about the user to be specified. This data
// may be from multiple sources whether from the exchange itself or third party providers as specified by
// the id field. A bid request can mix data objects from multiple providers. The specific data providers in
// use should be published by the exchange a priori to its bidders.
type Data struct {
	ID      string          `json:"id,omitempty"`
	Name    string          `json:"name,omitempty"`
	Segment []Segment       `json:"segment,omitempty"`
	Ext     json.RawMessage `json:"ext,omitempty"`
}

// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Value string          `json:"value,omitempty"`
	Ext   json.RawMessage `json:"ext,omitempty"`
}

// This object describes the network an ad will be displayed on. A Network is defined as the parent entity of
// the Channel object’s entity for the purposes of organizing Channels. Examples are companies that own
// and/or license a collection of content channels (Viacom, Discovery, CBS, WarnerMedia, Turner and others),
// or studio that creates such content and self-distributes content. Name is a human-readable field while
// domain and id can be used for reporting and targeting purposes.
type Network struct {
	ID     string          `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Domain string          `json:"domain,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// This object describes the channel an ad will be displayed on. A Channel is defined as the entity that curates
// a content library, or stream within a brand name for viewers. Examples are specific view selectable
// ‘channels’ within linear and streaming television (MTV, HGTV, CNN, BBC One, etc) or a specific stream of
// audio content commonly called ‘stations.’ Name is a human-readable field while domain and id can be used
// for reporting and targeting purposes
type Channel struct {
	ID     string          `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Domain string          `json:"domain,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}
