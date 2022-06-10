package openrtb

import (
	"encoding/json"
)

// This object is composed of a set of nodes where each node represents a specific entity that participates in
// the transacting of inventory. The entire chain of nodes from beginning to end represents all entities who
// are involved in the direct flow of payment for inventory
type SupplyChain struct {
	Complete int               `json:"complete,omitempty"`
	Node     []SupplyChainNode `json:"nodes,omitempty"`
	Version  string            `json:"ver,omitempty"`
	Ext      json.RawMessage   `json:"ext,omitempty"`
}

// This object is associated with a SupplyChain object as an array of nodes. These nodes define the identity of
// an entity participating in the supply chain of a bid request
type SupplyChainNode struct {
	ASI       string          `json:"asi,omitempty"`    // The canonical domain name of the SSP, Exchange, Header Wrapper, etc system that bidders connect to. This may be the operational domain of the system, if that is different than the parent corporate domain, to facilitate WHOIS and reverse IP lookups to establish clear ownership of the delegate system. This should be the same value as used to identify sellers in an ads.txt file if one exists
	SID       string          `json:"sid,omitempty"`    // The identifier associated with the seller or reseller account within the advertising system. This must contain the same value used in transactions (i.e. OpenRTB bid requests) in the field specified by the SSP/exchange. Typically, in OpenRTB, this is publisher.id. For OpenDirect it is typically the publisher’s organization ID.Should be limited to 64 characters in length.
	RequestId string          `json:"rid,omitempty"`    // The OpenRTB RequestId of the request as issued by this seller
	Name      string          `json:"name,omitempty"`   // The name of the company (the legal entity) that is paid for inventory transacted under the given seller_ID. This value is optional and should NOT be included if it exists in the advertising system’s sellers.json file
	Domain    string          `json:"domain,omitempty"` // The business domain name of the entity represented by this node. This value is optional and should NOT be included if it exists in the advertising system’s sellers.json file
	HP        int             `json:"hp,omitempty"`     // Indicates whether this node will be involved in the flow of payment for the inventory. When set to 1, the advertising system in the asi field pays the seller in the sid field, who is responsible for paying the previous node in the chain. When set to 0, this node is not involved in the flow of payment for the inventory.
	Ext       json.RawMessage `json:"ext,omitempty"`
}
