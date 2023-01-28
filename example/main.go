package main

import (
	"errors"
	"log"

	aojia "github.com/aishuaikang/go-aojiasoft"

	"os"
)

func main() {
	aj, err := CreateAJObj()
	if err != nil {
		log.Println(err)
		return
	}
	defer aj.Release()
	log.Printf("AJ插件版本:%s", aj.VerS())

	ret := aj.RegD("7d3359b407cd61f71e1f3fcbf65a630dbe6f20cade28c96fcb1b197e3d9619f9", 0)
	switch ret {
	case 1:
		log.Println("收费注册成功")
	case 0:
		log.Println("失败,错误未知")
	case -1:
		log.Println("获取机器码失败")
	case -2:
		log.Println("无法连接网络")
	case -3:
		log.Println("注册码不正确")
	case -4:
		log.Println("账户余额不足")
	case -5:
		log.Println("从本地发起的访问网络的请求失败")
	case -6:
		log.Println("读取服务器返回的数据失败")
	case -7, -8, -9, -10, -11, -12, -102:
		log.Println("使用插件的方式不正确")
	case -13, -14, -15, -16, -17, -18:
		log.Println("从服务器获取数据失败")
	default:
		log.Println("失败 (未知错误)")
	}
	data := aj.GetScreenDataBmpBytes(10, 10, 100, 100, "")
	os.WriteFile("test.bmp", data, os.ModePerm)
}

func CreateAJObj() (*aojia.AJsoft, error) {
	// 获取当前工作目录
	dir, _ := os.Getwd()

	// 设置dm.dll路径,并进行注册

	if ret := aojia.SetDllPathW(dir+"\\AoJia64.dll", 0); !ret {
		return nil, errors.New("AoJia64.dll 注册失败")
	}

	return aojia.NewAJsoft(), nil
}
