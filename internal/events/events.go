package events

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

const darwin = "darwin"

// https://developer.apple.com/library/archive/documentation/LanguagesUtilities/Conceptual/MacAutomationScriptingGuide/DisplayNotifications.html
func Notify(msg string) error {
	switch runtime.GOOS {
	case darwin:
		return notifyDarwin(msg)
	default:
		return errors.New("notify is not supported on this platform")
	}
}

func notifyDarwin(msg string) error {
	cmd := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf(`display notification %q`, msg),
	)

	return cmd.Run()
}

// https://developer.apple.com/library/archive/documentation/LanguagesUtilities/Conceptual/MacAutomationScriptingGuide/DisplayDialogsandAlerts.html#//apple_ref/doc/uid/TP40016239-CH15-SW1
func Alert(msg string) error {
	switch runtime.GOOS {
	case darwin:
		return alertDarwin(msg)
	default:
		return errors.New("alert is not supported on this platform")
	}
}

func alertDarwin(msg string) error {
	cmd := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf(`tell app "System Events" to display dialog %q`, msg),
	)

	return cmd.Run()
}
