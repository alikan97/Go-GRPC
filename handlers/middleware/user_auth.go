package middleware

import (
	"strings"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

/*
	User Authentication Middleware for securing private endpoints
*/

func AuthUser(s model.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				err := model.NewBadRequest("Invalid request parameters. See invalidArgs")

				c.JSON(err.Status(), gin.H{
					"error":       err,
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return
			}

			// otherwise error type is unknown
			err := model.NewInternal()
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		idTokenheader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenheader) < 2 {
			err := model.NewAuthorization("Must provide Authorzation header")

			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		user, err := s.ValidateIDToken(idTokenheader[1])
		if err != nil {
			err := model.NewAuthorization("Provided token is invalid")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
