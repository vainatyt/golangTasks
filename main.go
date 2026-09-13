package main
import "fmt"

func twoSum(nums []int, target int) []int {
    contain := make(map[int]int, len(nums))
    for idx, num := range nums{
        diff := target - num;
        if i, ok := contain[diff]; ok{
            return []int{idx, i}
        }
        contain[num] = idx
    }
    return nil
}

func test(nums[]int, target int){
    fmt.Printf("Start test with nums %v for target %d\n", nums, target)
    result := twoSum(nums, target)
    if result != nil{
        fmt.Printf("ans %v\n", result)
    } else {
        fmt.Println("no ans")
    }
}

func main(){
	testNums1 := []int{2,7,11,15}
	target1 := 9
	testNums2 := []int{3,2,4}
	target2 := 6
	testNums3 := []int{3,3}
	target3 := 6
	fmt.Println("===== Start tests =====")
	test(testNums1, target1)
    test(testNums2, target2)
    test(testNums3, target3)
    fmt.Println("===== End tests =====")
}
