package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.NoRoute(func(c *gin.Context) {
		target := "localhost:2121"
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.Transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Second,
				DualStack: true,
			}).DialContext,
			// MaxIdleConns:          100,
			// IdleConnTimeout: 5 * time.Second,
			// TLSHandshakeTimeout:   10 * time.Second,
			// ExpectContinueTimeout: 1 * time.Second,
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	})
	r.Run(":2020")
}
