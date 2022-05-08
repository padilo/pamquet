package infra

import (
	"log"
	"os"
	"os/user"
	"path"
)

func ConfigDir() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	dir := path.Join(u.HomeDir, "/.pomaquet")

	err = os.MkdirAll(dir, 0700)
	if err != nil {
		log.Fatalf("Couldn't create %v directory\n", dir)
	}
	return dir
}
