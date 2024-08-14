package cron

import (
	"fmt"

	. "files/global"

	"time"

	"github.com/spf13/viper"
)

func Init() {

	n := 0
	for n != 1 {

		CronJob()

		time.Sleep(5 * time.Second)
		setingskey := fmt.Sprintf("%s.cron", MicroServiceName)
		isCron := viper.GetBool(setingskey)
		if !isCron {
			n = 1
		}

	}

}

func CronJob() {
//Put your cron job codes here
}
