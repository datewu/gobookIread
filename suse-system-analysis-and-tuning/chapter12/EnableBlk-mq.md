Enable blk-mq I/O path for SCSI by default
Block multiqueue(blk-mq) is a multi-queue block I/O queueing mechanism.
Blk-mq uses percpu software queues to queue I/O requests. The software
queues are mapped to one or more hardware submission queues. Blk-mq 
significantly reduces lock contention. In particular blk-mq improves 
performance for devices that support a high number of input/output
operations per second (IOPS). Blk-mq is already the default for some 
devices, e.g. NVM Express devices. Currently blk-mq has no I/O scheduling
support (no CFQ, no deadline I/O scheduler). This lack of I/O scheduling
can cause significant performance degradation when spinning disks are
used. Therefore blk-mq is not enabled by default for SCSI devices.
If you have fast SCSI devices (e.g. SSDs) instead of spinning SCSI devices
attached to your system you might consider to switch to blk-mq for SCSI.
This is done using the kernel command line option
scsi_mod.use_blk_mq=1

