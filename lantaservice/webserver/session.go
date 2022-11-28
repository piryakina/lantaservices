package webserver

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type SessionContext struct {
	ID    int64
	Login string
}

type SessionManager struct {
	SessionKey  string
	MaxLifetime time.Duration
}

// Claims for auth
type Claims struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

// DeleteSession delete auth cookies
func DeleteSession(w http.ResponseWriter) {
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now(), Path: "/"}
	http.SetCookie(w, &deleteCookie)
}

// GetSession returning a user id as the string in cookies or returning error
func GetSession(r *http.Request) (SessionContext, error) {
	// if no Auth cookie is set then return a 404 not found page
	sess := SessionContext{}
	cookie, err := r.Cookie("Auth")
	if err != nil {
		return sess, err
	}
	// Return a token using the cookie
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return SigningKey, nil
	})
	if err != nil {
		return sess, err
	}
	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// fmt.Println("claim", claims.Username)
		return SessionContext{
			ID:    claims.ID,
			Login: claims.Login,
		}, nil
	}
	return sess, err
}

// SetSession set the token in the cookie if the password is correct
func SetSession(w http.ResponseWriter, user SessionContext) error {
	mySigningKey := SigningKey
	// expire the token and cookie
	expireToken := time.Now().Add(24 * time.Hour).Unix()
	expireCookie := time.Now().Add(24 * time.Hour)
	claims := Claims{
		user.ID,
		user.Login,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}
	cookie := http.Cookie{Name: "Auth", Value: ss, Expires: expireCookie, HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)
	return nil
}

// SigningKey salt
var SigningKey = []byte("salt")
