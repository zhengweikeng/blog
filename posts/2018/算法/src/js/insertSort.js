function insertSort(array) {
  for (let i = 1; i < array.length; i++) {
    const val = array[i]
    let j = i - 1

    while (val < array[j]) {
      array[j + 1] = array[j]
      if (j-- === 0) {
        break
      }
    }

    array[j + 1] = val
    console.log(array)
  }


  return array
}

const sortedArray = insertSort([5, 2, 4, 6, 1, 3])
