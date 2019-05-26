package IndexIterm

import (
	"Iterm"
	"fmt"
)

type IndexIterm struct {
	IndexKeyword string
	ItermList []Iterm.Iterm
}

var(
	DamnString = IndexIterm{}
)

func (indexItem *IndexIterm) InitIndex(inputIterm Iterm.SearchItermSet){
	indexItem.IndexKeyword = "null"
	indexItem.ItermList = inputIterm.ItermList
}

func (indexItem *IndexIterm) SetKeyword(inputKeyword string){
	indexItem.IndexKeyword = inputKeyword
}

func (indexItem *IndexIterm) GenerateIndex(inputIterm Iterm.SearchItermSet) {
	//innerCtr := inputIterm.Sum
	for _, iterm := range inputIterm.ItermList {
		fmt.Println(iterm)
	}
}

