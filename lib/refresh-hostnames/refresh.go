package refresh

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"strings"

	mEC2 "github.com/blinchik/go-aws/lib/manage-ec2"
)

var home, err = os.UserHomeDir()

func HostnamesRefresh() {

	summary := mEC2.DescribeAllMentionedTag("Name")

	sshFile := filepath.FromSlash(fmt.Sprintf("%s/.ssh/config", home))

	file, err := os.OpenFile(sshFile, os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sshConfig, err := ioutil.ReadAll(file)

	for i := 0; i < len(summary.InstanceId); i++ {

		if strings.Contains(string(sshConfig), *summary.TagValue[i]) {

			SubSSHConfig := strings.Split(string(sshConfig), *summary.TagValue[i])

			SubSSHConfig = strings.Split(SubSSHConfig[1], "HostName")

			reg := regexp.MustCompile("\\A.*")

			FirstLine := reg.FindString(SubSSHConfig[1])

			ip := strings.Replace(FirstLine, " ", "", -1)

			if summary.PublicIp[i] != nil {

				sshConfig = []byte(strings.Replace(string(sshConfig), ip, *summary.PublicIp[i], -1))

			} else {

				sshConfig = []byte(strings.Replace(string(sshConfig), ip, *summary.PrivateIpAddress[i], -1))

			}

		}

	}

	file, err = os.OpenFile(sshFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	w := bufio.NewWriter(file)

	_, err = w.WriteString(string(sshConfig))

	log.Println("HostNames are up-to-date")

	w.Flush()

}
