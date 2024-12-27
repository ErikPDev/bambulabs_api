package gcode_state

type GcodeState int

const (
	IDLE GcodeState = iota
	PREPARE
	RUNNING
	PAUSE
	FINISH
	FAILED
	UNKNOWN = -1
)

// String returns a human-readable description of the G-code state.
func (gs GcodeState) String() string {
	switch gs {
	case IDLE:
		return "The printer is idle."
	case PREPARE:
		return "The printer is preparing."
	case RUNNING:
		return "The printer is running."
	case PAUSE:
		return "The printer is paused."
	case FINISH:
		return "The printer has finished."
	case FAILED:
		return "The printer has failed."
	case UNKNOWN:
		return "The printer state is unknown."
	default:
		return "Invalid Gcode state."
	}
}

// GetGcodeState returns the description based on the provided state value.
func GetGcodeState(value int) GcodeState {
	switch GcodeState(value) {
	case IDLE, PREPARE, RUNNING, PAUSE, FINISH, FAILED, UNKNOWN:
		return GcodeState(value)
	default:
		return UNKNOWN
	}
}
