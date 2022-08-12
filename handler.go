package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirect(ctx *gin.Context) {
	scopeQuery, exist := ctx.GetQuery("scope")
	if !exist {
		ctx.JSON(http.StatusBadRequest, "scope is required.")
	}
	scope = scopeQuery
	OAuthEndpoint := fmt.Sprintf(authUrlFormat, client_id, redirect_uri, genScopeUrlEncode(), state, code_challenge)
	ctx.Redirect(http.StatusFound, OAuthEndpoint)
}

func code(ctx *gin.Context) {
	log.Println("code callback server is being requested.")

	errMsg, errExist := ctx.GetQuery("error")
	if errExist {
		log.Println("code callback error: ", errMsg)
		errDes, errExist := ctx.GetQuery("error_description")
		if errExist {
			log.Println("code callback error_description: ", errDes)
			tokenResponse = errDes
		} else {
			tokenResponse = "unkown error."
		}
		ctx.Redirect(302, "http://localhost:5001/display")
		return
	}

	code, codeExist := ctx.GetQuery("code")
	if !codeExist {
		log.Println("code not exist.")
		tokenResponse = "code not exist."
		ctx.Redirect(302, "http://localhost:5001/display")
		return
	}
	log.Printf("code: %v", code)

	err := getToken(ctx, fmt.Sprintf("%v", code))
	if err != nil {
		log.Println("get token err: ", err)
	}

	if tokenResponse != "" {
		ctx.Redirect(302, "http://localhost:5001/display")
	}

}

func display(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tokenResponse)
}
