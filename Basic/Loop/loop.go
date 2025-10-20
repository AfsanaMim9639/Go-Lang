//Loop--For(break continue)

package main

import "fmt"

func main() {
	numb := []int{1, 2, 3, 4, 5, 6}

	for index, value := range numb {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}


	for i:=1; i<=10; i++{
		if i%2==0{
			continue
		}
		fmt.Println("Odd Number:", i)
		if i==5{
			break
		}
	}


	//For loop acting like while loop



    i := 1              

    for i <= 5 {        
        fmt.Println(i)
        i++             
    }


}
