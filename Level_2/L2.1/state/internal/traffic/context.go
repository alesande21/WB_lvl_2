package traffic

type TrafficLightState interface {
	Handle()
}

type TrafficLight struct {
	currentState TrafficLightState
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{currentState: NewRedState()}
}

func (t *TrafficLight) ChangeState(state TrafficLightState) {
	t.currentState = state
}

func (t *TrafficLight) Request() {
	t.currentState.Handle()
}
