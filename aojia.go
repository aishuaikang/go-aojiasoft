package aojia

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var (
	aRegJDll     = syscall.NewLazyDLL("ARegJ.dll")
	_SetDllPathA = aRegJDll.NewProc("SetDllPathA")
	_SetDllPathW = aRegJDll.NewProc("SetDllPathW")
	_RegA        = aRegJDll.NewProc("RegA")
	_RegW        = aRegJDll.NewProc("RegW")
	_UnRegA      = aRegJDll.NewProc("UnRegA")
	_UnRegW      = aRegJDll.NewProc("UnRegW")
)

type AJsoft struct {
	aj       *ole.IDispatch
	IUnknown *ole.IUnknown
}

func NewAJsoft() *AJsoft {
	var com AJsoft
	var err error
	ole.CoInitialize(0)
	com.IUnknown, err = oleutil.CreateObject("AoJia.AoJiaD")
	if err != nil {
		panic(err)
	}
	com.aj, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(err)
	}
	return &com
}

// Release 释放
func (com *AJsoft) Release() {
	com.IUnknown.Release()
	com.aj.Release()
	ole.CoUninitialize()
}

/*
函数功能: 设置AoJia.dll或AoJia64.dll的完整路径,设置之后可以直接创建插件对象,不需要向系统注册表注册插件

函数原型: LONG WINAPI SetDllPathA(LPCSTR DllPath, LONG Type)

参数说明:
DllPath: AoJia.dll或AoJia64.dll的完整路径

Type: 备用参数,填0

返回值: 返回1表示成功,0表示失败
*/
func SetDllPathA(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathA.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}

/*
函数功能: 设置AoJia.dll或AoJia64.dll的完整路径,设置之后可以直接创建插件对象,不需要向系统注册表注册插件

函数原型: LONG WINAPI SetDllPathW(LPCWSTR DllPath, LONG Type)

参数说明:
DllPath: AoJia.dll或AoJia64.dll的完整路径

Type: 备用参数,填0

返回值: 返回1表示成功,0表示失败
*/
func SetDllPathW(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _SetDllPathW.Call(uintptr(unsafe.Pointer(_p0)), uintptr(mode))
	return ret != 0
}

/*
函数功能: 注册COM组件,也就是向系统注册表写入COM组件的信息

函数原型: LONG WINAPI RegA(LPCSTR DllPath)

参数说明:
DllPath: 要注册的COM组件的完整路径

返回值: 返回1表示成功,0表示失败
*/
func RegA(path string) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _RegA.Call(uintptr(unsafe.Pointer(_p0)))
	return ret != 0
}

/*
函数功能: 注册COM组件,也就是向系统注册表写入COM组件的信息

函数原型: LONG WINAPI RegW(LPCWSTR DllPath)

参数说明:
DllPath: 要注册的COM组件的完整路径

返回值: 返回1表示成功,0表示失败
*/
func RegW(path string) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _RegW.Call(uintptr(unsafe.Pointer(_p0)))
	return ret != 0
}

/*
函数功能: 反注册COM组件,也就是从系统注册表删除COM组件的信息

函数原型: LONG WINAPI UnRegA(LPCSTR DllPath)

参数说明:
DllPath: 要执行反注册的COM组件的完整路径

返回值: 返回1表示成功,0表示失败
*/
func UnRegA(path string) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _UnRegA.Call(uintptr(unsafe.Pointer(_p0)))
	return ret != 0
}

/*
函数功能: 反注册COM组件,也就是从系统注册表删除COM组件的信息

函数原型: LONG WINAPI UnRegW(LPCWSTR DllPath)

参数说明:
DllPath: 要执行反注册的COM组件的完整路径

返回值: 返回1表示成功,0表示失败
*/
func UnRegW(path string) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := _UnRegW.Call(uintptr(unsafe.Pointer(_p0)))
	return ret != 0
}
