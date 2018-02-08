package setup

import (
	"log"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func NewCommon() *Common {
	data := Common{}
	data.SetHost()
	data.SetCpu()
	data.SetDisk()
	data.SetDocker()
	data.SetLoad()
	data.SetMem()
	data.SetNet()
	data.SetEnv()
	return &data
}

type Common struct {
	HostInfoStat        *host.InfoStat
	HostUserStat        []host.UserStat
	HostTemperatureStat []host.TemperatureStat
	CpuTimesStat        []cpu.TimesStat
	CpuInfoStat         []cpu.InfoStat
	CpuCounts           int
	CpuPercent          []float64
	DiskUsageStat       []disk.UsageStat
	DiskPartitionStat   []disk.PartitionStat
	DiskIOCountersStat  map[string]disk.IOCountersStat
	// DockerCgroupMemStat       docker.CgroupMemStat
	DockerCgroupDockerStat []docker.CgroupDockerStat
	LoadAvgStat            *load.AvgStat
	LoadMiscStat           *load.MiscStat
	MemVirtualMemoryStat   *mem.VirtualMemoryStat
	MemSwapMemoryStat      *mem.SwapMemoryStat
	NetIOCountersStat      []net.IOCountersStat
	NetAddr                net.Addr
	NetConnectionStat      net.ConnectionStat
	NetProtoCountersStat   []net.ProtoCountersStat
	NetInterfaceStat       []net.InterfaceStat
	NetFilterStat          []net.FilterStat
	// Process                   process.Process
	// ProcessOpenFilesStat      process.OpenFilesStat
	// ProcessMemoryInfoStat     process.MemoryInfoStat
	// ProcessSignalInfoStat     process.SignalInfoStat
	// ProcessRlimitStat         process.RlimitStat
	// ProcessIOCountersStat     process.IOCountersStat
	// ProcessNumCtxSwitchesStat process.NumCtxSwitchesStat
	Env []string
}

func (this *Common) SetHost() error {
	infostat, err := host.Info()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.HostInfoStat = infostat

	userstat, err := host.Users()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.HostUserStat = userstat

	temperature, err := host.SensorsTemperatures()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.HostTemperatureStat = temperature
	return nil
}

func (this *Common) SetCpu() error {
	timestat, err := cpu.Times(false)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.CpuTimesStat = timestat

	infostat, err := cpu.Info()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.CpuInfoStat = infostat

	counts, err := cpu.Counts(true)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.CpuCounts = counts

	//获取cpu使用频率
	// percent, err := cpu.Percent(time.Second, false)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// this.CpuPercent = percent

	return nil
}

func (this *Common) SetDisk() error {
	usages := []disk.UsageStat{}

	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.DiskPartitionStat = partitions

	for _, data := range partitions {
		usage, err := disk.Usage(data.Mountpoint)
		if err != nil {
			log.Println(err.Error())
			break
			return err
		}
		// log.Println("usage", usage)
		usages = append(usages, *usage)
	}
	this.DiskUsageStat = usages

	iocounter, err := disk.IOCounters()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.DiskIOCountersStat = iocounter
	return nil
}

func (this *Common) SetDocker() error {
	cgdockerstat, err := docker.GetDockerStat()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.DockerCgroupDockerStat = cgdockerstat
	return nil
}

func (this *Common) SetLoad() error {
	avg, err := load.Avg()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.LoadAvgStat = avg

	misc, err := load.Misc()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.LoadMiscStat = misc
	return nil
}

func (this *Common) SetMem() error {
	swap, err := mem.SwapMemory()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.MemSwapMemoryStat = swap

	virtual, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.MemVirtualMemoryStat = virtual
	return nil
}

func (this *Common) SetNet() error {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.NetInterfaceStat = interfaces

	iocounterstat, err := net.IOCounters(false)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.NetIOCountersStat = iocounterstat

	filecounters, err := net.FilterCounters()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.NetFilterStat = filecounters

	protocounters, err := net.ProtoCounters([]string{"ip", "icmp", "icmpmsg", "tcp", "udp"})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	this.NetProtoCountersStat = protocounters
	return nil
}

// func (this *Common) SetProcess() {

// }

func (this *Common) SetEnv() error {
	this.Env = os.Environ()
	return nil
}
