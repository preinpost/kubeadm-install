package runner

import (
	"fmt"

	"github.com/preinpost/kubeadm-install/pkg/script"
	"github.com/preinpost/kubeadm-install/pkg/utils"
)

type scriptInfo struct {
	filename string
	content  string
}

func getScripts(ipAddr string) []scriptInfo {
	return []scriptInfo{
		{"01_vm-env-edit.sh", script.VmEnvEditScript},
		{"02_resolved-edit.sh", script.ResolvedEditScript},
		{"03_docker-install.sh", script.DockerInstallScript},
		{"04_containerd-edit.sh", script.ContainerdEditScript},
		{"05_iptables-setup.sh", script.IptablesSetupScript},
		{"06_kubeadm-install.sh", script.KubeadmInstallScript},
		{"07_kubeadm-init.sh", script.KubeadmInitScript(ipAddr)}, // ipAddr 사용
	}
}

// Run 함수는 scripts를 실행
func writeScripts(ipAddr string) {

	scriptList := getScripts(ipAddr)

	for _, script := range scriptList {
		if err := utils.WriteScript(script.filename, script.content); err != nil {
			err := fmt.Errorf("Error executing %s: %v\n", script.filename, err)
			panic(err)
		}
	}
}

func runScripts(ipAddr string) {

	scriptList := getScripts(ipAddr)

	for _, script := range scriptList {
		if err := utils.RunScript(script.filename); err != nil {
			err := fmt.Errorf("Error executing %s: %v\n", script.filename, err)
			panic(err)
		}
	}
}

func Run(ipAddr string, isExistRemainFlag bool) {
	writeScripts(ipAddr)
	runScripts(ipAddr)
	// remain flag 가 없을 때만 cleanup 수행
	if !isExistRemainFlag {
		utils.CleanUp()
	}
}
