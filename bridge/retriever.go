package bridge

import (
	"github.com/qredo/assets/libs/assets"
	"github.com/qredo/assets/libs/protobuffer"
	"github.com/qredo/assets/libs/watcher/qredochain"
	"github.com/tendermint/tendermint/rpc/client"
)

type Retriever interface {
	Start() (chan *qredochain.EventQueueItem, error)
	Stop()
}

type QredochainRetriever struct {
	sub       *qredochain.EventSubscription
	connector client.EventsClient
	out       chan *qredochain.EventQueueItem
}

func filterBridgeTransfer() func(item *qredochain.EventQueueItem) (bool, error) {
	return func(item *qredochain.EventQueueItem) (bool, error) {
		return item.Event.HistoryType != protobuffer.PBWalletHistoryDataType_HistoryBridgeTransfer, nil
	}
}

func NewRetriever(connector client.EventsClient) *QredochainRetriever {
	return &QredochainRetriever{
		connector: connector,
		out:       make(chan *qredochain.EventQueueItem),
	}
}

func (q *QredochainRetriever) Start() (chan *qredochain.EventQueueItem, error) {
	sub, err := qredochain.NewEventSubscription(q.connector, "bridge", q.out, assets.SUB_BRIDGE_TRANSFER, filterBridgeTransfer())
	if err != nil {
		return nil, err
	}
	q.sub = sub
	return q.out, nil
}

func (q *QredochainRetriever) Stop() {
	q.sub.Close()
}
