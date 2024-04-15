/*
Copyright 2023 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/testgrid/pb/api/v1"
	statuspb "github.com/GoogleCloudPlatform/testgrid/pb/test_status"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/encoding/protojson"
)

func importantSIGNodeDashboards() map[string][]string {
	return map[string][]string{
		"sig-node-release-blocking": []string{
			"ci-crio-cgroupv1-node-e2e-conformance",
			"ci-crio-cgroupv2-node-e2e-conformance",
			"ci-node-e2e",
			"node-kubelet-containerd-standalone-mode",
			"node-kubelet-containerd-standalone-mode-all-alpha",
			"node-kubelet-serial-containerd",
		},/*
		"sig-node-kubelet": []string{
			"gcp-kubelet-credential-provider",
			"kubelet-gce-e2e-fsquota-ubuntu",
			"kubelet-gce-e2e-arm64-ubuntu-serial",
			"kubelet-credential-provider",
			"kubelet-gce-e2e-lock-contention",
			"kubelet-serial-gce-e2e-cpu-manager",
			"kubelet-serial-gce-e2e-memory-manager",
			"kubelet-serial-gce-e2e-topology-manager",
			"kubelet-swap-conformance-fedora-serial",
			"kubelet-swap-conformance-ubuntu-serial",
			"kubelet-gce-e2e-swap-fedora",
			"kubelet-gce-e2e-swap-fedora-serial",
			"kubelet-gce-e2e-swap-ubuntu",
			"kubelet-gce-e2e-swap-ubuntu-serial",
		},
		"sig-node-containerd": []string{
			"cgroup-systemd-containerd-node-e2e",
			"cos-cgroupv1-containerd-e2e",
			"cos-cgroupv1-containerd-node-e2e",
			"cos-cgroupv1-containerd-node-features",
			"cos-cgroupv1-containerd-node-e2e-serial",
			"cos-cgroupv1-inplace-pod-resize-containerd-e2e-serial",
			"cos-cgroupv2-containerd-e2e",
			"cos-cgroupv2-containerd-node-e2e",
			"cos-cgroupv2-containerd-node-features",
			"cos-cgroupv2-containerd-node-e2e-serial",
			"cos-cgroupv2-inplace-pod-resize-containerd-e2e-serial",
			"e2e-cos-device-plugin-gpu",
			"node-e2e-features",
			"node-e2e-unlabelled",
			"ci-node-e2e",
			"node-kubelet-containerd-eviction",
			"node-kubelet-containerd-flaky",
			"node-kubelet-containerd-hugepages",
			"node-kubelet-containerd-performance-test",
			"node-kubelet-containerd-resource-managers",
		},
		"sig-node-cri-o": []string{
			"ci-crio-cdi-device-plugins",
			"ci-crio-cgroupv1-evented-pleg",
			"ci-crio-cgroupv1-node-e2e-conformance",
			"ci-crio-cgroupv1-node-e2e-eviction",
			"ci-crio-cgroupv1-node-e2e-features",
			"ci-crio-cgroupv1-node-e2e-flaky",
			"ci-crio-cgroupv1-node-e2e-hugepages",
			"ci-crio-cgroupv1-node-e2e-resource-managers",
			"ci-crio-cgroupv1-node-e2e-unlabelled",
			"ci-crio-cgroupv2-node-e2e-conformance",
			"ci-crio-cgroupv2-node-e2e-eviction",
			"node-kubelet-serial-crio",
			"kubelet-swap-conformance-fedora-serial",
			"kubelet-gce-e2e-swap-fedora",
			"kubelet-gce-e2e-swap-fedora-serial",
			"ci-node-e2e-crio-dra",
			"ci-node-e2e-crio-dra-features",
		},*/
	}
}

func main() {
	// map of testname to (total pass, total fails).
	type stat struct {
		pass  int
		fail  int
		empty int
	}
	testData := map[string]*stat{}

	for dashboard, tabs := range importantSIGNodeDashboards() {
		for _, tab := range tabs {
			ep := fmt.Sprintf("http://testgrid-data.k8s.io/api/v1/dashboards/%s/tabs/%s/rows", dashboard, tab)
			resp, err := http.Get(ep)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			tgData := &v1.ListRowsResponse{}
			err = protojson.Unmarshal(body, tgData)
			if err != nil {
				panic(err)
			}

			for _, row := range tgData.Rows {
				if !strings.Contains(row.Name, "sig-node") {
					continue
				}
				_, ok := testData[row.Name]
				if !ok {
					testData[row.Name] = &stat{pass: 0, fail: 0}
				}

				for _, result := range row.Cells {
					if result.Result == int32(statuspb.TestStatus_NO_RESULT) {
						testData[row.Name].empty += 1
					}
					if result.Result == int32(statuspb.TestStatus_PASS) {
						testData[row.Name].pass += 1
					}
					if result.Result == int32(statuspb.TestStatus_FAIL) || result.Result == int32(statuspb.TestStatus_FLAKY) {
						testData[row.Name].fail += 1
					}
				}
			}
		}
	}

	// Sort by flakiness
	tests := maps.Keys(testData)
	sort.Slice(tests, func(a, b int) bool {
		first := float32(testData[tests[a]].fail) / float32(testData[tests[a]].fail + testData[tests[a]].pass + testData[tests[a]].empty)
		second := float32(testData[tests[b]].fail) / float32(testData[tests[b]].fail + testData[tests[b]].pass + testData[tests[b]].empty)

		return first < second
	})

	fmt.Println("| Testname | Flaky% |")
	fmt.Println("|----------|--------|")
	for i := len(tests)-1; i >= 0; i-- {
		val := testData[tests[i]]
		flakyPercent := float32(val.fail) / float32(val.fail + val.pass + val.empty)*100
		if flakyPercent < 1 || val.fail == 0 {
			continue
		}
		fmt.Printf("| %s | %f |\n", tests[i], flakyPercent)
	}
}
