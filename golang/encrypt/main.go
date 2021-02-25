package main

import (
	"fmt"

	"encrypt/aes"
)

//$ echo "aaa" | openssl aes-256-cbc -e -pbkdf2 -pass pass:123456 -a
//U2FsdGVkX19F2XvW2xLy507tkHzILmTh0156RIMv46Q=
//$ echo "U2FsdGVkX19F2XvW2xLy507tkHzILmTh0156RIMv46Q=" | openssl aes-256-cbc -d -pbkdf2 -pass pass:123456 -a
//aaa

func main() {
	//str := "U2FsdGVkX1/Hdcmtgk68P6Y9oN6c6J/Prsz0JDZxdM0="
	str := "ATODdTwYNyrFTT054j2PZQ=="
	//var key = []byte("1234567812345678")

	en, err := aes.DeAesCode2Base64(str)
	if err != nil {
		fmt.Println("====>>>  err: ", err)
		return
	}
	fmt.Println(string(en))
}
