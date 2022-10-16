// 非递归方式实现二叉查找树

function createTree(key, value, left, right) {
  return {key, value, left, right, size: 1}
}

function search(tree, key) {
  let node = tree
  while(node != null) {
    if (node.key > key) {
      node = node.left
    } else if (node.key < key) {
      node = node.right
    } else return node.value
  }

  return null
}

function put(tree, key, value) {
  let currNode = tree
  let parentNode = null
  const node = createTree(key, value)

  while(true) {
    parentNode = currNode
    if (parentNode.key > key) {
      currNode = currNode.left
      if (currNode == null) {
        parentNode.left = node
        break
      }
    } else if (parentNode.key < key) {
      currNode = currNode.right
      if (currNode == null) {
        parentNode.right = node
        break
      }
    } else {
      parentNode.value = value
      break
    }
  }
  
  return tree
}

function getHeight(node) {
  if (node == null) {
    return 0
  }

  return Math.max(getHeight(node.left), getHeight(node.right)) + 1
}

const tree = createTree(10, 'a')
put(tree, 5, 'b')
put(tree, 6, 'c')
put(tree, 11, 'd')
put(tree, 1, 'e')
put(tree, 2, 'f')
put(tree, 19, 'g')
console.log(tree);
console.log(getHeight(tree));
