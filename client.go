package baiduyun

import (
	"net/http"
	"time"
)

type Client struct {
	appKey             string
	secret             string
	downloadTimeout    time.Duration
	downloadHTTPClient *http.Client
	accessToken        string
	refreshToken       string
	logger             Logger
}

func New(options ...Option) *Client {
	opt := new(option)
	for _, v := range options {
		v(opt)
	}
	if opt.downloadTimeout == 0 {
		opt.downloadTimeout = time.Minute * 120 // 2h
	}

	return &Client{
		appKey:          opt.appKey,
		secret:          opt.secret,
		downloadTimeout: opt.downloadTimeout,
		downloadHTTPClient: &http.Client{
			Timeout: opt.downloadTimeout,
		},
		accessToken:  opt.accessToken,
		refreshToken: opt.refreshToken,
		logger:       nil,
	}
}
