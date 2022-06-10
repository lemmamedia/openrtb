package openrtb

import (
	"encoding/json"
)

// Structured user agent information, which can be used when a client supports User-Agent Client Hints. If
// both device.ua and device.sua are present in the bid request, device.sua should be considered the more
// accurate representation of the device attributes. This is because the device.ua may contain a frozen or
// reduced user agent string due to deprecation of user agent strings by browsers
type UserAgent struct {
	Browsers     []BrandVersion  `json:"browsers,omitempty"`     // Each BrandVersion object (see Section 3.2.30) identifies a browser or similar software component. Implementers should send brands and versions derived from the Sec-CH-UA-Full-Version-List header
	PMPlatform   []BrandVersion  `json:"platform,omitempty"`     // A BrandVersion object (see Section 3.2.30) that identifies the user agent’s execution platform / OS. Implementers should send a brand derived from the Sec-CH-UA-Platform header, and version derived from the Sec-CH-UAPlatform-Version header *
	Mobile       int             `json:"mobile,omitempty"`       // 1 if the agent prefers a “mobile” version of the content, if available, i.e. optimized for small screens or touch input. 0 if the agent prefers the “desktop” or “full” content. Implementers should derive this value from the Sec-CH-UAMobile header *.
	Architecture string          `json:"architecture,omitempty"` // Device’s major binary architecture, e.g. “x86” or “arm”. Implementers should retrieve this value from the Sec-CH-UA-Arch header*.
	Bitness      string          `json:"bitness,omitempty"`      // Device’s bitness, e.g. “64” for 64-bit architecture. Implementers should retrieve this value from the Sec-CH-UA-Bitness header*.
	Model        string          `json:"model,omitempty"`        // Device model. Implementers should retrieve this value from the Sec-CH-UAModel header*.
	Source       int             `json:"source,omitempty"`       // The source of data used to create this object, List: User-Agent Source in AdCOM 1.0
	Ext          json.RawMessage `json:"ext,omitempty"`
}

// Further identification based on User-Agent Client Hints, the BrandVersion object is used to identify a
// device’s browser or similar software component, and the user agent’s execution platform or operating
// system
type BrandVersion struct {
	Brand  int             `json:"brand,omitempty"`   // A brand identifier, for example, “Chrome” or “Windows”. The value may be sourced from the User-Agent Client Hints headers, representing either the user agent brand (from the Sec-CH-UA-Full-Version header) or the platform brand (from the Sec-CH-UA-Platform header).
	Source []string        `json:"version,omitempty"` // A sequence of version components, in descending hierarchical order (major, minor, micro, …)
	Ext    json.RawMessage `json:"ext,omitempty"`
}
