package baiduyun

import (
	"io"
)

func (r *Client) DownloadFileID(fileID int64) (io.ReadCloser, error) {
	res, err := r.FileMetadata(&FileMetadataReq{
		FsIDs: []int64{fileID},
		DLink: ptrInt64(1),
	})
	if err != nil {
		return nil, err
	}
	return r.DownloadDLink(res[0].DLink)
}
