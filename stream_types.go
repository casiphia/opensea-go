/*
 * @Descripttion: X-club project
 * @Version: 1.0.0
 * @Author: wangxiaodiao
 * @Email: 413586280@qq.com
 * @Date: 2023-02-07 15:28:29
 * @LastEditors: wangxiaodiao
 * @LastEditTime: 2023-02-07 15:28:38
 * @FilePath: /opensea-go/stream_types.go
 */
package opensea

type EventType string

const (
	ItemMetadataUpdated EventType = "item_metadata_updated"
	ItemListed          EventType = "item_listed"
	ItemSold            EventType = "item_sold"
	ItemTransferred     EventType = "item_transferred"
	ItemReceivedOffer   EventType = "item_received_offer"
	ItemReceivedBid     EventType = "item_received_bid"
	ItemCancelled       EventType = "item_cancelled"
)

type Network string

const (
	MAINNET Network = "mainnet"
	TESTNET Network = "testnet"
)
