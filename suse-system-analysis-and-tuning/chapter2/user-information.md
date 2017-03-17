1. User Accessing Files: fuser
It can be useful to determine what processes or users are currently 
accessing certain files. Suppose, for example, you want to unmount a file
system mount at /mnt. umount returns "device is busy." The command fuser
can then be used to determine what processes are accessing the device.
fuser -v /mnt/**

When use with -k option, fuser will terminate processing the file as well**

2. Who Is Doing What: w
With the command w, find out who is logged into the system and what each
user is donig.
If any users of other systems have logged in remotely, the parameter -f
shows the computers from which they established the connection.
