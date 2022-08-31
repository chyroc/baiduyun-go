package baiduyun_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/chyroc/baiduyun-go"
)

// OAuth 流程
func Example_client_AuthURL() {
	cli := baiduyun.New(baiduyun.WithAppCredential("id", "secret"))

	// 生成 OAuth 授权页面的 URL
	fmt.Println(cli.AuthURL("http://127.0.0.1:3000"))

	// 上一步生成的 code 换取 token
	res, err := cli.AuthAccessToken("afe6db28ed86582262b074df3f3f6cb2", "http://127.0.0.1:3000")
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Printf("token: %#v\n", res)
	}
}

// 获取文件列表
func Example_client_FileList() {
	cli := baiduyun.New(
		baiduyun.WithAppCredential("id", "secret"),
		baiduyun.WithToken("access-token", "refresh-token"),
	)

	files, err := cli.FileList(&baiduyun.FileListReq{
		Dir:       &[]string{"/"}[0],
		Order:     nil,
		Desc:      nil,
		Start:     nil,
		Limit:     nil,
		Web:       nil,
		Folder:    nil,
		ShowEmpty: nil,
	})
	if err != nil {
		fmt.Println("err", err)
	} else {
		for _, file := range files {
			fmt.Println(file.FsID, file.Path)
		}
	}
}

// 获取文件元数据
func Example_client_FileMetadata() {
	cli := baiduyun.New(
		baiduyun.WithAppCredential("id", "secret"),
		baiduyun.WithToken("access-token", "refresh-token"),
	)

	files, err := cli.FileMetadata(&baiduyun.FileMetadataReq{
		FsIDs:     []int64{1, 2},
		DLink:     nil,
		Path:      nil,
		Thumb:     nil,
		Extra:     nil,
		NeedMedia: nil,
	})
	if err != nil {
		fmt.Println("err", err)
	} else {
		for _, file := range files {
			fmt.Println(file.FsID, file.Path)
		}
	}
}

// 搜索文件
func Example_client_FileSearch() {
	cli := baiduyun.New(
		baiduyun.WithAppCredential("id", "secret"),
		baiduyun.WithToken("access-token", "refresh-token"),
	)

	hasMore, files, err := cli.FileSearch(&baiduyun.FileSearchReq{
		Key:       "文本",
		Dir:       nil,
		Page:      nil,
		Num:       nil,
		Recursion: nil,
		Web:       nil,
	})
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("hasMore", hasMore)
		for _, file := range files {
			fmt.Println(file.FsID, file.Path)
		}
	}
}

// 下载文件
func Example_client_DownloadFileID() {
	cli := baiduyun.New(
		baiduyun.WithAppCredential("id", "secret"),
		baiduyun.WithDownloadTimeout(time.Minute*120),
		baiduyun.WithToken("access-token", "refresh-token"),
	)

	reader, err := cli.DownloadFileID(868217257594741)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	bs, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	err = ioutil.WriteFile("./example.mp4", bs, 0644)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}

// 上传文件
func Example_client_FileUpload() {
	cli := baiduyun.New(
		baiduyun.WithAppCredential("id", "secret"),
		baiduyun.WithToken("access-token", "refresh-token"),
	)

	f, err := os.Open("/filepath/1.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer f.Close()

	print(cli.FileUpload(&baiduyun.FileUploadReq{
		Name:  "/1.txt",
		File:  f,
		RType: nil,
	}))
}
