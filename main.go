package main

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
)

func art() {
	art := `
$$$$$$$\   $$$$$$\  $$$$$$$\                                                        
$$  __$$\ $$  __$$\ $$  __$$\                                                       
$$ |  $$ |\__/  $$ |$$ |  $$ |                                                      
$$$$$$$  | $$$$$$  |$$$$$$$  |                                                      
$$  ____/ $$  ____/ $$  ____/                                                       
$$ |      $$ |      $$ |                                                            
$$ |      $$$$$$$$\ $$ |                                                            
\__|      \________|\__|

in

   _____                      _      
  / ____|                    | |     
 | |     ___  _ __  ___  ___ | | ___ 
 | |    / _ \| '_ \/ __|/ _ \| |/ _ \
 | |___| (_) | | | \__ \ (_) | |  __/
  \_____\___/|_| |_|___/\___/|_|\___|


	`
	clearConsole()
	println(art)

}

func main() {
	art()
	killSwitch := false
	for {
		reader := bufio.NewReader(os.Stdin)
		println("1.Send a File\n2.Recieve a file\n3.Exit")
		input, _ := reader.ReadString('\n')
		if input[:len(input)-1] == "1" {
			SendFile()

		} else if input[:len(input)-1] == "2" {
			Recieve()
		} else {
			println("Ok Bye")
			killSwitch = true
		}
		if killSwitch != false {
			break
		}
	}

}

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows": // Windows
		cmd = exec.Command("cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
