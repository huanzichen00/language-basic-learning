"use strict";
// 第 7 课：泛型基础
//
// 本课目标：
// 1. 理解泛型是在“延后决定具体类型”
// 2. 会看懂函数上的 <T>
// 3. 会看懂泛型类型别名
// 4. 知道它在前端里常用于接口返回值和通用工具函数
//
// 抄写时建议你自己补这些内容：
// - 一个 identity<T> 函数
// - 一个把单个值包成数组的泛型函数
// - 一个 ApiResult<T> 类型
//
// 练习：
// 1. 写一个 getFirstItem<T>(items: T[]): T | undefined
// 2. 写一个 Pair<T> 类型，包含 first 和 second
// 3. 尝试定义 ApiResult<string[]>
// 4. 思考：为什么不用为 string、number、boolean 各写一套函数？
Object.defineProperty(exports, "__esModule", { value: true });
