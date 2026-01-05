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
	"log"
	"os"

	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/models"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	HttpVersion   uint16 `yaml:"httpVersion"`
	UseTLS        bool   `yaml:"useTLS"`
	Fqdn          string `yaml:"fqdn"`
	SbiPort       uint16 `yaml:"sbiPort"`
	OamPort       uint16 `yaml:"oamPort"`
	InitOnStartup bool   `yaml:"initOnStartup"`
	/* Custom configuration parameters */
	NetConfig *NetworkConfig `yaml:"simulationProfile"`
}

type NetworkConfig struct {
	//TODO add support for multiple snssai and dnns
	Snssai      models.Snssai `yaml:"slice" json:"slice"`
	Plmn        models.PlmnId `yaml:"plmn" json:"plmn"`
	Dnn         string        `yaml:"dnn" json:"dnn"`
	NumOfGnb    int           `yaml:"numOfgNB" json:"numOfgNB"`
	NumOfUe     int           `yaml:"numOfUe" json:"numOfUe"`
	ArrivalRate float32       `yaml:"arrivalRate" json:"arrivalRate"`
}

func InitConfig(configPath string) *AppConfig {

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("cannot read config file #%v ", err)
	}

	cfg := AppConfig{}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if cfg.InitOnStartup && cfg.NetConfig == nil {
		log.Fatalf("error: when initializing from startup, simulation profile must be defined in config file")
	}

	return &cfg
}

func (cfg *AppConfig) Dumps() string {
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(d)

}
