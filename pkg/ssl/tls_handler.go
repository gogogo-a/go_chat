package ssl

import (
	"github.com/gin-gonic/gin"
	// "github.com/unrolled/secure"
	// "kama_chat_server/pkg/zlog"
	// "strconv"
)

func TlsHandler(host string, port int) gin.HandlerFunc {
	// HTTPS 重定向已禁用，需要时取消下面注释
	// return func(c *gin.Context) {
	// 	secureMiddleware := secure.New(secure.Options{
	// 		SSLRedirect: true,
	// 		SSLHost:     host + ":" + strconv.Itoa(port),
	// 	})
	// 	err := secureMiddleware.Process(c.Writer, c.Request)

	// 	// If there was an error, do not continue.
	// 	if err != nil {
	// 		zlog.Fatal(err.Error())
	// 		return
	// 	}

	// 	c.Next()
	// }
	
	// 当前返回空中间件，不做 HTTPS 重定向
	return func(c *gin.Context) {
		c.Next()
	}
}
