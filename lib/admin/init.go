package admin

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type AdminInfo struct {
	Email string
	Name  string
}

func (a *AdminInfo) GetEmail() string {
	return a.Email
}

func (a *AdminInfo) GetName() string {
	return a.Name
}

func Info() *AdminInfo {
	command := "plesk bin admin -i"
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	reg := regexp.MustCompile(`([^:]+):\s(.+)`)
	var raw map[string]string = make(map[string]string)
	for _, line := range lines {
		match := reg.FindStringSubmatch(line)
		if len(match) == 3 {
			raw[match[1]] = match[2]
		}
	}

	// check if we have all the keys we need
	if _, ok := raw["email"]; !ok {
		log.Fatal("Could not find email in admin info")
	}
	if _, ok := raw["pname"]; !ok {
		log.Fatal("Could not find name in admin info")
	}

	return &AdminInfo{
		Email: raw["email"],
		Name:  raw["pname"],
	}
}
