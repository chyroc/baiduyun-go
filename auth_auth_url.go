package baiduyun

import (
	"fmt"
	"net/url"
)

// AuthURL 生成 OAuth 授权页面的 URL
//
// doc: https://pan.baidu.com/union/doc/al0rwqzzl
func (r *Client) AuthURL(redirectURI string) string {
	return fmt.Sprintf("https://openapi.baidu.com/oauth/2.0/authorize?"+
		"response_type=code&"+
		"client_id=%s&"+
		"redirect_uri=%s&"+
		"scope=basic,netdisk",
		url.QueryEscape(r.appKey),
		url.QueryEscape(redirectURI))
}
