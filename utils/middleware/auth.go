package middleware

import (
	e "errors"
	"log"
	"net/http"
	"strings"
	"thresher/utils/errors"
	"thresher/utils/firebase"

	"github.com/gin-gonic/gin"
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
		t,err := firebase.CheckFirebaseJWT(splitToken[1])
		if err != nil{
			RenderError(ctx,err)
			ctx.Abort()
		}
		ctx.Set("userId", t.UserId)
    }
}
func RenderError(ctx *gin.Context,err error) {
	if e, ok := err.(*errors.Error); ok {
		log.Println(err.Error())
		ctx.JSON(e.StatusCode, errors.ErrorResponse{StatusText: e.StatusText, Detail: e.Detail})
		return
	}
	log.Println(err.Error())
	ctx.JSON(http.StatusInternalServerError, e.New("unknown error"))
}
