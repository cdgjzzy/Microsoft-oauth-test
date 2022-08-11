package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getToken(ctx *gin.Context, code string) (err error) {

	log.Println("get token request")

	bodyFormat := "client_id=%s&code=%s&redirect_uri=%s&grant_type=authorization_code&code_verifier=%s"
	var body = strings.NewReader(fmt.Sprintf(bodyFormat, client_id, code, redirect_uri, code_verifier))
	client := &http.Client{}
	req, err := http.NewRequest("POST", tokenUri, body)
	req.Header.Add("Origin", "*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response error code:%v\nresponse body:%v", response.StatusCode, respBody)
	}
	fmt.Println("get token resp:", string(respBody))
	tokenResponse = string(respBody)
	return nil
}

func refreshToken(refresh_token string) error {

	log.Println("refresh token request")

	bodyFormat := "client_id=%s&refresh_token=%s&grant_type=refresh_token&client_secret=%s"
	var body = strings.NewReader(fmt.Sprintf(bodyFormat, client_id, refresh_token, client_secret))
	response, err := http.Post(tokenUri, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response error code:%v", response.StatusCode)
	}
	fmt.Println("get token resp:", string(respBody))
	tokenResponse = string(respBody)
	return nil
}
