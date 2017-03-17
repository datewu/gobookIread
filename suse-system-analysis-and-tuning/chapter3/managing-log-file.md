Log files under /var/log grow on a daily basis and quickly become very large.
logrotate is a tool that helps you manage log files and their growth.
It allows automatic rotation, removal, compression, and mailing of log files.
Log files can be handled periodically (daily, weekly, or monthly) or when 
exceeding a particular size.
logrotate is usually run as a daily cron job, and thus usually modifies log files
only once a day. However, exceptions occur when a log file is modified because of
its size, if logrotate is run multiple times a day, or if --force is enabled.
The main configuration file of logrotate is /etc/logrotate.conf. System packages
and programs that produce log files (for example, appche2) put their own 
configuration files in the /etc/logrotate.d/ directory. The content of 
/etc/logrotate.d/ is included via /etc/logrotate.conf.

logrotate is controlled through cron and is called daily by /etc/cron.daily/logrotate.
Use /var/lib/logrotate.status to find out when a particular file has been 
rotated lastly.

[root@nginx ~]# cat /var/lib/logrotate.status                                        
logrotate state -- version 2
"/var/log/vipnginx_worker/run.log" 2017-3-17
"/var/log/nginx/error.log" 2017-3-17
"/var/log/up/run.log" 2017-3-17
"/var/log/yum.log" 2017-1-1
"/var/log/dracut.log" 2014-1-23
"/var/log/wtmp" 2014-1-23
"/var/log/letsencrypt/letsencrypt.log" 2017-3-17
"/var/log/dhclient" 2017-3-12
"/var/log/spooler" 2017-3-12
"/var/log/monitor_admin/run.log" 2017-3-17
"/var/log/btmp" 2017-3-1
"/var/log/maillog" 2017-3-12
"/var/log/NetworkManager" 2017-3-12
"/var/log/audit/audit.log" 2017-3-17
"/var/log/wpa_supplicant.log" 2014-1-23
"/var/log/secure" 2017-3-12
"/var/log/nginx/access.log" 2017-3-17
"/var/log/etcd/etcd.log" 2017-3-17
"/var/log/ppp/connect-errors" 2014-1-23
"/var/log/messages" 2017-3-12
"/var/log/cron" 2017-3-12
[root@nginx ~]#
