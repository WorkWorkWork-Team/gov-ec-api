package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/WorkWorkWork-Team/gov-ec-api/handler"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	TestMiddleWareFunction gin.HandlerFunc
	TestApiKey             string = "key"
)

func NewGinTestContext() (*gin.Context, *httptest.ResponseRecorder, *gin.Engine) {
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	return c, res, r
}

var _ = Describe("Middleware", Label("unit"), func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		// Create MiddleWare for testing
		TestMiddleWareFunction = handler.ValidateAPIKey(TestApiKey)
	})
	Context("Authorization MiddleWare", func() {
		When("api key is invalid", func() {
			It("should reject request", func() {
				// Create GinContext containing request and data we want.
				c, _, _ := NewGinTestContext()
				c.Request = &http.Request{
					Header: http.Header{
						"Authorization": {fmt.Sprint("Bearer IMSURELYINVALID")},
					},
				}
				// Call MiddleWare
				TestMiddleWareFunction(c)
				// Expect data from context
				Expect(c.Writer.Status()).To(Equal(http.StatusUnauthorized))
			})
		})
		When("api key is valid", func() {
			It("should accept request", func() {
				// Create GinContext containing request and data we want.
				c, _, _ := NewGinTestContext()
				c.Request = &http.Request{
					Header: http.Header{
						"Authorization": {fmt.Sprint("Bearer ", TestApiKey)},
					},
				}
				// Call MiddleWare
				TestMiddleWareFunction(c)
				// Expect data from context
				Expect(c.Param("valid_test")).To(Equal("true"))
			})
		})
	})
})
