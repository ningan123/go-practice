

```bash
[root@ubuntu-22 chapter58-containerd-client]# go run main.go
2024/07/30 23:16:16 Successfully pulled docker.io/library/redis:alpine image
2024/07/30 23:16:16 Successfully created container with ID redis-server and snapshot with ID busybox
2024/07/30 23:16:16 Successfully created task redis-server
2024/07/30 23:16:16 Successfully started task redis-server
2024/07/30 23:16:16 

##====== redis server启动后输出的日志 ======##
1:C 30 Jul 2024 15:16:16.378 # WARNING Memory overcommit must be enabled! Without it, a background save or replication may fail under low memory condition. Being disabled, it can also cause failures without low memory condition, see https://github.com/jemalloc/jemalloc/issues/1328. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
1:C 30 Jul 2024 15:16:16.379 * oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 30 Jul 2024 15:16:16.379 * Redis version=7.4.0, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 30 Jul 2024 15:16:16.379 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
1:M 30 Jul 2024 15:16:16.380 # You requested maxclients of 10000 requiring at least 10032 max file descriptors.
1:M 30 Jul 2024 15:16:16.380 # Server can't set maximum open files to 10032 because of OS error: Operation not permitted.
1:M 30 Jul 2024 15:16:16.380 # Current maximum open files is 1024. maxclients has been reduced to 992 to compensate for low ulimit. If you need higher maxclients increase 'ulimit -n'.
1:M 30 Jul 2024 15:16:16.380 * monotonic clock: POSIX clock_gettime
1:M 30 Jul 2024 15:16:16.382 * Running mode=standalone, port=6379.
1:M 30 Jul 2024 15:16:16.383 * Server initialized
1:M 30 Jul 2024 15:16:16.383 * Ready to accept connections tcp
2024/07/30 23:16:26 

##====== redis进程处理sigterm信号 ======##
1:signal-handler (1722352586) Received SIGTERM scheduling shutdown...
1:M 30 Jul 2024 15:16:26.411 * User requested shutdown...
1:M 30 Jul 2024 15:16:26.411 * Saving the final RDB snapshot before exiting.
1:M 30 Jul 2024 15:16:26.416 * DB saved on disk
1:M 30 Jul 2024 15:16:26.416 # Redis is now ready to exit, bye bye...
2024/07/30 23:16:26 

##====== 获取到redis进程的退出状态码 ======##
busybox exited with status: 0, exitedAt: 2024-07-30 15:16:26.427559188 +0000 UTC
```