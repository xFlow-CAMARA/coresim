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

// Packet represents a network packet emitted by a traffic generator
type Packet struct {
	SizeBytes int       // Packet size in bytes
	Timestamp time.Time // Timestamp when generated
}

// TrafficGenerator defines the interface for all traffic generators
type TrafficGenerator interface {
	NextPacket(now time.Time) *Packet
}
