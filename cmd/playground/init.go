package main

import (
	"fmt"
	base62Manager "github.com/aryahmph/url-shortener/common/hash/base62"
)

func main() {
	hashManager := base62Manager.NewBase62Manager()
	fmt.Println(hashManager.Hash(8_000_002))
}
