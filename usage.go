package gost

import (
	"sync"
)

var hostUsageMap = map[string]uint{}
var hostUsageMapLock = sync.Mutex{}
var totalRequest = 0

func willVisitHost(host string) {
	hostUsageMapLock.Lock()
	defer hostUsageMapLock.Unlock()

	if visitTimes, found := hostUsageMap[host]; found {
		hostUsageMap[host] = visitTimes + 1
	} else {
		hostUsageMap[host] = 1
	}

	totalRequest += 1
}

func GetUsageData() map[string]uint {
	hostUsageMapLock.Lock()
	defer hostUsageMapLock.Unlock()
	if totalRequest > 100 {
		usageMap := hostUsageMap
		hostUsageMap = map[string]uint{}
		totalRequest = 0
		return usageMap
	}

	return nil
}
