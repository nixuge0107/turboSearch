package Rank

import (
	"IndexIterm"
	"Iterm"
	"fmt"
)

type SearchDocList struct {
	Org1                   IndexIterm.Org1
	SearchItermList      []Iterm.Iterm
	OrganizeIndexList    []IndexIterm.Org2
	SearchDocIdList      []uint16
	SearchIndexItermList []IndexIterm.Org2
	NumOfDocs              int
}

func (searchDocList *SearchDocList) InitSearchDoc(searchIterm []Iterm.Iterm, organizedIndex IndexIterm.Org1, numOfDocs int) {
	searchDocList.Org1              = organizedIndex
	searchDocList.SearchItermList   = searchIterm
	searchDocList.OrganizeIndexList = nil
	searchDocList.NumOfDocs         = numOfDocs
}

//获得包含所有关键词的indexitermlist
func (searchDocList *SearchDocList) GetOrganizeIndexListBySearch() {
	searchDocList.GetOrganizeIndexList()
	temp := searchDocList.GetDocListFromIndexIterms()
	searchDocList.SearchDocIdList = searchDocList.GetSameDocid(temp)
	searchDocList.GetSearchIndexListByDoclist()
	fmt.Println(searchDocList.SearchIndexItermList)
}

func (searchDocList *SearchDocList) GetOrganizeIndexList() {
	var organizedIndexIterms []IndexIterm.Org2
	for _, iterm := range searchDocList.SearchItermList {
		organizedIndexIterm := searchDocList.GetIndexItermBySearchKeyword(iterm.Keyword)
		organizedIndexIterms = append(organizedIndexIterms, organizedIndexIterm)
	}
	searchDocList.OrganizeIndexList = organizedIndexIterms
}

func (searchDocList *SearchDocList) GetIndexItermBySearchKeyword(keyword string) IndexIterm.Org2 {
	var organizedIndexItem IndexIterm.Org2
	for _, indexIterm := range searchDocList.Org1.Org2 {
		if keyword == indexIterm.Keyword {
			organizedIndexItem = indexIterm
		}
	}
	return organizedIndexItem
}

func (searchDocList *SearchDocList) GetDocListFromIndexIterms() [][]uint16 {
	var docs [][]uint16
	for _, iterms := range searchDocList.OrganizeIndexList {
		var doc []uint16
		for _, loc := range iterms.Org3 {
			doc = append(doc, loc.DocId)
		}
		docs = append(docs, doc)
	}
	return docs
}

func (searchDocList *SearchDocList) PrintIndexIterm() {
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
	var searchIndexItermList []IndexIterm.Org2
	for _, indexIterm := range searchDocList.OrganizeIndexList {
		var searchIndexIterm IndexIterm.Org2
		searchIndexIterm.Keyword = indexIterm.Keyword
		searchIndexIterm.Freq = indexIterm.Freq
		searchIndexIterm.SumLoc = indexIterm.SumLoc
		for _, location := range indexIterm.Org3 {
			for _, docid := range searchDocList.SearchDocIdList {
				if docid == location.DocId {
					searchIndexIterm.Org3 = append(searchIndexIterm.Org3, location)
				}
			}
		}

		searchIndexItermList = append(searchIndexItermList, searchIndexIterm)
	}
	searchDocList.SearchIndexItermList = searchIndexItermList
}
