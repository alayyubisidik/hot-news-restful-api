package helper

import (
	"errors"
	"hot_news_2/model/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("mysecretkey")

type MyCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user domain.User) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Username,
		user.FullName,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	
	return ss, err
}

func VerifyToken(tokenString string) (*MyCustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("metode tanda tangan tidak valid")
        }
        return mySigningKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, errors.New("token tidak valid")
    }
}
