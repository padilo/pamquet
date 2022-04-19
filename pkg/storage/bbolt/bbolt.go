package bbolt

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/padilo/pomaquet/pkg/app/pomodoro/model"
	"github.com/padilo/pomaquet/pkg/storage"
	bolt "go.etcd.io/bbolt"
)

type BboltStorage struct {
	db *bolt.DB
}

func NewBboltStorage() BboltStorage {
	bboltStorage := BboltStorage{}

	bboltStorage.connect()

	return bboltStorage
}

func (b BboltStorage) Settings() storage.SettingsStorage {
	return BboltSettingsStorage{}
}

func (b *BboltStorage) connect() {
	var err error

	b.db, err = bolt.Open(storage.ConfigDir()+"/data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	b.closeOnInterrupt()
}

func (b *BboltStorage) Close() error {
	return b.db.Close()
}

func (b *BboltStorage) closeOnInterrupt() {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-cancelChan
		_ = b.Close()
	}()
}

func NewBboltSettingsStorage(s BboltStorage) BboltSettingsStorage {
	return BboltSettingsStorage{
		db: s.db,
	}
}

type BboltSettingsStorage struct {
	db *bolt.DB
}

func (b BboltSettingsStorage) Save(settings model.Settings) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		var err error
		bucket, err := tx.CreateBucketIfNotExists([]byte("settings"))
		if err != nil {
			return err
		}
		settingsJson, err := json.Marshal(settings)
		if err != nil {
			return err
		}

		return bucket.Put([]byte("settings"), settingsJson)
	})
	if err != nil {
		log.Fatalf("Problem with database update: %+v", err)
	}
}

func (b BboltSettingsStorage) Get() model.Settings {
	var settings model.Settings

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("settings"))
		if bucket != nil {
			if settingsJson := bucket.Get([]byte("settings")); settingsJson != nil {
				return json.Unmarshal(settingsJson, &settings)
			}
		}

		settings = model.NewSettings()
		return nil
	})
	if err != nil {
		log.Fatalf("Problem with database: %v", err)
	}

	return settings
}
