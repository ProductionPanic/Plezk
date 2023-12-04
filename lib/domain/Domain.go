package domain

import (
	"fmt"
	"github.com/ProductionPanic/go-pretty"
	"plezk/lib/admin"
	"strconv"
)

type Domain struct {
	Id       string
	Name     string
	Children []Domain
}

type DomainInfo struct {
	DomainName    string
	SSL           bool
	DiskSpace     string
	HardDiskQuota int
	DomainId      int
	SshAccess     string
	Ip            string
	Description   string
	WwwRoot       string
	Php           bool
	FtpLogin      string
	FtpPassword   string
	TotalSize     string
	Status        string
	MailService   bool
	AccessToPlesk bool
	raw           map[string]interface{}
}

func (d *DomainInfo) getStringFromRaw(key string) string {
	if value, ok := d.raw[key]; ok {
		return value.(string)
	}
	return ""
}

func (d *DomainInfo) getBoolFromRaw(key string) bool {
	correctValues := []string{"true", "1", "On"}

	if value, ok := d.raw[key]; ok {
		for _, correctValue := range correctValues {
			if value.(string) == correctValue {
				return true
			}
		}
	}

	return false
}

func (d *DomainInfo) getIntFromRaw(key string) int {
	var strValue string
	if value, ok := d.raw[key]; ok {
		strValue = value.(string)
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0
	}
	return intValue
}

func (d *DomainInfo) init() {
	d.DomainName = d.getStringFromRaw("Domain name")
	d.SSL = d.getBoolFromRaw("SSL/TLS support")
	d.DiskSpace = d.getStringFromRaw("Disk space used by httpdocs")
	d.HardDiskQuota = d.getIntFromRaw("Hard disk quota")
	d.DomainId = d.getIntFromRaw("Domain ID")
	d.SshAccess = d.getStringFromRaw("SSH access to the server shell under subscription's system user")
	d.Ip = d.getStringFromRaw("IP address")
	d.Description = d.getStringFromRaw("Description")
	d.WwwRoot = d.getStringFromRaw("--WWW-Root--")
	d.Php = d.getBoolFromRaw("PHP support")
	d.FtpLogin = d.getStringFromRaw("FTP Login")
	d.FtpPassword = d.getStringFromRaw("FTP Password")
	d.TotalSize = d.getStringFromRaw("Total size")
	d.Status = d.getStringFromRaw("Domain status")
	d.MailService = d.getBoolFromRaw("Mail service")
	d.AccessToPlesk = d.getBoolFromRaw("Access to Plesk")
}

func (d *DomainInfo) itemToString(item string, value string) string {
	return pretty.Parse(fmt.Sprintf("[bold,blue]%s[][bold,cyan]:[] %s\n", item, value))
}

func (d *DomainInfo) GetInfoString() string {
	return pretty.Parse(`
[bold,blue]` + d.DomainName + `[]
		
[bold,cyan]ssl:[] ` + strconv.FormatBool(d.SSL) + `
[bold,cyan]disk space:[] ` + d.DiskSpace + `
[bold,cyan]hard disk quota:[] ` + strconv.Itoa(d.HardDiskQuota) + `
[bold,cyan]domain id:[] ` + strconv.Itoa(d.DomainId) + `
[bold,cyan]ssh access:[] ` + d.SshAccess + `
[bold,cyan]ip:[] ` + d.Ip + `
[bold,cyan]description:[] ` + d.Description + `
[bold,cyan]www root:[] ` + d.WwwRoot + `
[bold,cyan]php:[] ` + strconv.FormatBool(d.Php) + `
[bold,cyan]ftp login:[] ` + d.FtpLogin + `
[bold,cyan]ftp password:[] ` + d.FtpPassword + `
[bold,cyan]total size:[] ` + d.TotalSize + `
[bold,cyan]status:[] ` + d.Status + `
[bold,cyan]mail service:[] ` + strconv.FormatBool(d.MailService) + `
[bold,cyan]access to plesk:[] ` + strconv.FormatBool(d.AccessToPlesk) + `
	`)
}

func (d *DomainInfo) HasSshAccess() bool {
	sa := d.SshAccess
	if len(sa) == 0 {
		return false
	}
	first_char := sa[0]
	if first_char == '/' {
		return true
	}

	return false
}

func (d *Domain) Info() *DomainInfo {
	key_values := FetchDomainInfo(d)
	info := DomainInfo{
		raw: key_values,
	}
	info.init()
	return &info
}

func (d *Domain) Delete() {
	Delete(d.Name)
	pretty.Println("[bold,green]Domain deleted[]")
}

func (d *Domain) SSLEnable() {
	SetSsl(d.Name, admin.Info().GetEmail())
}

func (d *Domain) SSLDisable() {
	RemoveSsl(d.Name)
}

func (d *Domain) SshEnable() {
	SetSsh(d.Name)
}

func (d *Domain) SshDisable() {
	RemoveSsh(d.Name)
}
