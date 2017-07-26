package main

import (
	"fmt"
	"github.com/goburrow/modbus"
	_ "github.com/influxdata/surgemq"
	"log"
	"os"
	"time"
)

/*"github.com/kardianos/service"
"github.com/yuin/gopher-lua"

"wsbox/helper"
"wsbox/module/app"
"wsbox/module/registry"
"wsbox/module/winservice"*/
//"path/filepath"
//"os"
/*"bufio"
  "fmt"
  "net/http"
  _ "net/http/pprof"

  "runtime"
  "strings")*/

//rsrc -manifest nac.manifest -ico tb.ico -o nac.syso

//var logger service.Logger
/*
func runLua() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("app", app.Loader)
	L.PreloadModule("service", winservice.Loader)
	L.PreloadModule("registry", registry.Loader)
	f := helper.AbsFile(helper.ExeName + ".lua")
	log.Println(f)
	if err := L.DoFile(f); err != nil {
		panic(err)
	}
}*/
func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println(err) //这里的err其实就是panic传入的内容，55
		}
	}()
	fmt.Print("start iot!!!")
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0xFF
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		defer handler.Close()
		client := modbus.NewClient(handler)
		results, err := client.ReadDiscreteInputs(15, 2)
		fmt.Print(err)
		results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
		results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})
		fmt.Print(results)
	}
	fmt.Print(err)
	//
	//runLua()
	//return
}
