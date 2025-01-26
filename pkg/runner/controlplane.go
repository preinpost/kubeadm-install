package runner

import (
	"fmt"

	"github.com/preinpost/kubeadm-install/pkg/script"
	"github.com/preinpost/kubeadm-install/pkg/utils"
)

type scriptInfo struct {
	filename string
	content  string
	step     string
}

func getScripts(ipAddr string) []scriptInfo {
	return []scriptInfo{
		{"01_vm-env-edit.sh", script.VmEnvEditScript, "swap diabled"},
		{"02_resolved-edit.sh", script.ResolvedEditScript, "dns edited"},
		{"03_docker-install.sh", script.DockerInstallScript, "docker installed"},
		{"04_containerd-edit.sh", script.ContainerdEditScript, "containerd edited"},
		{"05_iptables-setup.sh", script.IptablesSetupScript, "iptables setup"},
		{"06_kubeadm-install.sh", script.KubeadmInstallScript, "kubeadm installed"},
		{"07_kubeadm-init.sh", script.KubeadmInitScript(ipAddr), "kubeadm init done"}, // ipAddr 사용
		{"08_kubeadm-init-after.sh", script.KubeadmControlplaneAfterInitScript, "kubeadm init after done"},
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
		fmt.Printf("==== %s ====\n", script.step)
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
