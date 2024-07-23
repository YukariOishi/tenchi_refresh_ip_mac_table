# tenchi_refresh_ip_mac_table
create ip mac table tool on laravel project for any OS.

- Install GO

https://go.dev/doc/install


- Export main.exe command

```
cd tenchi_refresh_ip_mac_table
SET GOOS=windows; SET GOARCH=amd64; go build maiRCH=amd64; go build main.go
```

- How to use main.exe

```
main.exe -d ./storage/framework/cache/data

e.g.)[wsl]
C:\Users\[user]\Desktop\main.exe -d \\wsl$\Ubuntu\home\[unix_user]\[laravel_project_name]\storage\framework\cache\data
```

- Changable

main.go

Cache retention period (seconds)

```
l36: fc := fmt.Sprintf("9999999999%s", s) // fc: file contents
```
