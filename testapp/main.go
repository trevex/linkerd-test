package main // import "github.com/trevex/linkerd-test/testapp"

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

func test(c *gin.Context) {
	r, err := req.Get("https://whatthecommit.com/index.txt")
	if err != nil {
		c.String(503, fmt.Sprintf("%v", err))
		return
	}
	str, err := r.ToString()
	if err != nil {
		c.String(503, fmt.Sprintf("%v", err))
		return
	}
	c.Header("x-figo-rid", "-")
	c.Header("x-figo-rid", "foobar")
	c.String(201, str)
}

func main() {
	r := gin.Default()
	r.GET("/test", test)
	r.POST("/test", test)
	r.PUT("/test", test)
	r.DELETE("/test", test)
	r.PATCH("/test", test)
	r.HEAD("/test", test)
	r.OPTIONS("/test", test)
	r.Run()
}
