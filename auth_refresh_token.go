package baiduyun

import (
	"fmt"
	"net/http"
)

func (r *Client) AuthRefreshToken(refreshToken string) (*Token, error) {
	url := fmt.Sprintf("https://openapi.baidu.com/oauth/2.0/token?"+
		"grant_type=refresh_token&"+
		"refresh_token=%s&"+
		"client_id=%s&"+
		"client_secret=%s", refreshToken, r.appKey, r.secret)
	resp := new(tokenResp)
	err := r.requestJSON(http.MethodGet, url, nil, resp)
	if err != nil {
		return nil, err
	} else if resp.ErrorDescription != "" {
		return nil, fmt.Errorf(resp.ErrorDescription)
	}

	r.accessToken = resp.AccessToken
	r.refreshToken = resp.RefreshToken

	return resp.Token, nil
}
