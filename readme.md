# Go 短网址服务

## Hash 算法

采用了 [Murmurhash3 算法](https://github.com/spaolacci/murmur3)

使用的 hashcode 是 32 位的，2^32 大概是 42.9 亿，如果用 0-9a-zA-Z 这 62 个字符去接收，`62^5 < 2^32 < 62^7`，62^6 能表达 568
亿。所以使用 32 位的 hashcode 最后生成的短域名是 6 位。

## 持久化

## 缓存

一般来说，短网址服务都是读多写少的服务，因此，缓存对于这类应用的性能提高是十分显著的。对于一般的小业务（不用多实例部署）来说，可能简单的本地缓存就够了。但是对于分布式的架构来说，多个应用实例的本地缓存之间容易存在缓存数据不一致的问题。所以一般会使用中心化的缓存，比如 Redis。

## 依赖框架

暂无，后续准备引入下列依赖：

- Web 框架：gin
- ORM 框架：gorm

主要是有了框架之后开发业务代码会方便很多。

## 后续计划引入的特性

- [ ] 短链有效期设置
- [ ] 短链长度的设置

## 待完善任务

- [ ] hash 冲突的处理
- [ ] 与前端通信数据结构的制定
