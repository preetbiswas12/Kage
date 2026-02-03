package provider

import (
	"github.com/preetbiswas12/Kage/provider/generic"
	"github.com/preetbiswas12/Kage/provider/mangadex"
	"github.com/preetbiswas12/Kage/provider/manganato"
	"github.com/preetbiswas12/Kage/provider/manganelo"
	"github.com/preetbiswas12/Kage/provider/mangapill"
	"github.com/preetbiswas12/Kage/source"
)

const CustomProviderExtension = ".lua"

var builtinProviders = []*Provider{
	{
		ID:   mangadex.ID,
		Name: mangadex.Name,
		CreateSource: func() (source.Source, error) {
			return mangadex.New(), nil
		},
	},
}

func init() {
	for _, conf := range []*generic.Configuration{
		manganelo.Config,
		manganato.Config,
		mangapill.Config,
	} {
		conf := conf
		builtinProviders = append(builtinProviders, &Provider{
			ID:   conf.ID(),
			Name: conf.Name,
			CreateSource: func() (source.Source, error) {
				return generic.New(conf), nil
			},
		})
	}
}
