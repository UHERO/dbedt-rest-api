package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var cache *data.Cache

func GetDimensionAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetDimensionByHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetDimensionKidsByHandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func CreateCache(pool *redis.Pool, ttl int) *data.Cache {
	cache = &data.Cache{Pool: pool, TTL: 60 * ttl} // TTL in seconds
	return cache
}

func CheckCache(c *data.Cache) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		url := GetFullRelativeURL(r)
		cachedVal, _ := c.GetCache(url)
		if cachedVal != nil {
			WriteResponse(w, cachedVal)
			return
		}
		next(w, r)
		return
	}
}

func WriteResponse(w http.ResponseWriter, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(payload)
	if err != nil {
		log.Printf("Response write FAILURE")
	}
}

func WriteCache(r *http.Request, c *data.Cache, payload []byte) {
	url := GetFullRelativeURL(r)
	err := c.SetCache(url, payload)
	if err != nil {
		log.Printf("Cache store FAILURE: %s", url)
		return
	}
}

func GetFullRelativeURL(r *http.Request) string {
	path := r.URL.Path
	if r.URL.RawQuery == "" {
		return path
	}
	return path + "?" + r.URL.RawQuery
}

func getIntParam(r *http.Request, name string) (intval int64, ok bool) {
	ok = true
	param, ok := mux.Vars(r)[name]
	if !ok {
		return
	}
	intval, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		ok = false
	}
	return
}

func getStrParam(r *http.Request, name string) (strval string, ok bool) {
	strval, ok = mux.Vars(r)[name]
	return
}
