/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"log"
	"os"

	pb "github.com/google/cadvisor/paas/grpcmonitor/monitor"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/grpcmonitor/monitor"
)

const (
	address = "localhost:50051"
	//address     = "10.12.136.137:22222"
	defaultName = "yunfan"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMonitorServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Send(context.Background(), &pb.Metric{Name: "2016-09-18", CpuPercent: 11, MemPercent: 22})
	if err != nil {
		log.Fatalf("could not send data: %v", err)
	}
	log.Printf("Hey Send: %s", name, r.Msg, r.Code)

	// test BatchSend
	/*
		var metric_items []*pb.Metric
		map1 := make(map[string]int32)
		map1["a"] = 1
		metric := pb.Metric{
			Uuid:        "10_1_1_2",
			Time:        111,
			Cpu:         1,
			Load:        2,
			Mem:         3,
			Swap:        4,
			Tcp:         5,
			NetIn:       map1,
			NetOut:      map1,
			NetInPk:     map1,
			NetOutPk:    map1,
			NetInErr:    map1,
			NetOutErr:   map1,
			DiscardPk:   1,
			IncomingPk:  1,
			RevSegments: 1,
			SndSegments: 1,
			BadSegments: 1,
			ActSegments: 1,
			Disk:        map1,
			DiskUtil:    map1,
			DiskRc:      map1,
			DiskWc:      map1,
			DiskRs:      map1,
			DiskWs:      map1,
		}
		metric_items = append(metric_items, &metric)
		metric_list := pb.MetricList{
			Items: metric_items,
		}
		r_bat, err := c.BatchSend(context.Background(), &metric_list)
		if err != nil {
			log.Fatalf("could not BatchSend Metric: %v", err)
		}
		log.Printf("BatchSend Metric: %s", name, r_bat.Msg, r_bat.Code)
	*/
}
