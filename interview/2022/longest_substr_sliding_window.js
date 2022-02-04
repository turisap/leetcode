const longestSubstr = s => {
  let res = []

  for (let i = 0; i < s.length; i++){
    const map = {}
    let start = i;
    let end = i + 1;

    map[s[start]] = true

    while(end < s.length && !map[s[end]]) {
      map[s[end]] = true
      end++
    }

    res = Object.keys(map).length > res.length ? Object.keys(map): res
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
