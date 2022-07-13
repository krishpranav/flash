package flash

import "fmt"

type OfflineAction int

const (
	OfflineActionDisablePage = iota

	OfflineActionDisbleForms
)

func getIsConnectedJavascript(
	action OfflineAction
) (js string) {
	switch action {
	case OfflineAction:
		return ""
	}

	panic(fmt.Sprintf("Invalid OfflineAction: %v", action))
}