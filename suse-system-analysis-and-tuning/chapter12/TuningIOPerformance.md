I/O scheduling controls how input/output operations will be submitted to
storage. SUSE Linux Enterprise Server offers various I/O algorithms, 
called elevators, suiting different workloads. Elevators can help to 
reduce seek operations and can prioritize I/O requests.
Choosing the best suited I/O elevator not only depends on the workload,
but on the hardware, too. Single ATA disk systems, SSDs, RAID arrays, or
network storage systems, for example, each require different tuning
strategies.

1. Switching I/O scheduling
SUSE Linux Enterprise Server picks a default I/O scheduler at boot-time,
which can be changed on the fly per block device. This makes it possible 
to set different algorithms, for example, for the device hosting the
system partition and the device hosting a database.
The default I/O scheduler is chosen for each device based on whether the
device reports to be rotational disk or not. For non-rotational disks
DEADLINE I/O scheduler is picked. Other devices default to CFQ(Completely
Fair Queuing). To change this default, use the following boot parameter:
`elevator=SCHEDULER`
Replace SCHEDULER with one of the values cfq, noop, or deadline.
To change the elevator for a specific device in the running system, run
the following command:
`echo SCHEDULER > /sys/block/DEVICE/queue/scheduler`
Here, SCHEDULER is one of cfq, noop, or deadline. DEVICE is the block
device(sda for example). Note that this change will not persist during 
reboot. For permanent I/O scheduler change for a particular device either
place the command switching the I/O scheduler into init scripts or add 
appropriate udev rule into /lib/udev/rules.d/.

NOTE: On IBM z Systems, the default I/O scheduler for a storage device is
set by the device driver.
