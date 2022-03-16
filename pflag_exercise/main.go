package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

//函数名带Var说明将flag的值绑定到变量
//函数名带P说明支持短选项
func main() {

	var (
		host string
		port int
		err  error
	)
	//1.结尾的Var表示将参数的值，绑定到变量
	pflag.StringVar(&host, "host", "127.0.0.1", "service host addresss")

	//2.支持短选项
	pflag.IntVarP(&port, "port", "P", 8080, "service host port")

	//3.支持废弃某选项
	pflag.StringVar(&host, "Host", "127.0.0.1", "service host addresss")
	pflag.CommandLine.MarkDeprecated("Host", "please use --host instead")

	//指定flag但并未指定值
	pflag.Lookup("host").NoOptDefVal = "localhost"

	//所有flag定义，调用pflag.Parse()
	pflag.Parse()

	//获取设置的flag参数
	host, err = pflag.CommandLine.GetString("host")
	if err != nil {
		fmt.Println("flagSet get host failed")
		return
	}

	port, err = pflag.CommandLine.GetInt("port")
	if err != nil {
		fmt.Println("flagSet get port failed")
		return
	}
	fmt.Printf("flagSet get,host:%s,port:%d\n", host, port)

	//获取命令行参数后面参数
	fmt.Printf("argument number is: %v\n", pflag.NArg())
	fmt.Printf("argument list is: %v\n", pflag.Args())
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))

	//输出变量
	fmt.Printf("host is: %v\n", host)
	fmt.Printf("port is: %v\n", port)
}
