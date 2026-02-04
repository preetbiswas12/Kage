package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/preetbiswas12/Kage/key"
	"github.com/preetbiswas12/Kage/provider"
	"github.com/spf13/viper"
)

func (b *statefulBubble) Init() tea.Cmd {
	if b.state == splashState {
		// Start progress bar animation (5 second duration)
		return tea.Batch(textinput.Blink, b.progressC.SetPercent(0.0), tea.Tick(time.Duration(5), func(t time.Time) tea.Msg {
			return t
		}))
	}

	if names := viper.GetStringSlice(key.DownloaderDefaultSources); b.state != historyState && len(names) != 0 {
		var providers []*provider.Provider

		for _, name := range names {
			p, ok := provider.Get(name)
			if !ok {
				b.raiseError(fmt.Errorf("provider %s not found", name))
				return nil
			}

			providers = append(providers, p)
		}

		b.setState(loadingState)
		return tea.Batch(b.startLoading(), b.loadSources(providers), b.waitForSourcesLoaded())
	}

	return tea.Batch(textinput.Blink, b.loadProviders())
}
