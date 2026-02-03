package custom

import (
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/source"
	lua "github.com/yuin/gopher-lua"
)

func (s *luaSource) PagesOf(chapter *source.Chapter) ([]*source.Page, error) {
	_, err := s.call(constant.ChapterPagesFn, lua.LTTable, lua.LString(chapter.URL))

	if err != nil {
		return nil, err
	}

	table := s.state.CheckTable(-1)
	pages := make([]*source.Page, 0)

	table.ForEach(func(k lua.LValue, v lua.LValue) {
		if k.Type() != lua.LTNumber {
			s.state.RaiseError("%s was expected to return a table with numbers as keys, got %s as a key", constant.ChapterPagesFn, k.Type().String())
		}

		if v.Type() != lua.LTTable {
			s.state.RaiseError("%s was expected to return a table with tables as values, got %s as a value", constant.ChapterPagesFn, v.Type().String())
		}

		page, err := pageFromTable(v.(*lua.LTable), chapter)

		if err != nil {
			s.state.RaiseError("%s", err.Error())
		}

		pages = append(pages, page)
	})

	return pages, nil
}
