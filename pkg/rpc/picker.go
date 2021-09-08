/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
	"context"
	"log"

	"github.com/serialx/hashring"
	"google.golang.org/grpc/balancer"
)

// PickKey is a context.Context Value key. Its associated value should
// be a *PickReq struct.
type PickKey struct{}

// PickReq is a context.Context Value.
type PickReq struct {
	Key     string
	Attempt int
}

var (
	_ balancer.Picker = &d7yPicker{}
)

type PickResult struct {
	Ctx context.Context
	SC  balancer.SubConn
}

func newD7yPicker(subConns map[string]balancer.SubConn) *d7yPicker {
	addrs := make([]string, 0)
	for addr := range subConns {
		addrs = append(addrs, addr)
	}
	return &d7yPicker{
		subConns:   subConns,
		hashRing:   hashring.New(addrs),
		needReport: false,
	}
}

func newD7yReportingPicker(subConns map[string]balancer.SubConn, reportChan chan<- PickResult) *d7yPicker {
	addrs := make([]string, 0)
	for addr := range subConns {
		addrs = append(addrs, addr)
	}
	return &d7yPicker{
		subConns:   subConns,
		hashRing:   hashring.New(addrs),
		needReport: true,
		reportChan: reportChan,
	}
}

type d7yPicker struct {
	subConns   map[string]balancer.SubConn // address string -> balancer.SubConn
	hashRing   *hashring.HashRing
	needReport bool
	reportChan chan<- PickResult
}

func (p *d7yPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	var ret balancer.PickResult
	pickReq, ok := info.Ctx.Value(PickKey{}).(*PickReq)
	if !ok {
		pickReq = &PickReq{
			Key:     info.FullMethodName,
			Attempt: 1,
		}
	}
	log.Printf("pick for %s, for %d time(s)\n", pickReq.Key, pickReq.Attempt)
	if targetAddr, ok := p.hashRing.GetNodes(pickReq.Key, pickReq.Attempt); ok {
		ret.SubConn = p.subConns[targetAddr[pickReq.Attempt-1]]
		if p.needReport {
			p.reportChan <- PickResult{Ctx: info.Ctx, SC: ret.SubConn}
		}
	}
	if ret.SubConn == nil {
		return ret, balancer.ErrNoSubConnAvailable
	}
	return ret, nil
}