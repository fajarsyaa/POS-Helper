package helper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecretKey = []byte("slash-fajar-test")

type CustomClaims struct {
	UserID int
	Email  string
	Role   string
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, email, role string) (string, error) {
	expirationTime := time.Now().Add(3 * time.Hour)

	claims := &CustomClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func TokenAuthMiddleware(requiredRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			err := gin.H{"error": "Authorization Required"}
			response := ResponseMessage("Failed", "Unauthorized", http.StatusUnauthorized, err)
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			err := gin.H{"error": "Token is required"}
			response := ResponseMessage("Failed", "Unauthorized", http.StatusUnauthorized, err)
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			err := gin.H{"error": "Invalid or expired token"}
			response := ResponseMessage("Failed", "Unauthorized", http.StatusUnauthorized, err)
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx.Set("userID", claims["UserID"])

			userRole, ok := claims["Role"].(string)
			if !ok {
				err := gin.H{"error": "Role not found in token"}
				response := ResponseMessage("Failed", "Unauthorized", http.StatusUnauthorized, err)
				ctx.JSON(http.StatusUnauthorized, response)
				ctx.Abort()
				return
			}

			roleAllowed := false
			for _, role := range requiredRoles {
				if userRole == role {
					roleAllowed = true
					break
				}
			}

			if !roleAllowed {
				err := gin.H{"error": "Insufficient permissions"}
				response := ResponseMessage("Failed", "Forbidden", http.StatusForbidden, err)
				ctx.JSON(http.StatusForbidden, response)
				ctx.Abort()
				return
			}
		} else {
			err := gin.H{"error": "Invalid or expired token"}
			response := ResponseMessage("Failed", "Unauthorized", http.StatusUnauthorized, err)
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
