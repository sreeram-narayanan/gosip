package txs

import (
	"fmt"

	"github.com/discoviking/fsm"
	"github.com/ghettovoice/gosip/core"
	"github.com/ghettovoice/gosip/log"
)

type Transaction interface {
	log.LocalLogger
	Origin() core.Request
	Receive(msg core.Message) error
	Delete() error
	IsInvite() bool
	IsAck() bool
	String() string
}

type transaction struct {
	logger   log.LocalLogger
	fsm      *fsm.FSM
	origin   core.Request
	lastResp core.Response
}

func (tx *transaction) String() string {
	if tx == nil {
		return "Transaction <nil>"
	}

	return fmt.Sprintf("Transaction %p [%s]", tx, tx.Origin().Short())
}

func (tx *transaction) Log() log.Logger {
	return tx.logger.Log()
}

func (tx *transaction) SetLog(logger log.Logger) {
	tx.logger.SetLog(logger.WithFields(map[string]interface{}{
		"tx": tx.String(),
	}))
}

func (tx *transaction) Origin() core.Request {
	return tx.origin
}

func (tx *transaction) IsInvite() bool {
	return tx.Origin().IsInvite()
}

func (tx *transaction) IsAck() bool {
	return tx.Origin().IsAck()
}