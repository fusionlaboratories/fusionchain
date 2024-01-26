package bridge

import (
	"context"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/qredo/assets/libs/logger"
	"github.com/qredo/assets/libs/protobuffer"
	"github.com/qredo/assets/libs/watcher/qredochain"
	"github.com/sirupsen/logrus"
)

var log = logger.Empty

type retryableQueueItem struct {
	*qredochain.EventQueueItem
	tries int64
	err   error
}

type fusionClient interface {
	Send(ctx context.Context, dest string, amount *big.Int) error
}

type Bridge struct {
	retrySleep time.Duration

	qredochain qredochain.KeySearcher

	waiting chan *retryableQueueItem

	retriever Retriever
	fusion    fusionClient // todo replace faucet use by directly communicating with fusion chain
	stopped   atomic.Bool

	maxTries int64

	stop chan struct{}
}

func NewBridge(qredochain qredochain.KeySearcher, retriever Retriever, fusion fusionClient) *Bridge {
	return &Bridge{
		retrySleep: 10,
		qredochain: qredochain,
		retriever:  retriever,
		fusion:     fusion,
		maxTries:   5,
	}
}

func (b *Bridge) Start() error {
	b.waiting = make(chan *retryableQueueItem)
	b.stop = make(chan struct{})

	out, err := b.retriever.Start()
	if err != nil {
		return err
	}

	go b.process()
	go b.receive(out)

	return nil
}

func (b *Bridge) receive(out <-chan *qredochain.EventQueueItem) {
	for {
		select {
		case <-b.stop:
			return
		case item := <-out:
			b.waiting <- &retryableQueueItem{
				EventQueueItem: item,
			}
		}
	}
}

func (b *Bridge) bridge(event *protobuffer.PBWalletHistoryData) error {
	address, err := qredochain.FusionAddress(b.qredochain, event.AssetID)
	if err != nil {
		return err
	}

	if err := b.fusion.Send(context.Background(), address, big.NewInt(event.AmountReceived)); err != nil {
		return err
	}
	return nil
}

func (b *Bridge) process() {
	for {
		select {
		case <-b.stop:
			return
		case item := <-b.waiting:
			log := log.WithFields(logrus.Fields{
				"assetID": fmt.Sprintf("%x", item.Event.GetAssetID()),
				"tries":   item.tries,
				"sent":    item.Event.AmountSent,
				"receive": item.Event.AmountReceived,
			})
			if item.tries > b.maxTries {
				log.WithFields(logrus.Fields{
					"error": item.err.Error(),
				}).Error("" +
					"bridge error")
				continue
			}
			go func(item *retryableQueueItem) {
				if item.tries > 1 {
					time.Sleep(b.retrySleep)
				}
				item.tries++
				if err := b.bridge(item.Event); err != nil {
					item.err = err
					log.WithFields(logrus.Fields{
						"error": item.err.Error(),
					}).Warn("bridge failed")
					if b.stopped.Load() {
						return
					}
					b.waiting <- item
					return
				}
				log.Info("bridge transfer")
			}(item)
		}
	}
}

func (b *Bridge) Stop() error {
	if b.stopped.Load() {
		return nil
	}
	b.stopped.Store(true)
	b.retriever.Stop()
	close(b.stop)
	close(b.waiting)
	return nil
}
