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

package utils

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type IPAllocator struct {
	availableIPs []string
	allocated    map[string]string // userID -> IP
	ipToUser     map[string]string // IP -> userID
}

func NewIpamService(subnet string, netmask string) *IPAllocator {
	_, ipnet, err := net.ParseCIDR(fmt.Sprintf("%s/%s", subnet, netmask))
	if err != nil {
		return nil
	}

	ips := []string{}
	for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// Remove network and broadcast address
	if len(ips) > 2 {
		ips = ips[1 : len(ips)-1]
	}

	allocator := &IPAllocator{
		availableIPs: ips,
		allocated:    make(map[string]string),
		ipToUser:     make(map[string]string),
	}

	return allocator
}

func (a *IPAllocator) AllocateIP(supi string, pduSessId int32) (string, error) {
	userString := fmt.Sprintf("%s-%d", supi, pduSessId)
	if ip, ok := a.allocated[userString]; ok {
		return ip, nil // user already has an IP
	}

	if len(a.availableIPs) == 0 {
		return "", errors.New("no available IP addresses")
	}

	ip := a.availableIPs[0]
	a.availableIPs = a.availableIPs[1:]
	a.allocated[userString] = ip
	a.ipToUser[ip] = userString

	return ip, nil
}

func (a *IPAllocator) ReleaseIP(supi string, pduSessId int32) error {
	userString := fmt.Sprintf("%s-%d", supi, pduSessId)
	ip, ok := a.allocated[userString]
	if !ok {
		return errors.New("user does not have an allocated IP")
	}

	delete(a.allocated, userString)
	delete(a.ipToUser, ip)
	a.availableIPs = append([]string{ip}, a.availableIPs...)

	return nil
}

func (a *IPAllocator) GetIP(supi string, pduSessId int32) (string, bool) {
	userString := fmt.Sprintf("%s-%d", supi, pduSessId)
	ip, ok := a.allocated[userString]
	return ip, ok
}

func (a *IPAllocator) GetUserStringOk(ip string) (string, int32, bool) {
	userString, ok := a.ipToUser[ip]
	if !ok {
		log.Printf("user not found %s, %+v", ip, a.ipToUser)
		return "", 0, false
	}
	userStringSplitted := strings.Split(userString, "-")
	user := userStringSplitted[0]
	pduSessId, err := strconv.Atoi(userStringSplitted[1])
	if err != nil {
		return "", 0, false
	}
	return user, int32(pduSessId), ok
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
