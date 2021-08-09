package game

type Summary struct {
	OnlineUser     int
	RegisteredUser int
}

var summary Summary

func prepareSummary() {
	summary.OnlineUser = len(g.activeChars)
	summary.RegisteredUser = len(g.userCharacters)
}

func getSummary() Summary {
	return summary
}
