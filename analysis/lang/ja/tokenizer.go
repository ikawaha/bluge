package ja

import (
	"unsafe"

	"github.com/blugelabs/bluge/analysis"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type JapaniseTokenizer struct{}

func (_ *JapaniseTokenizer) Tokenize(input []byte) analysis.TokenStream{
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	tokens := t.Analyze(*(*string)(unsafe.Pointer(&input)), tokenizer.Search)
	ret := make(analysis.TokenStream, 0, len(tokens))
	for _, v := range tokens {
		ret = append(ret, &analysis.Token{
			Start:        v.Start,
			End:          v.End,
			Term:         input[v.Start:v.End],
			PositionIncr: 1,
			Type:         analysis.Ideographic,
			KeyWord:      false,
		})
	}
	return ret
}

func Tokenizer() analysis.Tokenizer {
	return &JapaniseTokenizer{}
}