package main

import (
	"embed"
	"fmt"
	"github.com/kanelinweihao/OKXBRC20/app/boot"
	"github.com/kanelinweihao/OKXBRC20/app/utils/pack"
)

//go:embed resource/*
var FSEmbedResource embed.FS

//go:embed resource/ico/favicon.ico
var FSEmbedFavicon embed.FS

func init() {
	fmt.Println("--------")
	pack.SetFSEmbedResource(FSEmbedResource)
	pack.SetFSEmbedFavicon(FSEmbedFavicon)
	return
}

func main() {
	boot.Boot()
	return
}
