package main

import(
	"net"
	"fmt"
)

func main(){
	_, ipNet, err:= net.ParseCIDR("192.168.0.0/16")
	baseIP := ipNet.IP
	oldPrefix,_:= ipNet.Mask.Size()
	// fmt.Println(baseIP, oldPrefix)
	if err!=nil{
		panic(err)
	}
	newPrefix:= 24;
	numSubsets:= 1 << (newPrefix - oldPrefix)

	var subnets[] net.IPNet

	for i:=0; i<numSubsets; i++{
		subnetIP:= make(net.IP, len(baseIP))
		copy(subnetIP, baseIP)

		subnetIP[2] = byte(i)
		subnet:= net.IPNet{
			IP: subnetIP,
			Mask: net.CIDRMask(newPrefix ,32),//ffffff00
		}
		subnets = append(subnets, subnet)
	}

	areas:= []string{"home", "office", "cafe", "Park", "library"}

	for i, area := range areas{
		if(i < len(subnets)){
			fmt.Printf("%-10s : %s\n", area, subnets[i].String())
		}
	}


}