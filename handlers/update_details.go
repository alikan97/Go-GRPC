package handlers

import (
	"log"
	"net/http"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/gin-gonic/gin"
)

type detailsRequest struct {
	Name    string `json:"name" binding:"omitempty,max=40"`
	Email   string `json:"email" binding:"omitempty,email"`
	Website string `json:"website" binding:"omitempty,url"`
}

func (h *Handler) UpdateDetails(c *gin.Context) {
	user := c.MustGet("user").(*model.User) // cast var to a model.User type
	var req detailsRequest

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		UID:     user.UID,
		Name:    req.Name,
		Email:   req.Email,
		Website: req.Website,
	}

	ctx := c.Request.Context()
	err := h.UserService.UpdateDetails(ctx, u)

	if err != nil {
		log.Printf("Update user dfailed: %v\n", err)

		c.JSON(model.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
