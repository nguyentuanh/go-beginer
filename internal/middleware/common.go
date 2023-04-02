package middleware

import (
	"sync"
	"unsafe"

	"go-template/internal/constant"
	"go-template/pkg/l"
	tracer "go-template/pkg/trace"

	"github.com/gofiber/fiber/v2"
)

func getString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// NewLogging creates a new middleware handler
func NewLogging(ll l.Logger) fiber.Handler {
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		once.Do(func() {
			errHandler = c.App().Config().ErrorHandler
		})

		rid := getString(c.Context().Response.Header.Peek(fiber.HeaderXRequestID))
		// role := getString(c.Context().Request.Header.Peek(constant.XGapoRoleKey))
		// apikey := getString(c.Context().Request.Header.Peek(constant.XGapoAPIKey))
		// userID := getString(c.Context().Request.Header.Peek(constant.XGapoUserIDKey))
		// wsID := getString(c.Context().Request.Header.Peek(constant.XGapoWorkspaceIDKey))
		// lang := getString(c.Context().Request.Header.Peek(constant.XGapoLang))

		ctx := c.UserContext()

		xctx, span := tracer.StartSpan(ctx, c.Path())
		defer span.End()
		tracer.SetAttribute(span, "requestID", rid)
		// tracer.SetAttribute(span, "headers.role", role)
		// apiKeyLen := len(apikey)
		// if apiKeyLen > 3 {
		//     apikey = apikey[0:apiKeyLen-3] + "***"
		// }
		// tracer.SetAttribute(span, "headers.apikey", apikey)
		// tracer.SetAttribute(span, "headers.userID", userID)
		// tracer.SetAttribute(span, "headers.wsID", wsID)
		// tracer.SetAttribute(span, "headers.lang", lang)

		c.Locals(constant.SpanKey, xctx)
		c.SetUserContext(xctx)

		defer func() {
			tracer.SetAttribute(span, "code", c.Response().StatusCode())
		}()

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		return nil
	}
}
