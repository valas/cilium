// Copyright 2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/cilium/cilium/api/v1/models"
)

// startKubeProxyHealthHTTPService registers a handler function for the /healthz
// status HTTP endpoint exposed on addr. This endpoint equivalent to one exposed
// by kubeproxy here:
// https://github.com/kubernetes/kubernetes/blob/master/pkg/proxy/healthcheck/proxier_health.go
func (d *Daemon) startKubeProxyHealthHTTPService(addr string) {
	lc := net.ListenConfig{Control: setsockoptReuseAddrAndPort}
	// Specifying "tcp" will attempt to open both IPv4 and IPv6 sockets.
	ln, err := lc.Listen(context.Background(), "tcp", addr)
	if err != nil {
		log.WithError(err).Fatalf(
			"Unable to listen on %s port for pretend kube-proxy healthz", addr)
	}

	mux := http.NewServeMux()
	mux.Handle("/healthz", healthzHandler{d: d})
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		err := srv.Serve(ln)
		if errors.Is(err, http.ErrServerClosed) {
			log.Info("pretend kube-proxy healthz status server shutdown")
		} else if err != nil {
			log.WithError(err).Fatal("Unable to start pretend kube-proxy healthz server")
		}
	}()
	log.Infof("Started pretend kube-proxy healthz server on address %s", addr)
}

type healthzHandler struct {
	d *Daemon
}

func (h healthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isUnhealthy := func(sr *models.StatusResponse) bool {
		if sr.Cilium != nil {
			state := sr.Cilium.State
			return state != models.StatusStateOk && state != models.StatusStateDisabled
		}
		return false
	}

	statusCode := http.StatusOK
	// We piggy back here on Cilium daemon health. If Cilium is healthy, we can
	// reasonably assume that the node networking is ready.
	sr := h.d.getStatus(true)
	if isUnhealthy(&sr) {
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
}
