package domain

import (
	"encoding/xml"
	"fmt"
	"log"
	"os/exec"
	"plezk/lib/admin"
	"regexp"
	"strings"

	"github.com/ProductionPanic/go-input"
	"github.com/ProductionPanic/go-pretty"
)

type Resultset struct {
	XMLName xml.Name `xml:"resultset"`
	Rows    []Row    `xml:"row"`
}

type Row struct {
	XMLName xml.Name `xml:"row"`
	Fields  []Field  `xml:"field"`
}

type Field struct {
	XMLName xml.Name `xml:"field"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

func get_domains_by_parent(parentId string) []Domain {

	cmd := "plesk db 'select name, id from domains where parentDomainId = %s or parentDomainId is null' --xml"
	cmd = fmt.Sprintf(cmd, parentId)
	xmlOutput, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	var result Resultset
	err = xml.Unmarshal([]byte(xmlOutput), &result)
	if err != nil {
		log.Fatal(err)
	}
	var output []Domain
	for _, row := range result.Rows {
		output = append(output, Domain{Id: row.Fields[1].Value, Name: row.Fields[0].Value, Children: get_domains_by_parent(row.Fields[1].Value)})
	}
	return output
}

func List() []Domain {
	return get_domains_by_parent("0")
}

func ListFlat() []Domain {
	domains := List()
	var output []Domain
	for _, domain := range domains {
		output = append(output, domain)
		for _, childDomain := range domain.Children {
			output = append(output, childDomain)
		}
	}
	return output
}

func Delete(domain string) bool {
	command := fmt.Sprintf("plesk bin domain --remove %s", domain)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func Create(domain string) bool {
	pass, er := input.GetPassword("[bold,blue]Enter password for the domain:[]")
	if er != nil {
		log.Fatal(er)
	}
	command := fmt.Sprintf("plesk bin domain --create %[1]s -www-root %[1]s -php true -hosting true -ip %[2]s -login %[1]s -passwd %[3]s", domain, admin.Info().GetIp(), pass)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func CreateSubdomain(subdomain, domain string) bool {
	pass, er := input.GetText("[bold,blue]Enter password for the domain:[]")
	if er != nil {
		log.Fatal(er)
	}
	command := fmt.Sprintf("plesk bin subdomain --create %[1]s -domain %[2]s -www-root %[1]s.%[2]s -php true -hositng true -ip %[3]s -login %[1]s -passwd %[4]s", subdomain, domain, admin.Info().GetIp(), pass)
	_, err := exec.Command("bash", "-c", command).Output()
	return err == nil
}

func printRecursive(domains []Domain, depth int) {
	for i, domain := range domains {
		for i := 0; i < depth; i++ {
			pretty.Print("  ")
		}
		if depth > 0 {
			is_last := i == len(domains)-1
			if is_last {
				pretty.Print("[cyan]└ []")
			} else {
				pretty.Print("[cyan]├ []")
			}
		}
		fmt.Println(domain.Name)
		printRecursive(domain.Children, depth+1)
	}
}

func PrintList() {
	domains := List()
	pretty.Println("[bold,blue]Domains:[]")
	printRecursive(domains, 0)
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

func RemoveSsl(domain string) bool {
	// fetch domains from plesk
	bashCommand := fmt.Sprintf("plesk bin extension --exec letsencrypt cli.php --remove -d %s", domain)
	_, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		return false
	}
	return true
}

func SetSsh(domain string) bool {
	// fetch domains from plesk
	bashCommand := fmt.Sprintf("plesk bin subscription --update-web-server-settings %s -ssh true", domain)
	_, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		return false
	}
	return true
}

func RemoveSsh(domain string) bool {
	// fetch domains from plesk
	bashCommand := fmt.Sprintf("plesk bin subscription --update-web-server-settings %s -ssh false", domain)
	_, err := exec.Command("bash", "-c", bashCommand).Output()
	if err != nil {
		return false
	}
	return true
}

func FetchDomainInfo(domain *Domain) map[string]interface{} {
	cmd := "plesk bin domain --info %s"
	reg := "^\\s*([^:]+):\\s*(.*)$"
	cmd = fmt.Sprintf(cmd, domain.Name)
	strOutput, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	result = make(map[string]interface{})
	for _, line := range strings.Split(string(strOutput), "\n") {
		if strings.Contains(line, ":") {
			matches := regexp.MustCompile(reg).FindStringSubmatch(line)
			result[matches[1]] = matches[2]
		}
	}
	return result
}

func Get(do string) Domain {
	domains := ListFlat()
	for _, domain := range domains {
		if domain.Name == do {
			return domain
		}
	}
	return Domain{}
}
