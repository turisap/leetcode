const sum4 =(nums1, nums2, nums3, nums4) => {
  const map = {}

  for(let a of nums1){
    for(let b of nums2){
      if(map[a + b]){
	map[a + b] += 1
      } else {
	map[a + b] = 1
      }
    }
  }

  let res = 0

  for (let a of nums3){
    for(let b of nums4){
      const summ = a + b;

      if(map[-summ]){
	res += map[-summ]
      }
    }
  }

  return res
}

const r = sum4(
[-1,-1],
[-1,1],
[-1,1],
[1,-1]
)

console.log(r)
