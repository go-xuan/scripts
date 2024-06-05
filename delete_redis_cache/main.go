package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/redis/go-redis/v9"
)

/*
第一个name表示参数名称，在控制台的时候,提供给用户使用.
第二个value表示默认值，如果用户在控制台没有给该参数赋值的话,就会使用该默认值.
第三个usage表示使用说明和描述,在控制台中输入-arg的时候会显示该说明,类似-help
*/
func main() {
	var (
		db       = 7
		key      = "database-server:default_tenant:user:user:ssbm"
		password = "Init@123"
	)

	var hostIn = flag.String("host", "localhost", "输入redis连接IP，例如：-host=127.0.0.1")
	var portIn = flag.String("port", "6379", "输入redis连接端口，例如：-port=6379")
	var passwordIn = flag.String("password", password, "输入redis连接密码，例如：-password=123")
	var dbIn = flag.Int("db", db, "输入操作的数据库，例如：-db=1")
	var keyIn = flag.String("key", key, "输入需要删除的KEY，例如：-key=abc123")
	// 解析输入的参数
	flag.Parse()
	fmt.Println("host:", *hostIn)
	fmt.Println("port:", *portIn)
	fmt.Println("db:", *dbIn)
	fmt.Println("key:", *keyIn)

	var cmd = redis.NewClient(&redis.Options{Addr: *hostIn + ":" + *portIn, Password: *passwordIn, DB: *dbIn})
	var ctx = context.Background()
	ping, err := cmd.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(ping)
	cmd.Del(ctx, *keyIn)
	fmt.Printf("删除缓存成功！key=%s", *keyIn)
	if err = cmd.Close(); err != nil {
		panic(err)
	}
}
