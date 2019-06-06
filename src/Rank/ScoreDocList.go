package Rank

import (
	"IndexIterm"
	"math"
)

type ScoreDoc struct {
	DocId uint16
	Score float64
}

type ScoreDocList struct {
	ScoreLists []ScoreDoc
	DocList    SearchDocList
}

func (scoreList *ScoreDocList) InitScoreList(searchDocList SearchDocList) {
	scoreList.DocList = searchDocList
}

func (scoreList *ScoreDocList) GetScore() {
	for _, docid := range scoreList.DocList.SearchDocIdList {
		var itermLocations []IndexIterm.OrganizedLocation
		for _, indexiterm := range scoreList.DocList.SearchIndexItermList {
			for _, indexloc := range indexiterm.OrganizedLocation {
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
func (scoreList *ScoreDocList) GetScoreFromLocation(itermlocations []IndexIterm.OrganizedLocation) float64 {
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
