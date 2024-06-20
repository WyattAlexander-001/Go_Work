package main

import "fmt"

type insuranceStatus struct {
	HasInsurance bool
	IsTotaled    bool
	IsDented     bool
	IsBigDent    bool
}

func (s insuranceStatus) hasInsurance() bool {
	return s.HasInsurance
}

func (s insuranceStatus) isTotaled() bool {
	return s.IsTotaled
}

func (s insuranceStatus) isDented() bool {
	return s.IsDented
}

func (s insuranceStatus) isBigDent() bool {
	return s.IsBigDent
}

func getInsuranceAmount(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}

func main() {
	status := insuranceStatus{
		HasInsurance: true,
		IsTotaled:    false,
		IsDented:     false, // Will be 0 because of this
		IsBigDent:    false,
	}
	fmt.Println(getInsuranceAmount(status))
}
