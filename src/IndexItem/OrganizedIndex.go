package IndexItem

import (
	"Item"
	"fmt"
)

type Org1 struct {
	keyword                string
	Org2                 []Org2
	ItemList             []Item.Item
	SumorganizedLocation   uint16
}

type Org2 struct {
	Keyword                string
	Org3                 []Org3
	Freq                   float32
	SumLoc                 uint16
}

type Org3 struct {
	Location             []int
	DocId                  uint16
	SumDocLoc              uint16
}

func (org1 *Org1) InitIndex(inputItem Item.DocItemSet) {
	org1.keyword  = "null"
	org1.ItemList = inputItem.ItemList
}

func (org1 *Org1) SetKeyword(inputKeyword string) {
	org1.keyword = inputKeyword
}

func (org1 *Org1) GenerateIndex(inputItem Item.DocItemSet) {
	for i, item := range inputItem.ItemList {
		if i == 0 {
			Org3                      := Org3{}
			Org3.Location              = append(Org3.Location, item.Location)
			Org3.DocId                 = item.DocId
			Org3.SumDocLoc             = 1
			Org2                      := Org2{}
			Org2.Org3                  = append(Org2.Org3, Org3)
			Org2.Keyword               = item.Keyword
			Org2.SumLoc                = 1
			Org2.Freq                  = item.Freq
			org1.keyword               = "null"
			org1.ItemList              = inputItem.ItemList
			org1.Org2                  = append(org1.Org2, Org2)
			org1.SumorganizedLocation  = 1
			fmt.Println(item)
			fmt.Println("done")
			continue
		}
		if org1.isInIndex(item.Keyword) {
			indexWhere := org1.isInIndexWhere(item.Keyword)
			if org1.isInDocID(item.DocId, indexWhere) {
				docWhere := org1.isInDocIDWhere(item.DocId, indexWhere)
				org1.Org2[indexWhere].Org3[docWhere].Location = append(org1.Org2[indexWhere].Org3[docWhere].Location, item.Location)
				org1.Org2[indexWhere].Org3[docWhere].SumDocLoc++
			} else {
				Org3                         := Org3{}
				Org3.Location                 = append(Org3.Location, item.Location)
				Org3.SumDocLoc                = 1
				Org3.DocId                    = item.DocId
				org1.Org2[indexWhere].Org3    = append(org1.Org2[indexWhere].Org3, Org3)
				org1.Org2[indexWhere].SumLoc++
			}
		} else {
			Org3             := Org3{}
			Org3.Location     = append(Org3.Location, item.Location)
			Org3.DocId        = item.DocId
			Org3.SumDocLoc    = 1
			Org2             := Org2{}
			Org2.Org3         = append(Org2.Org3, Org3)
			Org2.Keyword      = item.Keyword
			Org2.SumLoc       = 1
			Org2.Freq         = item.Freq
			org1.keyword      = "null"
			org1.ItemList     = inputItem.ItemList
			org1.Org2         = append(org1.Org2, Org2)
			org1.SumorganizedLocation++
		}
		fmt.Println(item)
		fmt.Println("done")
	}
}

func (org1 *Org1) PrintIndexList() {
	for _, item := range org1.Org2 {
		fmt.Println("the keyword:" + item.Keyword)
		for _, jtem := range item.Org3 {
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

func (org1 *Org1) isInIndexWhere(inputKeyword string) int {
	for i, item := range org1.Org2 {
		if inputKeyword == item.Keyword {
			return i
		}
	}
	return 0xffff
}

func (org1 *Org1) isInDocIDWhere(inputDocID uint16, inputIndexWhere int) int {
	for i, item := range org1.Org2[inputIndexWhere].Org3 {
		if item.DocId == inputDocID {
			return i
		}
	}
	return 0xffff
}

func (org1 *Org1) isInIndex(inputKeyword string) bool {
	for _, item := range org1.Org2 {
		if inputKeyword == item.Keyword {
			return true
		}
	}
	return false
}

func (org1 *Org1) isInDocID(inputDocID uint16, inputIndexWhere int) bool {
	for _, item := range org1.Org2[inputIndexWhere].Org3 {
		if item.DocId == inputDocID {
			return true
		}
	}
	return false
}
