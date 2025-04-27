package main

import (
	"os"
	"time"
)

func updateWorkdayData() error {
	data, err := loadStandupData()
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	lastUpdate := fileInfo.ModTime().Truncate(24 * time.Hour)
	today := time.Now().In(time.Local).Truncate(24 * time.Hour)

	if lastUpdate.Before(today) {
		if len(data.Today) > 0 {
			data.LastWorkday = data.Today
			data.Today = []string{}
			err = saveStandupData(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}