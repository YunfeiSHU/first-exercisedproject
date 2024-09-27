1  设计model-----Data Model
model对应:定义数据结构和关系，通常对应数据库中的表或集合。
1.http层POST的请求与响应传递数据时的结构体---对应的接口（用接口存储方法）
2.http层GET，响应时的结构体----对应的接口
3.数据库中表---对应接口

2  设计api-----Controller Layer
controller：实现请求方法
middleware: jwt中间件
route：创建并接收指定路由的请求方法

补充：usecase层
由于repository层只对接User,api的相关操作与User的CRUD有关，
故usecase层：嵌套UserRepository,以此利用user的CRUD实现接口



3  database(避免重名，不用mysql):实现与数据库的交互，包括连接管理、查询等,作为抽象层
连接mysql数据库，定义接口和结构体

抽象层，是将具体操作，放在接口中实现

4  repository：给 mysql 提供对数据源的CRUD操作。
对数据库的users表的CRUD

repository 设计方法，实现对mysql中接口的对接

3，4总结：将数据库操作抽象成接口(mysql)，并在仓储（repository）层实现对这些接口的对接

5  Server：main

6  表示层：前端


设计internal：私人pkg
jwt_token:创建，验证，更新，更加tokne返回用户信息

config/bootstrap:
env配置文档，app(mysql+gin),...(redis)
