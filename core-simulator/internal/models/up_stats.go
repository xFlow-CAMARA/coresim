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

package models

import (
	"fmt"
	"time"
)

type UpStats struct {
	PduSessId         int32
	NumOfPackets      int64
	TotalBytes        int64
	NumUlPackets      int64
	NumDlPackets      int64
	TotalUlBytes      int64
	TotalDlBytes      int64
	LastUlUpdate      time.Time
	LastDlUpdate      time.Time
	LastUlSizeArrived int64
	LastDlSizeArrived int64
}

type UpStatsReport struct {
	UpStats
	UlBitrate    float64
	DlBitrate    float64
	UlPacketRate float64
	DlPacketRate float64
}

func NewUpStats(sessionId int32) *UpStats {
	return &UpStats{
		PduSessId:    sessionId,
		NumOfPackets: 0,
		TotalBytes:   0,
		TotalUlBytes: 0,
		TotalDlBytes: 0,
		NumUlPackets: 0,
		NumDlPackets: 0,
		LastDlUpdate: time.Now(),
		LastUlUpdate: time.Now(),
	}
}

func (stats *UpStatsReport) Dumps() string {
	return fmt.Sprintf("SessionId:      %d,\nPackets:        %d,\nBytes:          %d,\nUl Bitrate:     %.2f bps,\nUl Packet Rate: %.2f pps,\nDl Bitrate:     %.2f bps,\nDl Packet Rate: %.2f pps,\n",
		stats.PduSessId, stats.NumOfPackets, stats.TotalBytes, stats.UlBitrate, stats.UlPacketRate, stats.DlBitrate, stats.DlPacketRate)
}

func (stats *UpStats) NewPacket(ul bool, size int64, timestamp time.Time) {
	stats.NumOfPackets++
	stats.TotalBytes += size
	if ul {
		stats.NumUlPackets++
		stats.TotalUlBytes += size
		stats.LastUlSizeArrived = size
		stats.LastUlUpdate = timestamp
	} else {
		stats.NumDlPackets++
		stats.TotalDlBytes += size
		stats.LastDlSizeArrived = size
		stats.LastDlUpdate = timestamp
	}

}

func (stats *UpStats) GenerateReport() *UpStatsReport {
	return &UpStatsReport{
		UpStats:      *stats,
		DlBitrate:    float64(stats.LastDlSizeArrived) / float64(time.Since(stats.LastDlUpdate).Seconds()) * 8,
		DlPacketRate: 1.00 / float64(time.Since(stats.LastDlUpdate).Seconds()),
		UlBitrate:    float64(stats.LastUlSizeArrived) / float64(time.Since(stats.LastUlUpdate).Seconds()) * 8,
		UlPacketRate: 1.00 / float64(time.Since(stats.LastUlUpdate).Seconds()),
	}
}
