package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang-boilerplate/models"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)


func CreateTokens(userID int64) (*models.TokenDetails, error) {
	var err error
	tokenDetails := &models.TokenDetails{
		AccessUUID: uuid.NewV4().String(),
		RefreshUUID: uuid.NewV4().String(),
		AccessTokenExpiration: time.Now().Add(time.Minute * 15).Unix(),
		RefreshTokenExpiration: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	tokenDetails.AccessToken, err = createAccessToken(true, tokenDetails.AccessUUID, userID, tokenDetails.AccessTokenExpiration)
	if err != nil {
		return nil, err
	}

	tokenDetails.RefreshToken, err = createRefreshToken(tokenDetails.RefreshUUID, userID, tokenDetails.RefreshTokenExpiration)
	if err != nil {
		return nil, err
	}
	return tokenDetails, nil


}

func RetrieveTokenAccessDetails(r *http.Request) (*models.AccessDetails, error) {
	token, err := GetTokenFromRequest(r)

	if err != nil { return nil, err }
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok { return nil, err}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil { return nil, err }

		return &models.AccessDetails{
			AccessUUID: accessUUID,
			UserID: userID,
		}, nil
	}
	return nil, err
}

func GetTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	tokenString := retrieveToken(r)

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// before we run the parse we make sure the signing algo is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func IsTokenValid(r *http.Request) bool {
	token, err := GetTokenFromRequest(r)
	if err != nil { return false }

	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

func retrieveToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	tokenArray := strings.Split(bearerToken, " ")
	if len(tokenArray) == 2 {
		return tokenArray[1]
	} else {
		return ""
	}
}

func createAccessToken(isAuthorized bool, uuid string, userID int64, expiration int64) (string, error) {
	claims := jwt.MapClaims{
		"authorized": isAuthorized,
		"access_uuid": uuid,
		"user_id": userID,
		"expires_at": expiration,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token
	signedToken, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func createRefreshToken(refreshUUID string, userID int64, expiration int64) (string, error) {
	claims := jwt.MapClaims{
		"refresh_uuid": refreshUUID,
		"user_id": userID,
		"expires_at": expiration,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token
	signedToken, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}
