package handlers

import (
	"log"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// helper function to return false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	// bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
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
			return false
		}
		fallBack := model.NewInternal()
		c.JSON(fallBack.Status(), gin.H{"error": fallBack})
		return false
	}
	return true
}
