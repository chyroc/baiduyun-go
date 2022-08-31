package baiduyun

import (
	"io"
	"net/http"
	"net/url"
)

func (r *Client) DownloadDLink(dlink string) (io.ReadCloser, error) {
	token, err := r.getAuthToken()
	if err != nil {
		return nil, err
	}
	uriParsed, err := url.Parse(dlink)
	if err != nil {
		return nil, err
	}
	q := uriParsed.Query()
	q.Set("access_token", token)
	uriParsed.RawQuery = q.Encode()

	request, err := http.NewRequest(http.MethodGet, uriParsed.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "pan.baidu.com")

	resp, err := r.downloadHTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
