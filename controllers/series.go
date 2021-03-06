package controllers

import (
	"github.com/UHERO/dvw-api/data"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetTrendAirseatData(module string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		freq, ok := getStrParam(r, "frequency")
		indicators, ok := getHandleList(r, "i_list")
		if !ok {
			// do something
		}
		markets, _ := getHandleList(r, "m_list")
		destinations, _ := getHandleList(r, "d_list")
		seriesRes, err := data.GetTrendAirseatData(module, freq, indicators, markets, destinations)
		if err != nil {
			// do something
		}
		SendResponseData(w, r, SeriesResource{Data: seriesRes})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetHotelData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		freq, ok := getStrParam(r, "frequency")
		indicators, ok := getHandleList(r, "i_list")
		if !ok {
			// do something
		}
		categories, _ := getHandleList(r, "c_list")
		seriesRes, err := data.GetHotelData("HOTEL", freq, indicators, categories)
		if err != nil {
			// do something
		}
		SendResponseData(w, r, SeriesResource{Data: seriesRes})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetCharacteristicData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		freq, ok := getStrParam(r, "frequency")
		indicators, ok := getHandleList(r, "i_list")
		if !ok {
			// do something
		}
		groups, _ := getHandleList(r, "g_list")
		seriesRes, err := data.GetCharacteristicData("CHAR", freq, indicators, groups)
		if err != nil {
			// do something
		}
		SendResponseData(w, r, SeriesResource{Data: seriesRes})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func GetExpenditureData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		freq, ok := getStrParam(r, "frequency")
		indicators, ok := getHandleList(r, "i_list")
		if !ok {
			// do something
		}
		groups, _ := getHandleList(r, "g_list")
		categories, _ := getHandleList(r, "c_list")
		seriesRes, err := data.GetExpenditureData("EXP", freq, indicators, groups, categories)
		if err != nil {
			// do something
		}
		SendResponseData(w, r, SeriesResource{Data: seriesRes})
	}
}
