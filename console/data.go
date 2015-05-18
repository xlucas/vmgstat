package console

import (
	"math"
	"time"

	"github.com/cloudfoundry/gosigar"
	"github.com/xlucas/go-vmguestlib/vmguestlib"
)

type PrintFunc func(*Console, *Data, *Data, *vmguestlib.Session)

func AppendField(fields *map[string]PrintFunc, order *[]string, name string, value PrintFunc) {
	(*fields)[name] = value
	*order = append(*order, name)
}

type Data struct {
	CPUReservation     uint32    // CPU Guest
	CPULimit           uint32    // CPU Guest
	CPUShares          uint32    // CPU Guest
	CPUStolen          uint64    // CPU Guest
	CPUUsed            uint64    // CPU Guest
	HostCPUUsed        uint64    // CPU Host
	HostNumCPUCores    uint32    // CPU Host
	HostProcessorSpeed uint32    // CPU Host
	MemReservation     uint32    // Mem Guest
	MemLimit           uint32    // Mem Guest
	MemShares          uint32    // Mem Guest
	MemMapped          uint32    // Mem Guest
	MemActive          uint32    // Mem Guest
	MemOverhead        uint32    // Mem Guest
	MemBallooned       uint32    // Mem Guest
	MemSwapped         uint32    // Mem Guest
	MemShared          uint32    // Mem Guest
	MemSharedSaved     uint32    // Mem Guest
	MemUsed            uint32    // Mem Guest
	MemTargetSize      uint64    // Mem Guest
	HostMemSwapped     uint64    // Mem Host
	HostMemShared      uint64    // Mem Host
	HostMemUsed        uint64    // Mem Host
	HostMemPhys        uint64    // Mem Host
	HostMemPhysFree    uint64    // Mem Host
	HostMemKernOvhd    uint64    // Mem Host
	HostMemMapped      uint64    // Mem Host
	HostMemUnmapped    uint64    // Mem Host
	TimeElasped        uint64    // Timer
	CurrentTime        time.Time // Time
	MemorySize         uint64    // Memory Size of the VM
	SwapSize           uint64    // Memory Size of the VM
}

func (d *Data) Refresh(s *vmguestlib.Session) {
	m := new(sigar.Mem)
	w := new(sigar.Swap)
	_ = m.Get()
	_ = w.Get()
	d.CPUReservation, _ = s.GetCPUReservation()
	d.CPULimit, _ = s.GetCPULimit()
	d.CPUShares, _ = s.GetCPUShares()
	d.CPUStolen, _ = s.GetCPUStolen()
	d.CPUUsed, _ = s.GetCPUUsed()
	d.HostCPUUsed, _ = s.GetHostCPUUsed()
	d.HostNumCPUCores, _ = s.GetHostNumCPUCores()
	d.HostProcessorSpeed, _ = s.GetHostProcessorSpeed()
	d.MemReservation, _ = s.GetMemReservation()
	d.MemLimit, _ = s.GetMemLimit()
	d.MemShares, _ = s.GetMemShares()
	d.MemMapped, _ = s.GetMemMapped()
	d.MemActive, _ = s.GetMemActive()
	d.MemOverhead, _ = s.GetMemOverhead()
	d.MemBallooned, _ = s.GetMemBallooned()
	d.MemSwapped, _ = s.GetMemSwapped()
	d.MemShared, _ = s.GetMemShared()
	d.MemSharedSaved, _ = s.GetMemSharedSaved()
	d.MemUsed, _ = s.GetMemUsed()
	d.MemTargetSize, _ = s.GetMemTargetSize()
	d.HostMemSwapped, _ = s.GetHostMemSwapped()
	d.HostMemShared, _ = s.GetHostMemShared()
	d.HostMemUsed, _ = s.GetHostMemUsed()
	d.HostMemPhys, _ = s.GetHostMemPhys()
	d.HostMemPhysFree, _ = s.GetHostMemPhysFree()
	d.HostMemKernOvhd, _ = s.GetHostMemKernOvhd()
	d.HostMemMapped, _ = s.GetHostMemMapped()
	d.HostMemUnmapped, _ = s.GetHostMemUnmapped()
	d.TimeElasped, _ = s.GetTimeElapsed()
	d.CurrentTime = time.Now()
	d.MemorySize = m.Total / uint64(math.Pow(2, 20))
	d.SwapSize = w.Total / uint64(math.Pow(2, 20))
}

func PrintCPUUsed(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetCPUUsed(); err != nil {
		c.WriteNaCol()
	} else {
		c.WritePercentCol(float64((n.CPUUsed - o.CPUUsed) / (n.TimeElasped - o.TimeElasped)))
	}
}

func PrintCPUStolen(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetCPUStolen(); err != nil {
		c.WriteNaCol()
	} else {
		c.WritePercentCol(float64((n.CPUStolen - o.CPUStolen) / (n.TimeElasped - o.TimeElasped)))
	}
}

func PrintCurrentTime(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	c.WriteTimeCol(n.CurrentTime)
}

func PrintCPULimit(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetCPULimit(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.CPULimit)
	}
}

func PrintCPUReservation(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetCPUReservation(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.CPUReservation)
	}
}

func PrintCPUShares(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetCPUShares(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.CPUShares)
	}
}

func PrintHostCPUUsed(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetHostCPUUsed(); err != nil {
		c.WriteNaCol()
	} else {
		c.WritePercentCol(float64((n.HostCPUUsed - o.HostCPUUsed) / (n.TimeElasped - o.TimeElasped)))
	}
}

func PrintHostProcessorSpeed(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetHostProcessorSpeed(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.HostProcessorSpeed)
	}
}

func PrintHostNumCPUCores(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetHostNumCPUCores(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.HostNumCPUCores)
	}
}

func PrintMemActive(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemActive(); err != nil {
		c.WriteNaCol()
	} else {
		c.WritePercentCol(100.0 * float64(n.MemActive) / float64(n.MemorySize))
	}
}

func PrintMemBallooned(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemBallooned(); err != nil {
		c.WriteNaCol()
	} else {
		c.WritePercentCol(100.0 * float64(n.MemBallooned) / float64(n.MemorySize))
	}
}

func PrintMemLimit(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemLimit(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemLimit)
	}
}

func PrintMemMapped(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemMapped(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemMapped)
	}
}

func PrintMemOverhead(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemOverhead(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemOverhead)
	}
}

func PrintMemReservation(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemReservation(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemReservation)
	}
}

func PrintMemShares(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemShares(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemShares)
	}
}

func PrintMemShared(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemShared(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemShared)
	}
}

func PrintMemSharedSaved(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemSharedSaved(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemSharedSaved)
	}
}

func PrintMemSwapped(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemSwapped(); err != nil {
		c.WriteNaCol()
	} else {
		c.WritePercentCol(100 * float64(n.MemSwapped) / float64(n.SwapSize))
	}
}

func PrintMemUsed(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemUsed(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint32(n.MemUsed)
	}
}

func PrintMemTargetSize(c *Console, n *Data, o *Data, s *vmguestlib.Session) {
	if _, err := s.GetMemTargetSize(); err != nil {
		c.WriteNaCol()
	} else {
		c.WriteUint64(n.MemTargetSize)
	}
}
