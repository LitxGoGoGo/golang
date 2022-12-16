package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type BackMyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	GenerateToken()
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFBQUFBIiwiZXhwIjoxNjcwMjM0OTM5LCJpc3MiOiJMYWEiLCJuYmYiOjE2NzAyMzQ4Nzl9.k9mj-EGV6kpIM5so704BjJPdSrFNYXH_Obmbu7UdS4M"
	ParseToken(s)
}

func GenerateToken() {
	c := MyClaims{
		Username: "Litx",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			Issuer:    "Laa",
			ExpiresAt: time.Now().Unix() + 60*60*2,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//加密盐
	mySigningKey := []byte("asdwqda")

	s, e := token.SignedString(mySigningKey)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(token)
	fmt.Println(s)

}

func ParseToken(s string) {
	//s可以给前端使用了
	var myBackClaims BackMyClaims

	token, err := jwt.ParseWithClaims(s, &myBackClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("hallo123"), nil
	}) //返回的这个盐要一致才能解析
	fmt.Println(token)
	if claims, ok := token.Claims.(*BackMyClaims); ok && token.Valid {
		fmt.Println(12313)
		fmt.Println(claims.Username, claims.ExpiresAt, claims.NotBefore)
	} else {
		fmt.Println(err)
	}
}
