package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"message-sender/internal/pkg/apperrors"
	"message-sender/internal/pkg/logger"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *apperrors.AppError
			var e *apperrors.AppError
			if errors.As(err, &e) {
				appErr = e
			}

			logger.Error("Request Error", appErr)

			c.JSON(appErr.Code, appErr.ToResponse())
			c.Abort()
		}
	}
}
