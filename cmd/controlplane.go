package cmd

import (
	"fmt"

	"github.com/preinpost/kubeadm-install/pkg/runner"
	"github.com/spf13/cobra"
)

var ipAddr string
var isExistRemainTmpFlag = false

var controlplaneCmd = &cobra.Command{
	Use:   "controlplane",
	Short: "controlplane node 설치",
	Long:  `controlplane node 설치 합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if ipAddr == "" {
			return fmt.Errorf("ip 주소를 입력해주세요")
		}

		runner.Run(ipAddr, isExistRemainTmpFlag)

		return nil
	},
}

func init() {
	controlplaneCmd.Flags().StringVar(&ipAddr, "ip", "", "설치하기 위한 ip 주소를 입력해주세요")
	controlplaneCmd.Flags().BoolVar(&isExistRemainTmpFlag, "remain-script", false, "설치하기 위한 ip 주소를 입력해주세요")
	controlplaneCmd.MarkFlagRequired("ip")

	rootCmd.AddCommand(controlplaneCmd)

}
