function unique(arr) {
	return arr.filter((item, index, self) => (self.indexOf(item) === index));
}

const u = unique([1, 2, 3, 4, 5, 5, 6, 2, 3])
// filters out only unique items one line
