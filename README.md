

# golang的学习笔记

## singleton 单例（sync.Once）

## option 选项设计模式（函数选项, 接口选项-不对外暴露结构体）
why use
- 提供了默认值，同时允许灵活配置
- 可以轻松添加新的选项而不破坏现有代码
- 使用起来非常直观和可读

[Go语言设计模式之函数选项模式 | 李文周的博客](https://www.liwenzhou.com/posts/Go/functional-options-pattern/)

## builder 建造者模式
