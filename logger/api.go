package logger

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GinLog struct {
	ElapsedTime    int64
	ClientIp       string
	RequestUrl     string
	LoggerUserId   string `json:",omitempty"`
	RequestBody    string
	ResponseStatus int
	ResponseBody   string
	Header         map[string]string
}

type GinLogStart struct {
	StartTime  time.Time
	RequestUrl string
}

type ResponseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r ResponseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func GinLogger(c *gin.Context) {

	startTime := time.Now()

	w := &ResponseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w

	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := io.ReadAll(tee)
	c.Request.Body = io.NopCloser(&buf)

	var ginLogStart GinLogStart
	ginLogStart.RequestUrl = c.Request.Host + c.Request.URL.Path
	ginLogStart.StartTime = time.Now()
	if !strings.Contains(ginLogStart.RequestUrl, "kibana") {
		InfoLogWithType(c, ginLogStart, "GINSTART")
	}

	var ginLog GinLog

	ginLog.RequestBody = refineGinRequestBody(string(body))
	ginLog.ClientIp = c.ClientIP()
	ginLog.RequestUrl = c.Request.Host + c.Request.URL.Path
	ginLog.Header = make(map[string]string)
	for i, j := range c.Request.Header {
		if i == "X-Admin-Token" || i == "X-Login-Token" || i == "X-Apollo-Token" || i == "X-Internal-Token" || i == "X-Web-Token" {
			continue
		}
		ginLog.Header[i] = j[0]
	}
	c.Next()
	userId, ok := c.Get("userid")
	if ok {
		ginLog.LoggerUserId = userId.(string)
	}
	ginLog.ElapsedTime = time.Since(startTime).Milliseconds()
	ginLog.ResponseStatus = c.Writer.Status()
	ginLog.ResponseBody = refineGinResponseBody(c.Request.Host+c.Request.URL.Path, filterLogBody(w.body.String()))
	size := len(c.Request.Header["Content-Type"])
	if size == 1 {
		if !(strings.Contains(c.Request.Header["Content-Type"][0], "application/json") || strings.Contains(c.Request.Header["Content-Type"][0], "application/x-www-form-urlencoded")) {
			ginLog.RequestBody = c.Request.Header["Content-Type"][0]
		}
	}

	header := c.Writer.Header().Get("Content-Type")
	if !strings.Contains(header, "application/json") {
		ginLog.ResponseBody = filterLogResponseBody(w.body.String())
	}

	if strings.Contains(ginLog.RequestUrl, "kibana") {
		return
	}

	switch {
	case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
		{
			ErrorLogWithType(c, ginLog, "GINERROR")
		}
	case c.Writer.Status() >= http.StatusInternalServerError:
		{
			ErrorLogWithType(c, ginLog, "GINERROR")
		}
	default:
		select {
		case <-c.Request.Context().Done():
			InfoLogWithType(c, ginLog, "GINCANCEL")
		default:
			InfoLogWithType(c, ginLog, "GIN")
		}
	}
}

func refineGinResponseBody(url string, body string) string {
	if strings.Contains(url, "kibana") {
		return "kibana search"
	}
	if strings.Contains(url, "invoice") {
		return "invoice download"
	}
	return body
}

func refineGinRequestBody(body string) string {
	return body
}
