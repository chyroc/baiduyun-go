package baiduyun

import (
	"net/http"
)

// FileList 获取文件列表
//
// doc: https://pan.baidu.com/union/doc/nksg0sat9
func (r *Client) FileList(req *FileListReq) ([]*FileInfo, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return nil, err
	}

	url := "https://pan.baidu.com/rest/2.0/xpan/file"
	request := fileListReq2(req, token)
	resp := new(fileListResp)

	err = r.requestJSON(http.MethodGet, url, request, resp)
	if err != nil {
		return nil, err
	} else if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.List, nil
}

type FileListReq struct {
	Dir       *string `query:"dir"`       // 以 / 开头的绝对路径, 默认为 /, 需要 url-encode 编码
	Order     *string `query:"order"`     // 排序字段，默认为 name; time表示先按文件类型排序，后按修改时间排序; name表示先按文件类型排序，后按文件名称排序; size表示先按文件类型排序，后按文件大小排序。
	Desc      *int64  `query:"desc"`      // 排序方式，默认为0，即升序；1表示降序, 排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     *int64  `query:"start"`     // 分页开始位置，默认为0，表示从第一条记录开始
	Limit     *int64  `query:"limit"`     // 分页大小，默认为1000，建议最大不超过1000
	Web       *int64  `query:"web"`       // 值为1时，返回dir_empty属性和缩略图数据
	Folder    *int64  `query:"folder"`    // 0 返回所有，1 只返回文件夹，且属性只返回path字段
	ShowEmpty *int64  `query:"showempty"` // 0 不返回，1 返回dir_empty属性
}

type FileInfo struct {
	FsID           int64  `json:"fs_id"`                     // 文件在云端的唯一标识ID
	Path           string `json:"path,omitempty"`            // 文件的绝对路径
	ServerFilename string `json:"server_filename,omitempty"` // 文件名称
	Size           int64  `json:"size,omitempty"`            // 文件大小，单位B
	ServerMtime    int64  `json:"server_mtime"`              // 文件在服务器修改时间
	ServerCtime    int64  `json:"server_ctime,omitempty"`    // 文件在服务器创建时间
	LocalMtime     int64  `json:"local_mtime,omitempty"`     // 文件在客户端修改时间
	LocalCtime     int64  `json:"local_ctime,omitempty"`     // 文件在客户端创建时间
	IsDir          int64  `json:"isdir,omitempty"`           // 是否为目录，0 文件、1 目录
	Privacy        int64  `json:"privacy"`
	Category       int64  `json:"category"`            // 文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
	Md5            string `json:"md5,omitempty"`       // 云端哈希（非文件真实MD5），只有是文件类型时，该字段才存在
	DirEmpty       int64  `json:"dir_empty,omitempty"` // 该目录是否存在子目录，只有请求参数web=1且该条目为目录时，该字段才存在， 0为存在， 1为不存在
}

// == internal ==

type fileListReq struct {
	Method      string  `query:"method"`
	AccessToken string  `query:"access_token"`
	Dir         *string `query:"dir"`       // 以/开头的绝对路径, 默认为/, 需要 url-encode 编码
	Order       *string `query:"order"`     // 排序字段，默认为 name; time表示先按文件类型排序，后按修改时间排序; name表示先按文件类型排序，后按文件名称排序; size表示先按文件类型排序，后按文件大小排序。
	Desc        *int64  `query:"desc"`      // 排序方式，默认为0，即升序；1表示降序, 排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start       *int64  `query:"start"`     // 分页开始位置，默认为0，表示从第一条记录开始
	Limit       *int64  `query:"limit"`     // 分页大小，默认为1000，建议最大不超过1000
	Web         *int64  `query:"web"`       // 值为1时，返回dir_empty属性和缩略图数据
	Folder      *int64  `query:"folder"`    // 0 返回所有，1 只返回文件夹，且属性只返回path字段
	ShowEmpty   *int64  `query:"showempty"` // 0 不返回，1 返回dir_empty属性
}

func fileListReq2(r *FileListReq, token string) *fileListReq {
	res := &fileListReq{
		Method:      "list",
		AccessToken: token,
		Dir:         r.Dir,
		Order:       r.Order,
		Desc:        r.Desc,
		Start:       r.Start,
		Limit:       r.Limit,
		Web:         r.Web,
		Folder:      r.Folder,
		ShowEmpty:   r.ShowEmpty,
	}
	return res
}

type fileListResp struct {
	errnoErr
	GuidInfo  string      `json:"guid_info"`
	List      []*FileInfo `json:"list"`
	RequestId int64       `json:"request_id"`
	Guid      int         `json:"guid"`
}
