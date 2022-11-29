package model

import "github.com/shopspring/decimal"

type Side uint8

const (
	SideBuy Side = iota
	SideSell
)

// SaleKind the kind of sell order
// use only_english=true for filtering for only English Auctions
type SaleKind uint8

const (
	// SaleKindFixedOrMinBidPrice fixed-price sales or min-bid auctions
	SaleKindFixedOrMinBidPrice SaleKind = iota
	// SaleKindDecliningPrice declining-price Dutch Auctions
	SaleKindDecliningPrice
)

type HowToCall uint8

const (
	Call HowToCall = iota
	DelegateCall
)

type FeeMethod uint8

const (
	ProtocolFee FeeMethod = iota
	SplitFee
)

type AccountFees struct {
	Account     *Account `opensea:"account" json:"account"`
	BasisPoints string   `opensea:"basis_points" json:"basisPoints"`
}

type MakerAssetBundle struct {
	Assets []*Asset `opensea:"assets" json:"assets"`
}

type Order struct {
	CreatedDate      string           `opensea:"created_date" json:"createdDate"`
	ClosingDate      string           `opensea:"closing_date" json:"closingDate"`
	ListingTime      int              `opensea:"listing_time" json:"listingTime"`
	ExpirationTime   int              `opensea:"expiration_time" json:"expirationTime"`
	OrderHash        string           `opensea:"order_hash" json:"orderHash"`
	ProtocolData     *Protocol        `opensea:"protocol_data" json:"protocolData"`
	ProtocolAddress  string           `opensea:"protocol_address" json:"protocolAddress"`
	Maker            *Account         `opensea:"maker" json:"maker"`
	Taker            *Account         `opensea:"taker" json:"taker"`
	CurrentPrice     *decimal.Decimal `opensea:"current_price" json:"currentPrice"`
	MakerFees        []*AccountFees   `opensea:"maker_fees" json:"makerFees"`
	TakerFees        []*AccountFees   `opensea:"taker_fees" json:"takerFees"`
	Side             Side             `opensea:"side" json:"side"`
	OrderType        string           `opensea:"order_type" json:"orderType"`
	Cancelled        bool             `opensea:"cancelled" json:"cancelled"`
	Finalized        bool             `opensea:"finalized" json:"finalized"`
	MarkedInvalid    bool             `opensea:"marked_invalid" json:"markedInvalid"`
	ClientSignature  string           `opensea:"client_signature" json:"clientSignature"`
	RelayId          string           `opensea:"relay_id" json:"relayId"`
	CriteriaProof    *string          `opensea:"criteria_proof" json:"criteriaProof"`
	MakerAssetBundle []*Bundle        `opensea:"maker_asset_bundle" json:"makerAssetBundle"`
	TakerAssetBundle []*Bundle        `opensea:"taker_asset_bundle" json:"takerAssetBundle"`
}

func (o *Order) Hash() string {
	if o == nil {
		return ""
	}

	return o.OrderHash
}

type Metadata struct {
	Asset           *AssetAddress `opensea:"asset" json:"asset"`
	Schema          string        `opensea:"schema" json:"schema"`
	ReferrerAddress string        `opensea:"referrerAddress" json:"referrerAddress"`
}
