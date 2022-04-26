# log-template

> use to set golang read cfg.json template.
> log info into logfile.


# log vendor

>  确保 vendor/github.com/coreos/go-log/log/fields.go 文件被修改
> 1 修改 full_time 满足格式要求

```
"full_time":  time.Now().Format("2006-01-02 15:04:05.999"),  // time of log entry

```

> 2 修改 logger.verbose = true  由于属于内部变量无法外部修改

```
logger.verbose = true
```
> 3 logger example 

```
[2022-04-26 18:12:46.65] [DEBUG] [28062] [commands.go:33] >>> [main] msg=debug: yes
[2022-04-26 18:13:30.858] [DEBUG] [28224] [commands.go:33] >>> [main] msg=debug: yes
[2022-04-26 18:13:30.858] [INFO] [28224] [commands.go:33] >>> [main] msg=info: yes
[2022-04-26 18:13:30.858] [WARNING] [28224] [commands.go:33] >>> [main] msg=warning: yes
```
