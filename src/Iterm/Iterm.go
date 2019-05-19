package Iterm

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
)

var (
	dropwords = dropWords{}
)

type ItermSet struct {
	ItermList []Iterm
	Sum       uint16
}

type Iterm struct {
	Keyword  string
	Location int
	DocId    uint16
	Freq     float32
}

func (itermSet *ItermSet) Init() {
	dropwords.Init("src/data/drop_word.txt")
}

func (itermSet *ItermSet) AddIterm(content string, docid uint16) bool {
	words := itermSet.getItermWord(content)
	var distance = 0
	for _, word := range words {
		iterm := Iterm{}
		iterm.Keyword = string(word)
		iterm.Location = distance
		iterm.DocId = docid

		for _, singleword := range words {
			if singleword == word {
				iterm.Freq++
			}
		}

		distance = len(iterm.Keyword) + iterm.Location

		itermSet.ItermList = append(itermSet.ItermList, iterm)
		itermSet.Sum++
	}
	return true
}

func (itermSet *ItermSet) jiebaParticiples(content string) []string {
	var words []string
	use_hmm := true
	jieba := gojieba.NewJieba()
	defer jieba.Free()

	words = jieba.Cut(content, use_hmm)

	return words
}

func (itermSet *ItermSet) getItermWord(content string) []string {
	words := itermSet.jiebaParticiples(content)
	keywords := dropwords.DropWords(words)
	return keywords
}

//test
func (itermSet *ItermSet) AddIterm_test() {
	itermSet.AddIterm("一般为赋值表达式，给控制变量赋初值", 1)
	itermSet.AddIterm("关系表达式或逻辑表达式，循环控制条件", 2)

}

func (itermSet *ItermSet) PrintIterm_test() {
	fmt.Print(itermSet.Sum)
	fmt.Println("--------------------")
	for _, iterm := range itermSet.ItermList {
		fmt.Println(iterm)
	}
}
