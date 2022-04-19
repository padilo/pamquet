package storage

import (
	"github.com/padilo/pomaquet/pkg/app/pomodoro/model"
)

type SettingsStorage interface {
	Save(settings model.Settings)
	Get() model.Settings
}
