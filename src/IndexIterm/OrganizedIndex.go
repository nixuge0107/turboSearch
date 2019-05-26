package IndexIterm

import (
	"Iterm"
	"fmt"
)

type OrganizedIndex struct {
	keyword                 string
	organizedIndexItem    []OrganizedIndexItem
	ItemList              []Iterm.Iterm
	SumorganizedLocation    uint16
}

type OrganizedIndexItem struct {
	Keyword                 string
	organizedLocation     []OrganizedLocation
	Freq                    float32
	SumLoc                  uint16
}

type OrganizedLocation struct {
	Location              []int
	DocId                   uint16
	SumDocLoc               uint16
}

func (organizedIndex *OrganizedIndex) InitIndex(inputIterm Iterm.SearchItermSet){
	organizedIndex.keyword = "null"
	organizedIndex.ItemList = inputIterm.ItermList
}

func (organizedIndex *OrganizedIndex) SetKeyword(inputKeyword string){
	organizedIndex.keyword = inputKeyword
}

func (organizedIndex *OrganizedIndex) GenerateIndex(inputIterm Iterm.SearchItermSet) {
	for i, item := range inputIterm.ItermList {
		if i==0{
			organizedLocation                          := OrganizedLocation{}
			organizedLocation.Location                  = append(organizedLocation.Location, item.Location)
			organizedLocation.SumDocLoc                 = 1
			organizedLocation.DocId                     = item.DocId
			organizedIndexItem                         := OrganizedIndexItem{}
			organizedIndexItem.organizedLocation        = append(organizedIndexItem.organizedLocation, organizedLocation)
			organizedIndexItem.Keyword                  = item.Keyword
			organizedIndexItem.SumLoc                   = 1
			organizedIndexItem.Freq                     = item.Freq
			organizedIndex.keyword                      = "null"
			organizedIndex.ItemList                     = inputIterm.ItermList
			organizedIndex.organizedIndexItem           = append( organizedIndex.organizedIndexItem,  organizedIndexItem)
			organizedIndex.SumorganizedLocation         = 1
			fmt.Println( item )
			fmt.Println( "done" )
			continue
		}
		if organizedIndex.isInIndex(item.Keyword){
			indexWhere := organizedIndex.isInIndexWhere(item.Keyword)
			if organizedIndex.isInDocID(item.DocId, indexWhere) {
				docWhere  := organizedIndex.isInDocIDWhere(item.DocId, indexWhere)
				organizedIndex.organizedIndexItem[indexWhere].organizedLocation[docWhere].Location = append(organizedIndex.organizedIndexItem[indexWhere].organizedLocation[docWhere].Location, item.Location)
				organizedIndex.organizedIndexItem[indexWhere].organizedLocation[docWhere].SumDocLoc ++
			}else {
				organizedLocation                                                := OrganizedLocation{}
				organizedLocation.Location                                        = append(organizedLocation.Location, item.Location)
				organizedLocation.SumDocLoc                                       = 1
				organizedLocation.DocId                                           = item.DocId
				organizedIndex.organizedIndexItem[indexWhere].organizedLocation   = append(organizedIndex.organizedIndexItem[indexWhere].organizedLocation, organizedLocation)
				organizedIndex.organizedIndexItem[indexWhere].SumLoc ++
			}
		}else{
			organizedLocation                          := OrganizedLocation{}
			organizedLocation.Location                  = append(organizedLocation.Location, item.Location)
			organizedLocation.SumDocLoc                 = 1
			organizedLocation.DocId                     = item.DocId
			organizedIndexItem                         := OrganizedIndexItem{}
			organizedIndexItem.organizedLocation        = append(organizedIndexItem.organizedLocation, organizedLocation)
			organizedIndexItem.Keyword                  = item.Keyword
			organizedIndexItem.SumLoc                   = 1
			organizedIndexItem.Freq                     = item.Freq
			organizedIndex.keyword                      = "null"
			organizedIndex.ItemList                     = inputIterm.ItermList
			organizedIndex.organizedIndexItem           = append( organizedIndex.organizedIndexItem,  organizedIndexItem)
			organizedIndex.SumorganizedLocation ++
		}
		fmt.Println(item)
		fmt.Println( "done" )
	}
}

func (organizedIndex *OrganizedIndex) PrintIndexList(){
	for _, item := range organizedIndex.organizedIndexItem{
		fmt.Println("the keyword:" + item.Keyword )
		for _, jtem := range item.organizedLocation{
			fmt.Print("doc id ->")
			fmt.Print(jtem.DocId)
			fmt.Println("\n    Loc List:---")
			for _, ktem := range jtem.Location{
				fmt.Print("    -")
				fmt.Println(ktem)
			}
		}
	}
}

func (organizedIndex *OrganizedIndex) isInIndexWhere(inputKeyword string) int{
	for i, item := range organizedIndex.organizedIndexItem{
		if inputKeyword == item.Keyword{
			return i
		}
	}
	return 0xffff
}

func (organizedIndex *OrganizedIndex) isInDocIDWhere(inputDocID uint16, inputIndexWhere int) int{
	for i,item := range organizedIndex.organizedIndexItem[inputIndexWhere].organizedLocation{
		if item.DocId == inputDocID{
			return i
		}
	}
	return 0xffff
}

func (organizedIndex *OrganizedIndex) isInIndex(inputKeyword string) bool{
	for _, item := range organizedIndex.organizedIndexItem{
		if inputKeyword == item.Keyword{
			return true
		}
	}
	return false
}

func (organizedIndex *OrganizedIndex) isInDocID(inputDocID uint16, inputIndexWhere int) bool{
	for _,item := range organizedIndex.organizedIndexItem[inputIndexWhere].organizedLocation{
		if item.DocId == inputDocID{
			return true
		}
	}
	return false
}