package Iterm

import "fmt"

type SearchItermSet struct {
	ItermList []Iterm
	Sum       uint16
}

func (searchitermSet *SearchItermSet) AddIterm(content string, docid uint16) bool {
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

		searchitermSet.ItermList = append(searchitermSet.ItermList, iterm)
		searchitermSet.Sum++
	}
	return true
}

//test
func (searchitermSet *SearchItermSet) AddIterm_test() {
	searchitermSet.AddIterm("一般为赋值表达式，给控制变量赋初值", 1)
	searchitermSet.AddIterm("关系表达式或逻辑表达式，循环控制条件", 2)
}

func (searchitermSet *SearchItermSet) PrintIterm_test() {
	fmt.Print(searchitermSet.Sum)
	fmt.Println("--------------------")
	for _, iterm := range searchitermSet.ItermList {
		fmt.Println(iterm)
	}
}
