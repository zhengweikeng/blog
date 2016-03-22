angularjs中内置了很多指令（Directive），常用的如ng-repeat、ng-class等等，我们也可以自己定义指令，而且在项目中是非常普遍的。  

我们经常进行很多DOM操作，如点击事件，表单验证事件等等，angular不建议我们将这些操作放在控制层，也即Contrller处，而是应该在指令处进行DOM的操作。也算是为了减轻控制层的负担吧。  

<a href="#1">定义指令</a>
<a href="#2">定义指令的行为</a>
<a href="#3">restrict</a>
<a href="#4">priority</a>
<a href="#5">terminal</a>
<a href="#6">template</a>
<a href="#7">templateUrl</a>
<a href="#8">replace</a>
<a href="#9">scope</a>
<a href="#10">transclude</a>
<a href="#11">controllerAs</a>
<a href="#12">require</a>
<a href="#13">compile</a>
<a href="#14">link</a>

<span id="1"></span>
###### 定义指令
angular.module("myApp", []).directive(name, fn)
name (string) 指令的名称
fn (function) 该函数返回指令的全部行为  
当然，我们也可以往函数中注入各种服务，关于服务的定义可参考[此文](http://zhengweikeng.duapp.com/angularjsfu-wu-de-yong-fa/)  
请看示例：  
```
angular.module("myApp", []).
directive(name, function($http, $timeout, YOUR_SERVICE){
  // 指令的行为
})
```

<span id="2"></span>
###### 定义指令的行为
指令的行为可以是非常简单，也可以非常复杂，官方给了我们很多配置项，实际使用时只能具体问题具体分析了。   
指令的行为最终会返回一个object对象，对象以键值对的方式定义，如下所示：  
```
angular.module('myApp', [])
.directive('myDirective', function() {
  return {
    restrict: String,
    priority: Number,
    terminal: Boolean,
    template: String or Template Function:
      function(tElement, tAttrs) (...},
    templateUrl: String,
    replace: Boolean or String,
    scope: Boolean or Object,
    transclude: Boolean,
    controller: String or
    function(scope, element, attrs, transclude,  otherInjectables) { ... }, controllerAs: String,
    require: String,
    link: function(scope, iElement, iAttrs) { ... },
    compile: // 返回一个对象或连接函数,如下所示:
      function(tElement, tAttrs, transclude) {
        return {
          pre: function(scope, iElement, iAttrs, controller) { ... }, 
          post: function(scope, iElement, iAttrs, controller) { ... }
        }
        // 或者
        return function postLink(...) { ... }
     }
  };
});
```
看起来复杂，实际使用指令的时候只会使用到一部分。  

<span id="3"></span>
###### restrict （字符串）
该属性定义了属性在视图上的使用方式，默认值是A。备选值有：  
A：即Attribute，以属性的方式定义
E：即Element，以元素的方式定义
C：即Class，以类的方式定义
M：即Comment，以注释的方式定义
参见实例：
```
angular.module("myApp", [])
.directive("myDirecitve", function(){
  return {
    restrict: 'AECM'
  }
})

// 视图上使用属性的方式
<div my-direcitve></div>

// 视图上使用元素的方式
<my-direcitve></my-direcitve>

// 视图上使用类的方式
<div class="my-direcitve"></div>

// 视图上使用注释的方式
<!--directive:my-direcitve -->
```
一般情况下，我们会用AE，类和注释还是比较少用的，尤其是注释，一般不推荐使用。类的话是为了避免和我们定义样式时产生歧义。  
注意点可以看到我们定义指令时是用驼峰定义法定义指令的名称的，但是视图上使用是却用了-，这是因为HTML是不区分大小写的，因此angular则规定了视图上使用-分隔两个单词。

<span id="4"></span>
###### priority （数值型）
顾名思义，即定义指令的优先级，大部分情况下都不会去定义它，即默认为0。定义它可以让它优先于其他指令运行。ng-repeat指令则定义了优先级为1000。

<span id="5"></span>
###### terminal （布尔型）
用于告诉angular，优先级比该指令低的其他指令都停止执行，不过同级的指令依旧会执行。  
使用了terminal参数的例子是ngView和ngIf。ngIf的优先级略高于ngView,如果ngIf的表达式值为true,ngView就可以被正常执行,但如果ngIf表达式的值为false,由于ngView的优先 级较低就不会被执行。

<span id="6"></span>
###### template （字符串或者函数）
该属性可以用于定义嵌入到指令中的HTML文本。  
该属性可以是一个HTML字符串，也可以是一个函数。  
如果是函数，则该函数会接受两个参数，tElement和tAttrs，函数中会返回HTML字符串。   
HTML字符串会被angular正确解析，而且HTML中可以包含angular的表达式，使用{{}}包围即可，它会去访问指令的作用域中执行该表达式。  
另外需要注意的一点是，HTML字符串中的必须存在一个DOM根元素。  
请看示例：
```
angular.module("myApp", [])
.directive("myDirecitve", function(){
  temp = "<div><a href="{{link}}">click me</a><h1>When using two elements, wrap them in a parent element</h1></div>"
  return {
    restrict: 'AE',
    template: temp
  }
})

// 下面这种方式是会报错的
angular.module("myApp", [])
.directive("myDirecitve", function(){
  // 由于该HTML字符串没有根元素，因此会报错
  temp = "<a href="{{link}}">click me</a><h1>When using two elements, wrap them in a parent element</h1>"
  return {
    restrict: 'AE',
    template: temp
  }
})
```

<span id="7"></span>
###### templateUrl （字符串或函数）
实际中，我们不会使用template，因为我们HTML字符串不会像上面那么简单只有几行代码，因此我们会将它放在一个模板文件中。  
该属性的值同样也可以是个字符串或者函数。  
如果是字符串，则为模板的路径。  
如果是函数，则该函数也会接受两个参数，tElement和tAttrs，函数中会返回模板文件的路径。   
默认情况下,调用指令时会在后台通过Ajax来请求HTML模板文件。有两件事情是需要知道：  
1. 在本地开发时,需要在后台运行一个本地服务器,用以从文件系统加载HTML模板,否则会导致Cross Origin Request Script(CORS)错误。  
2. 模板加载是异步的,意味着编译和链接要暂停,等待模板加载完成。
还有一个提醒：  
如果页面都是通过AJAX的方式异步加载模板文件，将会大幅度的影响网页的加载速度，页面将会出现明显的卡顿现象。因此我们可以事先将模板文件都缓存起来，即换存在angular的$templateCache服务中。  
这里提供两个插件，一个是grunt的[grunt-angular-templates](https://github.com/ericclemmons/grunt-angular-templates)，另一个是gulp的[gulp-angular-templatecache](https://github.com/miickel/gulp-angular-templatecache)。

<span id="8"></span>
###### replace （布尔型）
该属性是个可选值，如果定义了该属性，则值必须为true，因为默认值为false。  
当为默认值false时，上述的HTML模板将会作为子元素插入到指令的元素内部。  
当值为true时，HTML模板将会替换掉指令元素。  
```
// 视图上的指令
<div my-direcitve></div>

// 定义指令
angular.module("myApp", [])
.directive("myDirecitve", function(){
  temp = "<div>some text...</div>"
  return {
    restrict: 'AE',
    template: temp //,
    // replace: true
  }
})

// 此时视图将为
<div my-direcitve><div>some text...</div></div>

// 如果放开指令中replace的注释的话，将会是：
<div>some text...</div>
```

<span id="9"></span>
###### scope （布尔型或对象）
默认情况下，指令是被赋予允许访问父DOM元素对应作用域的能力，此时scope即为即默认值，即为false。  
当该值为false时，将会直接访问父元素的作用域，此时指令中对作用域中元素的修改也会直接作用在父元素的作用域中。  
当该值为true时，则会继承父元素作用域，并创建一个新的作用域对象。
```
// 视图
<div ng-init="id='001'">
  id: {{id}}
  <div my-direcitve></div>
</div>

// 定义指令
angular.module("myApp", [])
.directive("myDirecitve", function(){
  temp = "<div>{{id}}</div>"
  return {
    restrict: 'AE',
    template: temp,
    //scope: true,
    link: function(scope){
      scope.id = 002;
    }
  }
})

// angular解析后的视图（去掉注释）
<div ng-init="id='001'">
  <!-- 同时改变了父作用域中变量的值 -->
  id: 002
  <div my-direcitve>002</div>
</div>

// 如果放开指令中的注释的话，则视图将会为（去掉注释）
<div ng-init="id='001'">
  <!-- 不会改变父作用域中变量的值 -->
  id: 001
  <!-- 拥有了独立的作用域 -->
  <div my-direcitve>002</div>
</div>
```
指令定义中scope属性最强大功能应该是定义**隔离作用域**。我们可以将scope的值定义为一个对象，如空对象{}，此时指令的作用域将会完全独立于外界。
```
<div ng-init="myProperty='wow,this is cool'">
  Surrounding scope: {{ myProperty }}
  <div my-inherit-scope-directive></div>
  <div my-directive></div>
</div>

angular.module('myApp', [])
.directive('myDirective', function() {
  return {
    restrict: 'A',
    template: 'Inside myDirective, isolate scope: {{ myProperty }}',
    // 定义了隔离作用域
    scope: {} 
  };
})
.directive('myInheritScopeDirective', function() {
  return {
    restrict: 'A',
    template: 'Inside myDirective, isolate scope: {{ myProperty }}',
    scope: true
  };
});

// 解析后的视图
<div ng-init="myProperty='wow,this is cool'">
  Surrounding scope: wow,this is cool
  <div my-inherit-scope-directive>Inside myDirective, isolate scope: wow,this is cool</div>
  <div my-directive>Inside myDirective, isolate scope: </div>
</div>
```
将scope定义为空对象也不是那么常见的，我们可以将它和指令外部的作用域进行数据绑定。  
可以进行如下三种方式的数据绑定。  
1. 本地作用域属性：使用@符号进行绑定，类似于将父作用域中的变量拷贝一份到指令的独立作用域中。  
2. 双向绑定：使用=符号进行绑定，此时在独立作用域中修改变量的值也会同步到父作用域中。  
3. 父级作用域绑定：通过&符号可以对父级作用域进行绑定,以便在其中运行函数。意味着对这个值进行设置时会生成一个指向父级作用域的包装函数。  
```
// 视图
<div ng-app="MyApp">
  <div class="container" ng-controller="myController">
    <div class="my-info">
      我的名字是：<span ng-bind="name"></span>
      <br/>
      我的年龄是：<span ng-bind="age"></span>
      <br />
    </div>
    <div my-directive my-name="{{name}}" age="age" change-my-age="changeAge()"></div>
  </div>
</div>

// 指令定义
angular.module("MyApp", [])
.controller("myController", function ($scope) {
  $scope.name = "Tom";
  $scope.age = 20;
  $scope.changeAge = function(){
    $scope.age = 10;
  }
})
.directive("myDirective", function () {
  var obj = {
    restrict: "AE",
    scope: {
      name: '@myName',
      age: '=',
      changeAge: '&changeMyAge'
     },
     replace: true,
     template: "<div class='my-directive'>" +
       "<h3>下面部分是我们创建的指令生成的</h3>" +
       "我的名字是：<span ng-bind='name'></span><br/>" +
       "我的年龄是：<span ng-bind='age'></span><br/>" +
       "修改名字：<input type='text' ng-model='name'><br/>"+
       "<button ng-click='changeAge()'>修改年龄</button>"+
       "</div>"
  }
  return obj;
});
```
由上述例子可以看到@，=，&后面的值即为视图元素中的属性名，在scope定义时是可以忽略的，如果这些前缀后面是没有值的，则值将会和scope中定义的属性名一样，即age将会去找age属性。   
1. @的用法，即数据是单向绑定，注意，属性的名字要用-将两个单词连接，而且因为是数据的单项绑定所以要通过使用{{}}来绑定数据，如果没有{{}}则会识别为字符串。
2. ＝的用法，即数据的双向绑定，不用使用{{}}
3. &的用法，放字符用于绑定函数，同理，属性的名字也要用-将多个个单词连接。

注意：  
在做项目的时候遇到一个问题，&的用法值得注意。  
&用于绑定一个父控制器的函数，若需要向该函数传递参数，则必须使用json的格式的参数。
```
angular.directive("myDirective", function(){
  ...
  scope: {
    toPage: "&"
  },
  link: function(scope, elem, attrs) {
    // 写法1，直接报错
    scope.show = function(){
      scope.toPage(1)
    }

    // 写法2，不会报错，但是到控制器那边的toPage函数时，参数值是undefined
    scope.show = function(){
      var page = {p: 1}
      scope.toPage(page)
    }

    // 写法3，正常执行，控制器的toPage函数的参数获取到值
    scope.show = function(){
      var page = {page: 1}
      scope.toPage(page)
    }
  }
})
.controller("myController", ["$scope",function($scope){
  $scope.toPage = function(page){
    // page的值为你在指令中传递的json中的value值，这里page为1
    console.log(page)  // 注意此时已经不是json了
  }
}])

// html
<div my-directive to-page="toPage(page)"></div>
```
没错，我们在传递参数给控制器的函数时，需要使用json，并且json中的Key值必须和模板template中写的参数名一致。而到了控制器的函数时，参数则为对应的value值。坑爹啊！！！

<span id="10"></span>
###### transclude （布尔值）
该属性默认值为false，一旦定义则必须设置为true。
使用它，我们可以将视图元素嵌入到我们指定的地方进去。嵌入通常用来创建可复用的组件,典型的例子是模态对话框或导航栏。  
注意：为了将作用域传递进去,scope参数的值必须通过{}或true设置成隔离作用域。如果没有设 置scope参数,那么指令内部的作用域将被设置为传入模板的作用域。
```
<div sideboxtitle="Links">
  <ul>
   <li>First link</li>
   <li>Second link</li>
  </ul>
</div>

angular.module('myApp', [])
.directive('sidebox', function() {
  return {
    restrict: 'EA',
    scope: {
      title: '@'
    },
    transclude: true,
    template: '<div class="sidebox">\
      <div class="content">\
      <h2 class="header">{{ title }}</h2>\
      <span class="content" ng-transclude>\
      </span>\
      </div>\
      </div>'
  }; 
});

// 解析后，上述视图中的元素将会被嵌入到包含ng-transclude指令的元素里面
<div class="sidebox">
  <div class="content">
    <h2 class="header">{{ title }}</h2>
    <span class="content" ng-transclude>
      <div sideboxtitle="Links">
        <ul>
          <li>First link</li>
          <li>Second link</li>
        </ul>
     </div>
    </span>
  </div>
</div>
```

<span id="10"></span>
###### controller （字符串或函数）
若该属性的值是字符串，则会以该字符串值去应用中查找该控制器。
```
angular.module('myApp', [])
.directive('myDirective',  function() {
  restrict: 'A',
  controller: 'myController'
})

angular.module('myApp', [])
.controller("myController", function($scope){
  ...
})
```
当然也可以自定义控制器，即将该属性的值定义为匿名函数
```
angular.module('myApp', [])
.directive('myDirective',  function() {
  return {
    restrict: 'A',
    controller: function($scope, $element, $attrs){
      ...
    }
  }
})
```
注意到上面控制器中$scope, $element, $attrs，指令中的控制器除了可以注入一些服务外，还可以注入这几个比较特殊的服务，$scope就是当前作用域，$element当前指令对应的元素，$attrs当前元素的属性组成的对象。 
 
指令的控制器和link函数可以进行互换。控制器主要是用来提供可在指令间复用的行为,但link函数只能在当前内部指令中定义行为,且无法在指令间复用。  
具体controller如何做共享，可以使用require属性，定义了它之后便可以访问其他指令的控制器。

<span id="11"></span>
###### controllerAs （字符串）
用于设置控制器的别名

<span id="12"></span>
###### require （字符串或数组）
require会将控制器注入到其值所指定的指令中，并作为当前指令的link函数的第四个参数。  
若该值为字符串，则会去查找该值对应的指令，将查找到的指令的控制器注入到当前指令中。  
可以附加一些前缀:  
1.不加前缀，只会在自身元素上查找控制器，若查找不到，会抛出错误。
```
angular.module('myApp', [])
.directive('myDirective',  function() {
  return {
    restrict: 'A',
    require: "ngModel"
  }
})

<!-- 指令会在本地作用域查找ng-model -->
<div my-directive ng-model="object"></div>
```
2.?前缀，如果在当前指令中没有找到所需要的控制器,会将null作为传给link函数的第四个参数。
3.^前缀，指令会在上游的指令链中查找require参数所指定的控制器。
4.?^前缀，指令会在上游的指令链中查找require参数所指定的控制器。

<span id="13"></span>
###### compile （对象或函数）
angular应用在启动后会经历两个阶段，一个是编译，一个是链接。  
编译阶段会遍历整个HTML文档，编译各个指令和模板，一旦编译阶段完成，便会调用编译函数，编译函数的参数包含有访问指令声明所在的元素(tElemente)及该元素其他属性(tAttrs)的方法。  
如果设置了compile函数,说明我们希望在指令和实时数据被放到DOM中之前进行DOM操作,在这个函数中进行诸如添加和删除节点等DOM操作是安全的。
```
compile: function(tEle, tAttrs, transcludeFn) {
  var tplEl = angular.element('<div>' +
             '<h2></h2>' +
             '</div>');
  var h2 = tplEl.find('h2'); 
  h2.attr('type', tAttrs.type); 
  h2.attr('ng-model', tAttrs.ngModel); 
  h2.val("hello"); 
  tEle.replaceWith(tplEl);
  return function(scope, ele, attrs) { // 连接函数
     ..
  }; 
}
```
DOM事件监听器的注册:这个操作应该在link函数中完成。  
注意：compile和link选项是互斥的。如果同时设置了这两个选项,那么会把compile所返回的函数当作链接函数,而link选项本身则会被忽略。就如上述代码中返回的就是link函数。

<span id="14"></span>
###### link
用法如下
```
...
link: function(scope, element, attrs){
  //...
}
...
```
如果有require时，可以将会有第四个参数，代表控制器或者所依赖的指
令的控制器
```
link: function(scope, element, attrs, SomeController) {
  // 在这里操作DOM,可以访问required指定的控制器 
}
```

关于compile和link的详解，可以看[此文](http://www.jb51.net/article/58229.htm)，了解compile和link这两个阶段，将会有助于理解angular的启动的整个过程
