package opensea

import (
	"context"

	"github.com/casiphia/gopkg/restgo"
	"github.com/casiphia/opensea-go/model"
)

// This endpoint is used to obtain the entire list of owners for an NFT. Results will also include the quantity owned.
func (c *Client) Owners(ctx context.Context, req *OwnersRequest) (*OwnersResponse, error) {
	var rsp, err = c.get(ctx, "/api/v1/asset/:asset_contract_address/:token_id/owners", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response OwnersResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type OwnersRequest struct {
	// Address of the contract for this NFT
	AssetContractAddress string `path:"asset_contract_address,required"`
	// Address of the contract for this NFT
	TokenID string `path:"token_id,required"`
	// How many results to show. The default value is 20 (if param is omitted), max is 50.
	Limit string `query:"limit,required"`
	// What field to order by. For now, we only support created_date.
	OrderBy string `query:"order_by"`
	// asc or desc. Descending is the default.
	OrderDirection string `query:"order_direction"`
	// The cursor token. Used for going forward / backward to the next/previous pages.
	Cursor string `query:"cursor"`
}

type OwnersResponse struct {
	Assets   []*model.Owner `opensea:"owners"`
	Next     string         `opensea:"next"`
	Previous string         `opensea:"previous"`
}
