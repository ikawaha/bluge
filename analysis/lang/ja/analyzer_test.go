package ja

import (
	"reflect"
	"testing"

	"github.com/blugelabs/bluge/analysis"
)

func TestSoraniAnalyzer(t *testing.T) {
	tests := []struct {
		input  []byte
		output analysis.TokenStream
	}{
		// stop word removal
		{
			input: []byte("ئەم پیاوە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 2,
					Start:        7,
					End:          17,
				},
			},
		},
		{
			input: []byte("پیاوە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          10,
				},
			},
		},
		{
			input: []byte("پیاو"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          8,
				},
			},
		},
	}

	analyzer := Analyzer()
	for _, test := range tests {
		actual := analyzer.Analyze(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %v, got %v", test.output, actual)
		}
	}
}
