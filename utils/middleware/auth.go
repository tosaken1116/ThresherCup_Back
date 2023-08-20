package middleware

import (
	"net/http"
	"strings"
	"thresher/utils/errors"
	j "thresher/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
        token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized,errors.ErrorResponse{StatusText: "Unauthorized",Detail: "token is required"})
			ctx.Abort()
		}
		splitToken := strings.Split(token," ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer"{
			ctx.JSON(http.StatusUnauthorized,errors.ErrorResponse{StatusText: "Unauthorized",Detail: "invalid token format"})
			ctx.Abort()
		}
		t,err := j.ValidateJwt(splitToken[1])
		if err != nil{
			ctx.JSON(http.StatusUnauthorized,errors.ErrorResponse{StatusText: "Unauthorized",Detail: "cannot decode token"})
			ctx.Abort()
		}
		if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid{
			userId := (*claims)["sub"].(string)
			ctx.Set("userId", userId)
		}else{
			ctx.JSON(http.StatusUnauthorized,errors.ErrorResponse{StatusText: "Unauthorized",Detail: "invalid token format"})
			ctx.Abort()
		}


    }
}