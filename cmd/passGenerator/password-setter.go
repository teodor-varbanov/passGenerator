package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func psSingleQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}

func setWinPassword(computername string, username string, secret string) error {

	adminUser := fmt.Sprintf(`%s\%s`, computername, localAdminUser)

	psScript := fmt.Sprintf(
		`$plainSec = [Console]::In.ReadLine()
		$adminUser = %s
		$adminPasswordPlain = %s
		$adminPassword = ConvertTo-SecureString $adminPasswordPlain -AsPlainText -Force
		$credentials = New-Object System.Management.Automation.PSCredential($adminUser, $adminPassword)
		Invoke-Command -ComputerName %s -Credential $credentials -ScriptBlock {
			param([string]$u,[string]$p)
			$password = ConvertTo-SecureString $p -AsPlainText -Force
			Set-LocalUser -Name $u -Password $password
} -ArgumentList %s, $plainSec`,
		psSingleQuote(adminUser),
		psSingleQuote(localAdminPassword),
		psSingleQuote(computername),
		psSingleQuote(username),
	)

	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", psScript)
	cmd.Stdin = strings.NewReader(secret + "\n")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to set password for user %s on %s: %w\nOutput: %s", username, computername, err, string(output))
	}

	return nil
}