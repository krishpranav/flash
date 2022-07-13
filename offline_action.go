package flash

import "fmt"

type OfflineAction int

const (
	OfflineActionDisablePage = iota
	OfflineActionDisableForms
	OfflineActionDoNothing
)

func getIsConnectedJavascript(action OfflineAction) (js string) {
	switch action {
	case OfflineActionDisablePage:
		return `
		var display = isConnected ? "none" : "flex";
		document.getElementById("disconnectedoverlay").style.display = display;
	`

	case OfflineActionDisableForms:
		return `
		document.getElementById("disconnectedoverlay").style.display = "none";
		if (!isConnected) {
			["input", "button", "textarea", "select"].forEach((selector) => {
				var matches = document.querySelectorAll(selector);
				for (var i = 0; i < matches.length; i++) {
					matches[i].disabled = true;
				}
			})
		}
	`

	case OfflineActionDoNothing:
		return ""
	}

	panic(fmt.Sprintf("Invalid OfflineAction: %v", action))
}