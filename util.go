package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/url"
	"strings"
)

var (
	code_verifier = "5d2309e5bb73b864f989753887fe52f79ce5270395e25862da6940d5" // 长度为43到128位的随机字符串
	scope         = ""
)

func genCodeChallenge() string {
	sum := sha256.Sum256([]byte(code_verifier))

	encodeingString := base64.RawURLEncoding.EncodeToString(sum[:]) // 结果不填充=,参数长度固定为43
	log.Println(encodeingString)
	return encodeingString
}

func genScopeUrlEncode() string {
	scopes := strings.Split(scope, " ")
	encodeScopes := make([]string, 0)
	for i := 0; i < len(scopes); i++ {
		encodeScopes = append(encodeScopes, url.QueryEscape(scopes[i]))
	}
	return strings.Join(encodeScopes, " ")
}
