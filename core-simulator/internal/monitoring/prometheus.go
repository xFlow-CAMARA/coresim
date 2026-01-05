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

package monitoring

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	UEsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ue_total",
			Help: "Total number of UEs by state",
		},
		[]string{"simulationId", "state"},
	)

	PduSessionsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pdu_sessions_total",
			Help: "Number of active PDU sessions",
		},
		[]string{"simulationId"},
	)

	TrafficBytes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ue_traffic_bytes_total",
			Help: "Total traffic bytes by direction",
		},
		[]string{"simulationId", "ueId", "direction"},
	)

	TrafficPackets = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ue_traffic_packets_total",
			Help: "Total traffic bytes by direction",
		},
		[]string{"simulationId", "ueId", "direction"},
	)
	TotalTraffic = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_traffic_bytes_total",
			Help: "Total traffic bytes by direction",
		},
		[]string{"simulationId", "direction"},
	)
	UEIPInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ue_ip_info",
			Help: "UE metadata mapping IMSI to IP address",
		},
		[]string{"simulationId", "imsi", "ip"},
	)
)

func init() {
	prometheus.MustRegister(UEsTotal, PduSessionsTotal, TrafficBytes, TrafficPackets, TotalTraffic, UEIPInfo)
	//prometheus.MustRegister(TotalTraffic)
}

func StartMetricsServer() {
	log.Printf("starting prometheus metrics server on :9090")
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9090", nil) // metrics on :9090
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not start metrics server: %s", err.Error())
		}
	}()
}
