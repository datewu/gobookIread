There are physical limitations to hardware that are encountered when 
large numbers of CPU and memory are required. For the purposes of this
chapter, the important limitation is that there is limited communication 
bandwidth between CPUs and the memory. One architecture modification that
was introduced to address this is Non-Uniform Memory Access (NUMA).
In this configuration, there are multiple nodes. Each of the nodes 
contains a subset of all CPUs and memory. The access speed to main memory
is determined by the location of the memory relative to the CPU. The 
performance of a workload depends on the application threads accessing 
data that is local to the CPU the thread is executing on. Automatic NUMA
Balancing is a new feature of SLE 12. Automatic NUMA Balancing migrates 
data on demand to memory nodes that are local to the CPU accessing that
data. Depending on the workload, this can dramatically boost performance
when using NUMA hardware.

1. Implementation
Automatic NUMA balancing happens in three basic steps:
  - A task scanner periodically scans a portion of a task's address space
      and marks the moemory to force a page fault when the data is next
      accessed.
  - The next access to the data will result in a NUMA Hinting Fault. Base
      on this fault, the data can be migrated to a memory node associated
      with the task accessing the memory.
  - To keep a task, the CPU it is using and the memory it is accessing 
      together, the scheduler groups tasks that share data.
The unmapping of data and page fault handling incurs overhead. However
commonly the overhead will be offset by threads accessing data associated
with the CPU.

2. Configuration
Static configuration has been the recommended way of tuning workloads on
NUMA hardward for some time. To do this, memory policies can be set with
numactl, taskset or cpusets.
NUMA-aware applications can use special APIs. In case where the static 
policies have already been created, automatic NUMA balancing should be
disabled as the data access should already be local.
numactl --hardware will show the memory configuration of the machine and
whether it supports NUMA or not.
Automatic NUMA balancing canbe enabled or disabled for the current 
session by writing 1 or 0 to /proc/sys/kernel/numa_balancing which will 
enable or disable the feature respectively. To permanently enable or 
disable it, use the kernel command line option numa_balancing=[enabled |
disabled].
If Automatic NUMA Balancing is enabled, the task scanner behavior can be
configured. The task scanner balances the overhead of Automatic NUMA 
Balancing with the amount of time it takes to identify the best placement
data.
  - numa_balancing_scan_delay_ms : The amount of CPU time a thread must
      consume before its data is scanned. This prevents creating overhead
      because of shor-lived process.
  - numa_balancing_scan_period_min_ms and numa_balancing_scan_period_max_ms
   : Controls how frequently a task's data is scanned. Depending on the
       locality of the faults the scan rate will increase or decrease.
       These settings control the min and max scan rates.
  - numa_balancing_scan_size_mb : Controls how much address space is 
      scanned when the task scanner is active.

3. Monitoring
The most important task is to assign metrics to your workload and measure
the performance with Automatic NUMA Balancing enabled and disabled to 
measure the impact. Profiling tools can be used to monitor local and 
remote memory accesses if the CPU supports such monitoring.
Automatic NUMA Balancing activity can be monitored via the following 
parameters in /proc/vmstat:
  - numa_pte_updates : The amount of base pages that were marked for
      NUMA hinting faults.
  - numa_huge_pte_updates : The amount of transparent huge pages that 
      were marked for NUMA hinting faults. In combination with 
      numa_pte_updates the total address space was marked can be
      calculated.
  - numa_hint_faults : Records how many NUMA hinting faults were trapped.
  - numa_hint_faults_local : Shows how many of the hinting faults were to
      local nodes. In combination with numa_hint_faults, the percentage
      of local versus remote faults can be calculated. A high percentage
      of local hinting faults indicates that the workload is closer to
      being converged.
  - numa_pages_migrated : Records how many pages were migrated because 
      they were misplaced. As migration is a copying operation, it 
      contributes the largest part of the overhead created by NUMA
      balancing.

4. Impact
Automatic NUMA Balancing takes away some of the pain when tuning 
workloads for high performance on NUMA machines. Where possible, it is 
still recommended to statically tune the workload to partition it within
each node. However, in all other cases, automatic NUMA balancing should
boost performance.

