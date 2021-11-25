package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resourcehealth/mgmt/2017-07-01/resourcehealth"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/mock"
)

type MockedResourceHealth struct {
	mock.Mock
}

type MockedResources struct {
	mock.Mock
}

func (mock *MockedResourceHealth) GetAvailabilityStatus(resourceURI string) (*resourcehealth.AvailabilityStatus, error) {
	args := mock.Called(resourceURI)
	return args.Get(0).(*resourcehealth.AvailabilityStatus), args.Error(1)
}

func (mock *MockedResourceHealth) GetAllAvailabilityStatuses() (*[]resourcehealth.AvailabilityStatus, error) {
	args := mock.Called()
	return args.Get(0).(*[]resourcehealth.AvailabilityStatus), args.Error(1)
}

func (mock *MockedResourceHealth) GetSubscriptionID() string {
	args := mock.Called()
	return args.Get(0).(string)
}

func (mock *MockedResourceHealth) GetLastRatelimitRemaining() string {
	args := mock.Called()
	return args.Get(0).(string)
}

func (mock *MockedResources) GetResources(resourceType string, resourceTags map[string]string) (*[]resources.GenericResource, error) {
	args := mock.Called(resourceType, resourceTags)
	return args.Get(0).(*[]resources.GenericResource), args.Error(1)
}

func CallExporter(collector ResourceHealthCollector) *httptest.ResponseRecorder {
	loadConfig("config/config_example.yml")
	req := httptest.NewRequest("GET", "/metrics", nil)
	rr := httptest.NewRecorder()
	registry := prometheus.NewRegistry()
	registry.MustRegister(&collector)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	handler.ServeHTTP(rr, req)
	return rr
}

func TestNewResourceHealthCollector_OK(t *testing.T) {
	session, err := NewAzureSession("subscriptionID")
	if err != nil {
		t.Errorf("Error occured %s", err)
	}
	_ = NewResourceHealthCollector(session)
}

func TestCollect_GetResources_Error(t *testing.T) {
	r := MockedResources{}
	rh := MockedResourceHealth{}
	collector := ResourceHealthCollector{
		resourceHealth: &rh,
		resources:      &r,
	}

	var asList []resourcehealth.AvailabilityStatus
	rh.On("GetAllAvailabilityStatuses").Return(&asList, nil)

	var resList []resources.GenericResource
	r.On("GetResources", mock.Anything, mock.Anything).Return(&resList, errors.New("Unit test Error"))

	rr := CallExporter(collector)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Wrong status code: got %v, want %v", status, http.StatusInternalServerError)
	}
}

func TestCollect_GetAllAvailabilityStatuses_Error(t *testing.T) {
	r := MockedResources{}
	rh := MockedResourceHealth{}
	collector := ResourceHealthCollector{
		resourceHealth: &rh,
		resources:      &r,
	}

	var resList []resources.GenericResource
	resourceID := "id"
	resourceType := "type"
	resList = append(resList, resources.GenericResource{
		ID:   &resourceID,
		Type: &resourceType,
	})
	r.On("GetResources", mock.Anything, mock.Anything).Return(&resList, nil)

	var asList []resourcehealth.AvailabilityStatus
	rh.On("GetAllAvailabilityStatuses", mock.Anything).Return(&asList, errors.New("Unit test Error"))

	rr := CallExporter(collector)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Wrong status code: got %v, want %v", status, http.StatusInternalServerError)
	}
}

func TestCollect_Collect_Ok(t *testing.T) {
	r := MockedResources{}
	rh := MockedResourceHealth{}
	collector := ResourceHealthCollector{
		resourceHealth: &rh,
		resources:      &r,
	}

	var resList []resources.GenericResource
	resourceID := "/subscriptions/my_subscription/resourceGroups/my_rg/providers/Microsoft.Compute/virtualMachines/my_instance"
	resourceType := "Microsoft.Compute/virtualMachines"
	resList = append(resList, resources.GenericResource{
		ID:   &resourceID,
		Type: &resourceType,
	})
	r.On("GetResources", "Microsoft.Compute/virtualMachines", mock.Anything).Return(&resList, nil)
	var emptyList []resources.GenericResource
	r.On("GetResources", "Microsoft.Web/serverfarms", mock.Anything).Return(&emptyList, nil)
	r.On("GetResources", "Microsoft.Web/sites", mock.Anything).Return(&emptyList, nil)

	var asList []resourcehealth.AvailabilityStatus
	asID1 := "/subscriptions/my_subscription/resourcegroups/my_rg/providers/microsoft.compute/virtualmachines/my_instance" + AvailabilityStatusIDSuffix
	asList = append(asList, resourcehealth.AvailabilityStatus{
		ID: &asID1,
		Properties: &resourcehealth.AvailabilityStatusProperties{
			AvailabilityState: resourcehealth.Unavailable,
		},
	})
	asID2 := "/subscriptions/my_subscription/resourcegroups/my_rg/providers/microsoft.compute/virtualmachines/my_instance2" + AvailabilityStatusIDSuffix
	asList = append(asList, resourcehealth.AvailabilityStatus{
		ID: &asID2,
		Properties: &resourcehealth.AvailabilityStatusProperties{
			AvailabilityState: resourcehealth.Unavailable,
		},
	})
	rh.On("GetAllAvailabilityStatuses", mock.Anything).Return(&asList, nil)
	rh.On("GetSubscriptionID").Return("my_subscription")
	rh.On("GetLastRatelimitRemaining").Return("99")

	rr := CallExporter(collector)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, want %v", status, http.StatusOK)
	}

	want := `# HELP azure_resource_health_availability_up Resource health availability that relies on signals from different Azure services to assess whether a resource is healthy
# TYPE azure_resource_health_availability_up gauge
azure_resource_health_availability_up{resource_group="my_rg",resource_name="my_instance",resource_type="Microsoft.Compute/virtualMachines",subscription_id="my_subscription"} 0
# HELP azure_resource_health_ratelimit_remaining_requests Azure subscription scoped Resource Health requests remaining (based on X-Ms-Ratelimit-Remaining-Subscription-Resource-Requests header)
# TYPE azure_resource_health_ratelimit_remaining_requests gauge
azure_resource_health_ratelimit_remaining_requests{subscription_id="my_subscription"} 99
# HELP azure_tag_info Tags of the Azure resource
# TYPE azure_tag_info gauge
azure_tag_info{resource_group="my_rg",resource_name="my_instance",resource_type="Microsoft.Compute/virtualMachines",subscription_id="my_subscription"} 1
`
	if rr.Body.String() != want {
		t.Errorf("Unexpected body: got %v, want %v", rr.Body.String(), want)
	}
}
