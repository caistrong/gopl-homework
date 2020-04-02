# Go语言圣经中文版 课后习题

本工程是《The Go Programming Language》(《Go语言圣经中文版》）这本书的课后习题代码我交的作业，只是我个人的实现思路，欢迎探讨，持续更新中...

[《Go语言圣经中文版》在线阅读](https://docs.hacknode.org/gopl-zh/)

[《Go语言圣经中文版》Github](https://github.com/golang-china/gopl-zh)

## 目录

- 1.入门
   - 1.修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_1/work1_1.go)
   - 2.修改 echo 程序，使其打印每个参数的索引和值，每个一行。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_2/work1_2.go)
   - 3.做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。(1.6 节讲解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测)。[-> code](https://github.com/caistrong/- gopl-homework/blob/master/src/chapter1/work1_3/work1_3.go)
   - 4.修改 dup2 ，出现重复的行时打印文件名称。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_4/work1_4.go)
   - 5.修改前面的Lissajous程序里的调色板，由黑色改为绿色。我们可以
用 color.RGBA{0xRR, 0xGG, 0xBB, 0xff} 来得到 #RRGGBB 这个色值，三个十六进制的字符串分 别代表红、绿、蓝像素。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_5/work1_5.go)
   - 6.修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex
的第三个参数，看看显示结果吧。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_6/work1_6.go)
   - 7.函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用 这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区 (例子中的b)来存储。记得处理io.Copy返回结果中的错误。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_7/work1_7.go)
   - 8.修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该 前缀。你可能会用到strings.HasPrefix这个函数。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_8/work1_8.go)
   - 9.修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_9/work1_9.go)
   - 10.找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个 URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一 致，修改本节中的程序，将响应结果输出，以便于进行对比。
[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_10/work1_10.go)
   - 11.在fatchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里 排名靠前的。如果一个网站没有回应，程序将采取怎样的行为？(Section8.9 描述了在这种 情况下的应对机制)。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_11/work1_11.go)
   - 12.修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/？ cycles=20 这个URL，这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字 可以调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter1/work1_12/work1_12.go)
   
- 2.程序结构
   - 1.向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝 对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_1/work2_1.go)
   - 2.写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的 话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对 应英尺和米，重量单位可以对应磅和公斤等。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_2/work2_2.go)
   - 3.重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。
(11.4节将展示如何系统地比较两个不同实现的性能)。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_3/work2_3.go)
   - 4.用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和
查表算法的性能差异。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_4/work2_4.go)
   - 5.表达式 x&(x-1) 用于将x的最低的一个非零的bit位清零。使用这个算法重写
PopCount函数，然后比较性能。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter2/work2_5/work2_5.go)
- 3.基础数据类型
   - 1.如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素 (虽然许多SVG渲染器会妥善处理这类问题)。修改程序跳过无效的多边形。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_1/work3_1.go)
   - 2.试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_2/work3_2.go)
   - 3.根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色 (#0000ff)。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_3/work3_3.go)
   - 4.参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返 回SVG数据给客户端。服务器必须设置Content-Type头部: "Content-Type", "image/svg+xml"。
[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_4/work3_4.go)
   - 5.实现一个彩色的Mandelbrot图像，使用image.NewRGBA创建图像，使用 color.RGBA或color.YCbCr生成颜色。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_5/work3_5.go)
   - 6.升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每 个像素分层四个子像素，实现它。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_6/work3_6.go)
   - 7.另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如z4 − 1 = 0。 每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_7/work3_7.go)
   - 8.通过提高精度来生成更多级别的分形。使用四种不同精度类型的数字实现相同的 分形:complex64、complex128、big.Float和big.Rat。(后面两种类型在math/big包声明。 Float是有指定限精度的浮点数;Rat是无效精度的有理数。)它们间的性能和内存使用对比如 何？当渲染图可见时缩放的级别是多少？
[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_8/work3_8.go)
   - 9.编写一个web服务器，用于给客户端生成分形的图像。运行客户端用过HTTP参数 参数指定x,y和zoom参数。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_9/work3_9.go)
   - 10.编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_10/work3_10.go)
   - 11.完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_11/work3_11.go)
   - 12.编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的 字符，但是对应不同的顺序。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_12/work3_12.go)
   - 13.编写KB、MB的常量声明，然后扩展到YB。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter3/work3_13/work3_13.go)
- 4.复合数据类型
   - 1.编写一个函数，计算两个SHA256哈希码中不同bit的数目。(参考2.6.2节的
PopCount函数)。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_1/work4_1.go)
   - 2.编写一个程序，默认打印标准输入的以SHA256哈希码，也可以通过命令行标准参
数选择SHA384或SHA512哈希算法。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_2/work4_2.go)
   - 3.重写reverse函数，使用数组指针代替slice。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_3/work4_3.go)
   - 4.编写一个rotate函数，通过一次循环完成旋转。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_4/work4_4.go)
   - 5.写一个函数在原地完成消除[]string中相邻重复的字符串的操作。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_5/work4_5.go)
   - 6.编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格(参考 unicode.IsSpace)替换成一个空格返回。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_6/work4_6.go)
   - 7.修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_7/work4_7.go)
   - 8.修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_8/work4_8.go)
   - 9.编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_9/work4_9.go)
   - 10.修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_10/work4_10.go)
   - 11.编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_11/work4_11.go)
   - 12.流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每 个链接(只下载一次)然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打 印和命令行输入的检索词相匹配的漫画的URL。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_12/work4_12.go)
   - 13.使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_13/work4_13.go)
   - 14.创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用 户信息。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter4/work4_14/work4_14.go)
- 5.函数
   - 1.修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_1/work5_1.go)
   - 2.编写函数，记录在HTML树中出现的同名元素的次数。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_2/work5_2.go)
   - 3.编写函数输出所有text结点的内容。注意不要访问`<script>`和`<style>`元素,因为
这些元素对浏览者是不可见的。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_3/work5_3.go)
   - 4.扩展vist函数，使其能够处理其他类型的结点，如images、scripts和style sheets。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_4/work5_4.go)
   - 5.实现countWordsAndImages。(参考练习4.9如何分词)[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_5/work5_5.go)
   - 6.修改gopl.io/ch3/surface (§3.2) 中的corner函数，将返回值命名，并使用bare
return。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_6/work5_6.go)
   - 7.完善startElement和endElement函数，使其成为通用的HTML输出器。要求:输出 注释结点，文本结点以及每个元素的属性(< a href='...'>)。使用简略格式输出没有孩子结点 的元素(即用 `<img/>` 代替 `<img></img>` )。编写测试，验证程序输出的格式正确。(详见11 章)[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_7/work5_7.go)
   - 8.修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止 forEachNoded的遍历。使用修改后的代码编写ElementByID函数，根据用户输入的id查找第 一个拥有该id元素的HTML元素，查找成功后，停止遍历。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_8/work5_8.go)
   - 9.编写函数expand，将s中的"foo"替换为f("foo")的返回值。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_9/work5_9.go)
   - 10.重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性 (结果不唯一)。
[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_10/work5_10.go)
   - 11.现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图 中的环。[-> code](https://github.com/caistrong/gopl-homework/blob/master/src/chapter5/work5_11/work5_11.go)
   - 12.gopl.io/ch5/outline2(5.5节)的startElement和endElement共用了全局变量 depth，将它们修改为匿名函数，使其共享outline中的局部变量。
   - 13.修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。只 保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页 面。
   - 14.使用breadthFirst遍历其他数据结构。比如，topoSort例子中的课程依赖关系(有向图),个人计算机的文件层次结构(树)，你所在城市的公交或地铁线路(无向图)。