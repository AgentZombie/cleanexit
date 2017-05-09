package cleanexit

import (
	"os"
	"os/signal"
)

var (
	cleanupFns = []func(){}
	SigInt     = os.Interrupt
)

func Cleanup() {
	for _, f := range cleanupFns {
		f()
	}
}

func Register(f func()) {
	cleanupFns = append(cleanupFns, f)
}

func OnSignals(sig ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	go func() {
		<-c
		Cleanup()
	}()
}
