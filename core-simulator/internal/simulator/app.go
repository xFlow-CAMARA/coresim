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
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/monitoring"
)

/* Simulation Controller code */

type SimulationStatus string

const (
	CONFIGURED SimulationStatus = "CONFIGURED"
	STARTED    SimulationStatus = "STARTED"
	STOPPED    SimulationStatus = "STOPPED"
	ERROR      SimulationStatus = "ERROR"
)

type SimulationStatusResponse struct {
	Status SimulationStatus
}

type CoreSimulatorApp struct {
	currentInstance *NetworkInstance
	status          SimulationStatus
	instanceMutex   sync.RWMutex
	server          *http.Server
	wg              sync.WaitGroup
	ctx             context.Context
	config          *AppConfig
}

func NewCoreSimulatorApp(configPath string) *CoreSimulatorApp {
	return &CoreSimulatorApp{
		currentInstance: nil,
		status:          STOPPED,
		instanceMutex:   sync.RWMutex{},
		wg:              sync.WaitGroup{},
		config:          InitConfig(configPath),
	}
}

func (app *CoreSimulatorApp) InitNewSimulation(config *NetworkConfig) error {

	if config == nil {
		return fmt.Errorf("no configuration provided, could not initialize")
	}

	app.instanceMutex.Lock()
	defer app.instanceMutex.Unlock()

	if app.currentInstance != nil {
		return fmt.Errorf("could not initialize the simulation instance, please stop or reset the current instance")
	}

	app.currentInstance = NewNetworkInstance(app.config.SbiPort, config)
	if app.currentInstance == nil {
		return fmt.Errorf("could not initialize the simulation instance")
	}

	err := app.currentInstance.InitNetworkInstance()
	if err != nil {
		return fmt.Errorf("could not initialize the simulation instance")
	}

	app.status = CONFIGURED
	return nil
}

func (app *CoreSimulatorApp) StartSimulation() error {
	app.instanceMutex.Lock()
	defer app.instanceMutex.Unlock()

	if app.currentInstance == nil {
		return fmt.Errorf("please configure the simulation via /configure")
	}

	// If already started, it's a restart - stop first
	if app.status == STARTED {
		if err := app.currentInstance.Stop(); err != nil {
			log.Printf("Warning: error stopping instance for restart: %s", err.Error())
		}
	}

	if err := app.currentInstance.Start(); err != nil {
		app.status = ERROR
		return fmt.Errorf("could not start the simulation instance")
	}

	app.status = STARTED
	return nil
}

func (app *CoreSimulatorApp) GetCurrentSimulationStatus() SimulationStatus {
	app.instanceMutex.Lock()
	defer app.instanceMutex.Unlock()

	return app.status
}

func (app *CoreSimulatorApp) StopSimulation() error {
	app.instanceMutex.Lock()
	defer app.instanceMutex.Unlock()

	if app.status == STOPPED || app.currentInstance == nil {
		return fmt.Errorf("no running instance")
	}

	if app.status == STARTED {
		if err := app.currentInstance.Stop(); err != nil {
			return fmt.Errorf("could not stop the simulation instance")
		}
	}

	// Don't set currentInstance to nil - keep it so we can restart
	app.status = STOPPED
	return nil
}

func (app *CoreSimulatorApp) Run() {

	var cancel context.CancelFunc
	app.ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	go app.listenShutdownEvent()
	log.Printf("running config: \n%s", app.config.Dumps())

	if app.config.InitOnStartup {
		log.Printf("bootstraping simulation instance")
		err := app.InitNewSimulation(app.config.NetConfig)
		if err != nil {
			log.Fatalf("could not initialize the simulator on startup")
		}
	}

	app.startHttpServer()
	monitoring.StartMetricsServer()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Printf("terminating...")

	cancel()
	app.wg.Wait()
}

func (app *CoreSimulatorApp) listenShutdownEvent() {
	defer func() {
		_ = recover()
		app.wg.Done()
	}()

	<-app.ctx.Done()
	app.stopHttpServer()
}
