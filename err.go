package baiduyun

import (
	"fmt"
)

type errnoErr struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}

var errnoMsg = map[int64]string{
	0:     "成功",
	2:     "参数错误",
	-10:   "云端容量已满",
	-9:    "文件或目录不存在",
	-8:    "文件或目录已存在",
	-7:    "文件或目录名错误或无权访问",
	-6:    "身份验证失败",
	6:     "不允许接入用户数据",
	10:    "创建文件失败",
	111:   "token 失效 或者 有其他异步任务正在执行",
	31034: "命中接口频控",
	31190: "文件不存在",
	42211: "图片详细信息查询失败",
	42212: "共享目录文件上传者信息查询失败，可重试",
	42213: "共享目录鉴权失败",
	42214: "文件基础信息查询失败",
}

func (e errnoErr) Err() error {
	if e.Errno == 0 {
		return nil
	}
	msg := e.Errmsg
	if msg == "" {
		msg = errnoMsg[e.Errno]
	}
	if msg == "" {
		msg = "未知错误"
	}
	return fmt.Errorf("%d: %s", e.Errno, msg)
}
