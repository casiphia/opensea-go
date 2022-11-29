package model

type Protocol struct {
	Parameters *Parameters `opensea:"parameters" json:"parameters"`
	Signature  string      `opensea:"signature" json:"signature"`
}

type Parameters struct {
	Offerer                         string           `opensea:"offerer" json:"offerer"`
	Offer                           *Offer           `opensea:"offer" json:"offer"`
	Consideration                   []*Consideration `opensea:"consideration" json:"consideration"`
	StartTime                       string           `opensea:"startTime" json:"startTime"`
	EndTime                         string           `opensea:"endTime" json:"endTime"`
	OrderType                       string           `opensea:"orderType" json:"orderType"`
	Zone                            string           `opensea:"zone" json:"zone"`
	ZoneHash                        string           `opensea:"zoneHash" json:"zoneHash"`
	Salt                            string           `opensea:"salt" json:"salt"`
	ConduitKey                      string           `opensea:"conduitKey" json:"conduitKey"`
	TotalOriginalConsiderationItems int              `opensea:"totalOriginalConsiderationItems" json:"totalOriginalConsiderationItems"`
	Counter                         int              `opensea:"counter" json:"counter"`
}

type Offer struct {
	ItemType             int    `opensea:"itemType" json:"itemType"`
	Token                string `opensea:"token" json:"token"`
	IdentifierOrCriteria string `opensea:"identifierOrCriteria" json:"identifierOrCriteria"`
	StartAmount          string `opensea:"startAmount" json:"startAmount"`
	EndAmount            string `opensea:"endAmount" json:"endAmount"`
}

type Consideration struct {
	ItemType             int    `opensea:"itemType" json:"itemType"`
	Token                string `opensea:"token" json:"token"`
	IdentifierOrCriteria string `opensea:"identifierOrCriteria" json:"identifierOrCriteria"`
	StartAmount          string `opensea:"startAmount" json:"startAmount"`
	EndAmount            string `opensea:"endAmount" json:"endAmount"`
	Recipient            string `opensea:"recipient" json:"recipient"`
}
