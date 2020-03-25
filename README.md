# Go语言圣经中文版 课后习题

本工程是《The Go Programming Language》(《Go语言圣经中文版》）这本书的课后习题代码我交的作业，只是我个人的实现思路，欢迎探讨，持续更新中...

[《Go语言圣经中文版》在线阅读](https://docs.hacknode.org/gopl-zh/)

[《Go语言圣经中文版》Github](https://github.com/golang-china/gopl-zh)

## 目录

1. 入门
   1. 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_1/work1_1.go)
   2. 修改 echo 程序，使其打印每个参数的索引和值，每个一行。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_2/work1_2.go)
   3. 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。(1.6 节讲解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。)[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_3/work1_3.go)
   4. 修改 dup2 ，出现重复的行时打印文件名称。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_4/work1_4.go)
   5. 修改前面的Lissajous程序里的调色板，由黑色改为绿色。我们可以
用 color.RGBA{0xRR, 0xGG, 0xBB, 0xff} 来得到 #RRGGBB 这个色值，三个十六进制的字符串分 别代表红、绿、蓝像素。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_5/work1_5.go)
   6. 修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex
的第三个参数，看看显示结果吧。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_6/work1_6.go)
   7. 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用 这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区 (例子中的b)来存储。记得处理io.Copy返回结果中的错误。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_7/work1_7.go)
   8. 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该 前缀。你可能会用到strings.HasPrefix这个函数。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_8/work1_8.go)
   9. 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_9/work1_9.go)
   10. 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个 URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一 致，修改本节中的程序，将响应结果输出，以便于进行对比。
[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_10/work1_10.go)
   11. 在fatchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里 排名靠前的。如果一个网站没有回应，程序将采取怎样的行为?(Section8.9 描述了在这种 情况下的应对机制)。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_11/work1_11.go)
   12. 修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/? cycles=20 这个URL，这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字 可以调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_12/work1_12.go)
   
2. 程序结构
   1. 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝 对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_1/work2_1.go)
   2. 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的 话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对 应英尺和米，重量单位可以对应磅和公斤等。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_2/work2_2.go)
   3. 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。
(11.4节将展示如何系统地比较两个不同实现的性能。)[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_3/work2_3.go)
   4. 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和
查表算法的性能差异。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_4/work2_4.go)
   5. 表达式 x&(x-1) 用于将x的最低的一个非零的bit位清零。使用这个算法重写
PopCount函数，然后比较性能。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_5/work2_5.go)
3. 基础数据类型
   1. 如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素 (虽然许多SVG渲染器会妥善处理这类问题)。修改程序跳过无效的多边形。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_1/work3_1.go)
   2. 试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案?[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_2/work3_2.go)
   3. 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色 (#0000ff)。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_3/work3_3.go)
   4. 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返 回SVG数据给客户端。服务器必须设置Content-Type头部:
[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_4/work3_4.go)
   5. 实现一个彩色的Mandelbrot图像，使用image.NewRGBA创建图像，使用 color.RGBA或color.YCbCr生成颜色。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_5/work3_5.go)
   6. 升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每 个像素分层四个子像素，实现它。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_6/work3_6.go)
   7. 另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如z4 − 1 = 0。 每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_7/work3_7.go)
   8. 通过提高精度来生成更多级别的分形。使用四种不同精度类型的数字实现相同的 分形:complex64、complex128、big.Float和big.Rat。(后面两种类型在math/big包声明。 Float是有指定限精度的浮点数;Rat是无效精度的有理数。)它们间的性能和内存使用对比如 何?当渲染图可见时缩放的级别是多少?
[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_8/work3_8.go)
   9. 编写一个web服务器，用于给客户端生成分形的图像。运行客户端用过HTTP参数 参数指定x,y和zoom参数。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_9/work3_9.go)
   10. 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_10/work3_10.go)
   11. 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_11/work3_11.go)
   12. 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的 字符，但是对应不同的顺序。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_12/work3_12.go)
   13. 编写KB、MB的常量声明，然后扩展到YB。[code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_13/work3_13.go)
4. 复合数据类型