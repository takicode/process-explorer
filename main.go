package main


import (
	"fmt"
	"os"

)

// func getPids()([]int, error){
//    entries, err:= os.ReadDir("/proc")
// 	if err != nil{
// 		return nil, err
// 	}

// 	for _, entry:= range entries{
// 			if entry.IsDir(){
// 				fmt.Print(entry.Name())
// 			}
// 	}
// }


func main(){
	// pids, err := getPids()
  // if err != nil{
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(pids)

	  entries, _:= os.ReadDir("/proc")
	// if err != nil{
	// 	return nil, err
	// }

	for _, entry:= range entries{
			if entry.IsDir(){
				pid, err := strconv.Atoi(entry.name())
				if err != nil{
					
				}
			}
	}
}