const input1 = [1, 2, 2, 1]
const input2 = [2, 2, 1]

const intersect = (nums1, nums2) => {
  const map = {}
  const res = []

  for(let number of nums1){
    if(map[number]){
      map[number]++
    } else {
      map[number] = 1
    }
  }

  for(let number of nums2){
    if(map[number]) {
      res.push(number)
      map[number]--
    }
  } 

  return res
}

console.log(intersect(input1, input2))
