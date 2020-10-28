package main
import (
  "fmt"
  "unicode/utf8"
  "github.com/kotaroooo0/gojaconv/jaconv"
)

type Seed struct {
  base string
  hebon string
  length int
}
type seeds []Seed

// keyword をプログラム内で扱われる独自の構造体に変換する
func generateSeed(keywords []string) []Seed {
  var seeds seeds
  for _ ,v := range keywords {
    hebon := jaconv.ToHebon(v)
    seed := Seed{v, hebon, utf8.RuneCountInString(v)}
    seeds = append(seeds, seed)
  }
  return seeds
}

// 与えられたキーワードを組み合わせて新しい名前を作る
func collaboration(seeds []Seed) string {
  collaboratedString := ""
  for _, v := range seeds {
    collaboratedString += v.hebon
  }
  return collaboratedString
}

func main() {
  var sample = []string{"おはよう", "こんにちは"}
  seeds := generateSeed(sample)
  name := collaboration(seeds)
  fmt.Println(name)
}