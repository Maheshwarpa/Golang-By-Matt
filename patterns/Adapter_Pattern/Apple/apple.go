package Apple

import (
	"fmt"
)

type Apple struct{}

func (a *Apple) ChargeAppleMobile() {
	fmt.Println("Charging the Apple Mobile")
}
