package opensea

import (
	"context"

	"github.com/casiphia/gopkg/restgo"
	"github.com/casiphia/opensea-go/model"
)

// This endpoint is used to fetch the set of active listings on a given NFT for the Seaport contract.
func (c *Client) OrdersListings(ctx context.Context, req *OrdersRequest) (*OrderResponse, error) {
	var rsp, err = c.get(ctx, "/v2/orders/{chain}/seaport/listings", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response OrderResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type OrdersRequest struct {
	// Specify the chain you would like to retrieve listings on. The options are "arbitrum", "avalanche", "ethereum", "klaytn", "matic" and "optimism".
	Chain string `path:"chain,required"`
	// Address of the contract for an NFT
	AssetContractAddress string `query:"asset_contract_address"`
	// Number of listings to retrieve
	Limit string `query:"limit,required"`
	// Filter by a list of token IDs for the order's asset.
	// Needs to be defined together with asset_contract_address.
	TokenIDs []string `query:"token_ids"`
	// Filter by the order maker's wallet address
	Maker string `query:"maker"`
	// Filter by the order taker's wallet address.
	// Orders open for any taker have the null address as their taker.
	Taker string `query:"taker"`
	// How to sort the orders. Can be created_date for when they were made,
	// or eth_price to see the lowest-priced orders first (converted to their ETH values).
	// eth_price is only supported when asset_contract_address and token_id are also defined.
	OrderBy string `query:"order_by"`
	// Can be asc or desc for ascending or descending sort.
	// For example, to see the cheapest orders, do order_direction asc and order_by eth_price.
	OrderDirection string `query:"order_direction,required"`
	// Only show orders listed after this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	ListedAfter string `query:"listed_after"`
	// Only show orders listed before this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	ListedBefore string `query:"listed_before"`
}

type OrderResponse struct {
	Orders   []*model.Order `opensea:"owners"`
	Next     string         `opensea:"next"`
	Previous string         `opensea:"previous"`
}

// This endpoint is used to fetch the set of active offers on a given NFT for the Seaport contract.
func (c *Client) OrdersOffers(ctx context.Context, req *OrdersRequest) (*OrderResponse, error) {
	var rsp, err = c.get(ctx, "/v2/orders/{chain}/seaport/offers", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response OrderResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
