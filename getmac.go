package main // import "getmac"

import (
	"io/ioutil"
	"log"
	"net"
)

func getMacAddr() ([]string, error) {
	macs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var results []string
	for _, mac := range macs {
		result := mac.HardwareAddr.String()
		if result != "" {
			results = append(results, result)
		}
	}
	return results, nil
}

func main() {
	macs, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}

	result := ""

	for _, mac := range macs {
		// fmt.Println(a)
		result += mac + "\r\n"
	}

	err = ioutil.WriteFile("mac.log", []byte(result), 0)
	if err != nil {
		log.Fatal(err)
	}
}
