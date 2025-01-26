package runner

import (
	"fmt"

	"github.com/preinpost/kubeadm-install/pkg/script"
	"github.com/preinpost/kubeadm-install/pkg/utils"
)

type scriptInfo struct {
	filename  string
	content   string
	contentFn func()
	step      string
	useFunc   bool
}

func getScripts(ipAddr string) []scriptInfo {
	return []scriptInfo{
		{filename: "01_vm-env-edit.sh", content: script.VmEnvEditScript, step: "swap disabled", useFunc: false},
		{filename: "02_resolved-edit.sh", content: script.ResolvedEditScript, step: "dns edited", useFunc: false},
		{filename: "03_docker-install.sh", content: script.DockerInstallScript, step: "docker installed", useFunc: false},
		{filename: "04_containerd-edit.sh", content: script.ContainerdEditScript, step: "containerd edited", useFunc: false},
		{filename: "05_iptables-setup.sh", content: script.IptablesSetupScript, step: "iptables setup", useFunc: false},
		{filename: "06_kubeadm-install.sh", content: script.KubeadmInstallScript, step: "kubeadm installed", useFunc: false},
		{filename: "07_kubeadm-init.sh", content: script.KubeadmInitScript(ipAddr), step: "kubeadm init done", useFunc: false}, // ipAddr 사용
		{filename: "08_kubeadm-init-after.sh", contentFn: script.KubeadmControlplaneAfterInitScript, step: "kubeadm init after done", useFunc: true},
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
		if script.useFunc {
			script.contentFn()
		} else {
			if err := utils.RunScript(script.filename); err != nil {
				err := fmt.Errorf("Error executing %s: %v\n", script.filename, err)
				panic(err)
			}
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
