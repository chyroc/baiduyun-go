package baiduyun

import (
	"encoding/json"
	"net/http"
)

// FileMetadata 查询文件信息
//
// doc: https://pan.baidu.com/union/doc/Fksg0sbcm
func (r *Client) FileMetadata(req *FileMetadataReq) ([]*FileMeta, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return nil, err
	}

	url := "https://pan.baidu.com/rest/2.0/xpan/multimedia"
	request := fileMetadataReq2(req, token)
	resp := new(fileMetadataResp)

	err = r.requestJSON(http.MethodGet, url, request, resp)
	if err != nil {
		return nil, err
	} else if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.List, nil
}

type FileMetadataReq struct {
	FsIDs     []int64 `query:"fsids"`     // 是	[414244021542671,633507813519281]	URL参数	文件id数组，数组中元素是uint64类型，数组大小上限是：100
	DLink     *int64  `query:"dlink"`     //	否	0	URL参数	是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
	Path      *string `query:"path"`      //	否	/123-571234	URL参数	查询共享目录或专属空间内文件时需要。; 共享目录格式： /uk-fsid; 其中uk为共享目录创建者id， fsid对应共享目录的fsid; 专属空间格式：/_pcs_.appdata/xpan/
	Thumb     *int64  `query:"thumb"`     //	否	0	URL参数	是否需要缩略图地址，0为否，1为是，默认为0
	Extra     *int64  `query:"extra"`     // 否	0	URL参数	图片是否需要拍摄时间、原图分辨率等其他信息，0 否、1 是，默认0
	NeedMedia *int64  `query:"needmedia"` // 否	0	URL参数	视频是否需要展示时长信息，0 否、1 是，默认0
}

type FileMeta struct {
	Category    int64    `json:"category"`     // 文件类型，含义如下：1 视频， 2 音乐，3 图片，4 文档，5 应用，6 其他，7 种子
	DLink       string   `json:"dlink"`        // 文件下载地址，参考下载文档进行下载操作
	Filename    string   `json:"filename"`     // 文件名
	FsID        int64    `json:"fs_id"`        // id
	Isdir       int64    `json:"isdir"`        // 是否是目录，为1表示目录，为0表示非目录
	Md5         string   `json:"md5"`          // md5
	OperatorID  int64    `json:"oper_id"`      // 操作者id
	Path        string   `json:"path"`         // 路径
	ServerCtime int64    `json:"server_ctime"` // 文件的服务器创建Unix时间戳，单位秒
	ServerMtime int64    `json:"server_mtime"` // 文件的服务器修改Unix时间戳，单位秒
	Size        int64    `json:"size"`         // 文件大小，单位字节
	Thumbs      struct { // 缩略图地址
		Icon string `json:"icon,omitempty"`
		Url1 string `json:"url1,omitempty"`
		Url2 string `json:"url2,omitempty"`
		Url3 string `json:"url3,omitempty"`
	} `json:"thumbs"`
	DateTaken int64 `json:"date_taken,omitempty"` // 图片拍摄时间
	Width     int64 `json:"width,omitempty"`      // 图片宽度
	Height    int64 `json:"height,omitempty"`     // 图片高度
}

// == internal ==

type fileMetadataReq struct {
	Method      string  `query:"method"`       //	是	filemetas	URL参数	本接口固定为filemetas
	AccessToken string  `query:"access_token"` //	是	12.a6b7dbd428f731035f771b8d15063f61.86400.1292922000-2346678-124328	URL参数	接口鉴权参数
	FsIDs       string  `query:"fsids"`        // 是	[414244021542671,633507813519281]	URL参数	文件id数组，数组中元素是uint64类型，数组大小上限是：100
	DLink       *int64  `query:"dlink"`        //	否	0	URL参数	是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
	Path        *string `query:"path"`         //	否	/123-571234	URL参数	查询共享目录或专属空间内文件时需要。; 共享目录格式： /uk-fsid; 其中uk为共享目录创建者id， fsid对应共享目录的fsid; 专属空间格式：/_pcs_.appdata/xpan/
	Thumb       *int64  `query:"thumb"`        //	否	0	URL参数	是否需要缩略图地址，0为否，1为是，默认为0
	Extra       *int64  `query:"extra"`        // 否	0	URL参数	图片是否需要拍摄时间、原图分辨率等其他信息，0 否、1 是，默认0
	NeedMedia   *int64  `query:"needmedia"`    // 否	0	URL参数	视频是否需要展示时长信息，0 否、1 是，默认0
}

func fileMetadataReq2(r *FileMetadataReq, token string) *fileMetadataReq {
	bs, _ := json.Marshal(r.FsIDs)
	return &fileMetadataReq{
		Method:      "filemetas",
		AccessToken: token,
		FsIDs:       string(bs),
		DLink:       r.DLink,
		Path:        r.Path,
		Thumb:       r.Thumb,
		Extra:       r.Extra,
		NeedMedia:   r.NeedMedia,
	}
}

type fileMetadataResp struct {
	errnoErr
	List []*FileMeta `json:"list"`
}
