package android

import (
	"fmt"
)

type Android struct{}

func (ad *Android) ChargeAppleMobile() {
	fmt.Println("Charging Android Mobile")
}
