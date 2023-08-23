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
// []int{2,4,7,9.13}
func binarySearch(s []int, target int) int {
	high := len(s) - 1
	low := 0
	
	for low <= high {
		mid := (high + low) / 2
		if s[mid] == target {
			return i
		} else if s[mid] < target {
			low = mid + 1
		} else s[mid] > target {
			high = mid - 1
		}
	}

	return -1
}
