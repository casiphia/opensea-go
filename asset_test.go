package opensea

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func TestClient_Asset(t *testing.T) {
	var cli = newClient()
	var asset, err = cli.Asset(ctx, &AssetRequest{
		AssetContractAddress: "0xf4910c763ed4e47a585e2d34baa9a4b611ae448c",
		TokenID:              "25777430808421588984778576081336814700450940661256628635588471382651290255361",
	})
	if err != nil {
		t.Error(err)
		return
	}
	pp.Default.SetColoringEnabled(false)
	pp.Println(asset)
}
