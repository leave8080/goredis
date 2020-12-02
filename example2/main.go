package main

import (
	"github.com/garyburd/redigo/redis"

	"fmt"
	"time"
	"strconv"
	"runtime"
)

var Address = "127.0.0.1:16379"
var Network = "tcp"
func GetRedis()  redis.Conn  {
	c, err := redis.Dial(Network, Address)
	if err != nil {
		return GetRedis()
	}
	err=c.Send("auth","123456")
	if err !=nil {
		fmt.Println(err)
	}
	return c
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i:=0;i<100;i++  {
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
		go do()
	}
	time.Sleep(time.Second*60*10)
}
func do()  {
	cnn:=GetRedis()
	defer   cnn.Close()
	redisLock("lock.foo",cnn,20,doFunc,"致远")

}
//lockKey锁的名称
//cnn       redis.Conn
//deadTime      锁默认消亡时间
//doFunc        参数名称
//param     方法参数
func redisLock(lockKey string,cnn redis.Conn,deadTime int,doFunc func(interface{}),param interface{})  {
	setnxTime:=time.Now().UTC().UnixNano()
	ex,err:=cnn.Do("SETNX",lockKey,setnxTime+int64(deadTime))
	if err==nil {
		if ex==int64(0) {
			//fmt.Println("存在锁:下来判断锁是否过期了")
			lock2,err:=cnn.Do("GET",lockKey)
			if lock2==nil {
				//fmt.Println("lock2=======为空ex",ex,ex==int64(0))
				redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
				return
			}
			if err!=nil {
				redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
				return
			}
			getTime, err :=strconv.ParseInt(string(lock2.([]uint8)), 10, 64)
			if getTime>setnxTime {
				//锁未过期
				//fmt.Println("锁没有过期：继续等吧")
				redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
				return
			}else {
				//锁已经过期
				time.Sleep(time.Millisecond*time.Duration(deadTime))//线程休眠
				getsettime:=time.Now().UTC().UnixNano()
				lock3,err:=cnn.Do("GETSET",lockKey,getsettime)
				if lock3==nil {
					//fmt.Println("lock3=======为空")
					redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
					return
				}
				getSetTime, err :=strconv.ParseInt(string(lock3.([]uint8)), 10, 64)
				if err!=nil {
					//fmt.Println("出问题了：去继续等吧")
					redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
					return
				}
				if getSetTime==getTime {//如果更改前的时间和已经过期的时间相同
					//获得锁直接操作数据
					//fmt.Println("锁过期：处理了死锁，可以直接操作数据")
					doFunc(param)
					cnn.Do("DEL",lockKey)//删除锁
					return
				}else{//更改前的时间和已经过期的时间不同
					//fmt.Println("判断后：没有死锁，继续等吧")
					redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
					return
				}
			}
		}else{
			//fmt.Println("不存在锁：可以操作数据")
			doFunc(param)
			cnn.Do("DEL",lockKey)//删除锁
			return
		}
	}else {
		redisLock(lockKey ,cnn ,deadTime ,doFunc ,param )
		return
	}
}
var count=0
func doFunc(str interface{})  {
	count+=1
	fmt.Println("操作数据中.............============================",count,str)
	return
}