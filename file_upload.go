package baiduyun

import (
	"bytes"
	"io"
	"io/ioutil"
)

type FileUploadReq struct {
	Name  string
	File  io.Reader
	RType *int64 `json:"rtype"` // 文件命名策略，默认为0。0 表示不进行重命名，若云端存在同名文件返回错误; 1 表示当path冲突时，进行重命名; 2 表示当path冲突且block_list不同时，进行重命名; 3 当云端存在同名文件时，对该文件进行覆盖
}

func (r *Client) FileUpload(req *FileUploadReq) error {
	bs, err := ioutil.ReadAll(req.File)
	if err != nil {
		return err
	}
	res, err := r.FileUploadSessionStart(&FileUploadSessionStartReq{
		Path:  req.Name,
		File:  bytes.NewReader(bs),
		RType: req.RType,
	})
	if err != nil {
		return err
	} else if res.ReturnType == 2 {
		return nil
	}

	for i, v := range splitBytes(bs, blockMaxSize) {
		err := r.FileUploadSessionAppend(&FileUploadSessionAppendReq{
			Path:     req.Name,
			UploadID: res.UploadID,
			PartSeq:  int64(i),
			File:     bytes.NewReader(v),
		})
		if err != nil {
			return err
		}
	}

	err = r.FileUploadSessionFinish(&FileUploadSessionFinishReq{
		Path:     req.Name,
		File:     bytes.NewReader(bs),
		UploadID: res.UploadID,
		RType:    req.RType,
	})
	if err != nil {
		return err
	}

	return nil
}
