# Java 实习面试复习路线

目标：你不是从零学，而是把“知道但容易忘”的知识重新串起来。每一轮都按这个节奏复习：

1. 先看八股答案骨架
2. 再跑小程序验证现象
3. 最后用一句话记忆钩子收口

## 第一阶段：Java 基础高频

这部分最适合先拿下，因为面试里出现频率高，而且能快速建立状态。

重点题：

- `==` 和 `equals()` 的区别
- `hashCode()` 为什么要和 `equals()` 一起重写
- String、StringBuilder、StringBuffer 的区别
- 基本类型和包装类型的区别
- 自动装箱与拆箱
- Integer 缓存
- Java 是值传递还是引用传递
- 重载和重写
- `final`、`finally`、`finalize`
- 接口和抽象类的区别

## 第二阶段：集合

重点题：

- ArrayList 和 LinkedList
- HashMap 原理
- HashMap 在 JDK 7 和 JDK 8 的区别
- ConcurrentHashMap 原理
- HashSet 为什么能去重
- fail-fast 和 fail-safe

## 第三阶段：并发

重点题：

- 线程和进程的区别
- `sleep()` 和 `wait()` 的区别
- synchronized 和 ReentrantLock
- volatile 的作用
- CAS 和 ABA
- 线程池参数和工作流程
- ThreadLocal 原理

## 第四阶段：JVM

重点题：

- JVM 内存模型 / 运行时数据区
- 堆、栈、方法区
- 对象创建过程
- 垃圾回收算法
- CMS 和 G1
- 类加载过程
- 双亲委派模型

## 第五阶段：Spring 常问

重点题：

- Spring IOC 和 AOP
- Bean 生命周期
- 循环依赖
- Spring 事务失效场景
- Spring Boot 自动配置原理

## 第六阶段：数据库与 Redis

重点题：

- 事务四大特性
- 隔离级别
- MVCC
- 索引结构
- Redis 持久化
- Redis 缓存穿透 / 击穿 / 雪崩

## 当前建议顺序

先把“Java 基础高频”拿稳，再进集合和并发。实习面试里，很多追问都从基础题一路延伸过去。

## 记忆方法

- 区分题：记“定义 + 关键差异 + 一句场景”
- 原理题：记“结构 + 过程 + 为什么这么设计”
- 场景题：记“出现原因 + 后果 + 解决办法”

## 本轮配套程序

先看并运行：

- `java/01-core-basics/CoreBasicsDemo.java`

运行方式：

```bash
javac java/01-core-basics/CoreBasicsDemo.java
java -cp java/01-core-basics CoreBasicsDemo
```
