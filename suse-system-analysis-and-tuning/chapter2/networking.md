1. Basic Network Diagnostics: ip
ip is a powerful tool to set up and control network interfaces. You can
also use it to quickly view basic statistics about network interfaces of
the system. For example, whether the interface is up or how many errors,
dropped packets, or packet collisions there are.
If you run ip with no additional parameter, it displays a help output. To
list all network interfaces, enter ip addr show (or abbreviated as ip a).
ip addr show up lists only running network interfaces.
ip -s link show device lists statistics for the specified interface only.

ip can also be used to show interfaces (link), routing tables (route),
and much more--refer to man 8 ip for details.

[root@nginx ~]# ip link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 52:54:51:c4:bd:5b brd ff:ff:ff:ff:ff:ff
21: tun0: <POINTOPOINT,MULTICAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UNKNOWN qlen 100
    link/[65534]
[root@nginx ~]# ip route
10.8.0.101 dev tun0  proto kernel  scope link  src 10.8.0.102
10.8.0.1 via 10.8.0.101 dev tun0
172.16.110.0/24 dev eth0  proto kernel  scope link  src 172.16.110.9  metric1
default via 172.16.110.1 dev eth0  proto static

2. Show the Network Usage of Processes: nethogs
In some cases, for example if the network traffic suddenly becomes very high, it is 
desirable to quickly find out which application(s) is/are causing the traffic.
nethogs, a tool with a design similar to top, shows incoming and outgoing traffic
for all relevant processes.
Like top, nethogs features interactive commands.
  - m: cycle between display modes (kb/s, kb, b, mb)
  - r: sort by RECEIVED
  - s: sort by SENT
  - q: quit

3. Ethernet Cards in Detail: ethtool
ethtool can display and change detailed aspects of your Ethernet network device. By
default it prints the current setting of the specified device.
[root@nginx ~]# ethtool lo
Settings for lo:
        Link detected: yes
[root@nginx ~]# ethtool eth0
Settings for eth0:
        Link detected: yes
[root@nginx ~]# ethtool tun0
Settings for tun0:
        Supported ports: [ ]
        Supported link modes:   Not reported
        Supported pause frame use: No
        Supports auto-negotiation: No
        Advertised link modes:  Not reported
        Advertised pause frame use: No
        Advertised auto-negotiation: No
        Speed: 10Mb/s
        Duplex: Full
        Port: Twisted Pair
        PHYAD: 0
        Transceiver: internal
        Auto-negotiation: off
        MDI-X: Unknown
        Current message level: 0xffffffa1 (-95)
                               drv ifup tx_err tx_queued intr tx_done rx_status pktdata hw wol 0xffff8000
        Link detected: yes

The following table shows ethtool options that you can use to query the device for
specific information:
  - -a: pause parameter information
  - -c: interrupt coalescing information
  - -g: Rx/Tx (receive/transmit) ring parameter information
  - -i: associated driver information
  - -k: offload information
  - -S: NIC and driver-specific statistics

4. Show the Network Status: ss
ss is a tool to dump socket statistics and replaces the netstat command.
To show a list of all connections use ss without parameters.
To show all network prots currently open, use ss -l.
When displaying network connections, you can specify the socket type to display:
TCP (-t) or UDP (-u) for example. The -p option shows the PID and name of the 
program to which each socket belongs.
The following example lists all TCP connections and the programs using these
connections. The -a option make sure all established connections (listening and
non-listening) are shown. The -p option shows the PID and name of the program to
which each socket belongs.
[root@nginx ~]# ss -tap | more
State      Recv-Q Send-Q      Local Address:Port          Peer Address:Port
LISTEN     0      16384                  :::58080                   :::*        users:(("monitor_api",29456,6))
LISTEN     0      16384                  :::53000                   :::*        users:(("monitor_admin",2814,3))
LISTEN     0      16384                  :::2379                    :::*        users:(("etcd",23064,5))
LISTEN     0      16384        172.16.110.9:2380                     *:*        users:(("etcd",23064,3))
LISTEN     0      16384                  :::jetdirect                 :::*        users:(("node_exporter",16651,6))
LISTEN     0      16384                  :::distinct                 :::*        users:(("linuxB",917,3))
LISTEN     0      511                     *:http                     *:*        users:(("nginx",31950,6),("nginx",31952,6))
LISTEN     0      128                    :::ssh                     :::*        users:(("sshd",22081,4))
LISTEN     0      128                     *:ssh                      *:*        users:(("sshd",22081,3))
LISTEN     0      511                     *:https                    *:*        users:(("nginx",31950,7),("nginx",31952,7))
ESTAB      0      0            172.16.110.9:54458          172.16.12.5:58080    users:(("monitor_admin",2814,22))
ESTAB      0      0            172.16.110.9:21791          172.16.11.3:58080    users:(("monitor_admin",2814,15))
TIME-WAIT  0      0            172.16.110.9:49529          172.16.10.3:58080
ESTAB      0      0            172.16.110.9:32650         172.16.110.4:2380     users:(("etcd",23064,24))
TIME-WAIT  0      0            172.16.110.9:49524          172.16.10.3:58080
TIME-WAIT  0      0            172.16.110.9:49547          172.16.10.3:58080
