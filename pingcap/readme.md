# 计算
## 方案一
![version1(../assert/ans1-preread.png)

### 设计缘由

- index为啥使用hash而不是key？

考虑key没有给出具体长度，同时不需要范围查询，使用key可能浪费空间

- value 为啥不放到index中？

考虑value没有给出具体长度， 如果value比较大，会有迁移成本， 同时导致index文件
利用效率低下

- hash 如何放置冲撞？

1: 如果hash在一个segment中，使用拉链法进行解决

2: 不在一个segment中，通过bloom filter 也能找到 ，不会存在丢失问题

- index和bloom filter为啥使用两个文件？

可以使用一个文件进行存储 解析


### 参数计算
由于没有说明key与value size大小范围，按照最极端情况进行计算，
一个kv大小： keysize（4b）+key（4b）+valuesize（4b）+value（4b）= 16b
1TB 最大的存储： 2^(40-4)=2^36 个kv
每个segment 大小： 64MB 
每个entry 大小： hash（8b）+ offset（8b）= 16b
一个segment存储：2^(26-4)=2^22个entry
总共需要： 2^(36-22)=2^12 个segment 文件进行存储

这边需要使用bloom过滤器，因此开始计算m，k，假定误报率fpp小于%1
m：bit大小 m=-(n*lnp)/(ln2)^2=-(n*lnp)/(ln2)^2=40202648=2^25bit=4Mb
k: hash次数 k=(m/n)*ln2=7

将2GB作为内存索引
一个支持 2^(11-3)=2^9 索引
公共需要 2^(12-9)=2^3 

### 工作量安排
1. 整体Server设计
2. 1T文件 使用mmap进行顺序读，解析，随机读 
3. 内存使用跳表设计，支持序列化 
4. LFU，LRU Cache设计实现 
5. Index 结构设计,存储，二分左边界查找
6. hash bloom 设计实现
7. 对外客户端服务编写

### 已实现的功能
1. 核心模块均有单元测试
2. server_test 完成一个100条kv存储 以及读取
3. cmd 支持mock 数据(test.log 是器生成的)

### 未完成
1. 对外客户端服务编写
2. 没有进行大量数据测试 ，目前使用100条数据测试
3. 配置参数不够友好 
4. 每个bloom 过滤器都需要算一次hash可以节省
5. index 可以进一步做压缩

### 架构存在的问题
1. 使用Hash做未Region方案，后续可以使用一致性hash,减少数据偏移
2. 虽然做了分区，但是Bloom Filter遍历仍然需要大量IO以及CPU
### 优点
1. 1T数据中的value没有迁移，大幅度减少硬盘空间（查找需要IO）
2. 顺序写性能很好， 落盘之后index文件是有序的，方便二分查找（O(logN)）
3. 使用Bloom Filter 快速查询key对应的索引 ，提高无效查询（额外的空间）


## 方案二
![version1(../assert/ans1-preread.png)
### 主要解决的问题
是不是可以不需要Bloom Filter 也能找到key对应的index 文件

### 一致性HASH或者虚拟一致性hash
难点：在与处理过程中需要进行spilt操作

### 架构存在的问题
1. 预处理过程中，每次spilt过程中需要rehash，同时里面的index不是有序的，
为了方便查找还需要进行一次排序（相对一而言，预处理时间比较长）
### 优点
1. 查找性能高，没有bloom filter消耗(IO CPU)
