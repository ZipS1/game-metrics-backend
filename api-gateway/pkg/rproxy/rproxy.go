package rproxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(remote)
		// proxy.Director = func(req *http.Request) {
		// 	req.Header = c.Request.Header
		// 	req.Host = remote.Host
		// 	req.URL.Scheme = remote.Scheme
		// 	req.URL.Host = remote.Host
		// 	req.URL.Path = c.Param("proxyPath")
		// }
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
