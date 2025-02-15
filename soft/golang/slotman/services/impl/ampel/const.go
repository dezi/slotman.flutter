package ampel

//goland:noinspection GoNameStartsWithPackageName
type AmpelState string

//goland:noinspection GoNameStartsWithPackageName
const (
	AmpelStateIdle          AmpelState = "ampel.idle"
	AmpelStateRaceStart     AmpelState = "ampel.race.start"
	AmpelStateRaceWaiting   AmpelState = "ampel.race.waiting"
	AmpelStateRaceSuspended AmpelState = "ampel.race.suspend"
	AmpelStateRaceRunning   AmpelState = "ampel.race.restart"
	AmpelStateRaceFinished  AmpelState = "ampel.race.finished"
)

var (
	pinsRed    = []int{0, 1, 2, 3, 4}
	pinsGreen  = []int{5, 6, 7, 8, 9}
	pinsYellow = []int{10, 11, 12, 13, 14}
)
