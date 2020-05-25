package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	// 模拟一些私有的数据
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	/*yzj := []int{2,7,1,6,4,3,11,15,17,5,8,9,12}
	fmt.Println("该切片的长度是：",len(yzj))

	for _,n := range yzj{
		go func(n int) { //定义一个匿名函数，并对该函数开启协程，每次循环都会开启一个协成，也就是说它开启了13个协程。
			time.Sleep(time.Duration(n) * time.Second) //表示每循环一次就需要睡1s，睡的总时间是由n来控制的，总长度是由s切片数组中最大的一个数字决定，也就是说这个协成最少需要17秒才会结束哟。
			fmt.Println(n)
		}(n)  //由于这个函数是匿名函数，所以调用方式就直接：（n）调用，不用输入函数名。
	}
	time.Sleep(17*time.Second) //主进程要执行的时间是12秒.*/

	// golang中time.After的理解
	//closeChannel
	/*c := make(chan int)
	timeout := time.After(time.Second * 2)
	t1 := time.NewTimer(time.Second * 3)  // 效果相同 只执行一次
	var i int
	go func() {
		for{
			select {
			case <-c:
				fmt.Println("channel sign")
				return
			case <-t1.C:    //代码段2
				fmt.Println("3s定时任务")
			case <-timeout:   //代码段1
				i++
				fmt.Println(i, "2s定时输出")
			case <-time.After(time.Second * 4):   //代码段3
				fmt.Println("4s timeout ...")
			default:                                  //代码段4
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 6)
	close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")*/

	//超时控制
/*	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
	}()

	select {
	case res := <- ch:
		fmt.Println(res)
		time.Sleep(time.Second * 10)
	case res := <- ch:
		fmt.Println(res)
	case <- time.After(time.Second * 4):
		fmt.Println("timeout")
	}*/

	//定义超时时间
	/* TimeOut := time.After(5 * time.Second)    //定义超时时间
	 NerverRings := make(chan int ,3)
	 NerverRings <- 1

	 RingsOccasionally := make(chan bool)
	 go func() {
	 	for {
			select {
			case value := <- NerverRings:
				fmt.Println("hahahah",value)
			case <- TimeOut:
				println("对不起，目前为止，NerverRings 并灭有接收到任何数据！程序以终止")
				RingsOccasionally <- true
				break
			}
		}
	 }()
	 <- RingsOccasionally   //从RingsOccasionally 这个channel 中获取数据，所以在获取数据之前，成功是出于阻塞状态的哟*/

	/*ch := make(chan int ,2)
	ch <- 1
	ch <- 2
	go func() {
		for data := range ch {
			fmt.Println(data)
			time.Sleep(time.Second * 2)
		}
	}()

time.Sleep(time.Second * 8)*/
/*	data := make(chan int)    //默认该channel 就是可读可写的哟

	go func() {
		write := chan<- int(data)  //此处的write 是一个单向的写入channel.
		write <- 150
		fmt.Println("hhe",write)
	}()
	read := <-chan int(data) 	// 此处的read就是一个单向的读取channel
	fmt.Println("读取", read)

	time.Sleep(time.Second * 5)*/


// 启动一个gin 服务并自定义中间键 并使用
  r := gin.New()
  r.Use(Logger())

  //在组中使用 gin.BasicAuth() 中间件
  //gin.Accounts 是 map[string]string 的快捷方式
  authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
  	"foo": "bar",
  	"austin": "1234",
  	"lena": "hello2",
  	"manu": "4321",
  }))

  // /admin/secrets 结尾
  //  点击 “localhost:8080/admin/secrets”
  authorized.GET("/secrets", func(c *gin.Context){
  	// 获取 user, 它是由 BasicAuth 中间件设置的
  	user := c.MustGet(gin.AuthUserKey).(string)
  	if secret, ok := secrets[user]; ok{
  		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
  })
  r.GET("/test", func(c *gin.Context){
  	example := c.MustGet("example").(string)

  	log.Println(example)
  })

   if r.Run(":8800") != nil {
   	fmt.Println("错误")
   }

}

//中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//设置简单的变量
		c.Set("example", "123456")


		//在请求之前

		c.Next()

		//在请求之后
		latency := time.Since(t)
		log.Print(latency)

		// 记录我们的访问状态
		status := c.Writer.Status()
		log.Println(status)
	}
}


