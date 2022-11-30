package opensea

import (
	"fmt"
	"testing"

	"github.com/casiphia/opensea-go/model/stream"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
)

func TestClient_NewStreamClient(t *testing.T) {
	client := NewStreamClient(MAINNET, "apikey", phx.LogInfo, func(err error) {
		fmt.Println("opensea.NewStreamClient err:", err)
	})
	if err := client.Connect(); err == nil {
		fmt.Println("client.Connect err:", err)
		return
	}

	client.OnItemListed("collection-slug", func(response any) {
		var itemListedEvent stream.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})

	client.OnItemSold("collection-slug", func(response any) {
		var itemSoldEvent stream.ItemSoldEvent
		err := mapstructure.Decode(response, &itemSoldEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemSoldEvent)
	})

	client.OnItemTransferred("collection-slug", func(response any) {
		var itemTransferredEvent stream.ItemTransferredEvent
		err := mapstructure.Decode(response, &itemTransferredEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemTransferredEvent)
	})

	client.OnItemCancelled("collection-slug", func(response any) {
		var itemCancelledEvent stream.ItemCancelledEvent
		err := mapstructure.Decode(response, &itemCancelledEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemCancelledEvent)
	})

	client.OnItemReceivedBid("collection-slug", func(response any) {
		var itemReceivedBidEvent stream.ItemReceivedBidEvent
		err := mapstructure.Decode(response, &itemReceivedBidEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemReceivedBidEvent)
	})
	client.OnItemReceivedOffer("collection-slug", func(response any) {
		var itemReceivedOfferEvent stream.ItemReceivedOfferEvent
		err := mapstructure.Decode(response, &itemReceivedOfferEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemReceivedOfferEvent)
	})

	client.OnItemMetadataUpdated("collection-slug", func(response any) {
		var itemMetadataUpdateEvent stream.ItemMetadataUpdateEvent
		err := mapstructure.Decode(response, &itemMetadataUpdateEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemMetadataUpdateEvent)
	})

	select {}
}
