package race

//goland:noinspection GoNameStartsWithPackageName
type RaceState string

//goland:noinspection GoNameStartsWithPackageName
const (
	RaceStateIdle          RaceState = "state.idle"
	RaceStateRaceStart     RaceState = "state.race.start"
	RaceStateRaceRunning   RaceState = "state.race.running"
	RaceStateRaceSuspended RaceState = "state.race.suspended"
	RaceStateRaceFinished  RaceState = "state.race.finished"
)
