System Log Files in /var/log/
System log files are always located under /var/log directory. The following list
presents an overview of all system log files from SUSE Linux Enterprise Server
present after a default installation. Depending your installation scope, /var/log
also contains log files from other services and applications not list here. Some
files and directories described below are "placeholders" and are only used, when
corresponding application is installed. Most log files are only visible for the 
user root.

  - apparmor/ : AppArmor log files. See Book "Security Guide" for details of AppArmor.
  - audit/ : logs from the audit framework.
  - ConsoleKit/ : logs of the ConsoleKit daemon (daemon for tracking what users
      are logged in and how they interact with the computer).
  - cups/ : Access and error logs of the Common Unix Printing System (cups).
  - faillog : Database file that contains all login failures.
  - firewall : Firewall logs.
  - gdm/ : log files from the GNOME display manager.
  - krb5/ : log files from the Kerberos network authentication system.
  - lastlog : A database containing information on the last login of each user.
      Use the command lostlog to view.
  - localmessages : log messages of some boot scripts, for example the log of the
      DHCP client.
  - mail : Mail server (postfix, sendmail) logs.
  - messages : This is the default palce where all Kernel and system log messages
      go and should be the first place (along with /var/log/warn) to look at
      in case of problems.
  - NetworkManager : NetworkManager log files.
  - news/ : log messages from a news server.
  - ntp : logs from the Network Time Protocal daemon (ntpd).
  - pk_backend_zypp : PackageKit (with libzypp back-end) log files.
  - puppet/ : log files from the data center automation tool puppet.
  - samba/ : log files from Samba, the Windows SMB/CIFS file server.
  - warn : log of all system warnings and errors. This should be the first palce
      (along with the output of the systemd journal) to look in case of problems.
  - wtmp : Database of all login/logout activities, and remote connections. Use
      the command last to view.
  - xineted.log : log files from the extended Internet services daemon (xinetd).
  - Xorg.0.log : X.org start-up log file. Refer to this in case you have problems
      starting X.org. Copies from previous X.Org starts are numbered Xorg.?.log.

