# SOLID原则

[toc]

## 依赖反转

SOLID的最后一个原则是Dependency Injection（DI），即依赖注入原则。在认识DI前，先了解下另一个常见的概念，Inversion Of Controll（IOC），即控制反转。

### 控制反转（Inversion Of Controll）

这里的控制，指的是对程序运行的控制。控制反转的核心就是要将控制逻辑和业务逻辑分开，不要在业务逻辑里写控制逻辑，这会导致控制逻辑依赖于业务逻辑，而是应该让业务逻辑依赖控制逻辑。

有个电灯和电器的例子，电灯是控制逻辑，电器是业务逻辑，不要在电器中实现开关，而是要把开关设计成一种协议，让电器都依赖于这种协议，降低代码复杂度，提高代码复用度。

换成程序代码的话，正常情况下，对函数调用都是知道其函数名、参数类型和返回值，属于主动的调用。当采用控制反转后，主调方就不是直接调用函数，而是使用框架代码间接的调用。

以一个用户注册后，给邮箱发送邮件为例子。在不使用反转控制的情况下，我们的代码可能会这么写。

```go
type UserRegister struct {
  Email 			string
  Password 		string
  Repassword 	string
}

func (u UserRegister) Regist() error {
  // 参数校验，密码md5，用户信息写入数据库
  ...
  
  if err := u.sendEmail(); err!=nil{
    return err
  }
  
  return nil
}

func (u UserRegister) sendEmail() error {
  fmt.Printf("send email to %s, subject:%s, content:%s", u.Email, "regist success", "regist success, please confirm!")
  return nil
}
```

从功能来看，好像这么写也没什么问题。这里的用户注册功能是一个业务逻辑，而发送邮件算是个控制逻辑，除了注册之外，可能其他业务功能都可能会用到发送邮件的功能。因此好的做法应该是将发送邮件这个逻辑剥离出来，让用户注册功能依赖发送邮件这个功能，而不是发送邮件依赖用户注册。

```go
type EmailSender struct {
  To 			[]string
  Cc 			[]string
  Subject string
  Content string
}

func (e EmailSender) Send() error {
  fmt.Printf("send email to %s, cc %s, subject:%s, content:%s", e.To, e.Cc, e.Subject, e.Content)
  return nil
}

type UserRegister struct {
  Email 			string
  Password 		string
  Repassword 	string
  emailSender EmailSender
}

func (u UserRegister) Regist() error {
  // 参数校验，密码md5，用户信息写入数据库
  ...
  
  if err := u.sendEmail(); err!=nil{
    return err
  }
  
  return nil
}

func (u UserRegister) sendEmail() error {
  e.emailSender = EmailSender{
    To: []string{u.Email},
    Subject: "regist success",
    Content: "regist success, please confirm!"
  }
  return e.emailSender.Send()
}
```

通过改造，我们的用户注册就依赖于发送邮件这个逻辑了，代码得到了解耦，邮件发送这个功能经过抽象，对外只提供了发送的协议，例如需要哪些字段，这样能被更多的地方复用，这便是控制反转。

### 依赖注入（Dependency Injection）

不通过 new() 的方式在类内部创建依赖类对象，而是将依赖的类对象在外部创建好之后，通过构造函数、函数参数等方式传递（或注入）给类使用。

针对上面用户注册的代码，我们稍微调整下，先来看下，不采用依赖注入的时候是怎样的。

```go
type EmailSender struct {
  To 			[]string
  Cc 			[]string
  Subject string
  Content string
}

func (e EmailSender) Send() error {
  fmt.Printf("send email to %s, cc %s, subject:%s, content:%s", e.To, e.Cc, e.Subject, e.Content)
  return nil
}

type UserRegister struct {
  Email 			string
  Password 		string
  Repassword 	string
  emailSender EmailSender
}

func NewUserRegister(email, password, repassword string) UserRegister {
  return UserRegister{
    Email: email,
    Passowrd: password,
    Repassword: repassword,
    emailSender: EmailSender{
    	To: []string{email},
    	Subject: "regist success",
    	Content: "regist success, please confirm!"
  	}
  }
}

func (u UserRegister) Regist() error {
  // 参数校验，密码md5，用户信息写入数据库
  ...
  
  if err := u.sendEmail(); err!=nil{
    return err
  }
  
  return nil
}

func (u UserRegister) sendEmail() error {
  return e.emailSender.Send()
}
```

Go没有构造函数，因此只能手动编写创建实例的函数，这里NewUserRegister创建了一个UserRegister的实例，内部创建EmailSender的实例。

接下来看下，采用依赖注入的方式，上述代码会如何调整。

```go
type EmailSender struct {
  To 			[]string
  Cc 			[]string
  Subject string
  Content string
}

func (e EmailSender) Send() error {
  fmt.Printf("send email to %s, cc %s, subject:%s, content:%s", e.To, e.Cc, e.Subject, e.Content)
  return nil
}

type UserRegister struct {
  Email 			string
  Password 		string
  Repassword 	string
  emailSender EmailSender
}

func NewUserRegister(email, password, repassword string, sender EmailSender) UserRegister {
  return UserRegister{
    Email: email,
    Passowrd: password,
    Repassword: repassword,
    emailSender: sender,
  }
}

func (u UserRegister) Regist() error {
  // 参数校验，密码md5，用户信息写入数据库
  ...
  
  if err := u.sendEmail(); err!=nil{
    return err
  }
  
  return nil
}

func (u UserRegister) sendEmail() error {
  return e.emailSender.Send()
}

s := EmailSender{
  To: []string{email},
  Subject: "regist success",
  Content: "regist success, please confirm!"
}
r := NewUserRegister("seed@123.com", "123", "123", s)
r.Regist()
```

这里EmailSender的实例通过参数的方式进行注入，不再是内部进行实例化，这样代码的灵活性就得到提高了，可以灵活的替换依赖的结构体。

更加好的方式，是将EmailSender替换成接口。

```go
type SmsSender interface {
  Send() error
}

type UserRegister struct {
  Email 			string
  Password 		string
  Repassword 	string
  smsSender 	SmsSender
}

func NewUserRegister(email, password, repassword string, sender SmsSender) UserRegister {
  return UserRegister{
    Email: email,
    Passowrd: password,
    Repassword: repassword,
    smsSender: sender,
  }
}

func (u UserRegister) sendEmail() error {
  return e.smsSender.Send()
}
```

通过接口的方式，对编写单元测试也更加的友好。

### 依赖反转原则（Dependency Inversion Principle）

高层模块（high-level modules）不要依赖低层模块（low-level）。高层模块和低层模块应该通过抽象（abstractions）来互相依赖。除此之外，抽象（abstractions）不要依赖具体实现细节（details），具体实现细节（details）依赖抽象（abstractions）。

对应于调用链，调用者属于高层，被调者属于低层，但是实际工作中，调用者依赖于被调者无可厚非，但是在开发框架时，我们便可以采用前面描述的控制反转来指导我们的设计。

例如我们要开发一个日志上报组件，当web服务写日志的时候，例如log.Info()，日志便能自动上报到远程日志收集的集群。这个日志组件只需要根据统一的日志上报协议，然后作为插件注入到web服务中，高层调用者，只需要如同往常写日志即可。