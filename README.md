# naive-admin-go

### 基于 [vue-naive-admin 2.0版](https://gitee.com/isme-admin/vue-naive-admin) 前端框架的 golan版本做服务端

### 使用 gin gorm mysql jwt session

[api documnet](https://apifox.com/apidoc/shared-ff4a4d32-c0d1-4caf-b0ee-6abc130f734a/api-134496720)

[api demo code](https://gitee.com/-/ide/project/isme-admin/isme-nest-serve/edit/main/-/src/modules/role/dto.ts)

[web demo](https://admin.isme.top/login?redirect=/)

## [api 接口](./api.md)


## 本地运行

#### 运行数据库服务
docker 方式
```shell
docker compose -f docker-compose-env.yaml 
```
或者自行修改 .env 文件 Mysql 连接参数，并导入 init.sql 数据库及表结构

#####  运行前端
```shell
cd vue-naive-front && npm install && npm run dev
```
##### 运行后端
```shell
go run main.go
```


最近在写[rust 版本的服务端](https://github.com/ituserxxx/rust_axum_web_api_demo)