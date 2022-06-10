package openrtb

import (
	"encoding/json"
)

// Extended identifiers support in the OpenRTB specification allows buyers to use audience data in real-time
// bidding. This object can contain one or more UIDs from a single source or a technology provider. The
// exchange should ensure that business agreements allow for the sending of this data.
type EID struct {
	Source string          `json:"source,omitempty"` // Source or technology provider responsible for the set of included IDs. Expressed as a top-level domain.
	UIDs   []UID           `json:"uids,omitempty"`
	Ext    json.RawMessage `json:"ext,omitempty"`
}

// This object contains a single user identifier provided as part of extended identifiers. The exchange should
// ensure that business agreements allow for the sending of this data
type UID struct {
	Id     string          `json:"id,omitempty"`    // The identifier for the user.
	AtType int             `json:"atype,omitempty"` // Type of user agent the ID is from. It is highly recommended to set this, as many DSPs separate app-native IDs from browser-based IDs and require a type value for ID resolution.
	Ext    json.RawMessage `json:"ext,omitempty"`
}
