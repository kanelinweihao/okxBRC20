package pack

import (
	"embed"
	"github.com/kanelinweihao/okxBRC20/app/utils/err"
	"github.com/kanelinweihao/okxBRC20/app/utils/file"
	"strings"
)

var fsEmbedResource embed.FS
var fsEmbedFavicon embed.FS

func SetFSEmbedFavicon(fsEmbedFaviconFromEmbed embed.FS) {
	fsEmbedFavicon = fsEmbedFaviconFromEmbed
	return
}

func GetFSEmbedFavicon() (fsEmbedFaviconFromEmbed embed.FS) {
	fsEmbedFaviconFromEmbed = fsEmbedFavicon
	return fsEmbedFaviconFromEmbed
}

func SetFSEmbedResource(fsEmbedResourceFromEmbed embed.FS) {
	fsEmbedResource = fsEmbedResourceFromEmbed
	return
}

func GetFSEmbedResource() (fsEmbedResourceFromEmbed embed.FS) {
	fsEmbedResourceFromEmbed = fsEmbedResource
	return fsEmbedResourceFromEmbed
}

func GetFSAndPath(path string) (fsResource embed.FS, pathEmbed string) {
	fsResource = GetFSEmbedResource()
	pathEmbed = file.GetFilePathEmbed(path)
	return fsResource, pathEmbed
}

func ReadFileEmbedAsArrayByte(path string) (arrByte []byte) {
	fsResource, pathEmbed := GetFSAndPath(path)
	arrByte, errRead := fsResource.ReadFile(pathEmbed)
	err.ErrCheck(errRead)
	return arrByte
}

func ReadFileEmbedAsString(path string) (str string) {
	arrByte := ReadFileEmbedAsArrayByte(path)
	str = string(arrByte)
	return str
}

func ReadFileEmbedAsArrayString(path string) (arrStr []string) {
	str := ReadFileEmbedAsString(path)
	arrStr = strings.Split(str, "\n")
	return arrStr
}
