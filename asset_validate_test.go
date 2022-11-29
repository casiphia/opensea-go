package opensea

import "testing"

func TestClient_AssetValidate(t *testing.T) {
	var cli = newClient()
	var assetV, err = cli.AssetValidate(ctx, &AssetValidateRequest{
		AssetContractAddress: "0xf4910c763ed4e47a585e2d34baa9a4b611ae448c",
		TokenID:              "25777430808421588984778576081336814700450940661256628635588471382651290255361",
		Validate:             "validate",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("respone:[%v]", assetV)
}
