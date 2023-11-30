package domain

import (
	"fmt"
	"github.com/ProductionPanic/go-input"
	"log"
	"os/exec"
	"plezk/lib/admin"
	"strings"
)

func List() []string {
	// fetch domains from plesk
	bashCommand := "plesk bin domain --list"
	output, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		log.Fatal(err)
	}
	// split output into lines
	lines := strings.Split(string(output), "\n")

	return lines
}

func Delete(domain string) bool {
	command := fmt.Sprintf("plesk bin domain --remove %s", domain)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func Create(domain string) bool {
	pass := input.GetText("[bold,blue]Enter password for the domain:[]")
	command := fmt.Sprintf("plesk bin domain --create %[1]s -www-root %[1]s -php true -hosting true -ip %[2]s -login %[1]s -passwd %[3]s", domain, admin.Info().GetIp(), pass)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func CreateSubdomain(subdomain, domain string) bool {
	pass := input.GetText("[bold,blue]Enter password for the domain:[]")
	command := fmt.Sprintf("plesk bin subdomain --create %[1]s -domain %[2]s -www-root %[1]s.%[2]s -php true -hositng true -ip %[3]s -login %[1]s -passwd %[4]s", subdomain, domain, admin.Info().GetIp(), pass)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func SetSsl(domain, admin string) bool {
	// fetch domains from plesk
	bashCommand := fmt.Sprintf("plesk bin extension --exec letsencrypt cli.php -d %s -m %s", domain, admin)
	_, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		return false
	}
	return true
}
