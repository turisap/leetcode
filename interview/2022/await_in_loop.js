const getUser = () => new Promise(res => setTimeout(() => res('user'), 1500))

const badFn = async (num) => {
  const users = []

  console.log('bad loop start')

  for(let i = 0; i < num; i++){
    const user = await getUser()
    users.push(user)

    // logs every 1.5s
    console.log('bod loop iteration') 
  }

  console.log('bad loop end')
  console.log(users)
}

badFn(5)
