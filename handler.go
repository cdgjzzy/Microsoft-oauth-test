package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, OAuthEndpoint)
}

func code(ctx *gin.Context) {
	log.Println("code callback server is being requested.")

	errMsg, errExist := ctx.GetQuery("error")
	if errExist {
		log.Println("code callback error: ", errMsg)
		errDes, errExist := ctx.GetQuery("error_description")
		if errExist {
			log.Fatalln("code callback error_description: ", errDes)
		} else {
			log.Fatal()
		}
	}

	code, codeExist := ctx.GetQuery("code")
	if !codeExist {
		log.Fatalln("code not exist.")
	}
	log.Printf("code: %v", code)

	err := getToken(ctx, fmt.Sprintf("%v", code))
	if err != nil {
		log.Fatalln("get token err: ", err)
	}

	if tokenResponse != "" {
		ctx.Redirect(302, "http://localhost:5001/display")
	}

}

func display(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tokenResponse)
}
