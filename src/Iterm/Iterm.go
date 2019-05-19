package Iterm

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
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

func (itermSet *ItermSet) AddIterm(content string, docid uint16) bool {
	words := itermSet.Participles(content)
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

func (itermSet *ItermSet) Participles(content string) []string {
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	words = x.Cut(content, use_hmm)

	return words
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
