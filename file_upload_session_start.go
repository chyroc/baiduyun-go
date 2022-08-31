package baiduyun

import (
	"encoding/json"
	"io"
	"net/http"
)

func (r *Client) FileUploadSessionStart(req *FileUploadSessionStartReq) (*FileUploadSessionStartResp, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return nil, err
	}

	req.Method = "precreate"
	req.AccessToken = token

	req_, err := req.to()
	if err != nil {
		return nil, err
	}

	resp := new(FileUploadSessionStartResp)

	err = r.requestURLEncode(http.MethodPost, "https://pan.baidu.com/rest/2.0/xpan/file", req_, resp)
	if err != nil {
		return nil, err
	} else if err := resp.Err(); err != nil {
		return nil, err
	}

	if len(resp.BlockList) == 0 {
		resp.BlockList = []int64{0}
	}

	return resp, nil
}

type FileUploadSessionStartReq struct {
	Method      string `query:"method"` // 本接口固定为precreate
	AccessToken string `query:"access_token"`

	Path string `json:"path"` // 上传后使用的文件绝对路径，需要urlencode
	File io.Reader
	// Size       int64   `json:"size"`        // 文件和目录两种情况：上传文件时，表示文件的大小，单位B；上传目录时，表示目录的大小，目录的话大小默认为0
	// IsDir      int64   `json:"isdir"`       // 是否为目录，0 文件，1 目录
	// BlockList  string  `json:"block_list"`  // 是	["98d02a0f54781a93e354b1fc85caf488", "ca5273571daefb8ea01a42bfa5d02220"]	RequestBody参数	文件各分片MD5数组的json串。block_list的含义如下，如果上传的文件小于4MB，其md5值（32位小写）即为block_list字符串数组的唯一元素；如果上传的文件大于4MB，需要将上传的文件按照4MB大小在本地切分成分片，不足4MB的分片自动成为最后一个分片，所有分片的md5值（32位小写）组成的字符串数组即为block_list。
	// AutoInit   int64   `json:"autoinit"`    //	是	1	RequestBody参数	固定值1
	RType *int64 `json:"rtype"` // 文件命名策略，默认为0。0 表示不进行重命名，若云端存在同名文件返回错误; 1 表示当path冲突时，进行重命名; 2 表示当path冲突且block_list不同时，进行重命名; 3 当云端存在同名文件时，对该文件进行覆盖
	// UploadID   *string `json:"uploadid"`    // 否	P1-MTAuMjI4LjQzLjMxOjE1OTU4NTg==	RequestBody参数	上传ID
	// ContentMD5 *string `json:"content-md5"` // 否	b20f8ac80063505f264e5f6fc187e69a	RequestBody参数	文件MD5，32位小写
	// SliceMD5   *string `json:"slice-md5"`   // 否	9aa0aa691s5c0257c5ab04dd7eddaa47	RequestBody参数	文件校验段的MD5，32位小写，校验段对应文件前256KB
	// LocalCTime *string `json:"local_ctime"` //	否	1595919297	RequestBody参数	客户端创建时间， 默认为当前时间戳
	// LocalMTime *string `json:"local_mtime"` // 否	1595919297	RequestBody参数	客户端修改时间，默认为当前时间戳
}

type filePrepareUploadReq struct {
	Method      string  `query:"method"` // 本接口固定为precreate
	AccessToken string  `query:"access_token"`
	Path        string  `json:"path"`        // 上传后使用的文件绝对路径，需要urlencode
	Size        int64   `json:"size"`        // 文件和目录两种情况：上传文件时，表示文件的大小，单位B；上传目录时，表示目录的大小，目录的话大小默认为0
	IsDir       int64   `json:"isdir"`       // 是否为目录，0 文件，1 目录
	BlockList   string  `json:"block_list"`  // 是	["98d02a0f54781a93e354b1fc85caf488", "ca5273571daefb8ea01a42bfa5d02220"]	RequestBody参数	文件各分片MD5数组的json串。block_list的含义如下，如果上传的文件小于4MB，其md5值（32位小写）即为block_list字符串数组的唯一元素；如果上传的文件大于4MB，需要将上传的文件按照4MB大小在本地切分成分片，不足4MB的分片自动成为最后一个分片，所有分片的md5值（32位小写）组成的字符串数组即为block_list。
	AutoInit    int64   `json:"autoinit"`    //	是	1	RequestBody参数	固定值1
	RType       *int64  `json:"rtype"`       // 文件命名策略，默认为0。0 表示不进行重命名，若云端存在同名文件返回错误; 1 表示当path冲突时，进行重命名; 2 表示当path冲突且block_list不同时，进行重命名; 3 当云端存在同名文件时，对该文件进行覆盖
	UploadID    *string `json:"uploadid"`    // 否	P1-MTAuMjI4LjQzLjMxOjE1OTU4NTg==	RequestBody参数	上传ID
	ContentMD5  *string `json:"content-md5"` // 否	b20f8ac80063505f264e5f6fc187e69a	RequestBody参数	文件MD5，32位小写
	SliceMD5    *string `json:"slice-md5"`   // 否	9aa0aa691s5c0257c5ab04dd7eddaa47	RequestBody参数	文件校验段的MD5，32位小写，校验段对应文件前256KB
	LocalCTime  *string `json:"local_ctime"` //	否	1595919297	RequestBody参数	客户端创建时间， 默认为当前时间戳
	LocalMTime  *string `json:"local_mtime"` // 否	1595919297	RequestBody参数	客户端修改时间，默认为当前时间戳
}

type FileUploadSessionStartResp struct {
	errnoErr
	Path       string  `json:"path"`
	UploadID   string  `json:"uploadid"`
	ReturnType int64   `json:"return_type"` // 1 文件在云端不存在，2 文件在云端已存在
	BlockList  []int64 `json:"block_list"`  // 需要上传的分片序号列表，索引从0开始
}

func (r *FileUploadSessionStartResp) Exist() bool {
	return r != nil && r.ReturnType == 2
}

func (r FileUploadSessionStartReq) to() (*filePrepareUploadReq, error) {
	bs, err := io.ReadAll(r.File)
	if err != nil {
		return nil, err
	}
	block := []string{}
	for _, v := range splitBytes(bs, blockMaxSize) {
		block = append(block, getMd5(v))
	}
	blockList, _ := json.Marshal(block)
	_ = blockList
	return &filePrepareUploadReq{
		Method:      r.Method,
		AccessToken: r.AccessToken,
		Path:        r.Path,
		Size:        int64(len(bs)),
		IsDir:       0,
		BlockList:   string(blockList),
		AutoInit:    1,
		RType:       r.RType,
		UploadID:    nil,
		// ContentMD5:  ptr.String(getMd5(bs)),
		// SliceMD5:    nil,
		// LocalCTime:  nil,
		// LocalMTime:  nil,
	}, nil
}
