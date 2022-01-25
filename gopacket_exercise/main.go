package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"os"
	"strings"
	"time"

	//"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

//返回待抓取的网口
func findAllDevs(ip string) string {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)

			if address.IP.String() == ip {
				return device.Name
			}
		}
	}
	return ""
}

func main() {
	var (
		deviceName  string        = ""
		snapShotLen int32         = 1024
		promiscuous bool          = false
		t           time.Duration = 1
		handle      *pcap.Handle
		err         error
		packetCount int = 0
	)

	deviceName = findAllDevs("192.168.17.72")

	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snapShotLen), layers.LinkTypeEthernet)
	defer f.Close()

	// open the device for capturing
	handle, err = pcap.OpenLive(deviceName, snapShotLen, promiscuous, t)
	if err != nil {
		log.Fatal("pcap OpenLive failed\n")
	}
	defer handle.Close()

	// start processing packets
	fmt.Println("linkType:", handle.LinkType())
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		printPacketInfo(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}
}

func printPacketInfo(packet gopacket.Packet) {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("enter ethernet detected")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Src Mac:", ethernetPacket.SrcMAC)
		fmt.Println("Dst Mac:", ethernetPacket.DstMAC)
	}

	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
	if ipv4Layer != nil {
		ipv4Packet, _ := ipv4Layer.(*layers.IPv4)
		fmt.Println("Src Ip:", ipv4Packet.SrcIP)
		fmt.Println("Dst Ip:", ipv4Packet.DstIP)
		fmt.Println("Ip version:", ipv4Packet.Version)
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcpPacket, _ := tcpLayer.(*layers.TCP)
		fmt.Println("Src port:", tcpPacket.SrcPort)
		fmt.Println("Dst port:", tcpPacket.DstPort)
	}

	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Printf("application payload:%s\n", applicationLayer.Payload())
		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}
}
