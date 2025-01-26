package main

import (
	"fmt"

	"github.com/preinpost/kubeadm-install/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("모든 스크립트가 성공적으로 실행되었습니다.")
}
