# 
# 项目介绍
本仓库为仿抖音app服务端开发。<br />[接口文档](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523)<br />[基础版抖音app使用说明](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)<br />`服务端部署地址`：[http://114.55.132.72:7080/](http://114.55.132.72:7080/)<br />极简版抖音项目划分为两大方向，`互动方向`和`社交方向`，两个方向均包含`基础功能`内容，在扩展功能上有所不同，具体内容见下表 

|  | **互动方向** |  | **社交方向** |  |
| --- | --- | --- | --- | --- |
| 基础功能项  | 视频 Feed 流、视频投稿、个人主页  |  |  |  |
| 基础功能项说明  | 视频Feed流：支持所有用户刷抖音，视频按投稿时间倒序推出 <br />视频投稿：支持登录用户自己拍视频投稿 <br />个人主页：支持查看用户基本信息和投稿列表，注册用户流程简化  |  |  |  |
| 方向功能项  | 喜欢列表  | 用户评论  | 关系列表  | 消息  |
| 方向功能项说明  | 登录用户可以对视频点赞，在个人主页喜欢Tab下能够查看点赞视频列表  | 支持未登录用户查看视频下的评论列表，登录用户能够发表评论  | 登录用户可以关注其他用户，能够在个人主页查看本人的关注数和粉丝数，查看关注列表和粉丝列表  | 登录用户在消息页展示已关注的用户列表，点击用户头像进入聊天页后可以发送消息  |

# 二、项目分工
| **团队成员** | **主要贡献** |
| --- | --- |
| | 数据库表搭建，项目互动接口和消息接口开发，代码审核及总体功能测试，总体功能bug修复及最终版项目审核 |
| | 负责`项目基础模块的接口`及测试，主要包括接口鉴权、对象存储上传下载中间件开发，数据库表设计，服务部署 |
| | 项目社交接口关注及粉丝接口开发，项目总体功能测试问题反馈 |

![](https://cdn.nlark.com/yuque/0/2023/png/2711588/1677461906801-cfa3ade0-38da-4100-858c-34986a7570f1.png#averageHue=%232f2d2c&clientId=u04fc5bcb-906d-4&from=paste&id=u5d2ce7e5&originHeight=766&originWidth=842&originalType=url&ratio=2&rotation=0&showTitle=false&status=done&style=none&taskId=u8f7231e0-8139-4d8a-907a-32abc6d95cf&title=)
# 项目实现
### 3.1 技术选型与相关开发文档
开发环境：win10+go1.19+mysql8.0+ffmpeg<br />web框架：go-gin<br />数据库操作：gorm
### 3.2 架构设计
**项目架构图：**<br />![](https://cdn.nlark.com/yuque/0/2023/png/2711588/1677461952076-55af0af5-5155-4849-9148-6cc8c087a2a2.png#averageHue=%23f5f5f5&clientId=u04fc5bcb-906d-4&from=paste&id=u210f90ef&originHeight=599&originWidth=710&originalType=url&ratio=2&rotation=0&showTitle=false&status=done&style=none&taskId=u1ba3f796-77bd-477f-9b4a-fc0c65fc10d&title=)
### 3.3 效果展示
![截屏2023-02-27 09.40.53.png](https://cdn.nlark.com/yuque/0/2023/png/2711588/1677462276221-f75626c3-15df-4235-ad90-2df518e2230d.png#averageHue=%23625f5a&clientId=u04fc5bcb-906d-4&from=drop&height=533&id=u2cf931e6&name=%E6%88%AA%E5%B1%8F2023-02-27%2009.40.53.png&originHeight=1124&originWidth=704&originalType=binary&ratio=2&rotation=0&showTitle=false&size=583481&status=done&style=none&taskId=u1c1fd69e-dcc9-45bf-abb2-fdab4abcdc8&title=&width=334)![截屏2023-02-27 09.41.35.png](https://cdn.nlark.com/yuque/0/2023/png/2711588/1677462294838-a6353dbc-0578-4bf2-9f82-ee88dd90b174.png#averageHue=%233e4b58&clientId=u04fc5bcb-906d-4&from=drop&height=526&id=u22d59fcc&name=%E6%88%AA%E5%B1%8F2023-02-27%2009.41.35.png&originHeight=1112&originWidth=694&originalType=binary&ratio=2&rotation=0&showTitle=false&size=802577&status=done&style=none&taskId=uc916f329-e25b-4364-808a-44a0c8013ba&title=&width=328)<br /> 
# 总结与反思
### 1. 视频表
存储视频唯一id，播放地址，封面地址，获赞数，评论数，自己是否点赞<br />Id 使用唯一id生成器、雪花算法<br />后面三个怎么存，表怎么设计，考虑数据量大，查询频繁，更新频繁
### 2. 上传视频 标题为表情
failed to insert Error 1366 (HY000) ---> 	修改数据库title字段 为utf8mbp4
### 3. 视频查找过慢 ---> 索引优化
### 4. 上传视频异步化 -> 可靠性
### 5. 客户端点赞按钮存在退出登录后可再次点赞问题，服务端已解决可重复点赞但不能存入重复数据。
### 6. 未使用redis等非关系型数据库技术进行缓存或预加载，所有请求均直接与数据库对接，查询加载效率低，无法应对高并发量。
后面可引入缓存

