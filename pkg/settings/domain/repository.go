package domain

type SettingsRepository interface {
	Save(settings Settings)
	Get() Settings
}
