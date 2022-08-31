package baiduyun

import (
	"encoding/json"
	"errors"
	"net/http"
)

// FileDelete 删除文件
//
// doc: https://pan.baidu.com/union/doc/mksg0s9l4
func (r *Client) FileDelete(req *FileDeleteReq) (int64, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return 0, err
	}

	url := "https://pan.baidu.com/rest/2.0/xpan/file"
	request := fileDeleteReq2(req, token)
	resp := new(fileDeleteResp)

	err = r.requestURLEncode(http.MethodPost, url, request, resp)
	if err != nil {
		return 0, err
	} else if err := resp.Err(); err != nil {
		return 0, err
	}

	return resp.TaskID, nil
}

type FileDeleteReq struct {
	Async    int64    // 0 同步，1 自适应，2 异步
	PathList []string // [{"path":"/test/123456.docx}"]
}

// == internal ==

type fileDeleteReq struct {
	Method      string `query:"method"`       //	是	filemetas	URL参数	本接口固定为filemanager
	AccessToken string `query:"access_token"` //	是	12.a6b7dbd428f731035f771b8d15063f61.86400.1292922000-2346678-124328	URL参数	接口鉴权参数
	Operate     string `query:"opera"`        // 可实现文件复制、移动、重命名、删除，依次对应的参数值为：copy、move、rename、delete
	Async       int64  `json:"async"`         // 0 同步，1 自适应，2 异步
	FileList    string `json:"filelist"`      // [{"path":"/test/123456.docx}"]
}

type deleteFile struct {
	Path string `json:"path"`
}

func fileDeleteReq2(r *FileDeleteReq, token string) *fileDeleteReq {
	res := []*deleteFile{}
	for _, v := range r.PathList {
		res = append(res, &deleteFile{Path: v})
	}
	bs, _ := json.Marshal(res)
	return &fileDeleteReq{
		Method:      "filemanager",
		AccessToken: token,
		Operate:     "delete",
		Async:       r.Async,
		FileList:    string(bs),
	}
}

type fileDeleteResp struct {
	errnoErr
	TaskID int64 `json:"taskid"` // 异步任务id, 当async=0或者2时返回

	Info []struct {
		Errno int64  `json:"errno"`
		Path  string `json:"path"`
	} `json:"info"`
}

func (r *fileDeleteResp) Err() error {
	if err := r.errnoErr.Err(); err != nil {
		if len(r.Info) > 0 {
			msg := ""
			for idx, v := range r.Info {
				if v.Errno == 0 {
					continue
				}
				e := errnoMsg[v.Errno]
				if e == "" {
					e = "未知错误"
				}
				if idx > 0 {
					msg += "; "
				}
				msg += v.Path + ": " + e
			}
			if msg != "" {
				return errors.New(msg)
			}
		}
		return err
	}
	return nil
}
