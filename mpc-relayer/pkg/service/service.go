package service

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/sirupsen/logrus"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/common"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/rpc"
)

type Service struct {
	keyringID uint64
	modules   []Module
	server    rpc.HTTPService
	log       *logrus.Entry
	keyDB     database.Database

	stop    chan struct{}
	stopped atomic.Bool
}

func New(keyRingID uint64, port int, logger *logrus.Entry, keyDB database.Database, modules ...Module) *Service {
	s := &Service{
		keyringID: keyRingID,
		log:       logger,
		keyDB:     keyDB,
		modules:   modules,
		stop:      make(chan struct{}, 1),
		stopped:   atomic.Bool{},
	}
	s.server = rpc.NewHTTPService(port, makeAPIHandlers(s), logger)
	return s
}

func (s *Service) Start() error {
	s.log.WithFields(logrus.Fields{
		"version":   common.FullVersion,
		"buildDate": common.Date,
	}).Info("starting mpc-relayer service")

	var errStr string
	for i, module := range s.modules {
		if err := module.Start(); err != nil {
			s.log.WithError(err).Error("cannot start module")
			errStr += fmt.Sprintf("%v : %v - ", i, err.Error())
		}
	}
	if errStr != "" {
		return fmt.Errorf("%v", errStr)
	}
	return nil
}

func (s *Service) Stop(sig os.Signal) error {
	s.log.WithFields(logrus.Fields{"signal": sig}).Info("received shutdown signal")
	if s.stopped.Load() {
		s.log.WithFields(logrus.Fields{"signal": sig}).Warn("already shutting down")
		return fmt.Errorf("already shutting down")
	}
	s.stopped.Store(true)
	close(s.stop)
	var errStr string
	for i, module := range s.modules {
		if err := module.Stop(); err != nil {
			s.log.WithError(err).Error("cannot stop module")
			errStr += fmt.Sprintf("%v : %v - ", i, err.Error())
		}
	}
	if errStr != "" {
		return fmt.Errorf("%v", errStr)
	}
	s.log.Info("mpc-relayer stopped")
	return nil
}
