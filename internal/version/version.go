package version

import "time"

var Version = "v0.0.1"
var BuildDate = "undefined"
var BuildOS = "undefined"
var GitCommit = "undefined"
var StartTime = time.Now()

func GetUptime() string {
	uptime := time.Since(StartTime)
	return uptime.String()
}
