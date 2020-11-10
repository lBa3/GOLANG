package main

import("net/http");
import ("github.com/gorilla/mux");

type Route struct{
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true);

	for _,route := range routes{
		router.Name(route.Name).Methods(route.Method).Path(route.Pattern).Handler(route.HandleFunc)
	}
	return router;
}

var routes = Routes{
	Route{"Index", "GET", "/", Index,},
	Route{"MovieList", "GET", "/peliculas", MovieList,},
	Route{"MovieShow", "GET", "/pelicula/{id}", MovieShow,},
	Route{"MovieAdd", "POST", "/pelicula-add", MovieAdd,},
	Route{"MovieUpdate", "PUT", "/pelicula-edit/{id}", MovieUpdate,},
	Route{"MovieDelete", "DELETE", "/pelicula-delete/{id}", MovieDelete,},
}