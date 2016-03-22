#### angularjs服务
angularjs的服务提供了一种能在应用的整个生命周期内保持数据的方法，它能够在控制器之间进行通信，并且保持数据的一致性。
服务是一种单例对象，在每个应用中只会被实例化一次，并且延迟加载（需要时才会被加载）。  
在控制器、指令、过滤器或者另外的服务中通过依赖声明的方式来使用服务。  
创建服务的方式有如下5种：  
1. factory()
2. service() 
3. constant()
4. value()  
5. provider()  

#### factory()方式
这是最简单最灵活的方式。它接受两个参数：  
name (字符串) ： 需要注册的服务名
getFn (函数)： 该函数会在Angularjs创建服务实例时被调用
```
angular.module("myApp", [])
.factory("myService", function(){
  return {
     ...
  }
})

// getFn也可以接受一个包含可被注入对象的数组或者函数
angular.module("myApp", [])
.factory("myService", ["$http", function($http){
  return {
     getUser: function(username){
       ...
     }
  }
}])
```

#### service()方式
使用该方式，可以为服务对象注册一个构造函数。可以知道angularjs内部会使用new的方式来实例化该对象。同样也接受两个参数。  
name (字符串) ： 需要注册的服务名  
construction (函数) 构造函数。
```
angular.module("myApp", [])
.service("myService", function($http){
  this.getUser = function(username){
     ...
  }
})
```
只有this之后的属性才能被当成服务的属性被外部controller或者指令使用。

#### constant()方式
可以将一个已经存在的变量值注册为一个服务。接受两个参数：   
name (字符串) 常量的名字  
value (常量) 常量的值
```
angular.module("myApp", []).constant("apiKey", "123456789")
```

#### value()方式
也可以注册一个常量，但是它不能注入到配置函数（config）中，constant可以。
```
angular.module("myApp", []).value("apiKey", "123456789")

// 以下会报错
angular.module("myApp", []).constant("apiKey", "123456789")
.config(function(apiKey){
  ..
})
```

#### provider()方式
上述创建服务的所有方式，内部最终都是使用$provider来进行创建的。我们也因此可以使用$provider的方式进行服务的创建。使用该种方式创建服务的好处就是可以进行服务的扩展，即可以在config()中扩展服务。  
提供者是一个具有$get()方法的对象，$injector通过调用$get方法创建服务实例。即$get方法返回的就是我们可以在控制器等使用的服务service。   
注意： $get必须是方法
  
当我们传递给提供者的函数就是$get()时，也就是我们看到的factory()模式。
```
angular.module("myApp", [])
.factory("mySerivce", function(){
  return {
     username: "auser"
  }
})

// 这与上面是等价的
angular.module("myApp", [])
.provider("mySerivce", {
  $get: function(){
    obj = {username: "auser"}
    return obj;
  }
})
```
接下来我们来详细看下这个方法，provider方法包含两个参数：  
name (字符串) 该参数将会是服务实例的名字，而name+Provider，如userProvider，将会成为服务的提供者。  
aProvider (对象/函数/数组)  
如果aProvider是对象，则该对象必须带有$get方法，而对象中的属性则可以在config中访问到，并且可以被修改。
```
myApp = angular.module("myApp", [])
myApp.provider("myService", {
  fav: "ball",
  "$get": function() {
    var obj = {
      sex: 'male',
      fav: this.fav
    }
    return obj;
  }
})

// 接下来尝试在config中修改对象值
myApp.config(["myService1Provider",
 function(myService1Provider){
    myService1Provider.fav = "fish"
}])

// 在controller中使用
.controller("mainController", ["myService1",
  function(myService1){
    var sex = myService1.sex;
    var fav = myService1.fav;
    console.log(sex);  // male
    console.log(fav);  // fish
}]);
```  

如果aProvider是函数，它将会通过依赖注入被调用，并且通过$get方法返回一个对象。  
```
angular.module("myApp", [])
.provider("mySerivce2", function(){
  var name = "Jim";
  var age = 10;
  this.backUrl = "http://localhost:8080/"
  this.setBackUrl = function(url){
    this.backUrl = url
  }
  var that = this;
  this.$get = function(){
     return {
       getBackUrl: function(){
         return that.backUrl
       }
       getAge: function(){
         return age;
       }
       getName: function(){
         return that.name;
       }
     }
  }
})

// 使用provider的好处就是我们可以在config中扩展该服务
angular.module("myApp")
.config(["myService2Provider", function(myService2Provider){
  myService2Provider.backUrl = "http://127.0.0.1:8080";
  myService2Provider.name = "Tom"
  // 在config中是无法访问私有属性的，如age,name这些
  // 因此上面虽然把name改为了Tom，但是它是类似于在provider中使用this.name
  // 实际上this中是没有name和age这个两个属性的
}])

// 此时我们便可以在controller等地方通过myService调用$get里的方法了
//我们只能调用$get方法返回的对象中的方法。如下所示：
angular.module("myApp")
.controller("mainController", ["$scope", "myService2", 
  function($scope, myService2){
    var url = myService2.getBackUrl();
    var age = myService2.getAge();
    var name = myService2.getName();
    console.log(url);  // http://127.0.0.1:8080
    console.log(age);  // 10
    console.log(name); // Tom
    // 注意到name是在config中修改后的值，因为我们在provider中getName我们返回的是that.name也就是this.name。因此config中的配置就生效了。
  }
])
```

如果aProvider是数组，会被当作一个带有行内依赖注入声明的函数来处理，数组的最后一个元素应该是个函数，可以返回一个带有$get方法的对象。  
需要注入其他服务时可以采用这种方式。
```
angular.module("myApp", [])
.provider("myService3", ["$http", function($http){
  $get: function(){
     return {...}
  }
}])
```

由此可以看得出，provider是最强大的。因此如果你的服务是需要开源出来，可供config中扩展的，就应该使用provider的方式创建服务。一般情况下，factory则足以应付需求了。

#### delegate装饰器
除了以上几种服务外，还有一种可以被称为装饰器的东西，它有$provide提供，可以在服务实例创建时对其进行拦截的功能，因此可以对服务进行扩展。  
它接受两个参数：  
name (字符串) 将要拦截的服务的名称   
decoration (函数)  
```
myApp.provider("myService4", function(){
  this.food = "apple"
  var that = this;
  this.$get = function() {
    return {
      getFood: function(){
        return that.food
      },
      getMoney: function(){
        return "no money"
      }
    }
  }
})

myApp.config(["$provide", function($provide){
  $provide.delegate("myService4", function($delegate){
    var getFood = function(){
      var food = $delegate.getFood()
      console.log("ok, this is delegate!!!!")
      return food;
    }
    return {
      getFood: getFood,
      // 不返回getMoney则调用服务时便没有该方法了
      getMoney: $delegate.getMoney
    }
 })
}])
.controller("mainController", ["myService4", function(myService4){
  var food = myService3.getFood()
  console.log(food);
  var money = myService3.getMoney()
  console.log(money);
}])

// 打印结果为
ok, this is delegate!!!!
apple
no money
```
