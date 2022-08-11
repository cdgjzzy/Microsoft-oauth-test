package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
)

var (
	code_verifier = "5d2309e5bb73b864f989753887fe52f79ce5270395e25862da6940d5" // 长度为43到128位的随机字符串
	
)

func genCodeChallenge() string {
	sum := sha256.Sum256([]byte(code_verifier))

	encodeingString := base64.RawURLEncoding.EncodeToString(sum[:]) // 结果不填充=,参数长度固定为43
	log.Println(encodeingString)
	return encodeingString
}

