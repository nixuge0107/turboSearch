package Iterm

import "fmt"

type DocItermSet struct {
	ItermList []Iterm
	Sum       uint16
}

func (docitermSet *DocItermSet) AddIterm(content string, docid uint16) bool {
	words := g_Iterm.getItermWord(content)
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

		docitermSet.ItermList = append(docitermSet.ItermList, iterm)
		docitermSet.Sum++
	}
	return true
}

//test
func (docitermSet *DocItermSet) AddIterm_test() {
	docitermSet.AddIterm("一般为赋值表达式，给控制变量赋初值", 1)
	docitermSet.AddIterm("关系表达式或逻辑表达式", 2)
	docitermSet.AddIterm("数学中的赋值，通过表达式来操作", 3)
}

func (docitermSet *DocItermSet) PrintIterm_test() {
	fmt.Print(docitermSet.Sum)
	fmt.Println("--------------------")
	for _, iterm := range docitermSet.ItermList {
		fmt.Println(iterm)
	}
}
