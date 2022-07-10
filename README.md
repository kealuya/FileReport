# FileReport

File Report System 文档发布系统

[TOC]



## 进度

### 2022-07-03

做了两个查询sql，自定义sql查询，一个使用orm默认的，感觉orm默认的update条件判断可能不好使，后续都使用自定义sql
进行了两个数据结构的定义
还有最重要的，凭啥我大半夜写的东西没push上去(先提个说明文档避免忘了这回事)

### 2022-07-05

分层稍微优化一下，对上传文档的逻辑进行了一个构思

### 2022-07-10

上传文件流程，先查product_info表是否存在项目，存在更新时间和人，不存在插条新的

查file_info表是否存在该文件，存在根据是否版本大更赋值大小版本号然后插入，不存在大小版本都是1插入

## 表

### file_info

| 字段名        | 中文名       | 类型    | 长度 | 空否 | 主键 | 默认值 | 主键 |
| ------------- | ------------ | ------- | ---- | ---- | ---- | ------ | ---- |
| id            | id           | integer |      |      |      |        |      |
| file_name     | 文档名称     | varchar |      |      |      |        |      |
| major_version | 大版本号     | tinyint |      |      |      |        |      |
| minor_version | 小版本号     | tinyint |      |      |      |        |      |
| product_name  | 所属项目     | varchar |      |      |      |        |      |
| create_time   | 首次上传时间 | varchar |      |      |      |        |      |
| creater       | 首次上传人   | varchar |      |      |      |        |      |
| modify_time   | 修改时间     | varchar |      |      |      |        |      |
| modifier      | 修改人       | varchar |      |      |      |        |      |

### product_info

| 字段名           | 中文名       | 类型    | 长度 | 空否 | 主键 | 默认值 | 主键 |
| ---------------- | ------------ | ------- | ---- | ---- | ---- | ------ | ---- |
| id               | id           | integer |      |      |      |        |      |
| product_name     | 所属项目     | varchar |      |      |      |        |      |
| last_update_time | 最后上传时间 | varchar |      |      |      |        |      |
| last_creater     | 最后上传人   | varchar |      |      |      |        |      |



## 视图

### product_status

| 字段名        | 中文名   | 类型    | 备注 |
| ------------- | -------- | ------- | ---- |
| product_name  | 项目名称 | TEXT    |      |
| file_number   | 文档总数 | INTEGER |      |
| person_number | 维护人数 | INTEGER |      |
| version_count | 维护次数 | INTEGER |      |

## 测试用例

### 上传文件

```
第一条
{
    "flieName":"aa",
    "majorUpdate":"no",
    "productName":"dsad",
    "creater":"zxx"
}
第二人上传同名文档
{
    "flieName":"aa",
    "majorUpdate":"no",
    "productName":"dsad",
    "creater":"xxx"
}
大版本升级
{
    "flieName":"aa",
    "majorUpdate":"yes",
    "productName":"dsad",
    "creater":"zxx"
}
另一个项目
{
    "flieName":"2a",
    "majorUpdate":"yes",
    "productName":"tt",
    "creater":"zxx"
}

```

