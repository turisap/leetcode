package main

func main() {

}

func twoSum(nums []int, target int) []int {
	var diffs = map[int]int{}

	for i := 0; i < len(nums); i++ {
		var diff = target - nums[i]

		if val, ok := diffs[diff]; ok {
			_ = val
			return []int{val, i}
		}

		diffs[nums[i]] = i
	}

	return []int{}
}

/*
const hashMap = {};

  for (let i = 0; i < input.length; i++) {
    const diff = target - input[i];

    if (typeof hashMap[diff] !== "undefined") {
      console.log(hashMap[diff], i);
      return [hashMap[diff], i];
    }

    hashMap[input[i]] = i;
  }

  return null;
*/
