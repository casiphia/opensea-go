package opensea

import (
	"context"

	"github.com/casiphia/gopkg/restgo"
	"github.com/casiphia/opensea-go/model"
)

// This endpoint is used to fetch the set of active listings on a given NFT for the Seaport contract.
func (c *Client) OrdersCreateListings(ctx context.Context, req *OrderCreateResponse) (*OrderResponse, error) {
	var rsp, err = c.post(ctx, "/v2/orders/{chain}/seaport/listings", restgo.ObjectParams(req)...)
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

type OrdersCreateRequest struct {
	// Specify the chain you would like to retrieve listings on. The options are "arbitrum", "avalanche", "ethereum", "klaytn", "matic" and "optimism".
	Chain string `path:"chain,required"`
	// Represents the details of the listing to be posted to Seaport. Refer to Orders Parameter Model for more details.
	OrderParameters *model.Protocol `form:"order_parameters,required"`
	Signature       string          `form:"signature,required"`
}

type OrderCreateResponse struct {
	Orders   []*model.Order `opensea:"owners"`
	Next     string         `opensea:"next"`
	Previous string         `opensea:"previous"`
}

// This endpoint is used to fetch the set of active offers on a given NFT for the Seaport contract.
func (c *Client) OrdersCreateOffers(ctx context.Context, req *OrderCreateResponse) (*OrderCreateResponse, error) {
	var rsp, err = c.post(ctx, "/v2/orders/{chain}/seaport/offers", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response OrderCreateResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
