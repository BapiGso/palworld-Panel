package backups

import (
	"palworld-panel/core/conf"
	"time"
)

func init() {
	go func() {
		ticker := time.NewTicker(time.Duration(10) * time.Minute)
		for {
			<-ticker.C
			backupsFrequency := conf.Config.Get("backupsFrequency").(int64)
			if backupsFrequency == 0 {
				continue
			}
			lastBackupTime := conf.Config.Get("lastBackupTime").(int64)
			if backupsFrequency > 0 && time.Now().Unix()-lastBackupTime > backupsFrequency {
				Compress()
			}
		}
	}()
}
