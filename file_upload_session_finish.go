package baiduyun

import (
	"encoding/json"
	"io"
	"net/http"
)

func (r *Client) FileUploadSessionFinish(req *FileUploadSessionFinishReq) error {
	token, err := r.getAuthToken()
	if err != nil {
		return err
	}

	req.Method = "create"
	req.AccessToken = token

	req_, err := req.to()
	if err != nil {
		return err
	}

	resp := new(fileUploadSessionFinishResp)

	err = r.requestURLEncode(http.MethodPost, "https://pan.baidu.com/rest/2.0/xpan/file", req_, resp)
	if err != nil {
		return err
	} else if err := resp.Err(); err != nil {
		return err
	}

	return nil
}

type FileUploadSessionFinishReq struct {
	Method      string `query:"method"` // 本接口固定为precreate
	AccessToken string `query:"access_token"`

	Path     string    `json:"path"` // 上传后使用的文件绝对路径，需要urlencode，需要与预上传precreate接口中的path保持一致
	File     io.Reader `json:"-"`
	UploadID string    `json:"uploadid"` // 预上传precreate接口下发的uploadid
	RType    *int64    `json:"rtype"`    // 0 为不重命名，返回冲突; 1 为只要path冲突即重命名; 2 为path冲突且block_list不同才重命名; 3 为覆盖，需要与预上传precreate接口中的rtype保持一致
	// local_ctime	int	否	1596009229	RequestBody参数	客户端创建时间(精确到秒)，默认为当前时间戳
	// local_mtime	int	否	1596009229	RequestBody参数	客户端修改时间(精确到秒)，默认为当前时间戳
	// zip_quality	int	否	70	RequestBody参数	图片压缩程度，有效值50、70、100
	// zip_sign	int	否	7d57c40c9fdb4e4a32d533bee1a4e409	RequestBody参数	未压缩原始图片文件真实md5
	// is_revision	int	否	0	RequestBody参数	是否需要多版本支持
	// 1为支持，0为不支持， 默认为0 (带此参数会忽略重命名策略)
	// mode	int	否	1	RequestBody参数	上传方式
	// 1 手动、2 批量上传、3 文件自动备份
	// 4 相册自动备份、5 视频自动备份
	// exif_info	string	否	{"height":3024,"date_time_original":"2018:09:06 15:58:58","model":"iPhone 6s","width":4032,"date_time_digitized":"2018:09:06 15:58:58","date_time":"2018:09:06 15:58:58","orientation":6,"recovery":0}	RequestBody参数	json字符串，orientation、width、height、recovery为必传字段，其他字段如果没有可以不传
}

type fileUploadSessionFinishReq struct {
	Method      string `query:"method"` // 本接口固定为precreate
	AccessToken string `query:"access_token"`

	Path      string `json:"path"`       // 上传后使用的文件绝对路径，需要urlencode，需要与预上传precreate接口中的path保持一致
	Size      int64  `json:"size"`       // 文件或目录的大小，必须要和文件真实大小保持一致，需要与预上传precreate接口中的size保持一致
	IsDir     int64  `json:"isdir"`      // 0 文件、1 目录，需要与预上传precreate接口中的isdir保持一致
	BlockList string `json:"block_list"` //	是	["7d57c40c9fdb4e4a32d533bee1a4e409"]	RequestBody参数	文件各分片md5数组的json串 需要与预上传precreate接口中的block_list保持一致，同时对应分片上传superfile2接口返回的md5，且要按照序号顺序排列，组成md5数组的json串。
	UploadID  string `json:"uploadid"`   // 预上传precreate接口下发的uploadid
	RType     *int64 `json:"rtype"`      // 0 为不重命名，返回冲突; 1 为只要path冲突即重命名; 2 为path冲突且block_list不同才重命名; 3 为覆盖，需要与预上传precreate接口中的rtype保持一致
	// local_ctime	int	否	1596009229	RequestBody参数	客户端创建时间(精确到秒)，默认为当前时间戳
	// local_mtime	int	否	1596009229	RequestBody参数	客户端修改时间(精确到秒)，默认为当前时间戳
	// zip_quality	int	否	70	RequestBody参数	图片压缩程度，有效值50、70、100
	// zip_sign	int	否	7d57c40c9fdb4e4a32d533bee1a4e409	RequestBody参数	未压缩原始图片文件真实md5
	// is_revision	int	否	0	RequestBody参数	是否需要多版本支持
	// 1为支持，0为不支持， 默认为0 (带此参数会忽略重命名策略)
	// mode	int	否	1	RequestBody参数	上传方式
	// 1 手动、2 批量上传、3 文件自动备份
	// 4 相册自动备份、5 视频自动备份
	// exif_info	string	否	{"height":3024,"date_time_original":"2018:09:06 15:58:58","model":"iPhone 6s","width":4032,"date_time_digitized":"2018:09:06 15:58:58","date_time":"2018:09:06 15:58:58","orientation":6,"recovery":0}	RequestBody参数	json字符串，orientation、width、height、recovery为必传字段，其他字段如果没有可以不传
}

func (r *FileUploadSessionFinishReq) to() (*fileUploadSessionFinishReq, error) {
	bs, err := io.ReadAll(r.File)
	if err != nil {
		return nil, err
	}
	block := []string{}
	for _, v := range splitBytes(bs, blockMaxSize) {
		block = append(block, getMd5(v))
	}
	x, _ := json.Marshal(block)
	return &fileUploadSessionFinishReq{
		Method:      r.Method,
		AccessToken: r.AccessToken,
		Path:        r.Path,
		Size:        int64(len(bs)),
		IsDir:       0,
		BlockList:   string(x),
		RType:       r.RType,
		UploadID:    r.UploadID,
	}, nil
}

type fileUploadSessionFinishResp struct {
	errnoErr
}
