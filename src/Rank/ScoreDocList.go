package Rank

type ScoreDoc struct {
	DocId int
	Score int
}

type ScoreDocList struct {
	scoreDocLists []ScoreDoc
	DocList       SearchDocList
}

func (scoreList *ScoreDocList) InitScoreList() {

}
