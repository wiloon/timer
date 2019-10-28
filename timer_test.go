package main

import (
	"github.com/wiloon/pingd-log/logconfig"
	"github.com/wiloon/pingd-utils/utils"
	"os"
	"testing"
)

func Test00(t *testing.T) {
	logconfig.Init()

	// remove db file
	_ = os.Remove("foo.db")

	systemTime := utils.StringToDateYMDHMS("2019-10-28 09:00:00")
	timerTask(systemTime)
	systemTime = utils.StringToDateYMDHMS("2019-10-28 10:00:00")
	timerTask(systemTime)

	systemTime = utils.StringToDateYMDHMS("2019-10-29 09:00:00")
	timerTask(systemTime)

	systemTime = utils.StringToDateYMDHMS("2019-10-30 09:00:00")
	timerTask(systemTime)
}
