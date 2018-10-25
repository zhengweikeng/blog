// 递归方式实习二叉查找树
function createTree(key, value, left, right) {
  return {key, value, left, right, size: 1}
}

function search(tree, key) {
  if (tree == null) {
    return null
  }

  if (tree.key < key) {
    return search(tree.right, key) 
  } else if(tree.key > key) {
    return search(tree.left, key) 
  } else return tree.value
}

function put(tree, key, value) {
  if (tree == null) {
    return createTree(key, value)
  }

  if (tree.key > key) {
    tree.left = put(tree.left, key, value)
  } else if (tree.key < key) {
    tree.right = put(tree.right, key, value)
  } else tree.value = value

  tree.size = size(tree.left) + size(tree.right) + 1

  return tree
}

function size(tree) {
  if (tree == null) {
    return 0
  }
  return tree.size
}

function min(node) {
  if (node.left == null) return node

  return min(node.left)
}

function max(node) {
  if (node.right == null) return node

  return max(node.right)
}

function floor(node, key) {
  if (node == null) {
    return null
  }

  if (key === node.key) {
    return node
  } else if (key < node.key) {
    return floor(node.left, key)
  } else {
    const n = floor(node.right, key)
    if (n != null) {
      return n
    } else return node
  }
}

function ceiling(node, key) {
  if (node == null) {
    return null
  }
console.log(node.key);
  if (key === node.key) {
    return node
  } else if (key < node.key) {
    const n = ceiling(node.left, key)
    if (n != null) {
      return n
    } else return node
  } else {
    return ceiling(node.right, key)
  }
}

function deleteMin(node) {
  if (node.left == null) {
    return node.right
  }

  node.left = deleteMin(node.left)
  node.size = size(node.left) + size(node.right) + 1
  return node
}

function deleteMax(node) {
  if (node.right == null) {
    return node.left
  }

  node.right = deleteMax(node.right)
  node.size = size(node.left) + size(node.right) + 1
  return node
}

function deleteKey(node, key) {
  if (key < node.key) {
    node.left = deleteKey(node.left, key)
  } else if (key > node.key) {
    node.right = deleteKey(node.right, key)
  } else {
    if (node.right == null) return node.left
    if (node.left == null) return node.right
    const t = node
    node = min(node.right)
    node.right = deleteMin(t.right)
    node.left = t.left
  }

  node.size = size(node.left) + size(node.right) + 1
  return node
}

const tree = createTree(10, 'a')
console.log(search(tree, 10));
put(tree, 5, 'b')
put(tree, 7, 'c')
put(tree, 11, 'd')
put(tree, 1, 'e')
put(tree, 2, 'f')
put(tree, 19, 'g')
// console.log(tree, size(tree));
// console.log(min(tree));
// console.log(floor(tree, 6));
// console.log(ceiling(tree, 6));
// console.log(deleteMin(tree));
// console.log(deleteMax(tree));