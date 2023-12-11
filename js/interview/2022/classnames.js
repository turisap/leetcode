const input = [
  'header',
  'menu-item',
  'menu-item',
  'menu-item',
  'article',
  'article',
  'article',
  'text-small',
  'menu-item',
  'text-small',
  'article-heading',
  'menu-item',
  'captions',
  'menu-item',
  'article-heading',
  'article-heading',
  'captions'
]

const sortClassnames = cl => {
  const clMap = {}
  const arrUnique  = []

  for(let className of cl){
    if(!clMap[className]){
      clMap[className] = 1
      arrUnique.push(className)
    } else {
      clMap[className]++
    }
  }

  return arrUnique.sort((a,b) => clMap[b] - clMap[a])
}

console.log(sortClassnames(input))
