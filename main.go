package main

import "github.com/nichuanfang/music-unblocker/v1/config"

func main() {
	c := config.LoadConfig("example.config.json")
	println(c.FfmpegPath)
}
