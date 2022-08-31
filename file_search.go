package baiduyun

import (
	"net/http"
)

func (r *Client) FileSearch(req *FileSearchReq) (bool, []*FileInfo, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return false, nil, err
	}

	url := "https://pan.baidu.com/rest/2.0/xpan/file"
	request := fileSearchReq2(req, token)
	resp := new(searchListResp)

	err = r.requestJSON(http.MethodGet, url, request, resp)
	if err != nil {
		return false, nil, err
	} else if err := resp.Err(); err != nil {
		return false, nil, err
	}
	return resp.HasMore != 0, resp.List, nil
}

type FileSearchReq struct {
	Key       string  `query:"key"`       //	是	"day"	URL参数	搜索关键字
	Dir       *string `query:"dir"`       // 否	/测试目录	URL参数	搜索目录，默认根目录
	Page      *int64  `query:"page"`      //	否	1	URL参数	页数，从1开始，缺省则返回所有条目
	Num       *int64  `query:"num"`       //	否	100	URL参数	默认为500，不能修改
	Recursion *int64  `query:"recursion"` //	否	1	URL参数	是否递归搜索子目录 1:是，0:否（默认）
	Web       *int64  `query:"web"`       //	否	0	URL参数	默认0，为1时返回缩略图信息
}

// == internal ==

type fileSearchReq struct {
	Method      string  `query:"method"`
	AccessToken string  `query:"access_token"`
	Key         string  `query:"key"`       //	是	"day"	URL参数	搜索关键字
	Dir         *string `query:"dir"`       // 否	/测试目录	URL参数	搜索目录，默认根目录
	Page        *int64  `query:"page"`      //	否	1	URL参数	页数，从1开始，缺省则返回所有条目
	Num         *int64  `query:"num"`       //	否	100	URL参数	默认为500，不能修改
	Recursion   *int64  `query:"recursion"` //	否	1	URL参数	是否递归搜索子目录 1:是，0:否（默认）
	Web         *int64  `query:"web"`       //	否	0	URL参数	默认0，为1时返回缩略图信息
}

func fileSearchReq2(r *FileSearchReq, token string) *fileSearchReq {
	return &fileSearchReq{
		Method:      "search",
		AccessToken: token,
		Key:         r.Key,
		Dir:         r.Dir,
		Page:        r.Page,
		Num:         r.Num,
		Recursion:   r.Recursion,
		Web:         r.Web,
	}
}

type searchListResp struct {
	errnoErr
	List    []*FileInfo `json:"list"`
	HasMore int         `json:"has_more"`
}
