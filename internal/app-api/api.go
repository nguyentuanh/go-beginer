package api

import (
	"unsafe"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"go-template/internal/app-api/handler/health"
	homehandler "go-template/internal/app-api/handler/home"
	postshandler "go-template/internal/app-api/handler/posts"
	userhandler "go-template/internal/app-api/handler/users"
	"go-template/internal/middleware"
	"go-template/pkg/l"
)

// Server ...
type Server struct {
	r *fiber.App
}

// HTTPErrorResponse ...
type HTTPErrorResponse struct {
	Status     string                 `json:"status"`
	Code       uint32                 `json:"code"`
	Message    string                 `json:"message"`
	DevMessage interface{}            `json:"dev_message" swaggerignore:"true"`
	Errors     map[string]interface{} `json:"errors"`
	RID        string                 `json:"rid"`
}

func getString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// New ...
func New(env string) *Server {
	r := fiber.New(fiber.Config{
		ErrorHandler:    customErrorHandler(env),
		WriteBufferSize: 15 * 4096,
		ReadBufferSize:  15 * 4096,
	})

	return &Server{
		r,
	}
}

// Listen ...
func (o Server) Listen(add string) error {
	return o.r.Listen(add)
}

// Middleware ...
func (o Server) Middleware(ll l.Logger) {
	o.r.Use(cors.New())
	o.r.Use(pprof.New())
	o.r.Use(requestid.New())
	o.r.Use(compress.New())
	o.r.Use(middleware.NewLogging(ll))
}

// InitHealth ...
func (o Server) InitHealth(healthHandler health.Controller) {
	o.r.Get("/health", healthHandler.Health)
	o.r.Get("/liveness", healthHandler.Liveness)
}

// InitMetrics ...
func (o Server) InitMetrics() {
	prometheus := fiberprometheus.New("go-template")
	prometheus.RegisterAt(o.r, "/metrics")
	o.r.Use(prometheus.Middleware)
}

// InitLogHandler ...
func (o Server) InitLogHandler() {
	o.r.Get("/log/level", adaptor.HTTPHandlerFunc(l.ServeHTTP))
	o.r.Put("/log/level", adaptor.HTTPHandlerFunc(l.ServeHTTP))
}

// InitHome ...
func (o Server) InitHome() {
	homeHandler := homehandler.New()
	w := o.r.Group("/home")
	w.Get("/index", homeHandler.Index)

}

// InitRouter
func (o Server) InitRouter() {
	// User router
	usersHandler := userhandler.New()
	o.r.Get("/users", usersHandler.GetAllUser)
	o.r.Post("/users", usersHandler.CreateUser)
	o.r.Get("/users/:id", usersHandler.GetUserById)
	o.r.Put("/users/:id", usersHandler.UpdateUser)
	o.r.Delete("/users/:id", usersHandler.DeleteUser)

	// Post router
	postsHandler := postshandler.New()
	o.r.Get("/users/:id/posts", postsHandler.GetPostsOfUser)
	o.r.Post("/users/:id/posts", postsHandler.CreatePost)
	o.r.Get("/users/:id/posts/:postId", postsHandler.GetPostDetail)
	o.r.Put("/users/:id/posts/:postId", postsHandler.UpdatePost)
	o.r.Delete("/users/:id/posts/:postId", postsHandler.DeletePost)
}
