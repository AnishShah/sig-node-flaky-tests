# sig-node-flaky-tests
Small script to find top flaky tests from http://testgrid.k8s.io/sig-node

## v1.30 release

| Testname | Flaky% |
|----------|--------|
|E2eNode Suite.[It] [sig-node] PriorityPidEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause PIDPressure  should eventually evict all of the correct pods | 59.401711 |
|E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager Static policy  should return the expected responses [NodeFeature:SidecarContainers] | 14.285715 |
|E2eNode Suite.[It] [sig-node] ImageGCNoEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 13.247864 |
|E2eNode Suite.[It] [sig-node] PriorityPidEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause PIDPressure; PodDisruptionConditions enabled [NodeFeature:PodDisruptionConditions]  should eventually evict all of the correct pods | 11.538462 |
|E2eNode Suite.[It] [sig-node] InodeEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 11.111112 |
|E2eNode Suite.[It] [sig-node] LocalStorageSoftEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 8.119658 |
|E2eNode Suite.[It] [sig-node] Density [Serial] [Slow] create a batch of pods [Flaky] latency/resource should be within limit when create 10 pods with 0s interval | 7.368421 |
|E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager Static policy  should account for resources of pods in terminal phase | 7.142858 |
|E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager Static policy  should return the expected responses | 7.142858 |
|E2eNode Suite.[It] [sig-node] OOMKiller for pod using more memory than node allocatable [LinuxOnly] [Serial] single process container without resource limits The containers terminated by OOM killer should have the reason set to OOMKilled | 6.591337 |
|E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] with the builtin rate limit values should hit throttling when calling podresources List in a tight loop | 6.403013 |
|E2eNode Suite.[It] [sig-node] LocalStorageEviction [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 5.128205 |
|E2eNode Suite.[It] [sig-node] POD Resources [Serial] [Feature:PodResources] [NodeFeature:PodResources] without SRIOV devices in the system with CPU manager None policy should return the expected responses | 4.708098 |
|E2eNode Suite.[It] [sig-node] PriorityLocalStorageEvictionOrdering [Slow] [Serial] [Disruptive] [NodeFeature:Eviction] when we run containers that should cause DiskPressure  should eventually evict all of the correct pods | 4.700855 |
|E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Keeps device plugin assignments across node reboots (no pod restart, no device plugin re-registration) | 4.519774 |
|E2eNode Suite.[It] [sig-node] Device Manager [Serial] [Feature:DeviceManager] [NodeFeature:DeviceManager] With sample device plugin [Serial] [Disruptive] should deploy pod consuming devices first but fail with admission error after kubelet restart in case device plugin hasn't re-registered | 4.143126 |
|E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should not report any pinning failures when the memorymanager allocation is expected to succeed | 3.856750 |
|E2eNode Suite.[It] [sig-node] Node Container Manager [Serial] Validate Node Allocatable [NodeFeature:NodeAllocatable] sets up the node and runs the test | 3.305785 |
|E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should report pinning failures when the memorymanager allocation is known to fail | 3.305785 |
|E2eNode Suite.[It] [sig-node] GracefulNodeShutdown [Serial] [NodeFeature:GracefulNodeShutdown] [NodeFeature:GracefulNodeShutdownBasedOnPodPriority] when gracefully shutting down with Pod priority should be able to gracefully shutdown pods with various grace periods | 3.201507 |
|E2eNode Suite.[It] [sig-node] CriticalPod [Serial] [Disruptive] [NodeFeature:CriticalPod] when we need to admit a critical pod should add DisruptionTarget condition to the preempted pod [NodeFeature:PodDisruptionConditions] | 2.830189 |
|E2eNode Suite.[It] [sig-node] Memory Manager Metrics [Serial] [Feature:MemoryManager] when querying /metrics should report zero pinning counters after a fresh restart | 2.203857 |
|E2eNode Suite.[It] [sig-node] [NodeFeature:SidecarContainers] Containers Lifecycle should terminate sidecars simultaneously if prestop doesn't exit [cos-stable01] | 1.683030 |
|E2eNode Suite.[It] [sig-node] [NodeFeature:SidecarContainers] Containers Lifecycle should terminate sidecars simultaneously if prestop doesn't exit [ubuntu01] | 1.680672 |
|E2eNode Suite.[It] [sig-node] CriticalPod [Serial] [Disruptive] [NodeFeature:CriticalPod] when we need to admit a critical pod should be able to create and delete a critical pod | 1.509434 |
|E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Keeps device plugin assignments across kubelet restarts (no pod restart, no device plugin restart) | 1.506591 |
|E2eNode Suite.[It] [sig-node] MirrorPodWithGracePeriod when create a mirror pod  and the container runtime is temporarily down during pod termination [NodeConformance] [Serial] [Disruptive] the mirror pod should terminate successfully | 1.320755 |
|E2eNode Suite.[It] [sig-node] Device Plugin [Feature:DevicePluginProbe] [NodeFeature:DevicePluginProbe] [Serial] DevicePlugin [Serial] [Disruptive] Can schedule a pod with a restartable init container [NodeFeature:SidecarContainers] | 1.132076 |
