package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/alikan97/Go-GRPC/model"
	"github.com/gin-gonic/gin"
)

type timeoutWriter struct {
	gin.ResponseWriter
	h    http.Header
	wbuf bytes.Buffer

	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
	code        int
}

/*
	Timeout Middleware for wrapping endpoints within a timed environment
*/

func Timeout(timeout time.Duration, errTimeout *model.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		tw := &timeoutWriter{ResponseWriter: c.Writer, h: make(http.Header)}
		c.Writer = tw

		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		finished := make(chan struct{})        // to indicate handler finished
		panicChan := make(chan interface{}, 1) // used to handle panics if we cant recover

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			c.Next()
			finished <- struct{}{}
		}()

		select {
		case <-panicChan:
			// Something messed up
			e := model.NewInternal()
			tw.ResponseWriter.WriteHeader(e.Status())
			eResp, _ := json.Marshal(gin.H{
				"error": e,
			})
			tw.ResponseWriter.Write(eResp)
			// It went well
		case <-finished:
			tw.mu.Lock()
			defer tw.mu.Unlock()

			dst := tw.ResponseWriter.Header()
			for k, vv := range tw.Header() {
				dst[k] = vv
			}

			tw.ResponseWriter.WriteHeader(tw.code)
			tw.ResponseWriter.Write(tw.wbuf.Bytes())
			// Timeed out
		case <-ctx.Done():
			tw.mu.Lock()
			defer tw.mu.Unlock()

			tw.ResponseWriter.Header().Set("Content-Type", "application/json")
			tw.ResponseWriter.WriteHeader(errTimeout.Status())

			eResp, _ := json.Marshal(gin.H{
				"error": errTimeout,
			})
			tw.ResponseWriter.Write(eResp)
			c.Abort()
			tw.SetTimedOut()
		}
	}
}

func (tw *timeoutWriter) Write(b []byte) (int, error) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	if tw.timedOut {
		return 0, nil
	}

	return tw.wbuf.Write(b)
}

// In http.ResponseWriter interface
func (tw *timeoutWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	tw.mu.Lock()
	defer tw.mu.Unlock()
	// We do not write the header if we've timed out or written the header
	if tw.timedOut || tw.wroteHeader {
		return
	}
	tw.writeHeader(code)
}

// set that the header has been written
func (tw *timeoutWriter) writeHeader(code int) {
	tw.wroteHeader = true
	tw.code = code
}

// Header "relays" the header, h, set in struct
// In http.ResponseWriter interface
func (tw *timeoutWriter) Header() http.Header {
	return tw.h
}

// SetTimeOut sets timedOut field to true
func (tw *timeoutWriter) SetTimedOut() {
	tw.timedOut = true
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}
