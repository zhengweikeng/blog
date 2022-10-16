## 创建组件
旧的做法：
```jsx
var Divider = React.createClass({
  render: function() {
    return (
      <div className="divider">
        <h2>Hello world</h2>
      </div>
    )
  }
})
```
es6的做法
```javascript
// Divider.js
import {components} from 'react'
class Divider extends components {
  render() {
    return (
      <div className="divider">
        <h2>Hello world</h2>
      </div>
    ) 
  }
}
export default Divider
```

使用旧的创建方式使用组件
```
<Divider />
```

使用es6的创建方式使用组件
```jsx
// app.js
import Divider from './Divider'
<Divider />


// app2.js
import {render} from 'react-dom'
import Divider from './Divider'
const el = document.getElementById("content")
render(
  <Divider />,
  el
) 

// app3.js
import {components} from 'react'
import Divider from './Divider'
class App extends components {
  render() {
    return (
      <Divider/>
    ) 
  }
}
export default App
```

## 组件动态值
jsx的{...}接收一个表达式，可将其中的渲染为动态值
```jsx
class Divider extends components {
  render() {
    const myClass = 'Divider'
    const content = 'content'
    return (
      <div className={myClass}>
        <h2>{content}</h2>
        <h2>{this.dateToString(new Date())}</h2>
      </div>
    ) 
  }
  
  dateToString(d) {
    return [
      d.getFullYear(), d.getMonth() + 1, d.getDate()
    ].join('-')
  }
}
```

## 组件属性
通过this.props可以访问组件的所有属性
```jsx
class Divider extends components {
  render() {
    return (
      <div className={this.props.className}>
        <h2>{this.props.content}</h2>
      </div>
    ) 
  }
}
<Divider className="myClass" content="hello world">
```
记住，不应该去修改props
```jsx
var div = <Divider className="myClass" content="hello world">
div.myClass = 'newClass'  // don't do this
```
通过es6还可以使用Spread Attributes
```jsx
const props = {className: 'myClass', content: 'hello world'}
<Divider {...props}>
```
可以为组件设置属性验证器
```jsx
var SurveyTableRow = React.createClass({
  propTypes: {
    survey: React.PropTypes.shape({
      id: React.PropType.number.isRequired
    }).isRequired,
    onClick: React.PropTypes.func
  },
  ...
})
```
采用es6的方式
```jsx
import {PropTypes, components} from 'React'
class SurveyTableRow extends components {
  propTypes: {
    survey: React.PropTypes.shape({
      id: React.PropType.number.isRequired
    }).isRequired,
    onClick: React.PropTypes.func
  },
}
```

## 子节点
this.props.children可以获取组件的子节点
```jsx
class Divider extends components {
  render() {
    return (
      <div className={this.props.className}>
        {this.props.children}
      </div>
    ) 
  }
}
<Divider className="myClass">
  <h2>hello world</h2>
</Divider>
```

## State
this.props可以访问组件自己的属性，但是不应该去修改它。如果需要修改，应该使用state。
```jsx
class Divider extends components {
  getInitialState() {
    return {
      showOptions: false
    }
  }
  
  handleClick() {
    this.setState({showOptions: true})
  }
  
  render() {
    let options = null
    if(this.state.showOptions) {
      options = (
        <ul className="options">
          <li>Test</li>  
          <li>Test2</li>  
          <li>Test3</li>  
        </ul>
      )
    }
    
    return (
      <div className="dropdown" onClick={this.handleClick}>
        <label>choose it</label>. {options}
      </div>
    ) 
  }
}
```
只要render的返回值有变化，VDOM则会更新，真实DOM也会更新。  
不要直接修改this.state，而是通过this.setState来修改。
