package ja

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/blugelabs/bluge/analysis/token"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		CharFilters: []analysis.CharFilter{
			NormalizeFilter(),
		},
		Tokenizer: Tokenizer(),
		TokenFilters: []analysis.TokenFilter{
			token.NewLowerCaseFilter(),
			StopWordsFilter(),
		},
	}
}
