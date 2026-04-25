package main


import (
	"os"
	"strconv"
	"fmt"
	"strings"
	"time"

)
// func getCmdline(pid int) string {
//     path := fmt.Sprintf("/proc/%d/cmdline", pid)

//     data, err := os.ReadFile(path)
//     if err != nil || len(data) == 0 {
//         return ""
//     }

//     cmd := strings.ReplaceAll(string(data), "\x00", " ")

//     return strings.TrimSpace(cmd)
// }

func countFDs(pid int)int{
	entries, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))
    if err != nil {
        return -1
    }
    return len(entries)
}

func parseMemory(pid int)int64{
    path := fmt.Sprintf("/proc/%d/statm", pid)

	data, err := os.ReadFile(path)

	if err != nil{
		return 0
	}

	fields := strings.Fields(string(data))
     
	if len(fields) < 2{
		return 0
	}

	pages ,err:= strconv.Atoi(fields[1])
     if err != nil {
        return 0
    }
	return int64(os.Getpagesize()* pages)/1024;
}


func parseStatus(pid int)(map[string]string, error){
path := fmt.Sprintf("/proc/%d/status", pid)

data, err := os.ReadFile(path)
		if err != nil{
			return nil, err
		}
fields := make(map[string]string)
lines := strings.Split(string(data), "\n")

for _, line := range lines{

	if !strings.Contains(line, ":"){
		continue
	}

	part := strings.SplitN(line, ":", 2)
    
    key := strings.TrimSpace(part[0])
	value := strings.TrimSpace(part[1])

    fields[key] = value
}
return fields, nil
}


func getPIDs()([]int, error){
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


func main() {
    for {
        fmt.Print("\033[H\033[2J") 

        pids, err := getPIDs()
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }

        fmt.Printf("Process Explorer — %s — %d processes\n\n",time.Now().Format("15:04:05"), len(pids))

        fmt.Printf("%-8s %-8s %-20s %-12s %-10s %-6s\n","PID", "PPID", "NAME", "STATE", "MEM(KB)", "FDs")
        fmt.Println(strings.Repeat("─", 70))

        for _, pid := range pids {
            status, err := parseStatus(pid)
            if err != nil {
                continue
            }
            ppid, _ := strconv.Atoi(strings.Fields(status["PPid"])[0])
            memKB := parseMemory(pid)
            fds := countFDs(pid)

            fmt.Printf("%-8d %-8d %-20s %-12s %-10d %-6d\n",pid, ppid, status["Name"], status["State"], memKB, fds)
        }

        time.Sleep(1 * time.Second)
    }
}