package IndexIterm

import (
	"Iterm"
	"fmt"
)

type OrganizedIndex struct {
	keyword              string
	OrganizedIndexItem   []OrganizedIndexItem
	ItemList             []Iterm.Iterm
	SumorganizedLocation uint16
}

type OrganizedIndexItem struct {
	Keyword           string
	OrganizedLocation []OrganizedLocation
	Freq              float32
	SumLoc            uint16
}

type OrganizedLocation struct {
	Location  []int
	DocId     uint16
	SumDocLoc uint16
}

func (organizedIndex *OrganizedIndex) InitIndex(inputIterm Iterm.DocItermSet) {
	organizedIndex.keyword = "null"
	organizedIndex.ItemList = inputIterm.ItermList
}

func (organizedIndex *OrganizedIndex) SetKeyword(inputKeyword string) {
	organizedIndex.keyword = inputKeyword
}

func (organizedIndex *OrganizedIndex) GenerateIndex(inputIterm Iterm.DocItermSet) {
	for i, item := range inputIterm.ItermList {
		if i == 0 {
			OrganizedLocation := OrganizedLocation{}
			OrganizedLocation.Location = append(OrganizedLocation.Location, item.Location)
			OrganizedLocation.SumDocLoc = 1
			OrganizedLocation.DocId = item.DocId
			OrganizedIndexItem := OrganizedIndexItem{}
			OrganizedIndexItem.OrganizedLocation = append(OrganizedIndexItem.OrganizedLocation, OrganizedLocation)
			OrganizedIndexItem.Keyword = item.Keyword
			OrganizedIndexItem.SumLoc = 1
			OrganizedIndexItem.Freq = item.Freq
			organizedIndex.keyword = "null"
			organizedIndex.ItemList = inputIterm.ItermList
			organizedIndex.OrganizedIndexItem = append(organizedIndex.OrganizedIndexItem, OrganizedIndexItem)
			organizedIndex.SumorganizedLocation = 1
			fmt.Println(item)
			fmt.Println("done")
			continue
		}
		if organizedIndex.isInIndex(item.Keyword) {
			indexWhere := organizedIndex.isInIndexWhere(item.Keyword)
			if organizedIndex.isInDocID(item.DocId, indexWhere) {
				docWhere := organizedIndex.isInDocIDWhere(item.DocId, indexWhere)
				organizedIndex.OrganizedIndexItem[indexWhere].OrganizedLocation[docWhere].Location = append(organizedIndex.OrganizedIndexItem[indexWhere].OrganizedLocation[docWhere].Location, item.Location)
				organizedIndex.OrganizedIndexItem[indexWhere].OrganizedLocation[docWhere].SumDocLoc++
			} else {
				OrganizedLocation := OrganizedLocation{}
				OrganizedLocation.Location = append(OrganizedLocation.Location, item.Location)
				OrganizedLocation.SumDocLoc = 1
				OrganizedLocation.DocId = item.DocId
				organizedIndex.OrganizedIndexItem[indexWhere].OrganizedLocation = append(organizedIndex.OrganizedIndexItem[indexWhere].OrganizedLocation, OrganizedLocation)
				organizedIndex.OrganizedIndexItem[indexWhere].SumLoc++
			}
		} else {
			OrganizedLocation := OrganizedLocation{}
			OrganizedLocation.Location = append(OrganizedLocation.Location, item.Location)
			OrganizedLocation.SumDocLoc = 1
			OrganizedLocation.DocId = item.DocId
			OrganizedIndexItem := OrganizedIndexItem{}
			OrganizedIndexItem.OrganizedLocation = append(OrganizedIndexItem.OrganizedLocation, OrganizedLocation)
			OrganizedIndexItem.Keyword = item.Keyword
			OrganizedIndexItem.SumLoc = 1
			OrganizedIndexItem.Freq = item.Freq
			organizedIndex.keyword = "null"
			organizedIndex.ItemList = inputIterm.ItermList
			organizedIndex.OrganizedIndexItem = append(organizedIndex.OrganizedIndexItem, OrganizedIndexItem)
			organizedIndex.SumorganizedLocation++
		}
		fmt.Println(item)
		fmt.Println("done")
	}
}

func (organizedIndex *OrganizedIndex) PrintIndexList() {
	for _, item := range organizedIndex.OrganizedIndexItem {
		fmt.Println("the keyword:" + item.Keyword)
		for _, jtem := range item.OrganizedLocation {
			fmt.Print("doc id ->")
			fmt.Print(jtem.DocId)
			fmt.Println("\n    Loc List:---")
			for _, ktem := range jtem.Location {
				fmt.Print("    -")
				fmt.Println(ktem)
			}
		}
	}
}

func (organizedIndex *OrganizedIndex) isInIndexWhere(inputKeyword string) int {
	for i, item := range organizedIndex.OrganizedIndexItem {
		if inputKeyword == item.Keyword {
			return i
		}
	}
	return 0xffff
}

func (organizedIndex *OrganizedIndex) isInDocIDWhere(inputDocID uint16, inputIndexWhere int) int {
	for i, item := range organizedIndex.OrganizedIndexItem[inputIndexWhere].OrganizedLocation {
		if item.DocId == inputDocID {
			return i
		}
	}
	return 0xffff
}

func (organizedIndex *OrganizedIndex) isInIndex(inputKeyword string) bool {
	for _, item := range organizedIndex.OrganizedIndexItem {
		if inputKeyword == item.Keyword {
			return true
		}
	}
	return false
}

func (organizedIndex *OrganizedIndex) isInDocID(inputDocID uint16, inputIndexWhere int) bool {
	for _, item := range organizedIndex.OrganizedIndexItem[inputIndexWhere].OrganizedLocation {
		if item.DocId == inputDocID {
			return true
		}
	}
	return false
}
