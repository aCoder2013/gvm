package main

import "fmt"
import (
	"gvm/cp"
	"strings"
	"gvm/classfile"
)

func main() {
	cmd := ParseCmd()
	if cmd.versionFlag {
		fmt.Println("openjdk version \"1.8.0_112-release\"")
		fmt.Println("OpenJDK Runtime Environment (build 1.8.0_112-release-b403)")
		fmt.Print("OpenJDK 64-Bit Server VM (build 25.112-b403, mixed mode)")
	} else if cmd.helpFlag || cmd.class == "" {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	clzPath := cp.Parse(cmd.cpOption)
	fmt.Printf("classpath : %s  class :%s args:%v\n", clzPath, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := clzPath.ReadClass(className)
	if err != nil {
		panic(err)
	}
	classfile.Parse(cmd.class, classData)
}
