const longestCommonPrefix = (strings) => {
  if(!strings.length) return ''
  
  // sort to speed up (the very first one will be the shortest
  // hence shorter while loop
  const strings = strings.sort()

  let res = strings[0]

  for(let str of strings){
    while(str.indexOf(res) !== 0 ){
      res = res.substr(0, res.length - 1)

      if(!res) return ''
    }
  }

  return res
}

console.log(longestCommonPrefix(["flower","flow","flight"]))
console.log(longestCommonPrefix(["flower"]))
console.log(longestCommonPrefix([]))

