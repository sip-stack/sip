package templates

import (
	"github.com/paulmanoni/sip/pkg/utils"
	"runtime"
)

func GenerateGinGonicProject(Title string) {
	//	project folder
	switch runtime.GOOS {
	case "windows":
		utils.ExecCommand("cmd", []string{"/c", "mkdir", Title}, true)
		utils.ExecCommand("cmd", []string{"/c", "cd", Title}, true)
		utils.ExecCommand("cmd", []string{"/c", "mkdir", "cmd"}, true)
	}
}
