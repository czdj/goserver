package iface

import (
	"time"
)

//负责网络相关功能的处理
type ITransport interface {
	Listen() error
	Connect() error
	OnNewConnect(transport ITransport)
	Read()
	Write()
	Send(msg interface{})error
	Close()
	OnClosed()
	IsStop() bool
	IsTimeout(tick *time.Timer) bool
}