package baiduyun

import (
	"fmt"
	"io"
	"net/http"
)

func (r *Client) FileUploadSessionAppend(req *FileUploadSessionAppendReq) error {
	token, err := r.getAuthToken()
	if err != nil {
		return err
	}

	req.Method = "upload"
	req.AccessToken = token
	req.Type = "tmpfile"

	resp := new(fileUploadSessionAppendResp)

	err = r.requestForm(http.MethodPost, "https://d.pcs.baidu.com/rest/2.0/pcs/superfile2", req, resp)
	if err != nil {
		return err
	} else if err := resp.Err(); err != nil {
		return err
	} else if resp.ErrorMsg != "" {
		return fmt.Errorf(resp.ErrorMsg)
	}

	return nil
}

type FileUploadSessionAppendReq struct {
	Method      string    `query:"method"` // 本接口固定为precreate
	AccessToken string    `query:"access_token"`
	Type        string    `query:"type"`     // 固定值 tmpfile
	Path        string    `query:"path"`     // 需要与上一个阶段预上传precreate接口中的path保持一致
	UploadID    string    `query:"uploadid"` // 上一个阶段预上传precreate接口下发的uploadid
	PartSeq     int64     `query:"partseq"`  // 文件分片的位置序号，从0开始，参考上一个阶段预上传precreate接口返回的block_list
	File        io.Reader `file:"file"`      // 是		RequestBody参数	上传的文件内容
}

type fileUploadSessionAppendResp struct {
	errnoErr
	ErrorMsg string `json:"error_msg"`
}
