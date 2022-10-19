package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/alikan97/Go-GRPC/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetAllAssets(c *gin.Context) {
	resp, err := h.Client.GetAllAsset(c, &emptypb.Empty{})

	if err != nil {
		fmt.Printf("Fucked it %v", err)
	}

	respContent := resp.GetResponse()

	// Empty request - proto out
	c.JSON(http.StatusOK, gin.H{
		"All-Assets": respContent,
	})
}

func (h *Handler) getAssetBySymbol(c *gin.Context) {
	symbolName := c.Query("symbolname")

	response, err := h.Client.GetAsset(c, &pb.GetAssetReq{Symbol: symbolName})

	if err != nil {
		e := fmt.Sprintf("Failed to get asset %s, error: %v", symbolName, err)
		c.JSON(http.StatusNotFound, gin.H{
			"Error": e,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Asset": response,
	})
}

func (h *Handler) getRecentTrades(c *gin.Context) {
	since, _ := strconv.ParseUint(c.Query("since"), 0, 64)
	symbolpair := c.Query("symbolpair")

	resp, err := h.Client.GetRecentTrades(c, &pb.RecentTradesReq{Since: since, Symbol: symbolpair})

	if err != nil {
		e := fmt.Sprintf("Failed find recent trades for symbol: %s from %v. Reason: %v", symbolpair, since, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": e,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"RecentTrades": resp,
	})
}

func (h *Handler) getCurrentQuote(c *gin.Context) {
	symbolpair := c.Query("symbolpair")

	resp, err := h.Client.GetCurrentQuote(c, &pb.GetQuoteReq{Symbol: symbolpair})

	if err != nil {
		e := fmt.Sprintf("Failed to get quote for symbol: %s. Reason: %v", symbolpair, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": e,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Quote": resp,
	})
}
