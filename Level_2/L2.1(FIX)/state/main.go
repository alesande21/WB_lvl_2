package main

import "state_pattern/internal/traffic"

func main() {
	trafficLight := traffic.NewTrafficLight()
	trafficLight.Request()
	trafficLight.ChangeState(traffic.NewYellowState())
	trafficLight.Request()
	trafficLight.ChangeState(traffic.NewGreenState())
	trafficLight.Request()
}
