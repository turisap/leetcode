var i = 10;
var array = [];

while (i--) {
    (function (i) {
        var i = i;
        array.push(function() {
            return i + i;
        });
    })(i);
}    

console.log([
    array[0](), // logs 18 NOT 20, 18 because on the first 
    array[1](), // logs 16 iteration  the loop i is already 9
])
