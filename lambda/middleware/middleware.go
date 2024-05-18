package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWTMiddleware(next func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		tokenString, err := extractTokenStringFromHeaders(request.Headers)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       "Invalid Authorization header",
				StatusCode: http.StatusBadRequest,
			}, err
		}

		if tokenString == "" {
			return events.APIGatewayProxyResponse{
				Body:       "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		claims, err := parseTokenString(tokenString)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		expires := claims["expires"].(float64)
		if int64(expires) < time.Now().Unix() {
			return events.APIGatewayProxyResponse{
				Body:       "Token expired",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}

		return next(request)
	}
}

func extractTokenStringFromHeaders(headers map[string]string) (string, error) {
	authorization, ok := headers["Authorization"]

	if !ok {
		return "", nil
	}

	authorizationSplit := strings.Split(authorization, "Bearer ")

	if len(authorizationSplit) != 2 {
		return "", errors.New("Invalid Authorization header")
	}

	return authorizationSplit[1], nil
}

func parseTokenString(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Invalid claims")
	}

	return claims, nil

}
