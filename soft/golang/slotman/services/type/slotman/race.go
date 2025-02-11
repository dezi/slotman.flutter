package slotman

import (
	"slotman/utils/simple"
	"time"
)

type RaceState string

const (
	RaceStateIdle          RaceState = "state.idle"
	RaceStateRaceWaiting   RaceState = "state.race.waiting"
	RaceStateRaceStarting  RaceState = "state.race.start"
	RaceStateRaceRunning   RaceState = "state.race.running"
	RaceStateRaceSuspended RaceState = "state.race.suspended"
	RaceStateRaceFinished  RaceState = "state.race.finished"
)

type Race struct {
	What string `json:"what,omitempty"`
	Mode string `json:"mode,omitempty"`

	Title string `json:"title"`

	Tracks int `json:"tracks"`
	Rounds int `json:"rounds"`
}

type RaceInfo struct {
	PilotUuid simple.UUIDHex `json:"pilotUuid"`

	Track    int `json:"track"`
	Rounds   int `json:"rounds"`
	Position int `json:"position"`

	ActRound float64 `json:"actRound"`
	TopRound float64 `json:"topRound"`

	ActSpeed float64 `json:"actSpeed"`
	TopSpeed float64 `json:"topSpeed"`

	LastRoundTime time.Time `json:"-"`
}
