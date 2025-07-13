package authentication

import (
	"net/http"
	"strings"

	"gostart/internal/auth/token"
	"gostart/internal/domain/authentication/entities"
	"gostart/internal/domain/authentication/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthenticationMiddleware struct {
	AuthenticationRepository repository.AuthenticationRepository
}

type contextKey string

const AuthUserContextKey = contextKey("authUser")

func SetAuthUser(context *gin.Context, authUser *entities.AuthUser) {
	context.Set(string(AuthUserContextKey), authUser)
}

func GetAuthUser(context *gin.Context) *entities.AuthUser {
	value, exists := context.Get(string(AuthUserContextKey))
	if !exists {
		// if we dont have request that has the value of an auth user (even anonymous) something is wrong
		// eg, bad actor call
		panic("missing user in request")
	}
	authUser, ok := value.(*entities.AuthUser)
	if !ok {
		panic("invalid auth user type in context")
	}
	return authUser
}

func (authenticationMiddleware *AuthenticationMiddleware) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Vary", "Authorization")
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			SetAuthUser(context, entities.AnonymousUser)
			context.Next()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			return
		}

		tokenString := headerParts[1]
		claims, err := token.VerifyJWTToken(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		authUserID, err := uuid.Parse(claims["sub"].(string))
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user id in token"})
			return
		}

		authUser, err := authenticationMiddleware.AuthenticationRepository.FindAuthUserByID(context.Request.Context(), authUserID)
		SetAuthUser(context, authUser)
		context.Next()
	}
}

func (authenticationMiddleware *AuthenticationMiddleware) RequireAuthUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := GetAuthUser(context)

		if user.IsAnonymous() {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "you must be logged in to access this route"})
			context.Abort()
			return
		}

		context.Next()
	}
}
