package opensea

import "testing"

func TestClient_Owners(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Owners(ctx, &OwnersRequest{
		AssetContractAddress: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb",
		TokenID:              "3",
		Limit:                "10",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("rsponse:[%v]", list)
}
