# sig-node-flaky-tests
Small script to find top flaky tests from http://testgrid.k8s.io/sig-node

## v1.30 release

### sig-node-release-blocking flaky tests

| Testname | Flaky% |
|----------|--------|
| E2eNode Suite.[It] [sig-node] Device Manager [Serial] [Feature:DeviceManager] [NodeFeature:DeviceManager] With sample device plugin [Serial] [Disruptive] should deploy pod consuming devices first but fail with admission error after kubelet restart in case device plugin hasn't re-registered | 14.444445 |
| E2eNode Suite.[It] [sig-node] PriorityPidEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause PIDPressure; PodDisruptionConditions enabled [NodeFeature:PodDisruptionConditions]  should eventually evict all of the correct pods | 13.333334 |
| E2eNode Suite.[It] [sig-node] PriorityPidEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause PIDPressure  should eventually evict all of the correct pods | 13.333334 |
| E2eNode Suite.[It] [sig-node] LocalStorageSoftEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 12.222222 |
| E2eNode Suite.[It] [sig-node] InodeEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 11.111112 |
| E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] with the builtin rate limit values should hit throttling when calling podresources List in a tight loop | 10.000000 |
| E2eNode Suite.[It] [sig-node] MirrorPodWithGracePeriod when create a mirror pod  and the container runtime is temporarily down during pod termination [NodeConformance] [Serial] [Disruptive] the mirror pod should terminate successfully | 10.000000 |
| E2eNode Suite.[It] [sig-node] LocalStorageEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 8.888889 |
| E2eNode Suite.[It] [sig-node] PriorityLocalStorageEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 8.888889 |
| E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Keeps device plugin assignments across node reboots (no pod restart, no device plugin re-registration) | 7.777778 |
| E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager None policy should return the expected responses | 7.777778 |
| E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should report zero pinning counters after a fresh restart | 7.777778 |
| E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should report pinning failures when the memorymanager allocation is known to fail | 6.666667 |
| E2eNode Suite.[It] [sig-node] GracefulNodeShutdown [Serial] [NodeFeature:GracefulNodeShutdown] [NodeFeature:GracefulNodeShutdownBasedOnPodPriority] when gracefully shutting down with Pod priority should be able to gracefully shutdown pods with various grace periods | 5.555556 |
| E2eNode Suite.[It] [sig-node] ImageGCNoEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 5.555556 |
| E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should not report any pinning failures when the memorymanager allocation is expected to succeed | 2.222222 |
| E2eNode Suite.[It] [sig-node] CriticalPod [Serial] [Disruptive] [NodeFeature:CriticalPod] when we need to admit a critical pod should add DisruptionTarget condition to the preempted pod [NodeFeature:PodDisruptionConditions] | 2.222222 |
| E2eNode Suite.[It] [sig-node] CriticalPod [Serial] [Disruptive] [NodeFeature:CriticalPod] when we need to admit a critical pod should be able to create and delete a critical pod | 2.222222 |
| E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Keeps device plugin assignments after kubelet restart and device plugin restart (no pod restart) | 2.222222 |
| E2eNode Suite.[It] [sig-node] Restart [Serial] [Slow] [Disruptive] Kubelet should correctly account for terminated pods after restart | 1.111111 |
| E2eNode Suite.[It] [sig-node] OOMKiller for pod using more memory than node allocatable [LinuxOnly] [Serial] single process container without resource limits The containers terminated by OOM killer should have the reason set to OOMKilled | 1.111111 |
| E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Can schedule a pod that requires a device | 1.111111 |

### Overall testgrid flaky tests

| Testname | Flaky% |
|----------|--------|
| E2eNode Suite.[It] [sig-node] PriorityPidEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause PIDPressure  should eventually evict all of the correct pods | 46.808510 |
| E2eNode Suite.[It] [sig-node] ImageGCNoEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 13.373860 |
| E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager Static policy  should return the expected responses | 12.500000 |
| E2eNode Suite.[It] [sig-node] InodeEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 10.942249 |
| E2eNode Suite.[It] [sig-node] LocalStorageSoftEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 10.942249 |
| E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager Static policy  should return the expected responses [NodeFeature:SidecarContainers] | 9.677419 |
| E2eNode Suite.[It] [sig-node] PriorityPidEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause PIDPressure; PodDisruptionConditions enabled [NodeFeature:PodDisruptionConditions]  should eventually evict all of the correct pods | 9.422492 |
| E2eNode Suite.[It] [sig-node] LocalStorageEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 8.206687 |
| E2eNode Suite.[It] [sig-node] PriorityLocalStorageEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 7.294833 |
| E2eNode Suite.[It] [sig-node] Density [Serial] [Slow] create a batch of pods [Flaky] latency/resource should be within limit when create 10 pods with 0s interval | 7.216495 |
| E2eNode Suite.[It] [sig-node] OOMKiller for pod using more memory than node allocatable [LinuxOnly] [Serial] single process container without resource limits The containers terminated by OOM killer should have the reason set to OOMKilled | 6.542056 |
| E2eNode Suite.[It] [sig-node] Device Manager [Serial] [Feature:DeviceManager] [NodeFeature:DeviceManager] With sample device plugin [Serial] [Disruptive] should deploy pod consuming devices first but fail with admission error after kubelet restart in case device plugin hasn't re-registered | 6.542056 |
| E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] with the builtin rate limit values should hit throttling when calling podresources List in a tight loop | 5.794393 |
| E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager None policy should return the expected responses | 5.420561 |
| E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Keeps device plugin assignments across node reboots (no pod restart, no device plugin re-registration) | 4.672897 |
| E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should report pinning failures when the memorymanager allocation is known to fail | 3.561644 |
| E2eNode Suite.[It] [sig-node] MirrorPodWithGracePeriod when create a mirror pod  and the container runtime is temporarily down during pod termination [NodeConformance] [Serial] [Disruptive] the mirror pod should terminate successfully | 3.177570 |
| E2eNode Suite.[It] [sig-node] CriticalPod [Serial] [Disruptive] [NodeFeature:CriticalPod] when we need to admit a critical pod should add DisruptionTarget condition to the preempted pod [NodeFeature:PodDisruptionConditions] | 2.990654 |
| E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should report zero pinning counters after a fresh restart | 2.602740 |
| E2eNode Suite.[It] [sig-node] Memory Manager [Disruptive] [Serial] [Feature:MemoryManager] with static policy when guaranteed pod memory request is bigger than free memory on each NUMA node should be rejected | 2.380952 |
| E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should not report any pinning failures when the memorymanager allocation is expected to succeed | 2.328767 |
| E2eNode Suite.[It] [sig-node] CriticalPod [Serial] [Disruptive] [NodeFeature:CriticalPod] when we need to admit a critical pod should be able to create and delete a critical pod | 2.056075 |
| E2eNode Suite.[It] [sig-node] GracefulNodeShutdown [Serial] [NodeFeature:GracefulNodeShutdown] [NodeFeature:GracefulNodeShutdownBasedOnPodPriority] when gracefully shutting down with Pod priority should be able to gracefully shutdown pods with various grace periods | 2.056075 |
| E2eNode Suite.[It] [sig-node] [NodeFeature:SidecarContainers] Containers Lifecycle should terminate sidecars simultaneously if prestop doesn't exit [ubuntu01] | 1.960784 |
| E2eNode Suite.[It] [sig-node] Node Container Manager [Serial] Validate Node Allocatable [NodeFeature:NodeAllocatable] sets up the node and runs the test | 1.538462 |
| E2eNode Suite.[It] [sig-node] MirrorPod when recreating a static pod it should launch successfully even if it temporarily failed termination due to volume failing to unmount [NodeConformance] [Serial] | 1.495327 |
| E2eNode Suite.[It] [sig-node] Topology Manager [Serial] [Feature:TopologyManager] With kubeconfig updated to static CPU Manager policy run the Topology Manager tests run Topology Manager policy test suite | 1.465201 |
| E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Can schedule a pod that requires a device | 1.308411 |
| E2eNode Suite.[It] [sig-node] [NodeFeature:SidecarContainers] Containers Lifecycle should terminate sidecars simultaneously if prestop doesn't exit [cos-stable01] | 1.262272 |
| E2eNode Suite.[It] [sig-node] Topology Manager Metrics [Serial] [Feature:TopologyManager] when querying /metrics should not report any admission failures when the topology manager alignment is expected to succeed | 1.204819 |
| E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Keeps device plugin assignments across pod restarts (no kubelet restart, no device plugin restart) | 1.121495 |
| E2eNode Suite.[It] [sig-node] SeccompDefault [Serial] [Feature:SeccompDefault] [LinuxOnly] with SeccompDefault enabled should use the default seccomp profile when unspecified | 1.121495 |
| E2eNode Suite.[It] [sig-node] [NodeFeature:SidecarContainers] Containers Lifecycle when using a restartable init container in a Pod with restartPolicy=Never when a restartable init container runs continuously should not restart a restartable init container [ubuntu01] | 1.120448 |
