Token认证模式

Token 是无状态会话管理方式,服务端不再存储信息 甚至不再使用Session , 逻辑如下
    1.客户端使用用户名,密码进行认证
    2.服务端验证用户名,密码正确后生成Token返回客户端
    3.客户端保存Token,访问需要认证的接口时在URL参数或HTTP Header中加入Token
    4.服务端通过解码Token进行鉴权,返回给客户端需要的数据

**优点**
    1.服务端不用存储和用户鉴权有关的信息,鉴权信息会被加密到Token中,服务端只需要读取Token中包含的鉴权信息即可
    2.避免了共享Session导致的不易扩展的问题
    3.不需要依赖Cookie
    4.使用CORS即可快速解决跨域问题

JWT介绍
    JWT 是一种基于JSON格式的Token标准.
    一个JWT Token由 **.** 分隔的三部分组成,这三部分依次是:
        1.头部(Header) : 存储了所使用的加密算法和Token类型
            {
                "alg" : "HS256",
                "typ" : "JWT"
            }
        2.负载(Payload) 提供了七个字段使用
            {
                iss(issuer) : 签发人
                exp:过期时间
                sub : 主题
                aud : 受宠
                nbf : 生效时间
                jat : 签发时间
                jti : 编号
            }
            负载也可以使用自己指定的字段,如:
            {   
                "sub" : "1221",
                'name' : "John Doe"
            }
        **JWT 默认是不加密的,任何人都可以看到,不能存放私密信息**
        3.签名 : 对前两部分的签名,防止数据被篡改
    头部和负载以JSON形式存在,这就是JWT中的JSON,三部分都分别经过了Base64编码,以**.**拼接成一个Token.

**Access Token**: 访问资源接口时所需要的Token.
**Refresh Token**: 当Access Token由于过期失效之后,使用Refresh Token 就可以获取到新的Assess Token,如果Refresh Token 失效了,用户就只能重新登录
