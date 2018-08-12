class P {
  constructor() {
    this.x = 1
    this.y = 3
  }
  print() {
    console.log(this.x, this.y)
  }
}

class C extends P {
  constructor() {
    super()
    this.x = 2
    // 和 this.z = 5 效果一样
    super.z = 5
  }

  say() {
    console.log(super.print()) // 2 3
    // 实际读取的是 P.prototype.z
    console.log(super.z) // undefined
    console.log(this.z) // 5
  }
}

const c = new C()
c.say()
