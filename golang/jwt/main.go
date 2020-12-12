package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {
	// 编码
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	expireTime := time.Now().Add(2 * time.Second)

	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 秒
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	fmt.Printf("ss: %s  || err: %v\n", ss, err)

	//	解码
	tokenString := ss

	deToken, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
	)

	if deClaims, ok := deToken.Claims.(*MyCustomClaims); ok && deToken.Valid {
		fmt.Printf("Foo: %v %v\n", deClaims.Foo, deClaims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println("err: ", err)
	}
}
