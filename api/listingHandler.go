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
func (l *ListingHandler) ListInefficientServersHandler(w http.ResponseWriter, r *http.Request) {

	ipData, err := l.IpDataService.GetServiceIpData()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Unable to get data from mock service " + err.Error()))
		return
	}

	inefficientHosts := l.IpDataService.GetInefficientServers(ipData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inefficientHosts)
}
