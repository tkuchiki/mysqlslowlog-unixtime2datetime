mysqlslowlog-unixtime2datetime
==============================

Convert a unixtime to datetime(mysql slowlog).

## Example

~~~~
$ cat mysqlslow.log
/rdsdbbin/mysql/bin/mysqld, Version: 5.6.17-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
# Time: 141005 13:05:00
# User@Host: hoge[hoge] @  [192.168.0.10]  Id: 7859392
# Query_time: 0.212722  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
use hoge;
SET timestamp=1412514300;
commit;
# User@Host: rdsadmin[rdsadmin] @ localhost [127.0.0.1]  Id: 916375
# Query_time: 0.393160  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
use mysql;
SET timestamp=1412514300;
flush logs;
# User@Host: hoge[hoge] @  [192.168.0.10]  Id: 7880380
# Query_time: 0.383446  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
use hoge;
SET timestamp=1412514300;
commit;
~~~~

~~~~
$ go run mysqlslowlog-unixtime2datetime.go mysqlslow.log
/rdsdbbin/mysql/bin/mysqld, Version: 5.6.17-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
# Time: 141005 13:05:00
# User@Host: hoge[hoge] @  [192.168.0.10]  Id: 7859392
# Query_time: 0.212722  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
use hoge;
SET timestamp=1412514300; (2014-10-05 22:05:00 +0900 JST)
commit;
# User@Host: rdsadmin[rdsadmin] @ localhost [127.0.0.1]  Id: 916375
# Query_time: 0.393160  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
use mysql;
SET timestamp=1412514300; (2014-10-05 22:05:00 +0900 JST)
flush logs;
# User@Host: hoge[hoge] @  [192.168.0.10]  Id: 7880380
# Query_time: 0.383446  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
use hoge;
SET timestamp=1412514300; (2014-10-05 22:05:00 +0900 JST)
commit;
~~~~
