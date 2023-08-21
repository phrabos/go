package loops

import "fmt"
import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipa IPAddr) String() string {
	str := ""
	for _, v := range ipa {
		intByte := int(v)
		strByte := strconv.Itoa(intByte)
		str += strByte + "."
	}
	return str
}

func Loops() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
