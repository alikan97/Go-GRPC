package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetAllAssets(c *gin.Context) {
	fmt.Printf("HITTTT")
	resp, err := h.Client.GetAllAsset(c, &emptypb.Empty{})
	if err != nil {
		fmt.Printf("Fucked it %v", err)
	}
	fmt.Printf("%v", resp)

	// Empty request - proto out
	c.JSON(http.StatusOK, gin.H{
		"req": "All assets",
	})
}
