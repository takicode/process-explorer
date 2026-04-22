package main


import (
	"os"
	"strconv"
	"fmt"

)

func getPids()([]int, error){
   entries, err:= os.ReadDir("/proc")
	if err != nil{
		return nil, err
	}
    
	var pids []int
	for _, entry:= range entries{
			if entry.IsDir(){
				pid, err := strconv.Atoi(entry.Name())
				if err == nil{
					pids = append(pids, pid)
				}
			}
	}
	return pids, nil
}


func main(){
	pids, err := getPids()
    if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
    fmt.Printf("found %d processes\n\n", len(pids))

	for _, pid := range pids{
		fmt.Println(pid)
	}

}