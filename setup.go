package aojia

/*
函数功能: 获取当前对象的唯一ID值

函数原型: 整数 GetAoJiaID()

参数说明: 无

返回值: 整数,返回当前对象的唯一ID值

示例:

TracePrint "对象ID: " & AJ.GetAoJiaID()
*/
func (com *AJsoft) GetAoJiaID() int {
	ret, _ := com.aj.CallMethod("GetAoJiaID")
	return int(ret.Val)
}

/*
函数功能: 获取已经创建的对象个数

函数原型: 整数 GetAoJiaNum()

参数说明: 无

返回值: 整数,返回已经创建的对象个数

示例:
TracePrint "对象数量: " & AJ.GetAoJiaNum()
*/
func (com *AJsoft) GetAoJiaNum() int {
	ret, _ := com.aj.CallMethod("GetAoJiaNum")
	return int(ret.Val)
}

/*
函数功能: 获取当前对象的函数执行过程中的最后错误

函数原型: 整数 GetLastError()

参数说明: 无

返回值: 整数,返回一个表示错误的整数,每一个函数的每次执行都会设置这个值,各数值的说明如下:

0: 说明函数执行过程中无错误
-1: 传递的参数不正确
-2: 图片解密失败,可能没有设置密码或密码不正确
-3: 字库解密失败,可能没有设置密码或密码不正确
-6: 无法正确获取窗口或窗口客户区坐标,可能是因为窗口被最小化或窗口句柄已经无效
-7: 内核程序加载失败
如果GetLastError()的返回值是0但函数没有执行成功则是因为未知的原因导致函数失败

示例:
AJD = AJ.GetLastError() : TracePrint AJD
*/
func (com *AJsoft) GetLastError() int {
	ret, _ := com.aj.CallMethod("GetLastError")
	return int(int32(ret.Val))
}

/*
函数功能: 获取当前对象的全局路径

函数原型: 字符串 GetPath()

参数说明: 无

返回值: 字符串,返回当前对象的全局路径

示例:
AJD = AJ.GetPath() : TracePrint AJD
*/
func (com *AJsoft) GetPath() string {
	ret, _ := com.aj.CallMethod("GetPath")
	return ret.ToString()
}

/*
函数功能: 收费注册

函数原型: 整数 RegD(RegC, Type)

参数说明:
RegC: 字符串,注册码

Type: 整数,备用参数,填0

返回值: 整数,返回值说明如下:

1: 收费注册成功
0: 失败,错误未知
-1: 获取机器码失败
-2: 无法连接网络
-3: 注册码不正确
-4: 账户余额不足
-5: 从本地发起的访问网络的请求失败
-6: 读取服务器返回的数据失败
-7,-8,-9,-10,-11,-12,-102: 使用插件的方式不正确
-13,-14,-15,-16,-17,-18: 从服务器获取数据失败
示例:
AJD = AJ.RegD("0123456789abcdeffedcba98765432100123456789abcdeffedcba9876543210", 0) : TracePrint AJD

除了VerS这个函数,其它所有函数都需要函数RegD返回1才能调用,每个对象在创建以后都需要调用函数RegD一次,如果函数RegD没有返回1就调用其它函数,那么调用进程会崩溃.

如果返回值是-2,-5,-6,还可以再调用插件的GetLastError函数获取一个由操作系统返回的错误码

*/
func (com *AJsoft) RegD(RegC string, Type int) int {
	ret, _ := com.aj.CallMethod("RegD", RegC, Type)
	return int(int32(ret.Val))
}

/*
函数功能: 设置前台模式时图色,文字,鼠标函数的坐标的有效范围

函数原型: 整数 SetDesktopArea(Hwnd, Dx1, Dy1, Dx2, Dy2, Flag, Type)

参数说明:
Hwnd: 整数,一个有效的窗口句柄

Dx1: 整数,区域左上角的横坐标,Dx1应该小于Dx2并且在[0, 桌面横向分辨率 - 1]的范围内

Dy1: 整数,区域左上角的纵坐标,Dy1应该小于Dy2并且在[0, 桌面纵向分辨率 - 1]的范围内

Dx2: 整数,区域右下角的横坐标,Dx2应该大于Dx1并且在[0, 桌面横向分辨率 - 1]的范围内

Dy2: 整数,区域右下角的纵坐标,Dy2应该大于Dy1并且在[0, 桌面纵向分辨率 - 1]的范围内

Flag: 整数,决定当Type的值为1或2把坐标设置成相对于窗口时,如果程序在运行过程中窗口被最小化或被遮挡激不激活窗口,只能取下面列出的值之一:

0: 表示不激活窗口,这时如果窗口被最小化,当Type的值为1有效范围只是窗口的标题栏范围,当Type的值为2相关函数都会不能正确的获取到坐标而执行失败
1: 表示激活窗口并在桌面最前面显示
Type: 整数,只能取下面列出的值之一:

0: 表示将有效范围设置成整个桌面,即坐标都相对于桌面左上角(0, 0),并且坐标值的有效范围是(0, 0, 桌面横向分辨率 - 1, 桌面纵向分辨率 - 1),这时忽略参数Hwnd, Dx1, Dy1, Dx2, Dy2, Flag
1: 表示将有效范围设置成Hwnd标识的窗口的区域,即坐标都相对于窗口左上角,并且坐标值的有效范围是(0, 0, 窗口宽度 - 1, 窗口高度 - 1),这时忽略参数Dx1, Dy1, Dx2, Dy2
2: 表示将有效范围设置成Hwnd标识的窗口的客户区,即坐标都相对于窗口客户区左上角,并且坐标值的有效范围是(0, 0, 窗口客户区宽度 - 1, 窗口客户区高度 - 1),这时忽略参数Dx1, Dy1, Dx2, Dy2
3: 表示将有效范围设置成参数Dx1,Dy1,Dx2,Dy2标识的区域,即坐标都相对于区域左上角(Dx1, Dy1),并且坐标值的有效范围是(0, 0, Dx2 - Dx1 - 1, Dy2 - Dy1 - 1),这个区域不包含右下角,这时忽略参数Hwnd, Flag
返回值: 整数,返回1表示成功,如果当前对象已经开启了后台或者参数不正确返回0表示失败

示例:
AJD = AJ.SetDesktopArea(0, 0, 0, 0, 0, 0, 0) : TracePrint AJD
AJD = AJ.SetDesktopArea(Hwnd, 0, 0, 0, 0, 1, 1) : TracePrint AJD
AJD = AJ.SetDesktopArea(Hwnd, 0, 0, 0, 0, 0, 2) : TracePrint AJD
AJD = AJ.SetDesktopArea(0, 100, 100, 1000, 1000, 0, 3) : TracePrint AJD

创建对象后默认的有效范围是整个桌面,如果之前调用过此函数设置成了其它区域,可以在需要的时候再次调用此函数并将参数Type设为0设置回整个桌面.在设置以后可以不限次数的更改成其它区域

如果将有效范围设置成Hwnd标识的窗口区域或客户区,在程序运行后即使窗口被移动或改变大小插件也能自动追踪窗口的位置的大小

在鼠标函数上,函数KQHouTai的设置能覆盖SetDesktopArea的设置

在图色函数上,函数KQHouTai的设置能覆盖SetDesktopArea的设置,而函数SetPCData的设置又能覆盖KQHouTai的设置
*/
func (com *AJsoft) SetDesktopArea(Hwnd, Dx1, Dy1, Dx2, Dy2, Flag, Type int) int {
	ret, _ := com.aj.CallMethod("SetDesktopArea", Hwnd, Dx1, Dy1, Dx2, Dy2, Flag, Type)
	return int(ret.Val)
}

/*
函数功能: 设置加载字库时解密字库所需的密码

函数原型: 整数 SetDictPw(Pw)

参数说明:
Pw: 字符串,密码,必须和使用EncryptFile进行加密的密码一样

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetDictPw("651368763CBA987609876543CFE0F1DC") : TracePrint AJD

所有字库共用一个密码,这个密码可以和图片的密码相同,也可以不同.加载字库时会判断字库有没有加密,如果字库没有加密不会进行解密,可以只对一部分字库进行加密,但加密过的字库只有设置了正确的密码才能加载成功.如果字库都没加密过,不需要调用此函数.
*/
func (com *AJsoft) SetDictPw(Pw string) int {
	ret, _ := com.aj.CallMethod("SetDictPw", Pw)
	return int(ret.Val)
}

/*
函数功能: 打开或关闭当前对象的错误提示,默认是关闭,在程序还在开发时可以打开帮助调试程序

函数原型: 整数 SetErrorMsg(Msg)

参数说明:
Msg: 整数,只能是0或1中的某一个值,0表示关闭,1表示打开

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetErrorMsg(1)
AJ.SetErrorMsg 0
*/
func (com *AJsoft) SetErrorMsg(Msg int) int {
	ret, _ := com.aj.CallMethod("SetErrorMsg", Msg)
	return int(ret.Val)
}

/*
函数功能: 设置所有图色查找和文字识别时需要排除的区域

函数原型: 整数 SetExcludeArea(Type, AreaD)

参数说明:
Type: 整数,只能是下面列出的值之一:

0: 清空已经设置的所有排除区域,此时忽略参数AreaD
1: 添加要排除的区域,此时参数AreaD表示所有要添加的排除区域及其颜色,排除区域左上角的横坐标不能大于右下角横坐标,左上角的纵坐标不能大于右下角纵坐标,否则函数会执行失败
2: 删除已经设置的排除区域,此时参数AreaD表示所有要删除的排除区域及其颜色,函数只会删除相匹配的区域
3: 改变已经设置的排除区域的颜色,此时参数AreaD表示所有要改变的排除区域及改变后的颜色,函数只会改变相匹配的区域的颜色
AreaD: 字符串,一个或多个要排除的区域及其颜色,格式为"x1,y1,x2,y2,BBGGRR|……",排除区域包含左上角和右下角,排除区域的宽度和高度分别是(x2 - x1 + 1)与(y2 - y1 + 1).

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetExcludeArea(1, "100,100,200,200,ED1C24|300,300,400,400,ED1C24|500,500,600,600,ED1C24") : TracePrint AJD
AJD = AJ.SetExcludeArea(2, "300,300,400,400,ED1C24|500,500,600,600,ED1C24") : TracePrint AJD
AJD = AJ.SetExcludeArea(3, "100,100,200,200,FF0000") : TracePrint AJD
AJD = AJ.SetExcludeArea(0, "") : TracePrint AJD

如果设置了排除区域,那么函数在截图之后会多执行一段把排除区域写成指定颜色的代码,并且在排除区域能找到这种颜色和这种颜色的纯色图片.
*/
func (com *AJsoft) SetExcludeArea(Type int, AreaD string) int {
	ret, _ := com.aj.CallMethod("SetExcludeArea", Type, AreaD)
	return int(ret.Val)
}

/*
函数功能: 设置当前对象是否开启使用全局字库

函数原型: 整数 SetGlobalDict(GD)

参数说明:
GD: 整数,只能是0或1中的某一个值,0表示关闭,1表示打开

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetGlobalDict(1) : TracePrint AJD
AJ.SetGlobalDict 0

全局字库在所有对象都可以访问的内存区域,这片内存区域是多线程安全的,任何时候都最多只有一个线程在写这片内存区域.每个对象都有一个只属于自己的变量描述是否开启使用全局字库,如果当前对象开启了全局字库,那所有和字库相关的操作都是对全局字库的操作.可以任意的开启和关闭以在使用当前对象自己的字库还是全局字库之间来回切换.开启后任何对象在任何时候都可以操作全局字库,插件从调用进程卸载后这片内存区域会自动释放.
*/
func (com *AJsoft) SetGlobalDict(GD int) int {
	ret, _ := com.aj.CallMethod("SetGlobalDict", GD)
	return int(ret.Val)
}

/*
函数功能: 设置当前对象的全局路径,设置以后,所有函数加载文件时的相对路径都相对于此路径,也可以继续使用绝对路径

函数原型: 整数 SetPath(Path)

参数说明:
Path: 字符串,要加载的文件的完整路径中不包含文件名的部分.也可以设置成相对当前进程路径的相对路径:

"\": 表示设置成调用插件函数的当前进程的可执行文件所在的路径,比如当前进程的可执行文件的路径是: "D:\按键精灵商业版\按键精灵2014\按键精灵2014.exe", 则全局路径设置成: "D:\按键精灵商业版\按键精灵2014"
"\Pic", "\Pic\1", "\Pic\1\……": 表示设置成相对当前进程路径更深的路径,比如当前进程的可执行文件的路径是: "D:\按键精灵商业版\按键精灵2014\按键精灵2014.exe", 则全局路径分别设置成: "D:\按键精灵商业版\按键精灵2014\Pic", "D:\按键精灵商业版\按键精灵2014\Pic\1", "D:\按键精灵商业版\按键精灵2014\Pic\1\……"

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetPath("C:\Pic")
AJD = AJ.SetPath("\")
AJD = AJ.SetPath("\Pic")
*/
func (com *AJsoft) SetPath(Path string) int {
	ret, _ := com.aj.CallMethod("SetPath", Path)
	return int(ret.Val)
}

/*
函数功能: 设置所有图色和文字相关函数要在其中查找和识别的图色数据的获取方式

函数原型: 整数 SetPCData(Type, PicName)

参数说明:
Type: 整数,指定图色数据的获取方式,只能是下面列出的值之一:

0: 表示从桌面或开启的后台窗口获取要在其中查找和识别的图色数据,在此情况下,如果没有调用函数KQHouTai设置只从某个窗口获取图色数据,那么就是从桌面获取图色数据
1: 表示从参数PicName指定的图片获取要在其中查找和识别的图色数据,设置后所有的图色和文字识别函数都是在这张图片中进行查找和识别,这时可以把图片看成是一个查找区域,这个查找区域左上角的坐标是(0, 0),右下角坐标是(图片宽度 - 1, 图片高度 - 1),调用函数指定区域(x1, y1, x2, y2)时,首先x1的值不能大于x2,y1的值不能大于y2,否则函数将不进行查找和识别而直接返回,然后x1的值如果小于0或大于等于图片的宽度会被较正为0,y1的值如果小于0或大于等于图片的高度会被较正为0,x2的值如果小于0或大于等于图片的宽度会被较正为(图片宽度 - 1),y2的值如果小于0或大于等于图片的高度会被较正为(图片高度 - 1).当获取方式是从桌面或开启的后台窗口时也一样,如果坐标的值不在有效范围内会被较正为有效范围的边界值
2: 和值设为1时完全一样,只是当Type设为2时,参数PicName是图片在内存中的地址
PicName: 字符串,这个参数的含义由Type决定,具体说明如下:

当Type的值设为0时这个参数被忽略
当Type的值设为1时表示包含路径的图片名称,如果设置了全局路径,则可以只给出图片名称,图片必须是24位位图格式,只支持一张图片.可以多次设置,每次设置不一样的图片,但只有最后一次设置的图片生效.设置后,不管有没有关闭图片缓存机制,在调用任何图色或文字识别函数时都会将设置的图片加入缓存,对象释放时会自动释放缓存的所有图片,如要提前释放,可以调用FreePic函数
当Type的值设为2时表示图片在内存中的地址,这个地址的数据要符合24位位图格式,即包含正确的位图文件头和位图信息头以及颜色数据,只能给出一个地址,可以多次设置,每次设置不一样的地址,但只有最后一次设置的地址生效,这个函数不会将这个地址的图片数据加入缓存

返回值: 整数,返回1表示成功,0表示失败

示例:

    AJD = AJ.SetPCData(1, "1.bmp") : TracePrint AJD
    AJ.SetPCData 0, ""
    AJD = AJ.SetPCData(2, "1988454252656")

    LONG Size; LLAJD = AJ.GetScreenDataBmp(0, 0, 1919, 1079, L"", Size);
    if (LLAJD != 0) {
      AJD = AJ.LoadPicM(LLAJD, L"Screen.bmp");
      S = AJ.AddPicAddr(L"", LLAJD);
      AJD = AJ.SetPCData(2, S); cout << "AJD: " << AJD << endl;
      LLAJD = AJ.GetScreenDataBmp(22, 722, 122, 922, L"", Size);
      if (LLAJD != 0) {
        AJD = AJ.LoadPicM(LLAJD, L"LLAJD.bmp"); cout << "AJD: " << AJD << endl;
        AJD = AJ.FindPic(0, 0, 1919, 1079, L"LLAJD.bmp", L"000000", 1.0, 0, 0, P, x, y);  SAS = P; cout << AJD << " , " << SAS << " -> " << x << " , " << y << endl;
      }
    }
  创建对象后默认的是从桌面获取图色数据,之后如果调用了函数KQHouTai设置只从某个窗口获取图色数据,那么就会只从这个窗口获取图色数据,之后如果再调用了函数SetPCData将图色数据的获取方式定向到了某张图片,那么就只会从这张图片获取图色数据,也就是在图色函数上SetPCData的设置能覆盖KQHouTai的设置.
*/
func (com *AJsoft) SetPCData(Type int, PicName string) int {
	ret, _ := com.aj.CallMethod("SetPCData", Type, PicName)
	return int(ret.Val)
}

/*
函数功能: 打开或关闭当前对象的图片缓存机制,默认是打开

函数原型: 整数 SetPicCache(PicD)

参数说明:
PicD: 整数,只能是0或1中的某一个值,0表示关闭,1表示打开

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetPicCache(0) : TracePrint AJD
AJ.SetPicCache 0

除非内存不够用,否则不要关闭.图片缓存的期限是被加载到缓存开始,直到对象释放.创建插件对象的程序运行结束都会自动释放对象,这时插件会自动释放对象的所有资源
*/
func (com *AJsoft) SetPicCache(PicD int) int {
	ret, _ := com.aj.CallMethod("SetPicCache", PicD)
	return int(ret.Val)
}

/*
函数功能: 设置加载图片时解密图片所需的密码

函数原型: 整数 SetPicPw(Pw)

参数说明:
Pw: 字符串,密码,必须和使用EncryptFile进行加密的密码一样

返回值: 整数,返回1表示成功,0表示失败

示例:
AJD = AJ.SetPicPw("651368763CBA987609876543CFE0F1DC") : TracePrint AJD

所有图片共用一个密码.加载图片时会判断图片有没有加密,如果图片没有加密不会进行解密,可以只对一部分图片进行加密,但加密过的图片只有设置了正确的密码才能加载成功.如果图片都没加密过,不需要调用此函数
*/
func (com *AJsoft) SetPicPw(Pw string) int {
	ret, _ := com.aj.CallMethod("SetPicPw", Pw)
	return int(ret.Val)
}

/*
函数功能: 设置当前对象的线程池中可用的线程数.每个对象都有一个线程池,线程池中默认的线程数为1

函数原型: 整数 SetThread(TN)

参数说明:
TN: 整数,要设置的线程池中可用的线程数,用ThNum表示处理器支持真正并行的最大线程数减1这个数值,则TN的取值范围是:[0, ThNum], 值的说明如下:

0: 将线程池中可用的线程数设为最大值,也就是ThNum这个数值,比如4核4线程会设为3, 8核16线程会设为15, 16核32线程会设为31
[1, ThNum]: 将线程池中可用的线程数设为大于等于1且小于等于ThNum的某一个数值.将参数设为1可以恢复插件默认的设置,即线程池中只有一个线程
返回值: 整数,返回1表示成功,如果参数不正确或处理器只支持1个线程则返回0表示失败.

示例:
AJD = AJ.SetThread(0) : TracePrint AJD
AJD = AJ.SetThread(1) : TracePrint AJD
AJD = AJ.SetThread(2) : TracePrint AJD
AJ.SetThread 3


插件中有一些函数可以启动线程池中的线程参与并发的运行,在这些函数没有运行时线程池中的线程不会运行,只有在调用这些函数时线程池中的线程才会启动运行,在函数运行完后线程会回到线程池中,不会占用CPU.可以调用函数GetCPU获取处理器支持真正并行的最大线程数.此函数可以在任何时候调用,可以调用多次,每次调用都可以传递不同的数值增加或减少线程池中可用的线程数
*/
func (com *AJsoft) SetThread(TN int) int {
	ret, _ := com.aj.CallMethod("SetThread", TN)
	return int(ret.Val)
}

/*
函数功能: 获取当前插件的版本号

函数原型: 字符串 VerS()

参数说明: 无

返回值: 字符串,返回当前插件的版本号

示例:
AJD = AJ.VerS() : TracePrint AJD

这个函数不需要进行收费注册就可以调用,可以用来测试对象是否创建成功

版本号就是插件的编译日期
*/
func (com *AJsoft) VerS() string {
	ver, _ := com.aj.CallMethod("VerS")
	return ver.ToString()
}
