function BubbleSort(arr, length) {
  for (let i = length - 1; i > 0; i--) {
    for (let j = 0; j < i; j++) {
      if (arr[j] > arr[j + 1]) {
        let temp = arr[j];
        arr[j] = arr[j + 1];
        arr[j + 1] = temp;
      }
    }
  }
}

const arr = [1, 5, 9, 3, 2, 8, 7];
BubbleSort(arr, arr.length);
console.log(arr);
