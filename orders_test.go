package opensea

import "testing"

func TestClient_Orders(t *testing.T) {
	var cli = newClient()
	var _, err = cli.OrdersListings(ctx, &OrdersRequest{
		AssetContractAddress: "0x622062a3d249cccee76864a148c89484a23a6144",
		Limit:                "1",
		OrderBy:              "created_date",
		OrderDirection:       "desc",
	})
	if err != nil {
		t.Error(err)
		return
	}
}
