package api

import (
	"encoding/json"
	service "mta-hosting-optimizer/service"
	"net/http"
)

type ListingHandler struct {
	IpDataService service.IpDataServiceInterface
}

func NewListingHandler(IpDataService service.IpDataServiceInterface) ListingHandler {
	return ListingHandler{
		IpDataService: IpDataService,
	}
}

type ListingResponse struct {
	InefficientServers []string
}

// ListInefficientServersHandler godoc
// @Summary ListInefficientServersHandler
// @Description List inefficient servers
// @Accept  json
// @Produce  json
// @Success 200 {object} ListingResponse
// @Failure 500
// @Router /hosts [get]

func (l *ListingHandler) ListInefficientServersHandler(w http.ResponseWriter, r *http.Request) {

	// Get Data from mock service
	ipData, err := l.IpDataService.GetServiceIpData()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Unable to get data from mock service " + err.Error()))
		return
	}

	// Get the GetInefficientServers from list of servers
	inefficientHosts := l.IpDataService.GetInefficientServers(ipData)
	w.Header().Set("Content-Type", "application/json")
	resp := ListingResponse{
		InefficientServers: inefficientHosts,
	}
	json.NewEncoder(w).Encode(resp.InefficientServers)
}
