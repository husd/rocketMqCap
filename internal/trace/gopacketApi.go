package trace

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
)

func captureByRule(r *rule) {

	//	获取 libpcap 的版本
	version := pcap.Version()
	fmt.Println(version)
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal("[error] 没有找到网卡设置，请检查。")
	}
	for _, d := range devices {
		go printDevice(d.Name, r)
	}
}

/**
 * 对gopacket的一些封装
 * @author hushengdong
 */
func printDevice(deviceName string, r *rule) {

	snapLen := int32(65535)
	filter := r.filter
	//打开网络接口，抓取在线数据
	handle, err := pcap.OpenLive(deviceName, snapLen, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("pcap open live failed: %v", err)
		return
	}
	// 设置过滤器
	if err := handle.SetBPFFilter(filter); err != nil {
		fmt.Printf("set bpf filter failed: %v", err)
		return
	}
	defer handle.Close()
	// 抓包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	//fmt.Printf("device:%v, snapLen:%v, port:%v\n", deviceName, snapLen, port)
	//fmt.Println("filter:", filter)
	for packet := range packetSource.Packets() {
		if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
			fmt.Println("unexpected packet")
			continue
		}
		// tcp 层
		tcp := packet.TransportLayer().(*layers.TCP)
		//fmt.Printf("device:%v filter:%s tcp:%v \n", deviceName,filter,tcp)
		// tcp payload，也即是tcp传输的数据
		//fmt.Printf("tcp payload:%v \n", tcp.Payload)
		r.ruleHandle(&tcp.Payload)
	}
}
