package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/cilium/cilium/api/v1/models"
	. "gopkg.in/check.v1"
)

// 'check' testing suite scaffolding.
type KubeProxyHealthTestSuite struct{}

var _ = Suite(&KubeProxyHealthTestSuite{})

// Injected fake service.
type FakeService struct {
	injectedCurrentTs    time.Time
	injectedLastUpdateTs time.Time
}

func (s *FakeService) GetCurrentTs() time.Time {
	return s.injectedCurrentTs
}

func (s *FakeService) GetLastUpdateTs() time.Time {
	return s.injectedLastUpdateTs
}

// Injected fake daemon.
type FakeDaemon struct {
	injectedStatusResponse models.StatusResponse
}

func (d *FakeDaemon) getStatus(blah bool) models.StatusResponse {
	return d.injectedStatusResponse
}

type healthzPayload struct {
	LastUpdated string
	CurrentTime string
}

func (s *KubeProxyHealthTestSuite) TestKubeProxyHealth(c *C) {
	lastUpdateTs := time.Unix(100, 0) // Fake 100 seconds after Unix.
	currentTs := time.Unix(200, 0)    // Fake 200 seconds after Unix.
	expectedTs := currentTs
	s.healthTestHelper(c, models.StatusStateOk, currentTs,
		lastUpdateTs, expectedTs, http.StatusOK)
	expectedTs = lastUpdateTs
	s.healthTestHelper(c, models.StatusStateWarning, currentTs,
		lastUpdateTs, expectedTs, http.StatusInternalServerError)
	s.healthTestHelper(c, models.StatusStateFailure, currentTs,
		lastUpdateTs, expectedTs, http.StatusInternalServerError)
	s.healthTestHelper(c, models.StatusStateDisabled, currentTs,
		lastUpdateTs, expectedTs, http.StatusInternalServerError)
}

func (s *KubeProxyHealthTestSuite) healthTestHelper(c *C, ciliumStatus string,
	currentTs time.Time, lastUpdateTs time.Time, expectedLastUpdateTs time.Time,
	expectedHttpStatus int) {

	// Create handler with injected behavior.
	h := healthzHandler{
		d: &FakeDaemon{injectedStatusResponse: models.StatusResponse{
			Cilium: &models.Status{State: ciliumStatus}}},
		svc: &FakeService{
			injectedCurrentTs:    currentTs,
			injectedLastUpdateTs: lastUpdateTs}}

	// Create a new request.
	req, err := http.NewRequest("GET", "/healthz", nil)
	c.Assert(err, IsNil)
	w := httptest.NewRecorder()

	// Serve.
	h.ServeHTTP(w, req)

	// Main return code meets expectations.
	c.Assert(w.Code, Equals, expectedHttpStatus,
		Commentf("expected status code %v, got %v", expectedHttpStatus, w.Code))

	// Timestamps meet expectations.
	var payload healthzPayload
	c.Assert(json.Unmarshal(w.Body.Bytes(), &payload), IsNil)
	layout := "2006-01-02 15:04:05 -0700 MST"
	lastUpdateTs, err = time.Parse(layout, payload.LastUpdated)
	currentTs, err = time.Parse(layout, payload.CurrentTime)
	c.Assert(lastUpdateTs.Equal(expectedLastUpdateTs), Equals, true)
}
