package words

import (
	"embed"
	"strings"

	"github.com/xinzhanguo/searchengine/searcher/utils"

	"github.com/wangbin/jiebago"
)

var (
	//go:embed data/*.txt
	dictionaryFS embed.FS
)

type Tokenizer struct {
	seg jiebago.Segmenter
}

func NewTokenizer(dictionaryPath string) *Tokenizer {
	file, err := dictionaryFS.Open("data/dictionary.txt")
	if err != nil {
		panic(err)
	}
	utils.ReleaseAssets(file, dictionaryPath)

	tokenizer := &Tokenizer{}

	err = tokenizer.seg.LoadDictionary(dictionaryPath)
	if err != nil {
		panic(err)
	}

	return tokenizer
}

func (t *Tokenizer) Cut(text string) []string {
	//不区分大小写
	text = strings.ToLower(text)
	//移除所有的标点符号
	text = utils.RemovePunctuation(text)
	//移除所有的空格
	text = utils.RemoveSpace(text)

	var wordMap = make(map[string]struct{})

	resultChan := t.seg.CutForSearch(text, true)
	var wordsSlice []string
	for {
		w, ok := <-resultChan
		if !ok {
			break
		}
		_, found := wordMap[w]
		if !found {
			//去除重复的词
			wordMap[w] = struct{}{}
			wordsSlice = append(wordsSlice, w)
		}
	}

	return wordsSlice
}
