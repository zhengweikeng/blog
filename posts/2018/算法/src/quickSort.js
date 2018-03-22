function partition(arr, lo, hi) {
  const item = arr[lo]
  let i = lo
  let j = hi + 1

  while (true) {
    while(arr[++i] < item) if(i === hi) break
    while(arr[--j] > item) if(j === lo) break
    if (i >= j) break
    
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
  }

  const tmp = arr[lo]
  arr[lo] = arr[j]
  arr[j] = tmp
  return j
}

function sort(arr, lo, hi) {
  if (lo >= hi) return

  const j = partition(arr, lo, hi)
  sort(arr, lo, j)
  sort(arr, j + 1, hi)
}

const array = [4, 1, 3, 2, 16, 9, 10, 14, 8, 7]
console.log('before: ', array)
sort(array, 0, array.length - 1)
console.log('after: ', array)