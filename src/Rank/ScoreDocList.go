package Rank

import (
	"IndexIterm"
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
}

func (scoreList *ScoreDocList) Bm25score(){
	fmt.Println(scoreList.DocList.SearchIndexItermList)
	N := scoreList.DocList.NumOfDocs
	// var score float64
	for i:=0 ; i < N ; i++ {
		temp := ScoreDoc{}
		temp.Score = 0
		temp.DocId = uint16(i)
		scoreList.ScoreLists = append(scoreList.ScoreLists, temp)
	}
	for _, item := range scoreList.DocList.OrganizeIndexList{
		for _, itemOrg3 := range item.Org3{
			// fmt.Println(item.SumLoc)
			// score := math.Log((float64(N)+0.5)/(float64(item.SumLoc)+0.5))
			// fmt.Println(score)
			DocID := itemOrg3.DocId
			
			for j, itemScore := range scoreList.ScoreLists{
				if DocID == itemScore.DocId{
					fmt.Println(j)
					fmt.Println("check")
					scoreList.ScoreLists[j].Score += math.Log((float64(N)+0.5)/(float64(item.SumLoc)+0.5))
					fmt.Println(math.Log((float64(N)+0.5)/(float64(item.SumLoc)+0.5)))
					fmt.Println(scoreList.ScoreLists)
				} 
			}
		}
	}
	fmt.Println(scoreList.ScoreLists)
}

func (scoreList *ScoreDocList) GetScore() {
	for _, docid := range scoreList.DocList.SearchDocIdList {
		var itermLocations []IndexIterm.Org3
		for _, indexiterm := range scoreList.DocList.SearchIndexItermList {
			for _, indexloc := range indexiterm.Org3 {
				if docid == indexloc.DocId {
					itermLocations = append(itermLocations, indexloc)
				}
			}
		}
		var score ScoreDoc
		score.DocId = docid
		score.Score = scoreList.GetScoreFromLocation(itermLocations)
		scoreList.ScoreLists = append(scoreList.ScoreLists, score)
	}
}

//随便写的积分规则。。。
func (scoreList *ScoreDocList) GetScoreFromLocation(itermlocations []IndexIterm.Org3) float64 {
	var distance float64 = 0
	var lastdistance float64 = 0
	var freq uint16 = 0
	for index, itermloc := range itermlocations {
		freq = freq + itermloc.SumDocLoc
		if index == 0 {
			lastdistance = float64(itermloc.Location[0])
			continue
		}
		distance = math.Abs(lastdistance-float64(itermloc.Location[0])) + distance
		lastdistance = float64(itermloc.Location[0])
	}
	return math.Log2(distance) * float64(freq)

}
