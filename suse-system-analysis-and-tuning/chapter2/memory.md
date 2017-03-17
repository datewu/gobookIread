1. Memory Usage: free
The utility free examines RAM and swap usage. Details of both free and used
memory and swap areas are shown.
The options -b, -k, -m, -g show the output in bytes, KB, MB or GB, respectively.
The parameter -s delay ensures that the display is refreshed every delay seconds
For example, free -s 1.5 produces an update every 1.5 seconds.

2. Detailed Memory Usage: /proc/meminfo
Use /proc/meminfo to get more detailed information on memory usage than with
free. Actually free use some of the data from this file. See an example output
from a 64-bit system below. Note that it slightly differs on 32-bit systems 
because of different memory management:

[root@nginx ~]# cat /proc/meminfo
MemTotal:        1020344 kB
MemFree:          399216 kB
Buffers:          150192 kB
Cached:           312060 kB
SwapCached:            0 kB
Active:           351360 kB
Inactive:         190808 kB
Active(anon):      16832 kB
Inactive(anon):    63288 kB
Active(file):     334528 kB
Inactive(file):   127520 kB
Unevictable:           0 kB
Mlocked:               0 kB
SwapTotal:       1048572 kB
SwapFree:        1048572 kB
Dirty:                44 kB
Writeback:             0 kB
AnonPages:         79948 kB
Mapped:            29032 kB
Shmem:               204 kB
Slab:              57524 kB
SReclaimable:      35488 kB
SUnreclaim:        22036 kB
KernelStack:        1160 kB
PageTables:         3144 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     1558744 kB
Committed_AS:     227996 kB
VmallocTotal:   34359738367 kB
VmallocUsed:       11784 kB
VmallocChunk:   34359703504 kB
HardwareCorrupted:     0 kB
AnonHugePages:      4096 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:        8184 kB
DirectMap2M:     1040384 kB

  - MemTotal: total amount of RAM
  - MemFree: amount of unused RAM
  - MemAvailable: estimate of how much memory is available for starting 
      new applications without swapping.
  - Buffers: File buffer cache in RAM containing file system metadata.
  - Cached: Page cache in RAM. This excludes buffer cache and swap cache,
      but includes Shmem memory.
  - SwapCached: Page cache for swapped-out memory.
  - Active,Active(anon),Active(file):
      Recently used memory that will not be reclaimed unless necessary or
      explicit request. Active is the sum of Active(anon) and Active(file):
        - Active(anon) tracks swap-backed memory. This includes private and
	    and shared anonymous mappings and private file pages after
	    copy-on-write.
	- Active(file) tracks other file system backed memory.
  - Inactive,Inactive(anon),Inactive(file):
      Less recently used memory that will usually be reclaimed first. Inactive
      is the sum of Inactive(anon) and Inactive(file):
        - Inactive(anon) tracks swap-backed memory. This includes private and
	    and shared anonymous mappings and private file pages after
	    copy-on-write.
	- Inactive(file) tracks other file system backed memory.
  - Unevictable: Amount of memory that cannot be reclaimed (for example, because
      it is Mlocked or used as a RAM disk).
  - Mlocked: Amount of memory that is backed by the mlock system call. mlock 
      allows processes to define which part of physical RAM their virtual 
      memory should be mapped to. However, mlock does not guarantee this
      placement.
  - SwapTotal: Amount of swap space.
  - SwapFree: Amount of unused swap space.
  - Dirty: Amount of memory waiting to be written to disk, because it contains
      changes compared to the backing storage. Dirty data can be either
      explicitly synchronized by the application or by the kernel after a short
      delay. A large amount of dirty data may take considerable time to write to
      disk resulting in stalls. The total amount of dirty data that can exist
      at any given time can be controlled with the sysctl parameters 
      vm.dirty_ratio or vm.dirty_bytes
  - Writeback: Amount of memory that is currently being written to disk.
  - Mapped: Memory claimed with the mmap system call.
  - Shmem: Memory shared between groups of processes, such as IPC data, tmpfs
      data, and shared anonymous memory.
  - Slab: Memory allocation for internal data structures of the kernel.
  - SReclaimable: Slab section that can be reclaimed, such as 
      csches (inode, dentry, etc.)
  - SUnreclaim: Slab section that cannot be reclaimed.
  - KernelStack: Amount of kernel space memory used by applications
      (through system calls).
  - PageTables: Amount of memory dedicated to page tables of all processes.
  - NFS_Unstable: NFS pages that have already been sent to the server, but
      are not yet committed there.
  - Bounce: Memory used for bounce buffers of block devices.
  - WritebackTmp: Memory used by FUSE for temporary writeback buffers.
  - CommitLimit: Amount of memory available to the system based on the 
      overcommit ratio setting. This is only enforced if strict overcommit 
      accounting is enable.
  - Committed_AS: An approximatin of the total amount of memory (RAM and
      swap) that the current workload would need in the worst case.
  - VmallocTotal: Amount of allocated kernel virtual address space.
  - VmallocUsed: Amount of used virtual address sapce.
  - VmallocChunk: The largest contiguous block of available kernel
      virtual address space.
  - HardwareCorrupted: Amount of failed memory (can only be detected when 
      using ECC RAM).
  - AnonHugePages: Anonymous hugepages that are mapped into user space page
      tables. These are allocated transparently for processes without being
      specifically requested, therefore they are also known as 
      transparent hugepages (THP).
  - HugePages_Total: Number of preallocated hugepages for use by SHM_HUGETLB
      and MAP_HUGETLB or througn the hugetlbfs file system, as defined in
      /proc/sys/vm/nr_hugepages.
  - HugePages_Free: Number of hugepages available.
  - HugePages_Rsvd: Number of hugepages that are committed.
  - HugePages_Surp: Number of hugepages available beyond HugePages_Total
      ("surplus"), as defined in /prc/sys/vm/nr_overcommit_hugepages.
  - Hugepagesize: Size of a hugepage-- on AMD64/Intel 64 the default 
      is 2048 KB.
  - DirectMap4k etc.
    Amount of kernel memory that is mapped to pages with a given size
    (in the example: 4kb).

3. Process Memory Usage: smaps
Exactly determining how much memory a certain process is consuming is not
possible with standard tools like top or ps. Use the smaps subsystem, 
introduced in kernel 2.6.14, if you need exact data. It can be found at 
/proc/pid/smaps and shows you the number of clean and dirty memory pages
the process with the ID PID is using at that time. It differentiates 
between shared and private memory, so you can see how much memory the 
process is using without including memoru shared with other processes.

Note: smaps is expensive to read. Therefore it is not recommended to monitor
it regularly, but only when closely monitoring a certain process.
