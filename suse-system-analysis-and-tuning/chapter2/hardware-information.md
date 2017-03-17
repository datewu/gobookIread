1. PCI Resources: lspci
The command lspci list the PCI resources. Using -v results in a more detailed list.
Information about device name resolution is obtained from the file /usr/share/pci.ids.
PCI IDs not listed in this file are marked "Unknown device."
The parameter -vv produces all the information that could be queried by the program
To view the pure numeric values, use the parameter -n.

2. USB Devices: lsusb
The command lsusb list all USB devices. With the option -v, print a more detailed
list. The detailed information is read from the directory /proc/bus/usb/. 

3. MCELog: Machine Check Exceptions (MCE)
The mcelog package logs and parse/translates Machine Check Exceptions (MCE) on 
hardware errors (also including memory errors). Formerly this has been done by a 
cron job executed hourly. Now hardware errors are immediately processed by an
mcelog daemon.
However, the mcelog service is not enabled by default, resulting in memory and CPU
errors also not being logged by default. In addition, mcelog has a new feature to
also handle predictive bad page offlining and automatic core offlining when
cache errors happen.

4. x86_64: dmidecode: DMI Table Decoder
dmidecode shows the machine's DMI table containing information such as serial 
numbers and BIOS revisions of the hardware.

[root@nginx ~]# dmidecode
# dmidecode 2.12
SMBIOS 2.4 present.
10 structures occupying 304 bytes.
Table at 0x000F09C0.

Handle 0x0000, DMI type 0, 24 bytes
BIOS Information
        Vendor: Bochs
        Version: Bochs
        Release Date: 01/01/2011
        Address: 0xE8000
        Runtime Size: 96 kB
        ROM Size: 64 kB
        Characteristics:
                BIOS characteristics not supported
                Targeted content distribution is supported
        BIOS Revision: 1.0

Handle 0x0100, DMI type 1, 27 bytes
System Information
        Manufacturer: QEMU
        Product Name: Standard PC (i440FX + PIIX, 1996)
        Version: pc-i440fx-2.0
        Serial Number: Not Specified
        UUID: F5281492-8182-3E5F-8045-89E1890385B9
        Wake-up Type: Power Switch
        SKU Number: Not Specified
        Family: Not Specified

Handle 0x0300, DMI type 3, 20 bytes
Chassis Information
        Manufacturer: Bochs
        Type: Other
        Lock: Not Present
        Version: Not Specified
        Serial Number: Not Specified

