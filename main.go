package main
 
import "fmt"
 
// func main() {
//     numbers := []int{2, 3}
//     v := append(numbers, 10)
//     fmt.Println(numbers)
//     _ = v
// }

// func main() {
//     n1 := []int{2, 3}
//     n1 = append(n1, 5.41)
//     fmt.Println(n1)
// }

// func main() {
//     n1 := []int{2, 3}
//     n1 = append(n1, 5.41)
//     fmt.Println(n1)
// }

// func main() {
//     for i := 1; i <= 50; i++ {
// 		if i%7 == 0{
// 			fmt.Println(i)
// 		}  else if i > 25 {
// 			break
// 		}
//     }
// }

// func main() {
//     count := 0
//     for i := 1; i <= 50; i++ {
//         if i%7 != 0 {
//             if count == 3 {
//                 break
//             }
//             continue
//         }
//         fmt.Println(i)
//         count++
//     }
// }

// func main()  {
// 	nums := [...]int{30, -1, -6, 90, -6}
// 	fmt.Println(nums)
// }

// func main() {
// 	var now = time.Now()
//     for years := 2001; years <= now.Year(); years++ {
// 		fmt.Println(years)
//     }
// }

// func main() {
// 	nums := [...]int{30, -1, -6, 90, -6}
// 	positiveCount := 0
// 	for _, pn := range nums {
// 		if pn >= 0 {
// 			positiveCount++
// 		}
// 	}
// 	fmt.Println(positiveCount)
// }

// func main()  {
// 	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
// 	anums := nums[1:6]
// 	fmt.Println(anums)
// 	fmt.Println("Jumlahnya : ", len(anums))
// }

// func main() {
// 	n1 := []int{2, 3}
//     n2 := []int{4, 5}
//     _, _ = n1, n2

// 	n3 := append(n1,n2...)
// 	fmt.Println(n3)
// }

// func main() {
//     numbers := []int{2, 3}
//     fmt.Println(numbers[2])
// }

func main() {
    for i := 1; i <= 10; i++ {
        if i%3 != 0 {
            break 
        }
        fmt.Println(i)
    }
}

