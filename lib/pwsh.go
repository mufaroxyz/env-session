package lib

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"syscall"
)

type Powershell struct {
	Path string
	Ver  int8
}

func GetPowershellInstallation() Powershell {
	powershellPath, err := FindExecutablePath("pwsh")
	if err != nil {
		fmt.Println(err)
		var hook = User32Init()
		MessageBox(hook, 0, "PowerShell not found, are you sure you have powershell installed?", "Error", MB_OK|MB_ICONSTOP|MB_SYSTEMMODAL)
		os.Exit(1)
	}

	var regex = `C:\\Program Files\\PowerShell\\(\d)\\pwsh.exe`
	var re = regexp.MustCompile(regex)
	var matches = re.FindStringSubmatch(powershellPath)
	var ver int8

	if len(matches) > 1 {
		ver = int8(matches[1][0] - '0')
	} else {
		ver = 7
	}

	var powershell = Powershell{
		Path: powershellPath,
		Ver:  ver,
	}

	return powershell
}

type PowershellArg struct {
	Env string
	Val string
}

func fmtPowershellEnvs(args []PowershellArg) string {
	var envs = ""

	for _, arg := range args {
		envs += fmt.Sprintf("Set-Alias -Name \"%s\" -Value \"%s\";", arg.Env, arg.Val)
	}

	return envs
}

// TODO silence the unnecessary logs from applications that are ran from powershell
func composeFullCommand(args []PowershellArg) string {
	var envs = fmtPowershellEnvs(args)

	var cmd = fmt.Sprintf(`& (Get-Process -Id $pid).Path -NoExit { Clear-Host; %s Write-Host "[env-session] Overrided environment variables: %x"; Write-Host ""; }`,
		IfThenElse(envs == "", "", envs+" "),
		len(args),
	)

	return cmd
}

func RunPowershellCommand(c Config) *exec.Cmd {
	var args []PowershellArg

	for key, val := range c {
		if key == "$powershell_path" {
			continue
		}

		args = append(args, PowershellArg{Env: key, Val: val})
	}

	var cmd = composeFullCommand(args)
	var powershell = exec.Command(c.Get("$powershell_path"), "-Command", cmd)
	fmt.Println(powershell.String())
	powershell.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	powershell.Stdout = os.Stdout
	powershell.Stderr = os.Stderr
	powershell.Stdin = os.Stdin
	err := powershell.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return powershell
}
