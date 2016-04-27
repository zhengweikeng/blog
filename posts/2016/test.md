1.作用在字符串对象上
 ```javascript
function log(str){
  console.log(str)
}
 ```
 
 2.作用域
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
function func(x) { 
  this.x = x; 
} 
func(5);


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
