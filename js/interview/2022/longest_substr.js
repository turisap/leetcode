const longestSubstr = s => {
  const chars = [...s]
  let res = []

  for (let i = 0; i < chars.length; i++){
    let left = i - 1;
    let right = i + 1
    let exporeLeft  = true
    let exporeRight  = true
    let currMax = [chars[i]]

     while (exporeRight && exporeLeft){
       const leftChar = chars[left]
       const rightChar = chars[right]

       if(currMax.includes(leftChar) || !leftChar) {
	 exporeLeft = false
       } else if(leftChar) {
	 currMax.push(leftChar)
	 left--
       }

       if(currMax.includes(rightChar) || !rightChar){
	 exporeRight = false
       } else if(rightChar) {
	 currMax.push(rightChar)
	 right++
       }
    }

    res = currMax.length > res.length ? currMax : res
  }

  return res.length
}


longestSubstr('')
longestSubstr('bbb')
longestSubstr('abc')
longestSubstr('a bc')
longestSubstr('a  bc')
longestSubstr('a Bbc')
longestSubstr(' aBbcC ')
longestSubstr('0123')
longestSubstr('000 123')
longestSubstr('abcabcbb')
longestSubstr('dvdf')

