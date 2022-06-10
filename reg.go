package openrtb

import (
	"encoding/json"
)

// Regulations object contains any legal, governmental, or industry regulations that apply to the request. The
// coppa flag signals whether or not the request falls under the United States Federal Trade Commission's
// regulations for the United States Children's Online Privacy Protection Act ("COPPA").
type Regulations struct {
	COPPA     int             `json:"coppa,omitempty"`      // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	GDPR      int             `json:"gdpr,omitempty"`       // Flag that indicates whether or not the request is subject to GDPR regulations 0 = No, 1 = Yes, omission indicates Unknown. Refer to Section 7.5 for more information
	UsPrivacy string          `json:"us_privacy,omitempty"` // Communicates signals regarding consumer privacy under US privacy regulation. See US Privacy String specifications. Refer to Section 7.5 for more information
	Ext       json.RawMessage `json:"ext,omitempty"`
}
