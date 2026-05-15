"use strict";
// 第 3 课：Promise 和 async / await
//
// 本课目标：
// 1. 知道 Promise 是“未来才会拿到的结果”
// 2. 知道 async 函数默认返回 Promise
// 3. 会用 await 等待异步结果
// 4. 知道 catch 在哪里处理错误
//
// 抄写时建议你自己补这些内容：
// - 一个 Post 类型
// - 一个 fakeFetchPosts(): Promise<Post[]>
// - 一个 async 函数
// - 一个 await 调用
// - 一个 catch 错误处理
//
// 练习：
// 1. 给 Post 增加 author 字段
// 2. 再写一个 fakeFetchUserName(): Promise<string>
// 3. 在 async 函数里同时打印用户名和文章列表
// 4. 思考：为什么前端请求接口时几乎一定会遇到 Promise？
Object.defineProperty(exports, "__esModule", { value: true });
