package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Get path of cmd in $PATH or Relative Path
func testLookUp() {
	fmt.Println("Get path of `ls`")
	path1, _ := exec.LookPath("ls")
	fmt.Println("->	result: ", path1)

	fmt.Println("Get path of non-exsit cmd `nocmd`")
	_, err := exec.LookPath("nocmd")
	if err != nil {
		fmt.Println("->	result error: ", err)
	}

	fmt.Println("Get path of script `test.sh` in current dir")
	path2, err2 := exec.LookPath("./test.sh")
	if err2 != nil {
		fmt.Println("->	result error: ", err2)
	} else {
		fmt.Println("-> result: ", path2)
	}
}

// Exec command
func testCommand() {
	fmt.Println("Begin exec ls with CombinedOutput")
	cmd1 := exec.Command("ls")
	out1, err1 := cmd1.CombinedOutput() // 运行命令，并返回标准输出和标准错误
	if err1 != nil {
		fmt.Println("->	result error:", err1)
	} else {
		fmt.Println("->	result: ")
		fmt.Println(string(out1))
	}

	fmt.Println("Begin exec ls with Output")
	cmd2 := exec.Command("ls")
	out2, err2 := cmd2.Output() // 运行命令并返回其标准输出
	if err2 != nil {
		fmt.Println("->	result error:", err2)
	} else {
		fmt.Println("-> result: ")
		fmt.Println(string(out2))
	}

	fmt.Println("Begin exec ls with Run")
	cmd3 := exec.Command("ls") //
	// Run没有标准输出, 通过这种方式展示命令的输出
	cmd3.Stdout = os.Stdout
	// 开始指定命令并且等待他执行结束，如果命令能够成功执行完毕，则返回nil，否则的话边会产生错误
	err3 := cmd3.Run()
	if err3 != nil {
		fmt.Println("->	result error:", err3)
	}

	fmt.Println("Begin exec ls /noexsit with Run")
	cmd4 := exec.Command("ls /noexsit")
	cmd4.Stdout = os.Stdout // 命令执行失败, 所以不会有输出
	err4 := cmd4.Run()
	if err4 != nil {
		fmt.Println("-> result error:", err4)
	}

	fmt.Println("Begin exec ping with Run")
	cmd5 := exec.Command("ping", "-c", "2", "-i", "5", "8.8.8.8")
	fmt.Println("->	start_time:", time.Now())
	err5 := cmd5.Run()
	if err5 != nil {
		fmt.Println("-> result error:", err5)
	}
	fmt.Println("->	end_time:", time.Now()) // 会等到执行结束再返回

	fmt.Println("Begin exec ping with Start")
	cmd6 := exec.Command("ping", "-c", "2", "-i", "5", "8.8.8.8")
	fmt.Println("->	start_time:", time.Now())
	err6 := cmd6.Start()
	if err6 != nil {
		fmt.Println("->	result error:", err6)
	}
	fmt.Println("->	end_time:", time.Now()) // 会立马返回

	fmt.Println("Begin exec ping with Start and Wait")
	cmd7 := exec.Command("ping", "-c", "2", "-i", "5", "8.8.8.8")
	fmt.Println("->	start_time:", time.Now())
	err7 := cmd7.Start()
	if err7 != nil {
		fmt.Println("->	result error:", err6)
	}
	cmd7.Wait()
	fmt.Println("->	end_time:", time.Now())
}

func main() {
	testLookUp()
	testCommand()
}
