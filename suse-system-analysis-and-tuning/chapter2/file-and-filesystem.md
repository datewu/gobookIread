1. Determine the File Type: file
The command file determines the type of a file or a list of files by checking 
/usr/share/misc/magic.
The parameter -f listfile specifies a file containing a list of file names to 
examine.
The -z allows file to look inside compressed files.
The parameter -i outputs a mime type string rather than the traditional description

2. File Systems and Their Usage: mount, df and du
The command mount shows which file system (device and type) is mounted at which 
mount point.
Obtain information about total usage of the file systems with the command df.
The parameter -h (or --human-readable) transforms the output into a form
understandable for common users.
Display the total size of all the files in a given directory and its subdirectories
with the command du. The parameter 0s suppresses the output of detailed information
and give only a total for each argument. -h again transforms the output into 
a human-readable form.

3. Additional Information about ELF (Executable and Linkable Format) Binaries
Read the content of binaries with readelf utility. This even works with ELF files
that were built for other hardware architectures:
[root@nginx ~]# readelf  --file-header /opt/monitor_admin/monitor_admin
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x462b60
  Start of program headers:          64 (bytes into file)
  Start of section headers:          456 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         7
  Size of section headers:           64 (bytes)
  Number of section headers:         23
  Section header string table index: 7
[root@nginx ~]# readelf  --file-header /bin/ls
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x4027e0
  Start of program headers:          64 (bytes into file)
  Start of section headers:          107352 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         8
  Size of section headers:           64 (bytes)
  Number of section headers:         29
  Section header string table index: 28
[root@nginx ~]#


4. File Properties: stat
The command stat displays file properties.
The parameter --file-system produces details of the properties of the 
file system in which the specified file is located.
