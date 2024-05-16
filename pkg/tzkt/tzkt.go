// Package tzkt provides a client for the TzKT API.
package tzkt

import (
	"context"
	"fmt"

	"github.com/dipdup-net/go-lib/tzkt/data"
	"github.com/dipdup-net/go-lib/tzkt/events"
)

const (
	operationTypeDelegation = "delegation"
)

// TzKT is the interface that wraps the basic methods to interact with the TzKT API.
//
//go:generate mockery --name=TzKT --output=mocks --filename=tzkt.go --outpkg=mocks
type TzKT interface {
	Listen() (Message, error)
	Close() error
}

type tzkt struct {
	driver       *events.TzKT
	ctxCanceller context.CancelFunc
}

// New creates a new TzKT client with an initial connection to the TzKT websocket.
// The client will subscribe to the delegation operations only.
// Call the Listen method to read the messages from the TzKT websocket.
func New(ctx context.Context) (TzKT, error) {
	var c Config
	var err error
	var t tzkt

	if c, err = NewConfig(); err != nil {
		return nil, err
	}

	t.driver = events.NewTzKT(c.Host)

	var newCTX context.Context
	newCTX, t.ctxCanceller = context.WithCancel(ctx)

	if err := t.driver.Connect(newCTX); err != nil {
		return nil, err
	}

	if err := t.driver.SubscribeToOperations("", operationTypeDelegation); err != nil {
		return nil, err
	}

	msg := <-t.driver.Listen()
	if msg.Type != events.MessageTypeState && msg.Channel != events.ChannelOperations {
		return nil, fmt.Errorf("invalid initial message type: %v with channel %s", msg.Type, msg.Channel)
	}
	msg = <-t.driver.Listen()
	if msg.Type != events.MessageTypeSubscribed && msg.Channel != events.MethodOperations {
		return nil, fmt.Errorf("invalid subscribe operation : %s", operationTypeDelegation)
	}

	return &t, nil
}

func (t *tzkt) Listen() (Message, error) {
	var response Message

	msg := <-t.driver.Listen()
	if msg.Channel != events.ChannelOperations {
		return Message{}, fmt.Errorf("invalid channel listenner: %s", msg.Channel)
	}

	response.Blocks = msg.State
	if msg.Type == events.MessageTypeData {
		response.Type = MessageTypeData

		list := msg.Body.([]any)
		delegates := make([]data.Delegation, len(list))
		for i, v := range list {
			delegates[i] = *(v.(*data.Delegation))
		}
		response.Delegations = delegates
	} else if msg.Type == events.MessageTypeReorg {
		response.Type = MessageTypeReorg
	}

	return response, nil
}

func (t *tzkt) Close() error {
	t.ctxCanceller()

	return t.driver.Close()
}
