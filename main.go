package main

import (
	"fmt"
	"github.com/kotaroooo0/gojaconv/jaconv"
	"math/rand"
	"time"
	"unicode/utf8"
)

type Seed struct {
	base    string
	hebon   string
	length  int
	english string
}
type seeds []Seed

// 生成する名前の最大長
const maxLength = 10

// 組み合わせる用語の数
const numberOfCollaborating = 2

// keyword をプログラム内で扱われる独自の構造体に変換する
func generateSeed(keywords []string) []Seed {
	var seeds seeds
	for _, v := range keywords {
		hebon := jaconv.ToHebon(v)
		seed := Seed{v, hebon, utf8.RuneCountInString(v), ""}
		seeds = append(seeds, seed)
	}
	return seeds
}

// 与えられたキーワードを組み合わせて新しい名前を作る
func collaboration(seeds []Seed) string {
	// 名前の長さを決める
	keywordLength := 0
	var nameLength int
	for _, v := range seeds {
		keywordLength += v.length
	}
	if keywordLength < 10 {
		nameLength = keywordLength
	} else {
		nameLength = maxLength
	}
	// 使う keyword を選ぶ
	// 現状はランダム
	rand.Seed(time.Now().UnixNano())
	leftIndex := rand.Intn(len(seeds))
	var rightIndex int
	for true {
		rand.Seed(time.Now().UnixNano())
		rightIndex = rand.Intn(len(seeds))
		if leftIndex != rightIndex {
			break
		}
	}

	// 名前を作成する
	leftWord := seeds[leftIndex].hebon
	rightWord := seeds[rightIndex].hebon
	var leftWordMaxLength int
	if len(leftWord) < nameLength {
		leftWordMaxLength = len(leftWord)
	} else {
		leftWordMaxLength = nameLength
	}
	rand.Seed(time.Now().UnixNano())
	leftWordLength := rand.Intn(leftWordMaxLength)
	rightWordStartIndex := len(rightWord) - (len(rightWord) - leftWordLength)
	return leftWord[:leftWordLength] + rightWord[rightWordStartIndex:]
}

func main() {
	var sample = []string{"おはよう", "こんにちは"}
	seeds := generateSeed(sample)
	name := collaboration(seeds)
	fmt.Println(name)
}
