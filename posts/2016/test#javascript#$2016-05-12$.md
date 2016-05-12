1.作用在字符串对象上
 ```javascript
function log(str){
  console.log(str)
}
 ```
 f能取到a,b吗？原理是什么？
 ```javascript
var F = function(){};
Object.prototype.a = function(){};
Function.prototype.b = function(){};
var f = new F()
 ```
 
 2.作用域
 ```javascript
if (!("a" in window)) {
  var a = 1;
}
alert(a);
 ```
 ```javascript
var x = "global value";
var getValue = function(){
  alert(x); 
  var x = "local value";
  alert(x);
}
getValue();
 ```
 
 3.this
```javascirpt
var x = 1
function func(x) { 
  this.x = x; 
} 
func(5);
alert(x)
```
```javascript
var User = {  
  count: 1,

  getCount: function() {
    return this.count;
  }
};
console.log(User.getCount());
var func2 = User.getCount;  
console.log(func2());
```
```javascript
function a() {
  alert(this);
}
a.call(null)
```
