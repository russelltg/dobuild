package libdobuild

import (
	"fmt"
)

func CreatePrinterTarget(toSay string) *Target {
	ret := Target{}

	ret.NeedsRebuild = func(*Target) bool { return true }
	ret.Build = func(*Target) bool {
		fmt.Print(toSay)
		return true
	}

	return &ret
}
