package tacho

import (
	"slotman/utils/log"
	"time"
)

func (sv *Service) GetTachosAttached() (tracks []bool) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	tracks = make([]bool, 8)

	for track := 0; track < len(tracks); track++ {
		pin := track * 2
		if sv.tachoStates[pin].time.Unix() > 0 {
			tracks[track] = true
		}
	}

	return
}

func (sv *Service) OnRaceStarted() {

	sv.mapsLock.Lock()

	log.Printf("OnRaceStarted...")

	for track := range sv.trackStates {
		trackState := sv.trackStates[track]
		trackState.RoundTs = time.Now()
		sv.trackStates[track] = trackState
	}

	sv.mapsLock.Unlock()
}
