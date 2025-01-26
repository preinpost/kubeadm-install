package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var OutDir = "./tmp"

func mkdir() {
	if err := os.Mkdir(OutDir, os.ModePerm); err != nil && !os.IsExist(err) {
		panic("디렉토리 생성에 실패했습니다.")
	}
}

func WriteScript(filename, content string) error {
	mkdir()

	fileSavePath := filepath.Join(OutDir, filename)
	err := os.WriteFile(fileSavePath, []byte(content), 0755)
	if err != nil {
		return fmt.Errorf("파일 작성 오류: %v", err)
	}

	return nil
}

func RunScript(filename string) error {
	fileSavePath := filepath.Join(OutDir, filename)

	cmd := exec.Command("bash", fileSavePath)
	cmd.Stdin = strings.NewReader("y\n")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CleanUp() {
	// 모든 스크립트가 성공적으로 실행된 후 *.sh 파일 삭제

	if err := os.RemoveAll(OutDir); err != nil {
		fmt.Printf("Error deleting ./out directory: %v\n", err)
		return
	}
}
