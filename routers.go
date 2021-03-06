/*
 * Recommendation API
 *
 * This is a recommendation API using k-means Clustering
 *
 * API version: 1.0.0
 * Contact: capela.nuno@ua.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/Nuno19/Recomendation_Service/1.0.0/",
		Index,
	},

	Route{
		"LoadItem",
		strings.ToUpper("Post"),
		"/Nuno19/Recomendation_Service/1.0.0/loadItem",
		LoadItem,
	},

	Route{
		"LoadItemList",
		strings.ToUpper("Post"),
		"/Nuno19/Recomendation_Service/1.0.0/loadItemList",
		LoadItemList,
	},

	Route{
		"SetClusterCount",
		strings.ToUpper("Post"),
		"/Nuno19/Recomendation_Service/1.0.0/setClusterNumber",
		SetClusterCount,
	},

	Route{
		"GetRecommended",
		strings.ToUpper("Get"),
		"/Nuno19/Recomendation_Service/1.0.0/getRecommended",
		GetRecommended,
	},
	Route{
		"GetTextRecommended",
		strings.ToUpper("Get"),
		"/Nuno19/Recomendation_Service/1.0.0/GetTextRecommended",
		GetTextRecommended,
	},
}
