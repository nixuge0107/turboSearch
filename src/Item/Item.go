package Item

import (
	"github.com/yanyiwu/gojieba"
)

var (
	g_Item      = Item{}
	g_Dropwords = dropWords{}
)

type Item struct {
	Keyword  string
	Location int
	DocId    uint16
	Freq     float32
}

func (item *Item) Init() {
	g_Dropwords.Init("src/data/drop_word.txt")
}

func (item *Item) getItemWord(content string) []string {
	words := item.jiebaParticiples(content)
	keywords := g_Dropwords.DropWords(words)
	return keywords
}

func (item *Item) jiebaParticiples(content string) []string {
	var words []string
	jieba := gojieba.NewJieba()
	defer jieba.Free()

	words = jieba.Cut(content, true)
	return words
}
