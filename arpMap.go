package main

import (
	"os/exec"
	"log"
	"strings"
)

func getArpMap() map[string]string{
	out, err := exec.Command("arp", "-an").Output()

	if err != nil {
		log.Fatal(err)
	}

		arpMap := make(map[string]string)

	listOfAdds := strings.Split(string(out), "\n")
	i := 0
	for i < len(listOfAdds) -1 {
		temp := strings.Split(listOfAdds[i], " ");
		arpMap[removePara(temp[1])] = temp[3]
		i++
	}
	return arpMap;

}

func removePara(s string) string{
	temp := strings.Split(s, "");
	rs := ""
	i := 1
	for i < len(temp) -1{
		rs = rs + temp[i]
		i++
	}
	return rs
}
