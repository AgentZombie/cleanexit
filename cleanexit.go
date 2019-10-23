package cleanexit

import (
	"os"
	"os/signal"
	"sync"
)

var (
	cleanupFns = []func(){}
	SigInt     = os.Interrupt
	once       = &sync.Once{}
	final      = func() { os.Exit(0) }
)

func Cleanup() {
	once.Do(cleanup)
}

func cleanup() {
	for _, f := range cleanupFns {
		f()
	}
	final()
}

func Register(f func()) {
	cleanupFns = append(cleanupFns, f)
}

func Finally(f func()) {
	final = f
}

func OnSignals(sig ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	go func() {
		<-c
		Cleanup()
	}()
}
