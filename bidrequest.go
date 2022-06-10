package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidReqNoID     = errors.New("openrtb: request ID missing")
	ErrInvalidReqNoImps   = errors.New("openrtb: request has no impressions")
	ErrInvalidReqMultiInv = errors.New("openrtb: request has multiple inventory sources") // has site and app
)

// BidRequest is the top-level bid request object contains a globally unique bid request or auction ID.  This "id"
// attribute is required as is at least one "imp" (i.e., impression) object.  Other attributes are
// optional since an exchange may establish default values.
type BidRequest struct {
	ID                 string            `json:"id"` // Unique ID of the bid request
	Impressions        []Impression      `json:"imp,omitempty"`
	Site               *Site             `json:"site,omitempty"`
	App                *App              `json:"app,omitempty"`
	Device             *Device           `json:"device,omitempty"`
	User               *User             `json:"user,omitempty"`
	Test               int               `json:"test,omitempty"`    // Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode
	AuctionType        int               `json:"at"`                // Auction type, where 1 = First Price, 2 = Second Price Plus. Exchange-specific auction types can be defined using values greater than 500.
	TMax               int               `json:"tmax,omitempty"`    // Maximum amount of time in milliseconds to submit a bid
	Seats              []string          `json:"wseat,omitempty"`   // Array of buyer seats allowed to bid on this auction
	BlockedSeats       []string          `json:"bseat,omitempty"`   // Array of buyer seats blocked to bid on this auction
	AllImpressions     int               `json:"allimps,omitempty"` // Flag to indicate whether exchange can verify that all impressions offered represent all of the impressions available in context, Default: 0
	Currencies         []string          `json:"cur,omitempty"`     // Array of allowed currencies
	Languages          []string          `json:"wlang,omitempty"`   // Array of languages for creatives using ISO-639-1-alpha-2
	BlockedCategories  []ContentCategory `json:"bcat,omitempty"`    // Blocked Advertiser Categories.
	CategoryTaxonomies int               `json:"cattax,omitempty"`  // The taxonomy in use for bcat. Refer to the AdCOM 1.0 list List: Category Taxonomies for values
	BlockedAdvDomains  []string          `json:"badv,omitempty"`    // Array of strings of blocked toplevel domains of advertisers
	BlockedApps        []string          `json:"bapp,omitempty"`    // Block list of applications by their platform-specific exchange-independent application identifiers. On Android, these should be bundle or package names (e.g., com.foo.mygame).  On iOS, these are numeric IDs.
	Source             *Source           `json:"source,omitempty"`  // A Source object that provides data about the inventory source and which entity makes the final decision
	Regulations        *Regulations      `json:"regs,omitempty"`
	Ext                json.RawMessage   `json:"ext,omitempty"`
}

// Validate the request
func (req *BidRequest) Validate() error {
	if req.ID == "" {
		return ErrInvalidReqNoID
	} else if 0 == len(req.Impressions) {
		return ErrInvalidReqNoImps
	} else if nil != req.Site && nil != req.App {
		return ErrInvalidReqMultiInv
	}

	for i := range req.Impressions {
		imp := req.Impressions[i]
		if err := (&imp).Validate(); err != nil {
			return err
		}
	}

	return nil
}
