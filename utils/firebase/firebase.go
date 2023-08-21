package firebase

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io"
	"net/http"
	"strings"
	"thresher/utils/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// reference to https://zenn.dev/takoyaki3/articles/a5f59a8c01d51a
type CustomClaims struct {
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Iss      string `json:"iss"`
	Aud      string `json:"aud"`
	AuthTime int64  `json:"auth_time"`
	UserId   string `json:"user_id"`
	Sub      string `json:"sub"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Email    string `json:"email"`
	jwt.StandardClaims
}



// JWTを検証する関数
func CheckFirebaseJWT(tokenString string)(*CustomClaims,error){

	// Googleの公開鍵を取得
	resp, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to get google public key","/utils/firebase/firebase")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to read google public key request","/utils/firebase/firebase")
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to marshal json","/utils/firebase/firebase")
	}

	// JWTのヘッダを解析し署名に用いられている鍵を取得
	parts := strings.Split(tokenString, ".")

	// decode the header
	headerJson, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to decode JWT header","/utils/firebase/firebase")
	}

	var header map[string]interface{}
	err = json.Unmarshal(headerJson, &header)
	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to unmarshall JWT header","/utils/firebase/firebase")
	}

	kid := header["kid"].(string)
	certString := result[kid].(string)
	block, _ := pem.Decode([]byte(certString))
	if block == nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to parse PEM block containing the public key","/utils/firebase/firebase")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to parse certificate","/utils/firebase/firebase")
	}

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	// 署名を検証
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return rsaPublicKey, nil
	})

	if err != nil {
		return nil,errors.New(http.StatusInternalServerError,"failed to parse token","/utils/firebase/firebase")
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if time.Unix(claims.Exp, 0).Before(time.Now()) {
			return nil,errors.New(http.StatusInternalServerError,"token is expired","/utils/firebase/firebase")
			} else {
			return claims, nil
		}
	} else {
		return nil,errors.New(http.StatusInternalServerError,"token is invalid","/utils/firebase/firebase")
	}
}