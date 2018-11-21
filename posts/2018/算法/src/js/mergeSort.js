class MergeSort {
  constructor (arr) {
    this.arr = arr
  }
  

  merge (left, mid, right) {
    let i = left
    let j = mid + 1
    const aux = []
    for (let k = left; k <= right; k++) {
      aux[k] = this.arr[k]
    }

    for (let k = left; k <= right; k++) {
      if (i > mid) this.arr[k] = aux[j++]
      else if (j > right) this.arr[k] = aux[i++]
      else if (aux[i] >= aux[j]) this.arr[k] = aux[j++]
      else this.arr[k] = aux[i++]
    }
  }

  sort (left, right) {
    if (left >= right) return
    
    const mid = left + Number.parseInt((right - left) / 2)
    this.sort(left, mid)
    this.sort(mid + 1, right)
    this.merge(left, mid, right)
  }
}


const array = [5,2,4,7,1,3,2,6]
const left = 0
const right = array.length - 1
const mergeSort = new MergeSort(array)
mergeSort.sort(0, right)
console.log(mergeSort.arr)
