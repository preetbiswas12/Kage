package integration

import (
	"github.com/preetbiswas12/Kage/integration/anilist"
	"github.com/preetbiswas12/Kage/source"
)

// Integrator is the interface that wraps the basic integration methods.
type Integrator interface {
	// MarkRead marks a chapter as read
	MarkRead(chapter *source.Chapter) error
}

var (
	Anilist Integrator = anilist.New()
)
