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

// IoTTraffic simulates periodic status updates from IoT devices
type IoTTraffic struct {
	PacketSize        int           // bytes
	HeartbeatInterval time.Duration // interval between updates

	lastPacketTime time.Time
}

// NewIoTTraffic creates an IoT traffic generator
func NewIoTTraffic(pktSize int, interval time.Duration) *IoTTraffic {
	return &IoTTraffic{
		PacketSize:        pktSize,
		HeartbeatInterval: interval,
	}
}

// NextPacket emits a packet periodically
func (i *IoTTraffic) NextPacket(now time.Time) *Packet {
	if now.Sub(i.lastPacketTime) >= i.HeartbeatInterval {
		i.lastPacketTime = now
		return &Packet{
			SizeBytes: i.PacketSize,
			Timestamp: now,
		}
	}
	return nil
}
