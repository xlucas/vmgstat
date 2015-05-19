# vmgstat

[![GitHub version](https://img.shields.io/github/release/xlucas/vmgstat.svg)](https://github.com/xlucas/vmgstat/releases/tag/v1.2.0)
[![Github licence](https://img.shields.io/github/license/xlucas/vmgstat.svg)](LICENSE)

VMware® vSphere Guest Statistics tool.

![ScreenShort](screenshot.png)

## Disclaimer
This is not an official VMware® product.

## Requirements
The requirements are the same than [go-vmguestlib](https://github.com/xlucas/go-vmguestlib) on which `vmgstat` is built.

## Usage
This tool uses golang flags library, thus options should be passed as `--<name>=<value>`. The `deflay` flags use golang's duration flag which means you can use `<T>ns`, `<T>us` `<T>ms`, `<T>s`, `<T>m`, `<T>h` as values where T is a positive number.

```
--color  : Use colors           Default value : true
--count  : Refresh count        Default value : 0 (unlimited)
--cpu    : Show cpu stats       Default value : true
--delay  : Refresh delay        Default value : 1s
--guest  : Stats from guest     Default value : true
--header : Header print step    Default value : 0 (only at start)
--host   : Stats from host      Default value : false
--memory : Show memory stats    Default value : false
```

## Metrics :

### Guest

* CPU
  * `CLimG` : maximum processing power in GHz available to the virtual machine.
  * `CResG` : minimum processing power in GHz available to the virtual machine.
  * `CShaG` : number of CPU shares allocated to the virtual machine.
  * `CStlG` : percentage of time the VM was runnable but not scheduled to run during the previous refresh delay divided by the vCPU count.
  * `CUseG` : percentage of time the VM was using host's CPUs during the previous refresh delay divided by the vCPU count.

* Memory
  * `MActG` : amount of memory in GB the virtual machine is actively using.
  * `MBalG` : amount of memory in GB that has been reclaimed from this virtual machine via the VMware® Memory Balloon mechanism.
  * `MLimG` : maximum amount of memory in GB that is available to the virtual machine.
  * `MMapG` : mapped memory size in GB of this virtual machine.
  * `MOvhG` : amount of overhead memory in GB associated with this virtual machine consumed on the host system.
  * `MResG` : minimum amount of memory in GB that is available to the virtual machine.
  * `MshaG` : number of memory shares allocated to the virtual machine.
  * `MShdG` : amount of physical memory in GB associated with this virtual machine that is copy-on-write (COW) shared on the host.
  * `MShsG` : estimated amount of physical memory in GB on the host saved from copy-on-write (COW) shared guest physical memory.
  * `MSwaG` : amount of memory in GB associated with this virtual machine that has been swapped by the host system.
  * `MTarG` : the memory target Size in GB.
  * `MUseG` : the estimated amount of physical host memory in GB currently consumed for this virtual machine's physical memory.

### Host
* CPU
  * `CUseH` : percentage of time the host was using its CPUs during the previous refresh delay divided by the CPU count on the host.
  * `CNumH` : the number of physical CPU cores on the host machine.
  * `CSpeH` : the host processor speed in GHz.

* Memory
  * `MOvhH` : total host kernel memory overhead, in GB.
  * `MMapH` : total mapped memory on the host, in GB.
  * `MPhyH` : total memory available to host OS kernel, in GB.
  * `MFreH` : total physical memory free on the host in GB.
  * `MShaH` : total COW (Copy-On-Write) memory on the host, in GB.
  * `MSwaH` : total memory swapped out on the host, in GB.
  * `MUnmH` : total unmapped memory on the host, in GB.
  * `MUseH` : total consumed memory in on the host, in GB.

## Contribute

Feel free to bring any contribution if you think that would be useful.
