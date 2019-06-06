package Rank

import (
	"IndexIterm"
	"Iterm"
	"fmt"
)

type SearchDocList struct {
	OrganizedIndex       IndexIterm.OrganizedIndex
	SearchItermList      []Iterm.Iterm
	OrganizeIndexList    []IndexIterm.OrganizedIndexItem
	SearchDocIdList      []uint16
	SearchIndexItermList []IndexIterm.OrganizedIndexItem
}

func (searchDocList *SearchDocList) InitSearchDoc(searchIterm []Iterm.Iterm, organizedIndex IndexIterm.OrganizedIndex) {
	searchDocList.OrganizedIndex = organizedIndex
	searchDocList.SearchItermList = searchIterm
	searchDocList.OrganizeIndexList = nil
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
	var organizedIndexIterms []IndexIterm.OrganizedIndexItem
	for _, iterm := range searchDocList.SearchItermList {
		organizedIndexIterm := searchDocList.GetIndexItermBySearchKeyword(iterm.Keyword)
		organizedIndexIterms = append(organizedIndexIterms, organizedIndexIterm)
	}
	searchDocList.OrganizeIndexList = organizedIndexIterms
}

func (searchDocList *SearchDocList) GetIndexItermBySearchKeyword(keyword string) IndexIterm.OrganizedIndexItem {
	var organizedIndexItem IndexIterm.OrganizedIndexItem
	for _, indexIterm := range searchDocList.OrganizedIndex.OrganizedIndexItem {
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
		for _, loc := range iterms.OrganizedLocation {
			doc = append(doc, loc.DocId)
		}
		docs = append(docs, doc)
	}
	return docs
}

func (searchDocList *SearchDocList) PrintIndexIterm() {
	for _, indexList := range searchDocList.OrganizeIndexList {
		fmt.Println(indexList.OrganizedLocation)
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
	var searchIndexItermList []IndexIterm.OrganizedIndexItem
	for _, indexIterm := range searchDocList.OrganizeIndexList {
		var searchIndexIterm IndexIterm.OrganizedIndexItem
		searchIndexIterm.Keyword = indexIterm.Keyword
		searchIndexIterm.Freq = indexIterm.Freq
		for _, location := range indexIterm.OrganizedLocation {
			for _, docid := range searchDocList.SearchDocIdList {
				if docid == location.DocId {
					searchIndexIterm.OrganizedLocation = append(searchIndexIterm.OrganizedLocation, location)
				}
			}
		}

		searchIndexItermList = append(searchIndexItermList, searchIndexIterm)
	}
	searchDocList.SearchIndexItermList = searchIndexItermList
}
