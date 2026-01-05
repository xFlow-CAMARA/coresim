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
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/components/core"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/components/ran"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/components/utils"
)

/* Network Instance Code*/

type NetworkInstance struct {
	// list of UEs
	ctx          context.Context
	UeList       map[string]*ran.Ue
	ueListMutex  sync.RWMutex
	Amf          *core.Amf
	Smf          *core.Smf
	Pcf          *core.Pcf
	config       *NetworkConfig
	ueGenContext context.Context
	ueGenCancel  context.CancelFunc
	ipam         *utils.IPAllocator
	sbiPort      uint16
	simId        string
	GnbList      []string
}

func NewNetworkInstance(sbiPort uint16, config *NetworkConfig) *NetworkInstance {
	return &NetworkInstance{
		ctx:          context.Background(),
		UeList:       make(map[string]*ran.Ue),
		ueListMutex:  sync.RWMutex{},
		config:       config,
		Amf:          nil,
		Smf:          nil,
		ueGenContext: nil,
		ueGenCancel:  nil,
		ipam:         nil,
		sbiPort:      sbiPort,
		simId:        uuid.NewString(),
	}
}

func (n *NetworkInstance) InitNetworkInstance() error {
	// initialize the ipam with a default subnet, in the future
	// this should be configurable and per slice as well
	n.ipam = utils.NewIpamService("12.1.0.0", "16")

	n.Amf = core.NewAmf(n.config.Plmn)
	n.Smf = core.NewSmf(n.config.Plmn, n.ipam)
	n.Pcf = core.NewPcf(n.config.Plmn, n.ipam)

	n.Amf.InitAmf()
	n.Smf.InitSmf()
	n.Pcf.InitPcf()

	//spawn the gNBs
	n.GnbList = generateNRCellIDsHex(uint64(n.config.NumOfGnb))

	/* enable corenetwork network service based interface */
	r := mux.NewRouter()

	// register amf events api
	n.Amf.RegisterNorthboundAPIs(r)
	// register smf events api
	n.Smf.RegisterNorthboundAPIs(r)
	// register pcf policy authorization api
	n.Pcf.RegisterNorthboundAPIs(r)

	go func() {
		h2server := &http2.Server{}
		h2chandler := h2c.NewHandler(r, h2server)

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", n.sbiPort),
			Handler: h2chandler,
		}

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not start 3GPP sbi server: %s", err.Error())
		}
	}()

	return nil

}

func (n *NetworkInstance) Start() error {
	n.ueGenContext, n.ueGenCancel = context.WithCancel(n.ctx)
	log.Printf("starting simulation %s", n.simId)

	go func() {
		select {
		case <-n.ueGenContext.Done():
			return
		default:
			for i := 0; i < n.config.NumOfUe; i++ {
				// generate a new UE with a unique IMSI

				arrTime := expRand(float64(n.config.ArrivalRate))
				time.Sleep(arrTime)

				imsi := fmt.Sprintf("%s%s00000%05d", n.config.Plmn.Mcc, n.config.Plmn.Mnc, i+1)
				imei := generateIMEI()
				// Generate unique MSISDN per UE based on index
				msisdn := fmt.Sprintf("+336%09d", 100000000+i)

				ue := ran.NewUserEquipment(n.ctx, ran.UeConfig{
					Imsi:   imsi,
					Msidn:  msisdn,
					Imei:   imei,
					Dnn:    n.config.Dnn,
					Snssai: n.config.Snssai,
					Type:   "Smartphone",
					Plmn:   n.config.Plmn,
				}, n.ipam, n.simId, n.GnbList)

				// if ue is not nil then start the UE and add it to the list
				if ue != nil {
					n.ueListMutex.Lock()
					ue.PowerUp()
					n.UeList[imsi] = ue
					n.ueListMutex.Unlock()
				}
			}
		}
	}()
	return nil
}

func (n *NetworkInstance) Stop() error {
	// stop the generation of UEs
	n.ueGenCancel()

	n.ueListMutex.Lock()
	defer n.ueListMutex.Unlock()
	// this function should always overtake the start function
	for imsi, ue := range n.UeList {
		//gracefully turn off the UE
		ue.TurnOff(true)
		delete(n.UeList, imsi)
	}

	return nil
}

func generateIMEI() string {

	// Generate TAC (Type Allocation Code) - 8 digits
	tac := fmt.Sprintf("%08d", rand.Intn(100000000))

	// Generate SNR (Serial Number) - 6 digits
	snr := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Concatenate TAC + SNR (14 digits so far)
	imei14 := tac + snr

	// Compute check digit using Luhn algorithm
	checkDigit := luhnCheckDigit(imei14)

	return imei14 + checkDigit
}

// luhnCheckDigit computes the Luhn check digit for a given numeric string.
func luhnCheckDigit(number string) string {
	sum := 0
	alt := true // start doubling from the rightmost digit
	for i := len(number) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(number[i]))
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	checkDigit := (10 - (sum % 10)) % 10
	return strconv.Itoa(checkDigit)
}

// exponential random variable with mean 1/λ
func expRand(lambda float64) time.Duration {
	u := rand.Float64()
	return time.Duration(-math.Log(1-u) / lambda * float64(time.Second))
}

func generateNRCellIDsHex(max uint64) []string {
	// NCI is 36 bits max (values: 0 .. 2^36-1)
	const maxNCI = (1 << 36) - 1

	if max > maxNCI {
		max = maxNCI
	}

	ncis := make([]string, 0, max)
	for i := uint64(0); i < max; i++ {
		// %09x ensures 9 hex digits, lower-case, zero-padded
		ncis = append(ncis, fmt.Sprintf("%09x", i))
	}
	return ncis
}
