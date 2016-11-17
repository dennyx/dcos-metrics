// Copyright 2016 Mesosphere, Inc.
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

package http

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dcos/dcos-metrics/producers"
	"github.com/gorilla/mux"
)

// /api/v0/node
func nodeHandler(p *producerImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var am []interface{}
		nodeMetrics, err := p.store.GetByRegex(producers.NodeMetricPrefix + ".*")
		if err != nil {
			httpLog.Error("/api/v0/node - %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for _, v := range nodeMetrics {
			am = append(am, v)
		}

		if len(am) != 0 {
			encode(am[0], w)
			return
		}

		httpLog.Error("/api/v0/node - no content in store.")
		http.Error(w, "No values found in store", http.StatusBadRequest)
	}
}

// /api/v0/containers
func containersHandler(p *producerImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cm := []string{}
		containerMetrics, err := p.store.GetByRegex(producers.ContainerMetricPrefix + ".*")
		if err != nil {
			httpLog.Error("/api/v0/containers - %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for _, c := range containerMetrics {
			if _, ok := c.(producers.MetricsMessage); !ok {
				httpLog.Error("/api/v0/contianers - unsupported message type.")
				http.Error(w, "Got unsupported message type.", http.StatusInternalServerError)
			}
			cm = append(cm, c.(producers.MetricsMessage).Dimensions.ContainerID)
		}

		encode(cm, w)
	}
}

// /api/v0/containers/{id}
func containerHandler(p *producerImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := strings.Join([]string{
			producers.ContainerMetricPrefix, vars["id"],
		}, producers.MetricNamespaceSep)

		containerMetrics, ok := p.store.Get(key)
		if !ok {
			httpLog.Error("/api/v0/containers/{id} - not found in store: %s", key)
			http.Error(w, "Key not found in store", http.StatusNoContent)
		}

		encode(containerMetrics, w)
	}
}

// /api/v0/containers/{id}/app/
func containerAppHandler(p *producerImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cid := vars["id"]
		key := strings.Join([]string{
			producers.AppMetricPrefix, cid,
		}, producers.MetricNamespaceSep)

		containerMetrics, ok := p.store.Get(key)
		if !ok {
			httpLog.Error("/api/v0/containers/{id}/app - not found in store: %s", key)
			http.Error(w, "Key not found in store", http.StatusNoContent)
		}

		encode(containerMetrics, w)
	}
}

func pingHandler(p *producerImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type ping struct {
			OK        bool   `json:"ok"`
			Timestamp string `json:"timestamp"`
		}

		encode(ping{OK: true, Timestamp: time.Now().UTC().Format(time.RFC3339)}, w)
	}
}

func notYetImplementedHandler(p *producerImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type fooData struct {
			Message string `json:"message"`
		}
		result := fooData{Message: "Not Yet Implemented"}
		encode(result, w)
	}
}

// -- helpers

func encode(v interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		httpLog.Error("Failed to encode value to JSON: %v", v)
		http.Error(w, "Failed to encode value to JSON", http.StatusInternalServerError)
	}
}
