package walk

import (
	"time"
	"path/filepath"
	"os"
)

var fileStats map[string]time.Time = make(map[string]time.Time)

func lastUpdated(path string) (time.Time, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return stat.ModTime(), nil
}

func walkPath(path string, info os.FileInfo, err error) error {
	if path[0] != byte(46) {
		time, err := lastUpdated(path)
		if err != nil {
			return err
		}
		fileStats[path] = time
	}
	return nil
}

func GetFileStats(root string) error {
	err := filepath.Walk(root, walkPath)
	if err != nil {
		return err
	}
	return nil
}

func CheckDiff() (bool, error) {
	modified := false
	for path, time := range fileStats {
		stat, err := os.Stat(path)
		if err != nil {
			return false, err
		}
		if time != stat.ModTime() {
			modified = true
			fileStats[path] = stat.ModTime()
		}
	}
	return modified, nil
}
