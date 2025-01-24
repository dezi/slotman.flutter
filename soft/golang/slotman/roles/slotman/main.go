package main

import (
	"slotman/services/impl/pilots"
	"slotman/services/impl/provider"
	"slotman/services/impl/proxy"
	"slotman/services/impl/turner"
	"slotman/utils/daemon"
	"slotman/utils/exitter"
	"slotman/utils/log"
)

func main() {
	daemon.Daemonize(startup)
}

func startup() {

	log.SetCallerLength(48)

	_ = provider.StartService()

	_ = proxy.StartService()
	_ = pilots.StartService()
	_ = turner.StartService()

	_ = exitter.StartService()

	_ = pilots.StopService()
	_ = turner.StopService()
	_ = proxy.StopService()

	_ = provider.StopService()
}
