package baiduyun

import (
	"fmt"
	"net/http"
	"net/url"
)

// AuthAccessToken OAuth 流程使用 code 换取 token
//
// doc: https://pan.baidu.com/union/doc/al0rwqzzl
func (r *Client) AuthAccessToken(code, redirectURI string) (*Token, error) {
	uri := fmt.Sprintf("https://openapi.baidu.com/oauth/2.0/token?"+
		"grant_type=authorization_code&"+
		"code=%s&"+
		"client_id=%s&"+
		"client_secret=%s&"+
		"redirect_uri=%s",
		url.QueryEscape(code),
		url.QueryEscape(r.appKey),
		url.QueryEscape(r.secret),
		redirectURI)
	resp := new(tokenResp)

	err := r.requestJSON(http.MethodGet, uri, nil, resp)
	if err != nil {
		return nil, err
	} else if resp.ErrorDescription != "" {
		return nil, fmt.Errorf(resp.ErrorDescription)
	}

	r.accessToken = resp.AccessToken
	r.refreshToken = resp.RefreshToken

	return resp.Token, nil
}

type tokenResp struct {
	*Token
	ErrorDescription string `json:"error_description"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
