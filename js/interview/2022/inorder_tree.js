var inorderTraversal = function(root) {
    if(!root) return []
     
    let current = root
    const stack = [current]
    const res = []

    
    while(stack.length) {
        
        if(current.left){
            stack.push(current.left)
            current = current.left
        } else {
            const node = stack.pop()
            res.push(node.val)

            if(node.right){
                current = node.right
                stack.push(node.right)
            }
            
        }
    }
    
    return res
};

const tree = {
    val: 1,
    left: {
        val: 2,
        left:{
            val: 4
        },
        right:{
            val: 5
        }
    },
    right: {
        val: 3,
        left: {
            val:6
        }
    }
}

console.log(inorderTraversal(tree))