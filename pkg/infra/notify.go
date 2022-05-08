package infra

import (
	"os/exec"

	"github.com/gen2brain/beeep"
)

func Notify(title string, message string) error {
	_ = exec.Command("play", "-q", "./assets/notify.ogg").Start()

	return beeep.Alert(title, message, "")
}
