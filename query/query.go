package query

import (
	"github.com/metafates/gache"
	"github.com/preetbiswas12/Kage/filesystem"
	"github.com/preetbiswas12/Kage/where"
)

type queryRecord struct {
	Rank  int    `json:"rank"`
	Query string `json:"query"`
}

var cacher = gache.New[map[string]*queryRecord](
	&gache.Options{
		Path:       where.Queries(),
		FileSystem: &filesystem.GacheFs{},
	},
)
