package openrtb

import (
	"encoding/json"
)

// Format object represents an allowed size (i.e., height and width combination) for a banner impression.
// These are typically used in an array for an impression where multiple sizes are permitted.
type Format struct {
	Width       int             `json:"w,omitempty"`      // Width in device independent pixels (DIPS).
	Height      int             `json:"h,omitempty"`      // Height in device independent pixels (DIPS).
	WidthRatio  int             `json:"wwratioomitempty"` // Relative width when expressing size as a ratio.
	HeightRatio int             `json:"hratio,omitempty"` // Relative height when expressing size as a ratio.
	wmin        int             `json:"wmin,omitempty"`   // The minimum width in device independent pixels (DIPS) at which the ad will be displayed the size is expressed as a ratio.
	Ext         json.RawMessage `json:"ext,omitempty"`
}
