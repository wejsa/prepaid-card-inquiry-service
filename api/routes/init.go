package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/serverkona1/prepaid-card-inquiry-service/pkg"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	initPingPongApi(r)
	return r
}

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	// Creates a router without any middleware by default
	//r := gin.New()
	r := gin.Default()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(pkg.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v3 := r.Group("/api/v3")
	{
		for _, route := range routes {
			v3.Handle(route.Method, route.Path, route.HandlerFunc)
		}
	}

	return r
}

var routes = Routes{
	Route{
		Name:        "PointParsInquiryasdfsd",
		Method:      "POST",
		Path:        "/point/pars/inquiry",
		HandlerFunc: PointParsInquiry,
	},
}
