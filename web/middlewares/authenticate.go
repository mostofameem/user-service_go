package middlewares

import (
	"base_service/config"
	"base_service/web/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

const (
	XAPIKEY = "x-api-key"
)

// Define a custom type for the context key
func GenerateToken(usr map[string]any) (string, string, error) {
	conf := config.GetConfig()
	expirationTime := time.Now().Add(5 * time.Minute).Unix()

	accessToken, err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"Id":  usr["id"],
			"exp": expirationTime,
		},
	).SignedString([]byte(conf.JwtSecret))
	if err != nil {
		log.Println(err.Error())
		return "", "", fmt.Errorf("error")
	}

	Time := time.Now().Add(7 * 24 * time.Hour).Unix()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"Id":  usr["id"],
			"exp": Time,
		},
	)

	refreshToken, err := token.SignedString([]byte(conf.JwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecret), nil
	})
}

func GenerateAccessTokenFromRefreshToken(claims jwt.Claims) (string, error) {
	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(config.GetConfig().JwtSecret))
	return accessTokenString, err
}

func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.SendError(w, http.StatusForbidden, fmt.Errorf("authorization header is missing"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, _ := ParseToken(tokenString)
		if token.Valid {
			next.ServeHTTP(w, r) // Token is valid, continue with the request
			return
		}
		// Token has expired, check for refresh token
		refreshHeader := r.Header.Get("Refresh-Token")
		if refreshHeader == "" {
			utils.SendError(w, http.StatusUnauthorized, fmt.Errorf("refresh token missing"))
			return
		}

		// Validate and parse the refresh token
		refreshString := strings.TrimPrefix(refreshHeader, "Bearer ")
		refreshToken, err := ParseToken(refreshString)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, fmt.Errorf("invalid refresh token: %v", err))
			return
		}

		if refreshToken.Valid {
			// Generate new access token
			claims := token.Claims.(jwt.MapClaims)
			claims["exp"] = time.Now().Add(1 * time.Minute).Unix()
			newToken, err := GenerateAccessTokenFromRefreshToken(claims)
			if err != nil {
				utils.SendError(w, http.StatusInternalServerError, fmt.Errorf("error generating new token: %v", err))
				return
			}
			log.Println(newToken)
			// Continue with the request
			next.ServeHTTP(w, r)
			return
		}

		utils.SendError(w, http.StatusUnauthorized, fmt.Errorf("refresh token invalid"))
	})
}

func GetUserIDFromToken(tokenStr string) (int, error) {
	conf := config.GetConfig()

	// Parse JWT
	var claims AuthClaims
	_, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.JwtSecret), nil
		},
	)
	if err != nil {
		return 0, err
	}

	// Return user ID from claims
	return claims.Id, nil
}

func unauthorizedResponse(w http.ResponseWriter) {
	utils.SendError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.GetConfig()

		if !conf.ApiKeyEnabled {
			next.ServeHTTP(w, r)
			return
		}

		// collect apiKey from header
		header := r.Header.Get(XAPIKEY)

		if header != conf.ApiKey {
			unauthorizedResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
