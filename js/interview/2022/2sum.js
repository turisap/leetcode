const twoSumm = (nums, target) => {
  const map = {}
    
  for(let i = 0; i < nums.length; i++){
    map[nums[i]] = i
  }

  for(let j = 0; j < nums.length; j++){
    const curr = nums[j]
    const diff = target - curr
    
    if(map[diff] && map[diff] !== j){
      return [map[diff], j]
    }
  }
}


console.log(twoSumm([3,3], 6))
console.log(twoSumm([2, 4, 5], 9))
