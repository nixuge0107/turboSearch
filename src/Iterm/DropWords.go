package Iterm

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type dropWords struct {
	dropwords []string
}

func (dropWord *dropWords) Init(file string) {
	dropWord.dropwords = dropWord.loadDropDictionary("src/data/drop_word.txt")
}

func (dropWord *dropWords) loadDropDictionary(file string) []string {

	var drop_words []string
	dictFile, err := os.Open(file)
	defer dictFile.Close()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("载入drop字典文件 \"%s\" \n", file)
	}

	reader := bufio.NewReader(dictFile)
	for {
		str, err := reader.ReadString('\n') // 循环读取一行
		if err == io.EOF {
			fmt.Println("drop字典载入完毕")
			return drop_words
		}
		if err != nil {
			fmt.Println("drop字典载入错误, err: ", err)
			return drop_words
		}
		str = strings.Replace(str, "\n", "", -1)
		drop_words = append(drop_words, str)
	}

	return drop_words
}

func (dropWord *dropWords) isDropWord(word string) bool {
	for _, dropword := range dropWord.dropwords {
		if strings.Compare(dropword, word) == 0 {
			return true
		}
	}
	return false
}

func (dropWord *dropWords) DropWords(words []string) []string {
	var keywords []string
	for _, word := range words {
		if dropWord.isDropWord(word) {
			continue
		}
		keywords = append(keywords, word)
	}
	return keywords
}
