package opensea

import (
	"context"

	"github.com/casiphia/gopkg/restgo"
)

// This endpoint is used to check the metadata URL on a given NFT.
func (c *Client) AssetValidate(ctx context.Context, req *AssetValidateRequest) (*AssetValidateResponse, error) {
	var rsp, err = c.get(ctx, "/api/v1/asset/:asset_contract_address/:token_id/validate", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response AssetValidateResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type AssetValidateRequest struct {
	// The NFT contract address for the assets
	AssetContractAddress string `path:"asset_contract_address,required"`
	// Address of the contract for this NFT
	TokenID string `path:"token_id,required"`
	// Address of an owner of the token.
	// If you include this, the response will include an ownership object that includes
	// the number of tokens owned by the address provided instead of the top_ownerships object
	// included in the standard response, which provides the number of tokens owned by each of
	// the 10 addresses with the greatest supply of the token.
	Validate string `path:"validate,required"`
}

type AssetValidateResponse struct {
	Valid    bool        `opensea:"valid" json:"valid"`
	TokenUri string      `opensea:"token_uri" json:"tokenUri"`
	Errors   interface{} `opensea:"errors" json:"errors"`
}
