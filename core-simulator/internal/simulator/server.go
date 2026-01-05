// Copyright 2025 EURECOM
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
//
// Contributors:
//   Giulio CAROTA
//   Thomas DU
//   Adlen KSENTINI

package simulator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *CoreSimulatorApp) handleInitSimulation(w http.ResponseWriter, r *http.Request) {

	config := &NetworkConfig{}

	if app.config.NetConfig != nil {
		// avoid cli config to override the default one
		config = app.config.NetConfig
	} else {
		if r.Body != nil {
			err := json.NewDecoder(r.Body).Decode(config)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "Missing request body", http.StatusBadRequest)
		}
	}

	err := app.InitNewSimulation(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(SimulationStatusResponse{
		Status: app.status,
	})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}

}

func (app *CoreSimulatorApp) handleStartSimulation(w http.ResponseWriter, r *http.Request) {
	err := app.StartSimulation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(SimulationStatusResponse{
		Status: app.status,
	})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}

func (app *CoreSimulatorApp) handleStatusSimulation(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(SimulationStatusResponse{
		Status: app.GetCurrentSimulationStatus(),
	})
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}

func (app *CoreSimulatorApp) handleStopSimulation(w http.ResponseWriter, r *http.Request) {
	err := app.StopSimulation()
	if err != nil {
		http.Error(w, "could not stop simulation", http.StatusInternalServerError)
		return
	}
}

func (app *CoreSimulatorApp) startHttpServer() {
	app.wg.Add(1)

	router := mux.NewRouter()

	router.HandleFunc("/core-simulator/v1/configure", app.handleInitSimulation)
	router.HandleFunc("/core-simulator/v1/start", app.handleStartSimulation)
	router.HandleFunc("/core-simulator/v1/status", app.handleStatusSimulation)
	router.HandleFunc("/core-simulator/v1/stop", app.handleStopSimulation)

	app.server = &http.Server{Addr: fmt.Sprintf(":%d", app.config.OamPort), Handler: router}

	go func() {
		defer func() {
			_ = recover()
			app.wg.Done()
		}()

		log.Printf("serving simulation api on :8081")
		// always returns error. ErrServerClosed on graceful close
		if err := app.server.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

}

func (app *CoreSimulatorApp) stopHttpServer() {
	if app.server != nil {
		err := app.server.Close()
		if err != nil {
			log.Default().Printf("could not stop nbi server")
		}
	}

}
