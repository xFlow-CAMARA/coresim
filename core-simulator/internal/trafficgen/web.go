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

package trafficgen

import "time"

// WebTraffic simulates bursty HTTP traffic (web browsing)
type WebTraffic struct {
	AvgBitrate    float64       // bits per second during burst
	PacketSize    int           // average packet size (bytes)
	BurstDuration time.Duration // duration of burst activity
	IdleDuration  time.Duration // duration of idle period

	lastPacketTime time.Time
	burstEndTime   time.Time
	inBurst        bool
}

// NewWebTraffic creates a new WebTraffic generator
func NewWebTraffic(bitrate float64, pktSize int, burst, idle time.Duration) *WebTraffic {
	return &WebTraffic{
		AvgBitrate:    bitrate,
		PacketSize:    pktSize,
		BurstDuration: burst,
		IdleDuration:  idle,
		inBurst:       true,
	}
}

// NextPacket emits the next packet or nil if no packet at this time
func (w *WebTraffic) NextPacket(now time.Time) *Packet {
	if now.After(w.burstEndTime) {
		// Switch between burst and idle
		w.inBurst = !w.inBurst
		if w.inBurst {
			w.burstEndTime = now.Add(w.BurstDuration)
		} else {
			w.burstEndTime = now.Add(w.IdleDuration)
		}
	}

	if w.inBurst {
		interval := time.Duration(float64(w.PacketSize*8)/w.AvgBitrate*1e9) * time.Nanosecond
		if now.Sub(w.lastPacketTime) >= interval {
			w.lastPacketTime = now
			return &Packet{
				SizeBytes: w.PacketSize,
				Timestamp: now,
			}
		}
	}
	return nil
}
