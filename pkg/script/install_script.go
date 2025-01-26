package script

import "fmt"

// 각 스크립트를 문자열로 정의합니다.
var VmEnvEditScript = `#!/bin/bash
swapoff -a
sed -i '/swap/s/^/#/' /etc/fstab
ufw disable
`
var ResolvedEditScript = `#!/bin/bash
sudo sed -i'.orig' -e 's/^#DNS=$/DNS=1.1.1.1/' /etc/systemd/resolved.conf
sudo systemctl restart systemd-resolved
`

var DockerInstallScript = `#!/bin/bash
sudo apt-get update && 
sudo apt-get install -y ca-certificates curl && 
sudo install -m 0755 -d /etc/apt/keyrings && 
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc && 
sudo chmod a+r /etc/apt/keyrings/docker.asc && 
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | 
sudo tee /etc/apt/sources.list.d/docker.list > /dev/null && 
sudo apt-get update && 
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin && 
sudo docker run --rm hello-world
`

var ContainerdEditScript = `#!/bin/bash
cat <<EOL | sudo tee /etc/containerd/config.toml
version = 2
[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
   [plugins."io.containerd.grpc.v1.cri".containerd]
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
          [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
            SystemdCgroup = true
EOL
sudo systemctl restart containerd
`

var IptablesSetupScript = `#!/bin/bash
modprobe br_netfilter
echo "br_netfilter" | sudo tee /etc/modules-load.d/k8s.conf
cat <<EOL | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOL
sudo sed -i'.orig' -e '/^#net\.ipv4\.ip_forward=1/s/^#//' /etc/sysctl.conf
sudo sysctl --system
`

var KubeadmInstallScript = `#!/bin/bash
sudo apt-get update &&
sudo apt-get install -y apt-transport-https ca-certificates curl gpg &&
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg &&
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list &&
sudo apt-get update &&
sudo apt-get install -y kubelet kubeadm kubectl &&
sudo apt-mark hold kubelet kubeadm kubectl &&
sudo systemctl enable --now kubelet
`

var KubeadmInitScript = func(ipAddr string) string {
	return fmt.Sprintf(`#!/bin/bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=%s --pod-network-cidr=10.244.0.0/16
`, ipAddr)
}