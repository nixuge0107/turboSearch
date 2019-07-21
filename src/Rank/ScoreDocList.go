package Rank

import (
	"math"
	"fmt"
)

type ScoreDoc struct {
	DocId uint16
	Score float64
}

type ScoreDocList struct {
	ScoreLists []ScoreDoc
	DocList      SearchDocList
}

func (scoreList *ScoreDocList) InitScoreList(searchDocList SearchDocList) {
	scoreList.DocList = searchDocList
	N := scoreList.DocList.NumOfDocs //表示所有文档的总数
	for i:=0 ; i < N ; i++ {
		temp := ScoreDoc{}
		temp.Score = 0
		temp.DocId = uint16(i)
		scoreList.ScoreLists = append(scoreList.ScoreLists, temp)
	}
}

func (scoreList *ScoreDocList) Bm25score(){
	N := scoreList.DocList.NumOfDocs //表示所有文档的总数
	for _, item := range scoreList.DocList.OrganizeIndexList{
		for _, itemOrg3 := range item.Org3{
			DocID := itemOrg3.DocId
			
			for j, itemScore := range scoreList.ScoreLists{
				if DocID == itemScore.DocId{
					scoreList.ScoreLists[j].Score += math.Log((float64(N)+0.5)/(float64(item.SumLoc)+0.5))
				} 
			}
		}
	}
	fmt.Println(scoreList.ScoreLists)
}

// func (scoreList *ScoreDocList) GetScore() {
// 	for _, docid := range scoreList.DocList.SearchDocIdList {
// 		var itemLocations []IndexItem.Org3
// 		for _, indexitem := range scoreList.DocList.SearchIndexItemList {
// 			for _, indexloc := range indexitem.Org3 {
// 				if docid == indexloc.DocId {
// 					itemLocations = append(itemLocations, indexloc)
// 				}
// 			}
// 		}
// 		var score ScoreDoc
// 		score.DocId = docid
// 		score.Score = scoreList.GetScoreFromLocation(itemLocations)
// 		scoreList.ScoreLists = append(scoreList.ScoreLists, score)
// 	}
// }

// //随便写的积分规则。。。
// func (scoreList *ScoreDocList) GetScoreFromLocation(itemlocations []IndexItem.Org3) float64 {
// 	var distance float64 = 0
// 	var lastdistance float64 = 0
// 	var freq uint16 = 0
// 	for index, itemloc := range itemlocations {
// 		freq = freq + itemloc.SumDocLoc
// 		if index == 0 {
// 			lastdistance = float64(itemloc.Location[0])
// 			continue
// 		}
// 		distance = math.Abs(lastdistance-float64(itemloc.Location[0])) + distance
// 		lastdistance = float64(itemloc.Location[0])
// 	}
// 	return math.Log2(distance) * float64(freq)

// }
