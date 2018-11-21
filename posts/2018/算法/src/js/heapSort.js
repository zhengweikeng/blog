function maxHeapify(arr, i, len = arr.length) {
  arr.unshift(arr[i - 1])

  let right = i * 2 + 1
  let max = i

  while (right < len + 1) {
    if (arr[right] > arr[0]) {
      max = right
    }
    if (arr[right - 1] > arr[max]) {
      max = right - 1
    }
    if (max === i) {
      break
    }

    arr[i] = arr[max]
    i = max
    right = i * 2 + 1
  }

  if (right === len + 1 && arr[0] < arr[right - 1]) {
    arr[i] = arr[right - 1]
    i = right - 1
  }

  arr[i] = arr[0]

  arr.shift()
}

function buildMaxHeap(arr) {
  for (let i = arr.length / 2; i >= 1; i--) {
    maxHeapify(arr, i)
  }
}

function sortHeap(arr) {
  let len = arr.length
  for (let i = len - 1; i >= 0; i--) {
    const tmp = arr[i]
    arr[i] = arr[0]
    arr[0] = tmp
    len -= 1
    
    maxHeapify(arr, 1, len)
  }
}

const array = [10, 14, 16, 8, 7, 9, 3, 2, 4, 1]
console.log('before: ', array)
maxHeapify(array, 1)
console.log('after: ', array)

const array2 = [4, 1, 3, 2, 16, 9, 10, 14, 8, 7]
console.log('before: ', array2)
buildMaxHeap(array2)
console.log('after: ', array2)
sortHeap(array2)
console.log('sort: ', array2)
