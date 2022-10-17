package yaml22json

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {

	jsonData := `{"apiVersion":"v1","kind":"Node","metadata":{"annotations":{"csi.volume.kubernetes.io/nodeid":"{\"com.tencent.cloud.csi.cbs\":\"ins-bzwza4jk\"}","node.alpha.kubernetes.io/ttl":"0","volumes.kubernetes.io/controller-managed-attach-detach":"true"},"creationTimestamp":"2022-09-20T11:57:54Z","labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/instance-type":"S3.MEDIUM4","beta.kubernetes.io/os":"linux","cloud.tencent.com/node-instance-id":"ins-bzwza4jk","failure-domain.beta.kubernetes.io/region":"hk","failure-domain.beta.kubernetes.io/zone":"300003","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"192.168.2.105","kubernetes.io/os":"linux","node.kubernetes.io/instance-type":"S3.MEDIUM4","topology.com.tencent.cloud.csi.cbs/zone":"ap-hongkong-3","topology.kubernetes.io/region":"hk","topology.kubernetes.io/zone":"300003"},"name":"192.168.2.105","resourceVersion":"3499832895","selfLink":"/api/v1/nodes/192.168.2.105","uid":"a16447c4-a86a-4c8e-96e0-08c20277d0bd"},"spec":{"podCIDR":"172.16.0.0/26","podCIDRs":["172.16.0.0/26"],"providerID":"qcloud:///300003/ins-bzwza4jk"},"status":{"addresses":[{"address":"192.168.2.105","type":"InternalIP"},{"address":"43.135.72.58","type":"ExternalIP"},{"address":"192.168.2.105","type":"Hostname"}],"allocatable":{"cpu":"1900m","ephemeral-storage":"47498714648","hugepages-1Gi":"0","hugepages-2Mi":"0","memory":"2607596Ki","pods":"61"},"capacity":{"cpu":"2","ephemeral-storage":"51539404Ki","hugepages-1Gi":"0","hugepages-2Mi":"0","memory":"3758572Ki","pods":"61"},"conditions":[{"lastHeartbeatTime":"2022-09-20T11:57:56Z","lastTransitionTime":"2022-09-20T11:57:56Z","message":"RouteController created a route","reason":"RouteCreated","status":"False","type":"NetworkUnavailable"},{"lastHeartbeatTime":"2022-10-17T03:08:27Z","lastTransitionTime":"2022-09-20T11:57:54Z","message":"kubelet has sufficient memory available","reason":"KubeletHasSufficientMemory","status":"False","type":"MemoryPressure"},{"lastHeartbeatTime":"2022-10-17T03:08:27Z","lastTransitionTime":"2022-09-20T11:57:54Z","message":"kubelet has no disk pressure","reason":"KubeletHasNoDiskPressure","status":"False","type":"DiskPressure"},{"lastHeartbeatTime":"2022-10-17T03:08:27Z","lastTransitionTime":"2022-09-20T11:57:54Z","message":"kubelet has sufficient PID available","reason":"KubeletHasSufficientPID","status":"False","type":"PIDPressure"},{"lastHeartbeatTime":"2022-10-17T03:08:27Z","lastTransitionTime":"2022-09-20T11:58:14Z","message":"kubelet is posting ready status","reason":"KubeletReady","status":"True","type":"Ready"}],"daemonEndpoints":{"kubeletEndpoint":{"Port":10250}},"images":[{"names":["hkccr.ccs.tencentyun.com/tkeimages/hyperkube@sha256:2a0e9400f6a9ece7e9225a4f1c38dd458f9c1563bd7ff1f14a809a93037b7261","hkccr.ccs.tencentyun.com/tkeimages/hyperkube:v1.22.5-tke.1"],"sizeBytes":184047984},{"names":["hkccr.ccs.tencentyun.com/tkeimages/tke-cni-agent@sha256:3994844a2edfec31e36d1fabcd2381e94ce7b0492a5010ed976fb6a7702a0b7a","hkccr.ccs.tencentyun.com/tkeimages/tke-cni-agent:v0.1.2"],"sizeBytes":71251871},{"names":["hkccr.ccs.tencentyun.com/tkeimages/monitor-agent@sha256:5cd31d2c843c7d5d430e45160b1acc7ebfacae67f0fa9c563549cc1654da41b5","hkccr.ccs.tencentyun.com/tkeimages/monitor-agent:v1.3.0-abe11a"],"sizeBytes":58900320},{"names":["hkccr.ccs.tencentyun.com/tkeimages/csi-tencentcloud-cbs@sha256:2c76fc98bf04e1412d7ce3a08be87ab7cc301b3f8655c14afd3f6cdacde72d6e","hkccr.ccs.tencentyun.com/tkeimages/csi-tencentcloud-cbs:v2.2.8"],"sizeBytes":41586124},{"names":["hkccr.ccs.tencentyun.com/tkeimages/tke-bridge-agent@sha256:70f1493795b2ceaa99d067a31d6bafd42b8171dac83e65a70defff1f419792ad","hkccr.ccs.tencentyun.com/tkeimages/tke-bridge-agent:v0.1.5"],"sizeBytes":22628290},{"names":["hkccr.ccs.tencentyun.com/tkeimages/apiserver-proxy@sha256:ae74732ec30ad10d6aa4daf6f9e355f62b77913244d1911744a3095d8890da8b","hkccr.ccs.tencentyun.com/tkeimages/apiserver-proxy:v1.3.5"],"sizeBytes":20941553},{"names":["hkccr.ccs.tencentyun.com/tkeimages/coredns@sha256:7a8fedc352f5e1002504b82a0230c07dff2657d23a2dc6a6db0b47d5b67830f3","hkccr.ccs.tencentyun.com/tkeimages/coredns:1.8.4"],"sizeBytes":13706014},{"names":["hkccr.ccs.tencentyun.com/tkeimages/csi-node-driver-registrar@sha256:588f77473b5e5e81d141bde83f3287e80ef7441ccf89774691c480254c4cb93a","hkccr.ccs.tencentyun.com/tkeimages/csi-node-driver-registrar:v2.0.1"],"sizeBytes":8414160},{"names":["hkccr.ccs.tencentyun.com/library/ip-masq-agent@sha256:655c4a06a15ca3252d395e9763505230edd3b8d181e16e039c7ba8fdfc1742e5","hkccr.ccs.tencentyun.com/library/ip-masq-agent:v2.6.2"],"sizeBytes":7466018},{"names":["hkccr.ccs.tencentyun.com/library/pause@sha256:5ab61aabaedd6c40d05ce1ac4ea72c2079f4a0f047ec1dc100ea297b553539ab","hkccr.ccs.tencentyun.com/library/pause:latest"],"sizeBytes":298560}],"nodeInfo":{"architecture":"amd64","bootID":"7bbce9fa-feaf-4493-ae26-a49312a5a154","containerRuntimeVersion":"containerd://1.3.4-1-g07bb01e5","kernelVersion":"5.4.119-19-0009.3","kubeProxyVersion":"v1.22.5-tke.1","kubeletVersion":"v1.22.5-tke.1","machineID":"d3bf44cf57154ba4bd061f8e5f3cdda2","operatingSystem":"linux","osImage":"TencentOS Server 3.1 (Final)","systemUUID":"d3bf44cf-5715-4ba4-bd06-1f8e5f3cdda2"}}}`
	yamlData, err := JSON2Yaml(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println(yamlData)

	yData := `apiVersion: v1
kind: Node
metadata:
  annotations:
    csi.volume.kubernetes.io/nodeid: '{"com.tencent.cloud.csi.cbs":"ins-bzwza4jk"}'
    node.alpha.kubernetes.io/ttl: "0"
    volumes.kubernetes.io/controller-managed-attach-detach: "true"
  creationTimestamp: "2022-09-20T11:57:54Z"
  labels:
    beta.kubernetes.io/arch: amd64
    beta.kubernetes.io/instance-type: S3.MEDIUM4
    beta.kubernetes.io/os: linux
    cloud.tencent.com/node-instance-id: ins-bzwza4jk
    failure-domain.beta.kubernetes.io/region: hk
    failure-domain.beta.kubernetes.io/zone: "300003"
    kubernetes.io/arch: amd64
    kubernetes.io/hostname: 192.168.2.105
    kubernetes.io/os: linux
    node.kubernetes.io/instance-type: S3.MEDIUM4
    topology.com.tencent.cloud.csi.cbs/zone: ap-hongkong-3
    topology.kubernetes.io/region: hk
    topology.kubernetes.io/zone: "300003"
  name: 192.168.2.105
  resourceVersion: "3499832895"
  selfLink: /api/v1/nodes/192.168.2.105
  uid: a16447c4-a86a-4c8e-96e0-08c20277d0bd
spec:
  podCIDR: 172.16.0.0/26
  podCIDRs:
  - 172.16.0.0/26
  providerID: qcloud:///300003/ins-bzwza4jk
status:
  addresses:
  - address: 192.168.2.105
    type: InternalIP
  - address: 43.135.72.58
    type: ExternalIP
  - address: 192.168.2.105
    type: Hostname
  allocatable:
    cpu: 1900m
    ephemeral-storage: "47498714648"
    hugepages-1Gi: "0"
    hugepages-2Mi: "0"
    memory: 2607596Ki
    pods: "61"
  capacity:
    cpu: "2"
    ephemeral-storage: 51539404Ki
    hugepages-1Gi: "0"
    hugepages-2Mi: "0"
    memory: 3758572Ki
    pods: "61"
  conditions:
  - lastHeartbeatTime: "2022-09-20T11:57:56Z"
    lastTransitionTime: "2022-09-20T11:57:56Z"
    message: RouteController created a route
    reason: RouteCreated
    status: "False"
    type: NetworkUnavailable
  - lastHeartbeatTime: "2022-10-17T03:08:27Z"
    lastTransitionTime: "2022-09-20T11:57:54Z"
    message: kubelet has sufficient memory available
    reason: KubeletHasSufficientMemory
    status: "False"
    type: MemoryPressure
  - lastHeartbeatTime: "2022-10-17T03:08:27Z"
    lastTransitionTime: "2022-09-20T11:57:54Z"
    message: kubelet has no disk pressure
    reason: KubeletHasNoDiskPressure
    status: "False"
    type: DiskPressure
  - lastHeartbeatTime: "2022-10-17T03:08:27Z"
    lastTransitionTime: "2022-09-20T11:57:54Z"
    message: kubelet has sufficient PID available
    reason: KubeletHasSufficientPID
    status: "False"
    type: PIDPressure
  - lastHeartbeatTime: "2022-10-17T03:08:27Z"
    lastTransitionTime: "2022-09-20T11:58:14Z"
    message: kubelet is posting ready status
    reason: KubeletReady
    status: "True"
    type: Ready
  daemonEndpoints:
    kubeletEndpoint:
      Port: 10250
  images:
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/hyperkube@sha256:2a0e9400f6a9ece7e9225a4f1c38dd458f9c1563bd7ff1f14a809a93037b7261
    - hkccr.ccs.tencentyun.com/tkeimages/hyperkube:v1.22.5-tke.1
    sizeBytes: 184047984
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/tke-cni-agent@sha256:3994844a2edfec31e36d1fabcd2381e94ce7b0492a5010ed976fb6a7702a0b7a
    - hkccr.ccs.tencentyun.com/tkeimages/tke-cni-agent:v0.1.2
    sizeBytes: 71251871
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/monitor-agent@sha256:5cd31d2c843c7d5d430e45160b1acc7ebfacae67f0fa9c563549cc1654da41b5
    - hkccr.ccs.tencentyun.com/tkeimages/monitor-agent:v1.3.0-abe11a
    sizeBytes: 58900320
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/csi-tencentcloud-cbs@sha256:2c76fc98bf04e1412d7ce3a08be87ab7cc301b3f8655c14afd3f6cdacde72d6e
    - hkccr.ccs.tencentyun.com/tkeimages/csi-tencentcloud-cbs:v2.2.8
    sizeBytes: 41586124
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/tke-bridge-agent@sha256:70f1493795b2ceaa99d067a31d6bafd42b8171dac83e65a70defff1f419792ad
    - hkccr.ccs.tencentyun.com/tkeimages/tke-bridge-agent:v0.1.5
    sizeBytes: 22628290
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/apiserver-proxy@sha256:ae74732ec30ad10d6aa4daf6f9e355f62b77913244d1911744a3095d8890da8b
    - hkccr.ccs.tencentyun.com/tkeimages/apiserver-proxy:v1.3.5
    sizeBytes: 20941553
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/coredns@sha256:7a8fedc352f5e1002504b82a0230c07dff2657d23a2dc6a6db0b47d5b67830f3
    - hkccr.ccs.tencentyun.com/tkeimages/coredns:1.8.4
    sizeBytes: 13706014
  - names:
    - hkccr.ccs.tencentyun.com/tkeimages/csi-node-driver-registrar@sha256:588f77473b5e5e81d141bde83f3287e80ef7441ccf89774691c480254c4cb93a
    - hkccr.ccs.tencentyun.com/tkeimages/csi-node-driver-registrar:v2.0.1
    sizeBytes: 8414160
  - names:
    - hkccr.ccs.tencentyun.com/library/ip-masq-agent@sha256:655c4a06a15ca3252d395e9763505230edd3b8d181e16e039c7ba8fdfc1742e5
    - hkccr.ccs.tencentyun.com/library/ip-masq-agent:v2.6.2
    sizeBytes: 7466018
  - names:
    - hkccr.ccs.tencentyun.com/library/pause@sha256:5ab61aabaedd6c40d05ce1ac4ea72c2079f4a0f047ec1dc100ea297b553539ab
    - hkccr.ccs.tencentyun.com/library/pause:latest
    sizeBytes: 298560
  nodeInfo:
    architecture: amd64
    bootID: 7bbce9fa-feaf-4493-ae26-a49312a5a154
    containerRuntimeVersion: containerd://1.3.4-1-g07bb01e5
    kernelVersion: 5.4.119-19-0009.3
    kubeProxyVersion: v1.22.5-tke.1
    kubeletVersion: v1.22.5-tke.1
    machineID: d3bf44cf57154ba4bd061f8e5f3cdda2
    operatingSystem: linux
    osImage: TencentOS Server 3.1 (Final)
    systemUUID: d3bf44cf-5715-4ba4-bd06-1f8e5f3cdda2`

	yamlD, _ := Yaml2JSON(yData)
	fmt.Println(yamlD)
}
