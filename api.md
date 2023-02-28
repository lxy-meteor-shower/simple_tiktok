
### Register
```shell
curl -X POST http://localhost:8080/douyin/user/register/?"username=lxy&password=123456"


```

### Login
```shell
curl -X POST http://localhost:8080/douyin/user/login/?"username=lxy&password=123456"
```

### UserInfo
```shell
curl http://localhost:8080/douyin/user/?"user_id=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc1Njg0MjIsImlkIjoxLCJvcmlnX2lhdCI6MTY3NzU2NDgyMn0.96SHxeWzQ22FYcBs8O-LNP9i65P_X37wrva3VjUzE9s"
```

### Feed
```shell
curl http://localhost:8080/douyin/feed/?"latest_time=1677560960&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc1Njg0MjIsImlkIjoxLCJvcmlnX2lhdCI6MTY3NzU2NDgyMn0.96SHxeWzQ22FYcBs8O-LNP9i65P_X37wrva3VjUzE9s"
```

### Publish
```shell
curl -X POST http://localhost:8080/douyin/publish/action/?"title=11111&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc1Njg0MjIsImlkIjoxLCJvcmlnX2lhdCI6MTY3NzU2NDgyMn0.96SHxeWzQ22FYcBs8O-LNP9i65P_X37wrva3VjUzE9s"
```
