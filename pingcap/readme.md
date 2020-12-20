# 计算

由于没有说明key与value size大小范围，按照最极端情况进行计算，
一个kv大小： keysize（4b）+key（4b）+valuesize（4b）+value（4b）= 16b
1TB 最大的存储： 2^(40-4)=2^36 个kv
每个segment 大小： 64MB 
每个entry 大小： hash（8b）+ offset（8b）= 16b
一个segment存储：2^(26-4)=2^22个entry
总共需要： 2^(36-22)=2^12 个segment 文件进行存储

这边需要使用bloom过滤器，因此开始计算m，k，假定误报率fpp小于%5
m：bit大小 m=-(n*lnp)/(ln2)^2=-(n*lnp)/(ln2)^2=40202648=2^25bit=8Mb
k: hash次数 k=(m/n)*ln2=6

将2GB作为内存索引
一个支持 2^(11-3)=2^9 索引
公共需要 2^(12-9)=2^3 




## 预处理
![预处理](../assert/ans1-preread.png)


## 方案
## 存在的问题
1. hash 冲突
2. 误判率
