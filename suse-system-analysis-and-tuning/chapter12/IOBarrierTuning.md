Most file systems(such as XFS, Ext3, Ext4, or reserfs) send write barriers
to disk after fsync or during transaction commits. Write barriers enforce
proper ordering of writes, making volatile disk write caches safe to use
(at some performance penalty). If your disks are battery-backed in one 
way or another, disabling barriers can safely improve performance.
Sending write barriers can be disabled using the nobarrier mount option.
WARNING: Disabling Barriers Can Lead to Data Loss
Disabling barriers when disks cannot guarantee caches are properly written
in case of power failure can lead to server file system corruption and
data loss.
