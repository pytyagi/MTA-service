package api_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mta-hosting-optimizer/api"
	"mta-hosting-optimizer/mocks"
	service "mta-hosting-optimizer/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestListingHandler(t *testing.T) {

	MockIpDataService := mocks.IpDataServiceInterface{}
	apiresp := make([]service.IpConfig, 0)
	apiresp = append(apiresp, service.IpConfig{
		Id:       1,
		Hostname: "mta-prod-test-5",
		Ip:       "127.0.0.1",
		Active:   true,
	}, service.IpConfig{
		Id:       2,
		Hostname: "mta-prod-test-6",
		Ip:       "127.0.0.2",
		Active:   true,
	}, service.IpConfig{
		Id:       3,
		Hostname: "mta-prod-test-3",
		Ip:       "127.0.0.3",
		Active:   true,
	})

	serverList := make([]string, 0)
	serverList = append(serverList, "mta-prod-5", "mta-prod-6", "mta-prod-3")
	MockIpDataService.Mock.On("GetServiceIpData").Return(&apiresp, nil)
	MockIpDataService.Mock.On("GetInefficientServers", &apiresp).Return(serverList, nil)
	req := httptest.NewRequest(http.MethodGet, "/hosts", nil)
	rr := httptest.NewRecorder()

	h := api.ListingHandler{IpDataService: &MockIpDataService}
	h.ListInefficientServersHandler(rr, req)

	expectedResponse := make([]string, 0)
	expectedResponse = append(expectedResponse, "mta-prod-5", "mta-prod-6", "mta-prod-3")

	res := rr.Result()
	defer res.Body.Close()
	respdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	actualResponse := make([]string, 0)
	_ = json.Unmarshal(respdata, &actualResponse)
	assert.DeepEqual(t, expectedResponse, actualResponse)
}

func TestListingHandlerError(t *testing.T) {

	MockIpDataService := mocks.IpDataServiceInterface{}
	apiresp := make([]service.IpConfig, 0)
	apiresp = append(apiresp, service.IpConfig{
		Id:       1,
		Hostname: "mta-prod-test-5",
		Ip:       "127.0.0.1",
		Active:   true,
	}, service.IpConfig{
		Id:       2,
		Hostname: "mta-prod-test-6",
		Ip:       "127.0.0.2",
		Active:   true,
	}, service.IpConfig{
		Id:       3,
		Hostname: "mta-prod-test-3",
		Ip:       "127.0.0.3",
		Active:   true,
	})

	serverList := make([]string, 0)
	serverList = append(serverList, "mta-prod-5", "mta-prod-6", "mta-prod-3")
	MockIpDataService.Mock.On("GetServiceIpData").Return(nil, errors.New("error in service"))
	MockIpDataService.Mock.On("GetInefficientServers", &apiresp).Return(serverList, nil)
	req := httptest.NewRequest(http.MethodGet, "/hosts", nil)
	rr := httptest.NewRecorder()

	h := api.ListingHandler{IpDataService: &MockIpDataService}
	h.ListInefficientServersHandler(rr, req)

	expectedResponse := make([]string, 0)
	expectedResponse = append(expectedResponse, "mta-prod-5", "mta-prod-6", "mta-prod-3")

	res := rr.Result()
	defer res.Body.Close()
	respdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	actualResponse := make([]string, 0)
	_ = json.Unmarshal(respdata, &actualResponse)
	fmt.Println(string(respdata))
	assert.Error(t, fmt.Errorf(string(respdata)), "Unable to get data from mock service error in service")
}
