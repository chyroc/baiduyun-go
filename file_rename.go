package baiduyun

import (
	"encoding/json"
	"net/http"
)

// FileRename 重命名文件
//
// doc: https://pan.baidu.com/union/doc/mksg0s9l4
func (r *Client) FileRename(req *FileRenameReq) (int64, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return 0, err
	}

	url := "https://pan.baidu.com/rest/2.0/xpan/file"
	request := fileRenameReq2(req, token)
	resp := new(fileDeleteResp)

	err = r.requestURLEncode(http.MethodPost, url, request, resp)
	if err != nil {
		return 0, err
	} else if err := resp.Err(); err != nil {
		return 0, err
	}

	return resp.TaskID, nil
}

type FileRenameReq struct {
	Async       int64 // 0 同步，1 自适应，2 异步
	PathList    []*FileRenameReqPath
	OnDuplicate *string //	遇到重复文件的处理策略: fail(默认，直接返回失败)、newcopy(重命名文件)、overwrite、skip
}

type FileRenameReqPath struct {
	Path    string `json:"path"`
	NewName string `json:"newname"`
}

// == internal ==

type fileRenameReq struct {
	Method      string  `query:"method"`       //	是	filemetas	URL参数	本接口固定为filemanager
	AccessToken string  `query:"access_token"` //	是	12.a6b7dbd428f731035f771b8d15063f61.86400.1292922000-2346678-124328	URL参数	接口鉴权参数
	Operate     string  `query:"opera"`        // 可实现文件复制、移动、重命名、删除，依次对应的参数值为：copy、move、rename、delete
	Async       int64   `json:"async"`         // 0 同步，1 自适应，2 异步
	FileList    string  `json:"filelist"`      // [{"path":"/test/123456.docx}"]
	OnDuplicate *string `json:"ondup"`         //	遇到重复文件的处理策略: fail(默认，直接返回失败)、newcopy(重命名文件)、overwrite、skip
}

func fileRenameReq2(r *FileRenameReq, token string) *fileRenameReq {
	bs, _ := json.Marshal(r.PathList)
	return &fileRenameReq{
		Method:      "filemanager",
		AccessToken: token,
		Operate:     "rename",
		Async:       r.Async,
		FileList:    string(bs),
		OnDuplicate: r.OnDuplicate,
	}
}
