package main

import (
	"fmt"
	//"gopkg.in/urfave/cli.v1"

	"os"
	"path/filepath"
	"strings"

	_ "github.com/blocktree/OpenWallet/test/environment"
	"github.com/blocktree/OpenWallet/test/tech"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func main() {
	/*app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
	  fmt.Println("boom! I say!")
	  return nil
	}
	err := app.Run(os.Args)
	if err != nil {
	  fmt.Println(err)
	}*/

	//fmt.Println("change dir err:",err)
	//tech.TestNewWallet("peter2", "12345678")
	tech.TestBatchCreateAddr()
	//tech.TestBitInt()

	fmt.Println("done ... ")
}