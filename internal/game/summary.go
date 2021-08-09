package game

import (
	"sync"
)

type Summary struct {
	OnlineUser     int
	RegisteredUser int
}

type SummaryManager struct {
	sync.RWMutex

	summary Summary
}

var summaryManager SummaryManager

func prepareSummary() {
	summaryManager.summary.OnlineUser = len(g.activeChars)
	summaryManager.summary.RegisteredUser = len(g.userCharacters)
}

func reloadSummary() {
	summaryManager.Lock()
	defer summaryManager.Unlock()

	prepareSummary()
}

func getSummary() Summary {
	summaryManager.RLock()
	defer summaryManager.RUnlock()

	return summaryManager.summary
}
