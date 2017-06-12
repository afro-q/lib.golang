package routing

import (
	"net/http"
	
	"github.com/quinlanmorake/lib.golang/endpoints"
)

type RequestHandler struct { }

var requestMap map[string]http.HandlerFunc

func SetupRoutes(routingMap map[string]http.HandlerFunc) {
	requestMap = routingMap	
}

func setCorsHeaders(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	responseWriter.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")   
}

func setResponseHeaderToJson(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}

func WrapMiddlewareAndHandle(responseWriter http.ResponseWriter, request *http.Request, handler http.HandlerFunc) {
	setResponseHeaderToJson(responseWriter)
	
	requestMap[request.URL.String()](responseWriter, request)
}

func (*RequestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	setCorsHeaders(responseWriter)

	if (request.Method == "OPTIONS") {
		return
	}
	
	if handler, handlerFound := requestMap[request.URL.String()]; handlerFound {
		handler(responseWriter, request)
		return
	}

	routeNotFound := endpoints.InvalidRouteResult(request.URL.String())
	endpoints.WriteResponse(responseWriter, routeNotFound)
}