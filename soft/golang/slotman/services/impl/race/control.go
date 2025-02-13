package race

import (
	"slotman/services/iface/ampel"
	"slotman/services/iface/pilots"
	"slotman/services/iface/server"
	"slotman/services/iface/speedi"
	"slotman/services/iface/speedo"
	"slotman/services/iface/tacho"
	"slotman/services/iface/teams"
	"slotman/services/type/slotman"
	"time"
)

func (sv *Service) DoControlTask() {
	sv.checkServices()
	sv.checkLooper()

	//log.Printf("Race raceState=%s", sv.raceState)
}

func (sv *Service) checkServices() {

	if sv.servicesReady {
		return
	}

	//
	// Yield time for target services to register.
	//

	time.Sleep(time.Second)

	var tryErr error

	sv.srv, tryErr = server.GetInstance()
	if tryErr != nil {
		return
	}

	sv.amp, tryErr = ampel.GetInstance()
	if tryErr != nil {
		return
	}

	sv.sdi, tryErr = speedi.GetInstance()
	if tryErr != nil {
		return
	}

	sv.sdo, tryErr = speedo.GetInstance()
	if tryErr != nil {
		return
	}

	sv.tco, tryErr = tacho.GetInstance()
	if tryErr != nil {
		return
	}

	sv.plt, tryErr = pilots.GetInstance()
	if tryErr != nil {
		return
	}

	sv.tms, tryErr = teams.GetInstance()
	if tryErr != nil {
		return
	}

	sv.servicesReady = true

	//
	// Test fake race start.
	//

	sv.OnAmpelClickLong()

	sv.trackStates[0] = slotman.TrackStateReady
	sv.trackStates[1] = slotman.TrackStateReady
}

func (sv *Service) checkLooper() {

	if !sv.servicesReady || sv.looperStarted {
		return
	}

	sv.looperStarted = true

	go sv.looper()
}
