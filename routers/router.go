package routers

import (
	"github.com/UHERO/dvw-api/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func CreateRouter(apiName string) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	prefix := "/" + apiName

	apiRouter := mux.NewRouter().StrictSlash(false).PathPrefix(prefix).Subrouter()
	SetRoutes(apiRouter)
	router.PathPrefix(prefix).Handler(negroni.New(
		negroni.HandlerFunc(controllers.CORSOptionsHandler),
		//negroni.HandlerFunc(controllers.ValidApiKey(applicationRepository)),
		negroni.HandlerFunc(controllers.CheckCache()),
		negroni.Wrap(apiRouter),
	))
	return router
}

func SetRoutes(r *mux.Router) {
	r.HandleFunc("/dimensions/{module:[a-z]+}",           controllers.GetModuleDimensions()).Methods("GET")
	r.HandleFunc("/{dimension:[a-z]+}/all",               controllers.GetDimensionAll()).Methods("GET")
	r.HandleFunc("/{dimension:[a-z]+}/{handle}/children", controllers.GetDimensionKidsByHandle()).Methods("GET")
	r.HandleFunc("/{dimension:[a-z]+}/{handle}",          controllers.GetDimensionByHandle()).Methods("GET")
}
