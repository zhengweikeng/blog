> 学习react的时候，看了这篇[react-patterns](https://github.com/planningcenter/react-patterns)文档，颇有收获，所以做了以下笔记。删掉了原文最后一部分

React
=====

## 目录

1. Organization
  1. [Component Organization](#component-organization)
  1. [Formatting Props](#formatting-props)
1. Patterns
  1. [Computed Props](#computed-props)
  1. [Compound State](#compound-state)
  1. [prefer-ternary-to-sub-render](#prefer-ternary-to-sub-render)
  1. [View Components](#view-components)
  1. [Container Components](#container-components)
1. Anti-patterns
  1. [Compound Conditions](#compound-conditions)
  1. [Cached State in render](#cached-state-in-render)
  1. [Existence Checking](#existence-checking)
  1. [Setting State from Props](#setting-state-from-props)
1. Practices
  1. [Naming Handle Methods](#naming-handler-methods)
  1. [Naming Events](#naming-events)
  1. [Using PropTypes](#using-proptypes)
  1. [Using Entities](#using-entities)
1. Gotchas
  1. [Tables](#tables)
1. Libraries
  1. [classnames](#classnames)

---

## Component Organization
我们知道定义一个组件，方法就是`React.createClass()`

但是ES6之后，有了Class，即类的概念，React也提供了这种方式，我们也应该提倡使用这种方式来定义组件

* class definition(类定义)
  * constructor
    * event handlers(事件处理函数)
    * initial state(state设置默认值)
  * 'component' lifecycle events(组件生命周期)
  * getters(属性的getter方法，用于做计算属性)
  * render(渲染函数)
* defaultProps(组件类静态方法，props的默认值)
* proptypes(组件类静态方法，props的配置对象)

```javascript
Class Person extends React.Component {
  constructor (props) {
    super(props)

    this.state = { smiling: false }

    this.handleClick = () => {
      this.setState({smiling: !this.state.smiling})
    }
  }

  componentWillMount () {
    // 此处可以添加事件监听(Flux Store, WebSocket, document, etc...)
  }

  componentDidMount () {
    // React.getDOMNode()
  }

  componentWillUnmount () {
    // 移除事件监听(Flux Store, WebSocket, document, etc...)
  } 

  get smillingMessage () {
    return (this.state.smiling) ? 'is smiling' : ''
  }

  render () {
    return (
      <div onClick={this.handleClick}>
        {this.props.name} {this.smillingMessage}
      </div>
    )
  }
}

Person.defaultsProps = {
  name: 'Guest'
}

Person.propTypes = {
  name: React.PropTypes.string
}
```
通过这种方式构造出来的组件，结构清晰。

**[⬆ 回到顶部](#目录)**

## Formatting Props
格式化html中属性的书写

```html
// bad
<Person
 firstName="Michael" />

// good
<Person firstName="Michael" />

```

```html
// bad
<Person firstName="Michael" lastName="Chan" occupation="Designer" favoriteFood="Drunken Noodles" />

// good
<Person
 firstName="Michael"
 lastName="Chan"
 occupation="Designer"
 favoriteFood="Drunken Noodles" />
```

**[⬆ 回到顶部](#目录)**

## Computed Props
计算属性，即有些值是需要通过一些变量计算变化而来。

我们应该采用getter的方式来实现计算属性，将计算过程抽取出来

```javascript
// bad
firstAndLastName () {
  return `${this.props.firstName} ${this.props.lastName}`
}

// good
get fullName () {
  return `${this.props.firstName} ${this.props.lastName}`
}
```

**[⬆ 回到顶部](#目录)**

## Compound State
有时需要根据多个state的值来返回一个结果，这时也应该使用getter的方式，并且应该带有动词含义的前缀来命名getter的名字

```javascript
// bad
happyAndKnowIt () {
  return this.state.happy && this.state.knowsIt
}

// good
get isHappyAndKnowsIt () {
  return this.state.happy && this.state.knowsIt
}
```

**[⬆ 回到顶部](#目录)**

## Prefer Ternary to Sub-render
采用三元表达式

```javascript
// bad
renderSmilingStatement () {
  return <strong>{(this.state.isSmiling) ? " is smiling." : ""}</strong>;
},

render () {
  return <div>{this.props.name}{this.renderSmilingStatement()}</div>;
}
```

```javascript
render () {
  return (
    <div>
      {this.props.name}
      {(this.state.smiling)
        ? <span>is smiling</span>
        : null
      }
    </div>
  )
}
```

**[⬆ 回到顶部](#目录)**

## View Components
不要去创建那种既包含布局又包含域组件的一次性组件

```javascript
// bad
class PeopleWrappedInBSRow extends React.Component {
  render () {
    return (
      <div className="row">
        <People people={this.state.people} />
      </div>
    );
  }
}
```

```javascript
// good
class BSRow extends React.Component {
  render () {
    return <div className="row">{this.props.children}</div>
  }
}

class SomeView extends React.Component {
  render () {
    return (
      <BSRow>
        <People people={this.state.people} />
      </BSRow>
    )
  }
}
```

**[⬆ 回到顶部](#目录)**

## Container Components
> A container does data fetching and then renders its corresponding
> sub-component. That's it. &mdash; Jason Bonta

所谓的容器组件简单来说就是承担了数据获取的组件。这部分组件应该要将其从组件中抽取出来，单独成一个容器组件，负责数据的生成。

#### Bad

```javascript
// CommentList.js

class CommentList extends React.Component {
  getInitialState () {
    return { comments: [] };
  }

  componentDidMount () {
    $.ajax({
      url: "/my-comments.json",
      dataType: 'json',
      success: function(comments) {
        this.setState({comments: comments});
      }.bind(this)
    });
  }

  render () {
    return (
      <ul>
        {this.state.comments.map(({body, author}) => {
          return <li>{body}—{author}</li>;
        })}
      </ul>
    );
  }
}
```

#### Good

```javascript
// CommentList
class CommentList extends React.Component {
  render () {
    return (
      <ul>
        {this.props.comments.map(({body, author}) => {
          return <li>{body}-{author}</li>
        })}
      </ul>
    )
  }
}

class CommentListContainer extends React.Component {
  getInitialState () {
    return { comments: [] }
  }

  componentDidMount () {
    $.ajax({
      url: "/my-comments.json",
      dataType: 'json',
      success: function(comments) {
        this.setState({comments: comments});
      }.bind(this)
    })
  }

  render () {
    return <CommentList comments={this.state.comments} />
  }
}
```

[Read more](https://medium.com/@learnreact/container-components-c0e67432e005)  
[Watch more](https://www.youtube.com/watch?v=KYzlpRvWZ6c&t=1351)

**[⬆ 回到顶部](#目录)**

## Cached State in `render`
不要将状态保留在render函数中

```javascript
// bad
render () {
  let name = `Mrs. ${this.props.name}`;

  return <div>{name}</div>;
}

// good
render () {
  return <div>{`Mrs. ${this.props.name}`}</div>;
}
```

```javascript
get fancyName () {
  return `Mrs. ${this.props.name}`
}

render () {
  return <div>{this.fancyName}</div>
}
```

其实就是尽量的使用计算属性

**[⬆ 回到顶部](#目录)**

## Compound Conditions
不要在render中使用复杂的条件判断，应该抽取成getter函数

```javascript
// bad
render () {
  return <div>{if (this.state.happy && this.state.knowsIt) { return "Clapping hands" }}</div>
}
```

```javascript
// good
get isTotesHappy () {
  return this.state.happy && this.state.knowsIt
}

render () {
  return <div>{this.isTotesHappy} && "Clapping hands"</div>;
}
```

更好的方式，应该是使用容器组件去管理状态，将状态state作为props传递到子组件中去

**[⬆ 回到顶部](#目录)**

## Existence Checking
不要去检查属性props中变量的存在性，应该是去使用defaultProps的方式为其设置默认值

```javascript
// bad
render () {
  if (this.props.person) {
    return <div>{this.props.person.firstName}</div>;
  } else {
    return null;
  }
}
```

```javascript
// good
class MyComponent extends React.Component {
  render() {
    return <div>{this.props.person.firstName}</div>;
  }
}

MyComponent.defaultProps = {
  person: {
    firstName: 'Guest'
  }
}
```

**[⬆ 回到顶部](#目录)**

## Setting State from Props
如果你没有明确的目的，不要尝试去将props设置成state，这是反设计模式的，会导致state的来源不一

```javascript
// bad
getInitialState () {
  return {
    items: this.props.items
  };
}
```

```javascript
// good
getInitialState () {
  return {
    items: this.props.initialItems
  };
}
```

**[⬆ 回到顶部](#目录)**

## Naming Handler Methods
为表单的事件监听函数正确命名

命名规则应该是如下的：

- 以`handle`开头
- 以事件的类型结尾 (eg, `Click`, `Change`)
- 采用现在时语法

```javascript
// bad
punchABadger () { /*...*/ },

render () {
  return <div onClick={this.punchABadger} />;
}
```

```javascript
// good
handleClick () { /*...*/ },

render () {
  return <div onClick={this.handleClick} />;
}
```

如果需要区分多个同类型的事件，可以在handle和事件类型名称中间加一些额外的描述

例如`handleNameChange`和`handleAgeChange`

如果你的组件中有多个事件监听，你最好仔细思考一下，你的组件是不是写复杂了，需不需要拆分出来

**[⬆ 回到顶部](#目录)**

## Naming Events
有时我们也可以为组件提供事件

```javascript
class Owner extends React.Component {
  handleDelete () {
    // handle Ownee's onDelete event
  }

  render () {
    return <Ownee onDelete={this.handleDelete} />;
  }
}

class Ownee extends React.Component {
  render () {
    return <div onChange={this.props.onDelete} />;
  }
}

Ownee.propTypes = {
  onDelete: React.PropTypes.func.isRequired
};
```

**[⬆ 回到顶部](#目录)**

## Using PropTypes
定制属性props配置文件是一个好习惯，当我们提供了不恰当的属性时会有警告。

同时也可以为其他看你组件代码的人更快的了解你的组件需要的属性和属性定义

```javascript
MyValidatedComponent.propTypes = {
  name: React.PropTypes.string
};
```

**[⬆ 回到顶部](#目录)**

## Using Entities
对一些特殊的字符，应该使用`String.fromCharCode()`

```javascript
// bad
<div>PiCO · Mascot</div>

// nope
<div>PiCO &middot; Mascot</div>

// good
<div>{'PiCO ' + String.fromCharCode(183) + ' Mascot'}</div>

// better
<div>{`PiCO ${String.fromCharCode(183)} Mascot`}</div>
```

**[⬆ 回到顶部](#目录)**

## Tables
在react中，并不会去帮你在table中生成tbody，而是直接插入tr。为了避免隐含的问题，最好手动的去插入tbody

```javascript
// bad
render () {
  return (
    <table>
      <tr>...</tr>
    </table>
  );
}

// good
render () {
  return (
    <table>
      <tbody>
        <tr>...</tr>
      </tbody>
    </table>
  );
}
```

**[⬆ 回到顶部](#目录)**

## classnames
当你的组件中有较为复杂的class样式变化时，可以使用这个库[classNames](https://www.npmjs.com/package/classnames)

```javascript
// bad
get classes () {
  let classes = ['MyComponent'];

  if (this.state.active) {
    classes.push('MyComponent--active');
  }

  return classes.join(' ');
}

render () {
  return <div className={this.classes} />;
}
```

```javascript
// good
render () {
  let classes = {
    'MyComponent': true,
    'MyComponent--active': this.state.active
  };

  return <div className={classnames(classes)} />;
}
```