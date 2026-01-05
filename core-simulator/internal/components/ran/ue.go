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

package ran

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/giuliocarot0/gitc"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/components/utils"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/models"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/monitoring"
	"gitlab.eurecom.fr/open-exposure/coresim/core-simulator/internal/trafficgen"
)

// A Ue represents a User Equipment (UE) in the context of the core network simulator.
// It detains all the information concerning is current stauts and configuration.

type Ue struct {
	ctx       context.Context
	cancelFun context.CancelFunc

	// identifiers
	Imsi  string
	Msidn string
	Imei  string

	// status variables
	ueState          models.UeState
	PlmnId           models.PlmnId
	CurrentCellId    string
	PduSessions      map[int32]models.PduSessionInfo
	RmStatus         models.RmState
	CmStatus         models.CmState
	UserLocation     models.UserLocation
	HasUplinkData    bool
	LastActivityTime time.Time
	statusMutex      sync.RWMutex
	accessType       models.AccessType
	ipManager        *utils.IPAllocator
	defautlDnn       string
	defaultSnssai    models.Snssai

	//stats variables
	UpStats    map[int32]*models.UpStats
	statsMutex sync.RWMutex
	//
	// configuration variables
	UpStatsReport    bool
	IncactivityTimer time.Duration

	//simulation variables
	Profile string
	simId   string
	gnbList []string
}

type UeConfig struct {
	Imsi   string
	Msidn  string
	Imei   string
	Dnn    string
	Snssai models.Snssai
	Type   string
	Plmn   models.PlmnId
}

// NewUserEquipement creates a Ue instance with the provided configuration
// It takes a UeConfig type as input
// It retruns a pointer to the initialized Ue instance
func NewUserEquipment(ctx context.Context, cfg UeConfig, ipManager *utils.IPAllocator, simulationId string, gnbs []string) *Ue {
	ueCtx, ueCancelFunc := context.WithCancel(ctx)

	return &Ue{
		ctx:              ueCtx,
		cancelFun:        ueCancelFunc,
		PlmnId:           cfg.Plmn,
		Imsi:             cfg.Imsi,
		Msidn:            cfg.Msidn,
		Imei:             cfg.Imei,
		Profile:          cfg.Type,
		defautlDnn:       cfg.Dnn,
		defaultSnssai:    cfg.Snssai,
		RmStatus:         models.RmStateDeregistered,
		CmStatus:         models.CmStateIdle,
		PduSessions:      make(map[int32]models.PduSessionInfo),
		UpStats:          make(map[int32]*models.UpStats),
		statusMutex:      sync.RWMutex{},
		statsMutex:       sync.RWMutex{},
		HasUplinkData:    false,
		LastActivityTime: time.Now(),
		// set ACCESS TYPE to 3GPP by default
		accessType: models.ACCESSTYPE__3_GPP_ACCESS,
		ipManager:  ipManager,
		simId:      simulationId,
		gnbList:    gnbs,
	}
}

// Register sets the registration status of the UE to registered.
// It triggers the registration event which is logged by the simulation core.
func (ue *Ue) Register() {
	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()

	ue.CurrentCellId = ue.pickRandomNRCellID()

	ue.RmStatus = models.RmStateRegistered
	log.Printf("[%s] successfully registered to the network, cellId: %s", ue.Imsi, ue.CurrentCellId)

	/*prepare gitc message for AMF*/
	msg := &models.UeToAmfMsg{
		EventType:     models.AMFEVENTTYPEANYOF_REGISTRATION_STATE_REPORT,
		TimeStamp:     time.Now(),
		RmState:       ue.RmStatus,
		CmState:       ue.CmStatus,
		Supi:          ue.Imsi,
		Gpsi:          ue.Msidn,
		PlmnId:        ue.PlmnId,
		CurrentCellId: ue.CurrentCellId,
		AccessType:    ue.accessType,
	}
	if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
		log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
	}

	/*prepare gitc message for AMF*/
	msg2 := &models.UeToAmfMsg{
		EventType:     models.AMFEVENTTYPEANYOF_LOCATION_REPORT,
		TimeStamp:     time.Now(),
		RmState:       ue.RmStatus,
		CmState:       ue.CmStatus,
		Supi:          ue.Imsi,
		Gpsi:          ue.Msidn,
		PlmnId:        ue.PlmnId,
		CurrentCellId: ue.CurrentCellId,
		AccessType:    ue.accessType,
	}
	if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg2); err != nil {
		log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
	}

	monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.RmStateRegistered)).Inc()
	monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.RmStateDeregistered)).Dec()

}

// Attach enables the UE signaling connection with the AMF on the first registration
// It triggers to connectivity report event which is logged by the simulation core.
// It checks if the UE is registered before allowing the attach.
// If the UE is not registered, it logs an error message and does not proceed with the attach.
// It also starts the inactivity monitor to manage the UE's activity status.
func (ue *Ue) Attach(inactivityTimer time.Duration) {
	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()

	if ue.RmStatus != models.RmStateRegistered {
		log.Printf("UE %s is not registered to the network, cannot attach", ue.Imsi)
		return
	}
	ue.CmStatus = models.CmStateConnected
	log.Printf("[%s] successfully attached to the network", ue.Imsi)

	/*prepare gitc message for AMF*/
	msg := &models.UeToAmfMsg{
		EventType:     models.AMFEVENTTYPEANYOF_CONNECTIVITY_STATE_REPORT,
		TimeStamp:     time.Now(),
		RmState:       ue.RmStatus,
		CmState:       ue.CmStatus,
		Supi:          ue.Imsi,
		Gpsi:          ue.Msidn,
		PlmnId:        ue.PlmnId,
		CurrentCellId: ue.CurrentCellId,
		AccessType:    ue.accessType,
	}
	if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
		log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
	}

	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateConnected)).Inc()
	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateIdle)).Dec()

	/*start here the Inactivity Timer*/
	go ue.inactivityMonitor(inactivityTimer)
}

// It kills the UE RF.
// To trigger the loss of connectivity event, lossOfConnection must be set to true.
func (ue *Ue) LossOfConnection(isGracefully bool) {

	log.Printf("[%s] connection lost", ue.Imsi)

	/*prepare gitc message for AMF*/
	msg := &models.UeToAmfMsg{
		EventType:     models.AMFEVENTTYPEANYOF_LOSS_OF_CONNECTIVITY,
		TimeStamp:     time.Now(),
		RmState:       ue.RmStatus,
		CmState:       ue.CmStatus,
		Supi:          ue.Imsi,
		Gpsi:          ue.Msidn,
		PlmnId:        ue.PlmnId,
		CurrentCellId: ue.CurrentCellId,
		AccessType:    ue.accessType,
	}
	if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
		log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
	}

	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()
	ue.CmStatus = models.CmStateIdle
	ue.RmStatus = models.RmStateDeregistered

	if isGracefully {
		msg := &models.UeToAmfMsg{
			EventType:     models.AMFEVENTTYPEANYOF_REGISTRATION_STATE_REPORT,
			TimeStamp:     time.Now(),
			RmState:       ue.RmStatus,
			CmState:       ue.CmStatus,
			Supi:          ue.Imsi,
			Gpsi:          ue.Msidn,
			PlmnId:        ue.PlmnId,
			CurrentCellId: ue.CurrentCellId,
			AccessType:    ue.accessType,
		}
		if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
			log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
		}
	}

	monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.RmStateRegistered)).Dec()
	monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.RmStateDeregistered)).Inc()

	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateConnected)).Dec()
	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateIdle)).Inc()

	for i := range len(ue.PduSessions) {
		ue.ReleasePduSession(int32(i + 1))
	}

}

// NewPduSession establishes a new PDU Session for the UE.
// It accepts a sessionId, the DNN (Data Network Name), and SNSSAI (Single Network Slice Selection Assistance Information).
// It updates the PDU Sessions map and logs the establishment of the session.
// If the UE is not attached to the network, it logs an error message and does not establish the session.
// It also initializes the uplink data statistics for the session if enabled.
func (ue *Ue) NewPduSession(sessionId int32, dnn string, snssai models.Snssai, enableReport bool) {
	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()
	ue.LastActivityTime = time.Now()

	if ue.CmStatus != models.CmStateConnected {
		log.Printf("[%s] ue is not attached to the network, cannot establish PDU Session", ue.Imsi)
		return
	}

	ip, err := ue.ipManager.AllocateIP(ue.Imsi, sessionId)
	if err != nil {
		return
	}

	pduCtx, pduCancelFunc := context.WithCancel(ue.ctx)

	ue.PduSessions[sessionId] = models.PduSessionInfo{
		Id:           sessionId,
		Ipv4:         ip,
		Snssai:       snssai,
		Dnn:          dnn,
		Ctx:          pduCtx,
		CtxCancelFun: pduCancelFunc,
	}

	log.Printf("[%s] PDU Session %d established (dnn=%s, snssai=%+v, ip=%s", ue.Imsi, sessionId, dnn, snssai, ip)
	ue.HasUplinkData = false

	/*prepare gitc message for SMF*/
	msg := &models.UeToSmfMsg{
		EventType:   models.SMFEVENTANYOF_PDU_SES_EST,
		TimeStamp:   time.Now(),
		Dnn:         dnn,
		Snssai:      snssai,
		PduSessType: models.PDUSESSIONTYPEANYOF_IPV4,
		UeAddress:   ip,
		Supi:        ue.Imsi,
		Gpsi:        ue.Msidn,
		PlmnId:      ue.PlmnId,
		PduSessId:   sessionId,
		AccessType:  ue.accessType,
	}

	if err := gitc.Send(ue.Imsi, "SMF", models.UeToSmfType, msg); err != nil {
		log.Printf("Error sending UeToSmfType for UE %s: %v", ue.Imsi, err)
	}

	if enableReport {
		ue.UpStats[sessionId] = models.NewUpStats(sessionId)
		go ue.userplaneReport(pduCtx, sessionId)

	}

	monitoring.PduSessionsTotal.WithLabelValues(ue.simId).Inc()
	monitoring.UEIPInfo.WithLabelValues(ue.simId, ue.Imsi, ip).Set(1)

}

// ReleasePduSession releases the specified PDU Session for the UE.
// It accepts a sessionId as input and removes the session from the PDU Sessions map.
// It does not perform any additional actions or checks.
func (ue *Ue) ReleasePduSession(sessionId int32) {
	pduSess, exists := ue.PduSessions[sessionId]
	if !exists {
		log.Printf("[%s] invalid pduSessionId %d, cannot release", ue.Imsi, sessionId)
	}

	err := ue.ipManager.ReleaseIP(ue.Imsi, sessionId)
	if err != nil {
		return
	}

	/*prepare gitc message for SMF*/
	msg := &models.UeToSmfMsg{
		EventType:   models.SMFEVENTANYOF_PDU_SES_REL,
		TimeStamp:   time.Now(),
		Dnn:         pduSess.Dnn,
		Snssai:      pduSess.Snssai,
		PduSessType: models.PDUSESSIONTYPEANYOF_IPV4,
		UeAddress:   pduSess.Ipv4,
		Supi:        ue.Imsi,
		Gpsi:        ue.Msidn,
		PlmnId:      ue.PlmnId,
		PduSessId:   sessionId,
		AccessType:  ue.accessType,
	}

	monitoring.UEIPInfo.DeleteLabelValues(ue.simId, ue.Imsi, pduSess.Ipv4)
	monitoring.PduSessionsTotal.WithLabelValues(ue.simId).Dec()

	pduSess.CtxCancelFun()
	delete(ue.PduSessions, sessionId)
	log.Printf("[%s] released pduSessionId %d", ue.Imsi, sessionId)

	if err := gitc.Send(ue.Imsi, "SMF", models.UeToSmfType, msg); err != nil {
		log.Printf("Error sending UeToSmfType for UE %s: %v", ue.Imsi, err)
	}

}

// Sleep puts the UE into idle mode, allowing it to conserve resources when not actively communicating.
// It updates the CM status to idle and logs the activation of idle mode.
// It locks the status mutex to ensure thread safety during the update.
// If the UE is not registered, it logs an error message and does not proceed with the sleep operation.
// It also updates the LastActivityTime to the current time to reflect the transition to idle mode.
// This method is typically called when the UE is not actively communicating with the network.
func (ue *Ue) Sleep(force bool) {
	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()

	ue.CmStatus = models.CmStateIdle
	log.Printf("[%s] successfully activated idle mode", ue.Imsi)

	/*prepare gitc message for AMF*/
	msg := &models.UeToAmfMsg{
		EventType:     models.AMFEVENTTYPEANYOF_CONNECTIVITY_STATE_REPORT,
		TimeStamp:     time.Now(),
		RmState:       ue.RmStatus,
		CmState:       ue.CmStatus,
		Supi:          ue.Imsi,
		Gpsi:          ue.Msidn,
		PlmnId:        ue.PlmnId,
		CurrentCellId: ue.CurrentCellId,
		AccessType:    ue.accessType,
	}
	if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
		log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
	}

	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateIdle)).Inc()
	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateConnected)).Dec()

}

// WakeUp wakes up the UE from idle mode, allowing it to re-establish connectivity with the core network for signaling.
// It updates the CM status to connected and logs the service request.
// It locks the status mutex to ensure thread safety during the update.
// If the UE is not registered, it logs an error message and does not proceed with the wake-up operation.
// This method is typically called when the UE needs to re-establish connectivity after being in idle mode.
// It also updates the LastActivityTime to the current time to reflect the activity.
func (ue *Ue) WakeUp(isPaging bool) {
	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()

	if ue.RmStatus != models.RmStateRegistered {
		log.Printf("[%s] cannot page the ue, not registered", ue.Imsi)
		return
	}
	if ue.CmStatus != models.CmStateConnected {
		ue.LastActivityTime = time.Now()
		if isPaging {
			log.Printf("[%s] paging", ue.Imsi)
		} else {
			log.Printf("[%s] service request", ue.Imsi)
		}
		ue.CmStatus = models.CmStateConnected
		/*prepare gitc message for AMF*/
		msg := &models.UeToAmfMsg{
			EventType:     models.AMFEVENTTYPEANYOF_CONNECTIVITY_STATE_REPORT,
			TimeStamp:     time.Now(),
			RmState:       ue.RmStatus,
			CmState:       ue.CmStatus,
			Supi:          ue.Imsi,
			Gpsi:          ue.Msidn,
			PlmnId:        ue.PlmnId,
			CurrentCellId: ue.CurrentCellId,
			AccessType:    ue.accessType,
		}
		if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
			log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
		}
	}

	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateIdle)).Dec()
	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateConnected)).Inc()
}

// StartTrafficSession starts a traffic session for the specified PDU Session ID.
// It accepts a sessionId and a trafficProfile as input.
// It checks if the PDU Session exists in the PDU Sessions map and logs an error message if it does not.
// It initializes a timer for the traffic session duration and starts a goroutine to generate traffic.
// The traffic generation is simulated using a traffic generator, which can be customized based on the traffic profile.
// The goroutine continuously generates packets until the timer expires or the context is done.
// It updates the uplink statistics for the session, including the number of packets, total bytes, packet rate, and bitrate.
// If the timer expires, it logs the end of the traffic session and resets the uplink data status.
func (ue *Ue) StartTrafficSession(sessionId int32, ul bool, trafficProfile string, durationSec uint) {
	pduSess, exists := ue.PduSessions[sessionId]
	if !exists {
		log.Printf("[%s] invalid pduSessionId %d, cannot start traffic", ue.Imsi, sessionId)
	}

	var timerChannel <-chan time.Time

	if durationSec > 0 {
		timer := time.NewTimer(time.Duration(10) * time.Second)
		timerChannel = timer.C
	} else {
		timerChannel = nil
	}

	go func(ctx context.Context) {
		var trafficGen trafficgen.TrafficGenerator
		switch trafficProfile {
		case "web":
			trafficGen = trafficgen.NewWebTraffic(2e6, 1200, 20*time.Second, 40*time.Second)
		case "video":
			trafficGen = trafficgen.NewVideoTraffic(8e6, 1300)
		case "iot":
			trafficGen = trafficgen.NewIoTTraffic(1000, 15*time.Second)
		case "sip":
			trafficGen = trafficgen.NewVoIPTraffic(600, 50)
		default:
			log.Printf("[%s] unknown traffic profile %s, using default web traffic", ue.Imsi, trafficProfile)
			// Default to web traffic if no profile is specified
			trafficGen = trafficgen.NewWebTraffic(2e6, 1200, 6*time.Second, 10*time.Second)
		}
		//trafficGen := trafficgen.NewVideoTraffic(8e6, 1300)
		//trafficGen := trafficgen.NewIoTTraffic(1000, 15*time.Second)

		for {
			select {
			case <-timerChannel:
				log.Printf("[%s] traffic session ended for UE %d", ue.Imsi, sessionId)
				return

			case <-ctx.Done():
				log.Printf("[%s] traffic session cancelled for UE %d", ue.Imsi, sessionId)
				return
			default:
				//log.Printf("[%s] traffic session ongoing for session %d", ue.Imsi, sessionId)
				now := time.Now()
				pkt := trafficGen.NextPacket(now)

				if pkt != nil {

					ue.WakeUp(false)
					ue.statusMutex.Lock()
					ue.LastActivityTime = now
					ue.statusMutex.Unlock()

					ue.statsMutex.Lock()
					if stats, exists := ue.UpStats[sessionId]; exists {
						stats.NewPacket(ul, int64(pkt.SizeBytes), now)
						if ul {
							monitoring.TrafficPackets.WithLabelValues(ue.simId, ue.Imsi, "UL").Inc()
							monitoring.TrafficBytes.WithLabelValues(ue.simId, ue.Imsi, "UL").Add(float64(pkt.SizeBytes))
							monitoring.TotalTraffic.WithLabelValues(ue.simId, "UL").Add(float64(pkt.SizeBytes))
						} else {
							monitoring.TrafficPackets.WithLabelValues(ue.simId, ue.Imsi, "DL").Inc()
							monitoring.TrafficBytes.WithLabelValues(ue.simId, ue.Imsi, "DL").Add(float64(pkt.SizeBytes))
							monitoring.TotalTraffic.WithLabelValues(ue.simId, "UL").Add(float64(pkt.SizeBytes))
						}
					}
					ue.statsMutex.Unlock()

				}

				time.Sleep(1 * time.Millisecond) // Simulate ongoing traffic
			}
		}
	}(pduSess.Ctx)
}

// DoHandover performs a handover for the UE to a target cell.
// It updates the current cell ID and logs the handover event.
// It locks the status mutex to ensure thread safety during the update.
// It also updates the LastActivityTime to the current time to reflect the activity.
// This method is typically called when the UE needs to switch to a different cell, such as during a handover process.
// It is assumed that the handover process is managed by the network and does not require additional procedures
func (ue *Ue) DoHandover(targetCellId string) {
	ue.statusMutex.Lock()
	defer ue.statusMutex.Unlock()

	// freeze all the curent procedures

	log.Printf("[%s] handover to cell %s", ue.Imsi, targetCellId)

	/*prepare gitc message for AMF*/
	msg := &models.UeToAmfMsg{
		EventType:     models.AMFEVENTTYPEANYOF_LOCATION_REPORT,
		TimeStamp:     time.Now(),
		RmState:       ue.RmStatus,
		CmState:       ue.CmStatus,
		Supi:          ue.Imsi,
		Gpsi:          ue.Msidn,
		PlmnId:        ue.PlmnId,
		CurrentCellId: ue.CurrentCellId,
		AccessType:    ue.accessType,
	}
	if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
		log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
	}

	ue.LastActivityTime = time.Now()
	ue.CurrentCellId = targetCellId

}

/* ue inactivity monitor routine*/
func (ue *Ue) inactivityMonitor(inactivityTimer time.Duration) {
	for {
		select {
		case <-ue.ctx.Done():
			return
		case <-time.After(inactivityTimer):
			maxTolleratedInactivity := time.Now().Add(-inactivityTimer)

			ue.statusMutex.Lock()

			if ue.CmStatus == models.CmStateConnected && ue.LastActivityTime.Before(maxTolleratedInactivity) {
				log.Printf("[%s] ue inactivity timer expired, idle mode", ue.Imsi)
				ue.CmStatus = models.CmStateIdle
				ue.ueState = models.Idle
				/*prepare gitc message for AMF*/
				msg := &models.UeToAmfMsg{
					EventType:     models.AMFEVENTTYPEANYOF_CONNECTIVITY_STATE_REPORT,
					TimeStamp:     time.Now(),
					RmState:       ue.RmStatus,
					CmState:       ue.CmStatus,
					Supi:          ue.Imsi,
					Gpsi:          ue.Msidn,
					PlmnId:        ue.PlmnId,
					CurrentCellId: ue.CurrentCellId,
					AccessType:    ue.accessType,
				}
				if err := gitc.Send(ue.Imsi, "AMF", models.UeToAmfType, msg); err != nil {
					log.Printf("Error sending UeToAmfMsg for UE %s: %v", ue.Imsi, err)
				}
			}

			//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateIdle)).Inc()
			//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateConnected)).Dec()

			ue.statusMutex.Unlock()
		}
	}
}

/* pdu session qos monitoring routine */
func (ue *Ue) userplaneReport(ctx context.Context, pduSessId int32) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("[%s] stopped userplane report for PDU Session %d", ue.Imsi, pduSessId)
			return
		case <-time.After(5 * time.Second):
			ue.statusMutex.Lock()
			report := ue.UpStats[pduSessId].GenerateReport()
			session := ue.PduSessions[pduSessId]

			/*prepare gitc message for SMF*/
			msg := &models.UeToSmfMsg{
				EventType:   models.SMFEVENTANYOF_QOS_MON,
				TimeStamp:   time.Now(),
				Dnn:         session.Dnn,
				Snssai:      session.Snssai,
				PduSessType: models.PDUSESSIONTYPEANYOF_IPV4,
				UeAddress:   session.Ipv4,
				Supi:        ue.Imsi,
				Gpsi:        ue.Msidn,
				PlmnId:      ue.PlmnId,
				PduSessId:   pduSessId,
				AccessType:  ue.accessType,
				UpReport:    report,
			}
			if err := gitc.Send(ue.Imsi, "SMF", models.UeToSmfType, msg); err != nil {
				log.Printf("Error sending UeToSmfType for UE %s: %v", ue.Imsi, err)
			}
			ue.statusMutex.Unlock()
			//log.Printf("[%s] created userplane report for PDU Session %d\n", ue.Imsi, pduSessId)

		}
	}
}

// The UeSimulationRouting handles one Ue equipment and orchestrate its behaviors
// Based on a Markovian Process, it will change the states inside the users to simulate
// real world beahvior in terms of signalling and user plane traffic
func (ue *Ue) PowerUp() {

	monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.RmStateDeregistered)).Inc()
	//monitoring.UEsTotal.WithLabelValues(ue.simId, string(models.CmStateIdle)).Inc()

	err := gitc.StartTask(ue.Imsi, func(msg gitc.Message) {
		log.Printf("UE Received message from the network")
	}, 1024)
	if err != nil {
		log.Printf("Error starting GITC task for UE %s: %v", ue.Imsi, err)
		return
	}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				ue.statusMutex.Lock()
				var procedure models.UeProcedure
				ue.ueState, procedure = NextState(ue.ueState)
				//log.Printf("%d, %s", ue.ueState, procedure)
				ue.statusMutex.Unlock()

				switch procedure {
				case models.Registration:
					ue.Register()
				case models.Attach:
					ue.Attach(10 * time.Second)
				case models.PduSessionEstablishement:
					ue.NewPduSession(1, ue.defautlDnn, ue.defaultSnssai, true)
					ue.StartTrafficSession(1, false, "video", 0) // 0 = infinite duration
					ue.StartTrafficSession(1, true, "sip", 0)    // 0 = infinite duration
				case models.PduSessionFailure:
					// no actions for the UE
				case models.PduSessionRelease:
					ue.ReleasePduSession(1)
				case models.LossOfConnection:
					// Kill RF is loss of connection
					ue.LossOfConnection(false)
				case models.Sleep:
					// sleep is handled by inactivity timer
				case models.Paging:
					// Idle->Connected is handled here for paging, while service request is handled by the traffic routine
					ue.WakeUp(true)
					ue.StartTrafficSession(1, false, "sip", 0)
				case models.HandoverSuccessful:

					ue.DoHandover(ue.pickRandomNRCellID())
				case models.HandoverFailure:
					// Loss of connection if HO fails
					ue.LossOfConnection(false)
				case models.HandoverInitiated:
					// No actions for the UE
				default:

				}
			case <-ue.ctx.Done():
				return
			}
		}
	}()
}

func (ue *Ue) TurnOff(isGracefully bool) {
	ue.cancelFun()
	ue.LossOfConnection(isGracefully)
}

func (ue *Ue) pickRandomNRCellID() string {
	if len(ue.gnbList) <= 1 {
		return ""
	}

	for {
		idx := rand.Intn(len(ue.gnbList))
		if ue.gnbList[idx] != ue.CurrentCellId {
			return ue.gnbList[idx]
		}
	}
}
