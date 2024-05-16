package tzkt

import "github.com/dipdup-net/go-lib/tzkt/data"

const (
	MessageTypeData MessageType = iota
	MessageTypeReorg
)

type Message struct {
	Blocks      uint64
	Type        MessageType
	Delegations []data.Delegation
}

type MessageType uint8
