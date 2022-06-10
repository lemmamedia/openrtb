package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidVideoNoMIMEs     = errors.New("openrtb: video has no mimes")
	ErrInvalidVideoNoLinearity = errors.New("openrtb: video linearity missing")
	ErrInvalidVideoNoProtocols = errors.New("openrtb: video protocols missing")
)

// Video object must be included directly in the impression object if the impression offered
// for auction is an in-stream video ad opportunity.
type Video struct {
	MIMEs           []string            `json:"mimes,omitempty"`          // Content MIME types supported.
	MinDuration     int                 `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	MaxDuration     int                 `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	StartDelay      StartDelay          `json:"startdelay,omitempty"`     // Indicates the start delay in seconds
	MaxSeq          int                 `json:"maxseq,omitempty"`         //Indicates the maximum number of ads that may be served into a “dynamic” video ad pod (where the precise number of ads is not predetermined by the seller). See Section 7.6 for more details
	PodDur          int                 `json:"poddur,omitempty"`         //Indicates the total amount of time in seconds that advertisers may fill for a “dynamic” video ad pod
	Protocols       []Protocol          `json:"protocols,omitempty"`      // Video bid response protocols
	Protocol        Protocol            `json:"protocol,omitempty"`       // Video bid response protocols DEPRECATED
	Width           int                 `json:"w,omitempty"`              // Width of the player in pixels
	Height          int                 `json:"h,omitempty"`              // Height of the player in pixels
	PoDid           int                 `json:"podid,omitempty"`          // Unique identifier indicating that an impression opportunity belongs to a video ad pod. If multiple impression opportunities within a bid request share the same podid, this indicates that those impression opportunities belong to the same video ad pod
	PodSeq          int                 `json:"podseq,omitempty"`         // The sequence (position) of the video ad pod within a content     stream
	RqdDurs         int                 `json:"rqddurs,omitempty"`        // Precise acceptable durations for video creatives in seconds. This field specifically targets the Live TV use case where non-exact ad durations would result in undesirable ‘dead air’. This field is mutually exclusive with minduration and maxduration; if rqddurs is specified, minduration and maxduration must not be specified and vice versa
	Placement       VideoPlacement      `json:"placement,omitempty"`      // Video placement type
	Linearity       VideoLinearity      `json:"linearity,omitempty"`      // Indicates whether the ad impression is linear or non-linear
	Skip            int                 `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	SkipMin         int                 `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable
	SkipAfter       int                 `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled
	Sequence        int                 `json:"sequence,omitempty"`       // Default: 1 DEPRECATED
	SlotInPod       int                 `json:"slotinpod,omitempty"`      // For video ad pods, this value indicates that the seller can guarantee delivery against the indicated slot position in the pod.
	MinCPMPerSec    float32             `json:"mincpmpersec,omitempty"`   // Minimum CPM per second. This is a price floor for the “dynamic” portion of a video ad pod, relative to the duration of bids an advertiser may submit.
	BlockedAttrs    []CreativeAttribute `json:"battr,omitempty"`          // Blocked creative attributes
	MaxExtended     int                 `json:"maxextended,omitempty"`    // Maximum extended video ad duration
	MinBitrate      int                 `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps
	MaxBitrate      int                 `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps
	BoxingAllowed   *int                `json:"boxingallowed,omitempty"`  // If exchange publisher has rules preventing letter boxing
	PlaybackMethods []VideoPlayback     `json:"playbackmethod,omitempty"` // List of allowed playback methods
	Delivery        []ContentDelivery   `json:"delivery,omitempty"`       // List of supported delivery methods
	Position        AdPosition          `json:"pos,omitempty"`            // Ad Position
	CompanionAds    []Banner            `json:"companionad,omitempty"`
	APIs            []APIFramework      `json:"api,omitempty"` // List of supported API frameworks
	CompanionTypes  []CompanionType     `json:"companiontype,omitempty"`
	Ext             json.RawMessage     `json:"ext,omitempty"`
}

type jsonVideo Video

// Validate the object
func (v *Video) Validate() error {
	if len(v.MIMEs) == 0 {
		return ErrInvalidVideoNoMIMEs
	} else if v.Linearity == 0 {
		return ErrInvalidVideoNoLinearity
	} else if v.Protocol == 0 && len(v.Protocols) == 0 {
		return ErrInvalidVideoNoProtocols
	}
	return nil
}

// GetBoxingAllowed returns the boxing-allowed indicator
func (v *Video) GetBoxingAllowed() int {
	if v.BoxingAllowed != nil {
		return *v.BoxingAllowed
	}
	return 1
}

// MarshalJSON custom marshalling with normalization
func (v *Video) MarshalJSON() ([]byte, error) {
	v.normalize()
	return json.Marshal((*jsonVideo)(v))
}

// UnmarshalJSON custom unmarshalling with normalization
func (v *Video) UnmarshalJSON(data []byte) error {
	var h jsonVideo
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*v = (Video)(h)
	v.normalize()
	return nil
}

func (v *Video) normalize() {
	if v.Sequence == 0 {
		v.Sequence = 1
	}
	if v.Linearity == 0 {
		v.Linearity = VideoLinearityLinear
	}
}
