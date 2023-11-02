package domain

import (
	"fmt"
	"log"
	"os/exec"
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
	command := fmt.Sprintf("plesk bin domain --create %[1]s -www-root %[1]s", domain)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func CreateSubdomain(subdomain, domain string) bool {
	command := fmt.Sprintf("plesk bin subdomain --create %[1]s -domain %[2]s -www-root %[1]s.%[2]s", subdomain, domain)
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
