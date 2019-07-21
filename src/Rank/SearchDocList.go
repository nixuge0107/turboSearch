package Rank

import (
	"IndexItem"
	"Item"
	"fmt"
)

type SearchDocList struct {
	Org1                   IndexItem.Org1
	SearchItemList       []Item.Item
	OrganizeIndexList    []IndexItem.Org2
	SearchDocIdList      []uint16
	SearchIndexItemList  []IndexItem.Org2
	NumOfDocs              int
}

func (searchDocList *SearchDocList) InitSearchDoc(searchItem []Item.Item, organizedIndex IndexItem.Org1, numOfDocs int) {
	searchDocList.Org1              = organizedIndex
	searchDocList.SearchItemList    = searchItem
	searchDocList.OrganizeIndexList = nil
	searchDocList.NumOfDocs         = numOfDocs
}

//获得包含所有关键词的indexitemlist
func (searchDocList *SearchDocList) GetOrganizeIndexListBySearch() {
	searchDocList.GetOrganizeIndexList()
	temp := searchDocList.GetDocListFromIndexItems()
	searchDocList.SearchDocIdList = searchDocList.GetSameDocid(temp)
	searchDocList.GetSearchIndexListByDoclist()
	fmt.Println(searchDocList.SearchIndexItemList)
}

func (searchDocList *SearchDocList) GetOrganizeIndexList() {
	var organizedIndexItems []IndexItem.Org2
	for _, item := range searchDocList.SearchItemList {
		organizedIndexItem := searchDocList.GetIndexItemBySearchKeyword(item.Keyword)
		organizedIndexItems = append(organizedIndexItems, organizedIndexItem)
	}
	searchDocList.OrganizeIndexList = organizedIndexItems
	fmt.Println("notice")
	fmt.Println(searchDocList.OrganizeIndexList)
}

func (searchDocList *SearchDocList) GetIndexItemBySearchKeyword(keyword string) IndexItem.Org2 {
	var organizedIndexItem IndexItem.Org2
	for _, indexItem := range searchDocList.Org1.Org2 {
		if keyword == indexItem.Keyword {
			organizedIndexItem = indexItem
		}
	}
	return organizedIndexItem
}

func (searchDocList *SearchDocList) GetDocListFromIndexItems() [][]uint16 {
	var docs [][]uint16
	for _, items := range searchDocList.OrganizeIndexList {
		var doc []uint16
		for _, loc := range items.Org3 {
			doc = append(doc, loc.DocId)
		}
		docs = append(docs, doc)
	}
	return docs
}

func (searchDocList *SearchDocList) PrintIndexItem() {
	for _, indexList := range searchDocList.OrganizeIndexList {
		fmt.Println(indexList.Org3)
	}
}

//util
func (searchDocList *SearchDocList) GetSameDocid(doclists [][]uint16) []uint16 {
	var last_docs []uint16
	for index, doclist := range doclists {
		if index == 0 {
			last_docs = doclist
		} else {
			last_docs = searchDocList.getSame(doclist, last_docs)
		}
	}
	return last_docs
}

func (searchDocList *SearchDocList) getSame(this []uint16, last []uint16) []uint16 {
	var docs []uint16
	for _, this_docid := range this {
		for _, last_docid := range last {
			if this_docid == last_docid {
				docs = append(docs, this_docid)
				break
			}
		}
	}
	return docs
}

func (searchDocList *SearchDocList) GetSearchIndexListByDoclist() {
	var searchIndexItemList []IndexItem.Org2
	for _, indexItem := range searchDocList.OrganizeIndexList {
		var searchIndexItem IndexItem.Org2
		searchIndexItem.Keyword = indexItem.Keyword
		searchIndexItem.Freq    = indexItem.Freq
		searchIndexItem.SumLoc  = indexItem.SumLoc
		for _, location := range indexItem.Org3 {
			for _, docid := range searchDocList.SearchDocIdList {
				if docid == location.DocId {
					searchIndexItem.Org3 = append(searchIndexItem.Org3, location)
				}
			}
		}

		searchIndexItemList = append(searchIndexItemList, searchIndexItem)
	}
	searchDocList.SearchIndexItemList = searchIndexItemList
}
