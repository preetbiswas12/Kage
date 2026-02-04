package tui

type state int

const (
	splashState state = iota + 1
	scrapersInstallState
	errorState
	loadingState
	historyState
	sourcesState
	searchState
	mangasState
	chaptersState
	anilistSelectState
	confirmState
	readState
	downloadState
	downloadDoneState
)
