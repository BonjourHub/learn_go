package main

import (
  "encoding/json"
  "fmt"
  "math/cmplx"
  "hello/student"
)

func main() {
  fmt.Println("hellow world!")
  // 变量
  learnVar()

  // 数据类型
  learnDatatype()

  // 数组 切片 map
  learSetMap()

  // 数组类型转换
  transDataType()

  // 流程控制
  learnProcessControl()

  // 指针
  learnPoint()

  // 函数定义
  learnFunction()

  // 方法 结构体 接口
  learnStructInterface()

  // 自定义包
  callCustomPackge()

  // 内置包：Json编码
  learnJsonEncode()
}

//// 变量声明
func learnVar()  {
  fmt.Println("=== 变量声明 ===")
  var a int // 默认赋值为0
  var b  = 1 // 初始化变量为1 自动确定类型为int
  str := "另一种声明方式" // := 是一个声明语句 只能在函数体内使用

  fmt.Print(a, b, str, "\n")
}


//// 数据类型
func learnDatatype()  {
  fmt.Println("=== 数据类型 ===")
  // string
  var str string = "字符串类型"
  var str1  = "字符串类型"
  var a int = 12 // 整型
  var a1 int16  = 345 // 整型
  var flo float32 = 98.9; // 浮点型
  var x complex128 = cmplx.Sqrt(-5 + 12i) // 复数类型

  fmt.Println(str, str1, a, a1, flo, x)
}


//// 数组 切片 maps
func learSetMap()  {
  fmt.Println("=== 数组 ====")
  var arr [5]int // 一维数组 固定数组长度 不能扩展 
  var multiD [2][3]int // 多维数组

  fmt.Println("=== 切片 ===")
  var mArr []float32 // 切片定义方式：不指定数组长度 可扩展 类似iOS可变数组
  var mArr1 []int = make([]int, 5, 10)
  mArr2 := make([]int, 5, 10) // 初始化长度为5，最大容量为10
  // append 给数组增加值，当容量不够的时候，自动扩充容量
  mArr2 = append(mArr2, 1, 2, 3)
  // copy 复制一个新切片: 创建一个更大容量的切片3 将mArr2复制到mArr3
  mArr3 := make([]int, 15)
  copy(mArr3, mArr2)

  fmt.Println(arr,  multiD, mArr, mArr1, mArr2)


  fmt.Println("=== 切片裁剪 ===")
  /// 切片数组裁剪 创建子切片
  numbers := []int {1, 2, 3, 4} // index-> 0 1 2 3
  fmt.Println(numbers) // -> [1, 2, 3, 4]
  // 创建子切片
  slice1 := numbers[2:]
  fmt.Println(slice1)  // -> [3, 4] 半开半闭区间/左闭右开区间
  slice2 := numbers[:3]
  fmt.Println(slice2) // -> [1, 2, 3]
  slice3 := numbers[1:4]
  fmt.Println(slice3) // -> [2, 3, 4]


  fmt.Println("=== 创建map字典 ===")
  dic := make(map[string]int) // key为string类型 value为int类型
  dic["age"] = 18 // adding key/value
  dic["tel"] = 110110
  fmt.Println(dic, len(dic))
  fmt.Println(dic["age"]) // 访问字典


  fmt.Println("=== Range枚举数组 ===")
  // 枚举
  for i, num := range numbers { // 如果不需要使用下标 可以使用 _ 替换, i num 接收多返回值
    fmt.Printf("枚举：index:%d value:%d \n", i, num)
  }

}


//// 数据类型转换
func transDataType()  {
  fmt.Println("=== 数据类型转换 ===")
  a := 1.1
  b := int(a)
  fmt.Println(b)
}


//// 流程控制
func learnProcessControl()  {
  /// if else 语句
  fmt.Println("=== if else ===")
  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }

  /// switch语句 可以是字符串
  i := 2
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  default:
    fmt.Println("none")
  }

  /// for循环语句
  j := 0
  sum := 0
  for j < 10  { // 类似于while循环
    sum += 1
    j++
  }

  for k := 0; k < 10; k++  {
    sum += k;
    fmt.Print("for循环：", sum, k)
  }


  /// for 无限循环
  //for {
  //  println("无限xunhuan")
  //}

  print("\n")

}

//// 指针
func learnPoint()  {
  fmt.Println("=== 指针 ===")
  var ap *int // 指向整数类型的指针
  a := 1212
  ap = &a
  fmt.Println(*ap)

  /*
  1. 函数参数传值，实际是复制意味这开辟更多内存
  2. 使用指针传值，节省内存，可在函数/调用者内修改指针中的值
  */
  // learnModPointValue(&a)
  learnModPointValue(ap)
  fmt.Println("修改后的指针值", *ap, a)
}

func learnModPointValue(i *int)  {
  *i += 10;
}


//// 函数
func learnFunction()  {
  fmt.Println("=== 函数 ===")
  fmt.Println("加法函数:" ,add(2, 8))
  fmt.Println("加法函数:" ,add1(3, 8))
  sum, message := add2(4, 8)
  fmt.Println("多返回值加法：", sum, message)
}

func add(a int, b int) int { // 返回值类型为int
  sum := a + b
  return sum
}

func add1(a int, b int) (sum int)  { // 函数中预先定义返回值
  sum = a + b
  return // sum 自动返回 无须重复定义
}

func add2(a int, b int) (int, string)  { // 多返回值类型函数定义
  sum := a + b
  return sum, "success"
}


///// 方法 结构体 接口
func learnStructInterface()  {
  fmt.Println("=== 结构体 ===")
  // 结构体实例化
  // 方式1：指定属性和值
  p := person{name: "Bob", age: 42, gender: "Male"}
  // 方式2：指定值
  p1 := person{"Lili", 18, "Man"}
  fmt.Println("点号访问p:", p.name, p.age, p.gender) // 访问
  fmt.Println("点号访问p1", p1.name, p1.age, p1.gender)
  pp := &person{"Sanm", 17, "Man"}
  fmt.Println("指针访问", pp.name, pp.age, pp.gender)

  // 方法
  fmt.Println("=== 方法 ===")
  pp.describe()
  p1.describe()
  pp.setAge(18)
  pp.describe()
  pp1 := p1.setName("Susan") // 不会修改p1的name 因为传入的不是指针而是值类型
  pp1.describe()

  // 接口
  debugInterface()
}

// 定义 person 结构体
type person struct {
  name string
  age int
  gender string
}

// 定义方法
func (p *person) describe()  {
  fmt.Printf("%v is %v years old. \n", p.name, p.age) // Printf 格式化字符串
}

func (p *person) setAge(age int)  {
  p.age = age
}

func (p person) setName(name string) person {
  p.name = name;
  return p
}

// 定义接口
func debugInterface()  {
  fmt.Println("=== 定义接口 ===")

  var a animal
  a = snake{Poisonous:true}
  fmt.Println(a.description()) // => Poisonous: true
  a = cat{Sound:"Meow!!!"}
  fmt.Println(a.description()) // => Sound: Meow!!!

  /*
  类似于iOS中的协议代理
  animal接口interface声明了一个方法
  实际运行过程中调用的是实例自己的方法
  */
}
type animal interface {
  description() string
}

type snake struct {
  Type string
  Poisonous bool
}

type cat struct {
  Type string
  Sound string
}

func (s snake) description() string  {
  return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string  {
  return fmt.Sprintf("Sound: %v", c.Sound)
}



/*
 注意：默认workspace $HOME/go
 包会自动去默认workspace中pkg中索引
 自定义$GOPATH 看官网文档
 go build hello.go 编译输出.o文件
 go install hello.go 在bin目录生成可执行文件
*/
///// 包
func callCustomPackge()  {
 fmt.Println("=== 自定义包 ===")
s := student.Description("Milap")
fmt.Println(s)
}

/*
 json.Marshal 将go对象编码为json串（go对象：map struct等）
 json.MarshalIndent 将go对象编码为整齐格式的字符串

 json.Unmarshal 将json串解码为Go数据结构
*/
//// Json Encode
func learnJsonEncode()  {
  fmt.Println("=== Json ====")
  mapA := map[string]int{"apple": 5, "lettuce": 7}
  mapB, _:= json.Marshal(mapA)
  fmt.Println(string(mapB))
}


// 通道 类似管道
