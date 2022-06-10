package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidAudioNoMIMEs = errors.New("openrtb: audio has no mimes")
)

// Audio object must be included directly in the impression object
type Audio struct {
	MIMEs          []string            `json:"mimes"`                  // Content MIME types supported.
	MinDuration    int                 `json:"minduration,omitempty"`  // Minimum video ad duration in seconds
	MaxDuration    int                 `json:"maxduration,omitempty"`  // Maximum video ad duration in seconds
	PodDur         int                 `json:"poddur,omitempty"`       //Indicates the total amount of time in seconds that advertisers may fill for a “dynamic” video ad pod
	Protocols      []Protocol          `json:"protocols,omitempty"`    // Video bid response protocols
	StartDelay     StartDelay          `json:"startdelay,omitempty"`   // Indicates the start delay in seconds
	PoDid          int                 `json:"podid,omitempty"`        // Unique identifier indicating that an impression opportunity belongs to a video ad pod. If multiple impression opportunities within a bid request share the same podid, this indicates that those impression opportunities belong to the same video ad pod
	PodSeq         int                 `json:"podseq,omitempty"`       // The sequence (position) of the video ad pod within a content     stream
	RqdDurs        int                 `json:"rqddurs,omitempty"`      // Precise acceptable durations for video creatives in seconds. This field specifically targets the Live TV use case where non-exact ad durations would result in undesirable ‘dead air’. This field is mutually exclusive with minduration and maxduration; if rqddurs is specified, minduration and maxduration must not be specified and vice versa
	Sequence       int                 `json:"sequence,omitempty"`     // Default: 1
	SlotInPod      int                 `json:"slotinpod,omitempty"`    // For video ad pods, this value indicates that the seller can guarantee delivery against the indicated slot position in the pod.
	MinCPMPerSec   float32             `json:"mincpmpersec,omitempty"` // Minimum CPM per second. This is a price floor for the “dynamic” portion of a video ad pod, relative to the duration of bids an advertiser may submit.
	BlockedAttrs   []CreativeAttribute `json:"battr,omitempty"`        // Blocked creative attributes
	MaxExtended    int                 `json:"maxextended,omitempty"`  // Maximum extended video ad duration
	MinBitrate     int                 `json:"minbitrate,omitempty"`   // Minimum bit rate in Kbps
	MaxBitrate     int                 `json:"maxbitrate,omitempty"`   // Maximum bit rate in Kbps
	Delivery       []ContentDelivery   `json:"delivery,omitempty"`     // List of supported delivery methods
	CompanionAds   []Banner            `json:"companionad,omitempty"`
	APIs           []APIFramework      `json:"api,omitempty"`
	CompanionTypes []CompanionType     `json:"companiontype,omitempty"`
	MaxSequence    int                 `json:"maxseq,omitempty"`   // The maximumnumber of ads that canbe played in an ad pod.
	Feed           FeedType            `json:"feed,omitempty"`     // Type of audio feed.
	Stitched       int                 `json:"stitched,omitempty"` // Indicates if the ad is stitched with audio content or delivered independently
	VolumeNorm     VolumeNorm          `json:"nvol,omitempty"`     // Volume normalization mode.
	Ext            json.RawMessage     `json:"ext,omitempty"`
}

type jsonAudio Audio

// Validate the object
func (a *Audio) Validate() error {
	if len(a.MIMEs) == 0 {
		return ErrInvalidAudioNoMIMEs
	}
	return nil
}

// MarshalJSON custom marshalling with normalization
func (a *Audio) MarshalJSON() ([]byte, error) {
	a.normalize()
	return json.Marshal((*jsonAudio)(a))
}

// UnmarshalJSON custom unmarshalling with normalization
func (a *Audio) UnmarshalJSON(data []byte) error {
	var h jsonAudio
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*a = (Audio)(h)
	a.normalize()
	return nil
}

func (a *Audio) normalize() {
	if a.Sequence == 0 {
		a.Sequence = 1
	}
}
