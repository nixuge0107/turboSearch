package Iterm

import (
	"github.com/yanyiwu/gojieba"
)

var (
	g_Iterm     = Iterm{}
	g_Dropwords = dropWords{}
)

type Iterm struct {
	Keyword  string
	Location int
	DocId    uint16
	Freq     float32
}

func (iterm *Iterm) Init() {
	g_Dropwords.Init("src/data/drop_word.txt")
}

func (iterm *Iterm) getItermWord(content string) []string {
	words := iterm.jiebaParticiples(content)
	keywords := g_Dropwords.DropWords(words)
	return keywords
}

func (iterm *Iterm) jiebaParticiples(content string) []string {
	var words []string
	jieba := gojieba.NewJieba()
	defer jieba.Free()

	words = jieba.Cut(content, true)
	return words
}
