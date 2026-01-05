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

// VideoTraffic simulates steady video streaming traffic
type VideoTraffic struct {
	Bitrate    float64       // bits per second
	PacketSize int           // average packet size (bytes)
	Interval   time.Duration // inter-packet interval

	lastPacketTime time.Time
}

// NewVideoTraffic creates a new VideoTraffic generator
func NewVideoTraffic(bitrate float64, pktSize int) *VideoTraffic {
	pktPerSec := bitrate / float64(pktSize*8)
	interval := time.Duration(1e9/pktPerSec) * time.Nanosecond
	return &VideoTraffic{
		Bitrate:    bitrate,
		PacketSize: pktSize,
		Interval:   interval,
	}
}

// NextPacket emits the next packet or nil if no packet at this time
func (v *VideoTraffic) NextPacket(now time.Time) *Packet {
	if now.Sub(v.lastPacketTime) >= v.Interval {
		v.lastPacketTime = now
		return &Packet{
			SizeBytes: v.PacketSize,
			Timestamp: now,
		}
	}
	return nil
}
