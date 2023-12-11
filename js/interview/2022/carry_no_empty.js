function foo(value) {
	var acc = value;
	function addNext(next) {
		acc += next;
		return addNext;
	}
	addNext.toString = addNext.valueOf = function() {
		return acc;
	}
	return addNext;
}

console.log(foo(2)(3)(4))
