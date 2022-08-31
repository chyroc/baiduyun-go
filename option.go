package baiduyun

import (
	"time"
)

type option struct {
	appKey          string
	secret          string
	downloadTimeout time.Duration
	accessToken     string
	refreshToken    string
}

type Option func(o *option)

func WithAppCredential(appKey, secret string) Option {
	return func(o *option) {
		o.appKey = appKey
		o.secret = secret
	}
}

func WithDownloadTimeout(timeout time.Duration) Option {
	return func(o *option) {
		o.downloadTimeout = timeout
	}
}

func WithToken(accessToken, refreshTokens string) Option {
	return func(o *option) {
		o.accessToken = accessToken
		o.refreshToken = refreshTokens
	}
}
