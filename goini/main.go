package main

import (
	"fmt"
	"time"
	"github.com/go-ini/ini"
)

type Note struct {
	Content string
	Cities []string
}
type Person struct {
	Name string
	Age int `ini:"age"`
	Male bool 
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}
// 从ini文件映射到struct
func map_ini_to_struct () {
	cfg_map,err := ini.Load("mapconfig.ini")
	if err != nil {
		fmt.Println("load err:",err)
	}

	p := new(Person)

	err = cfg_map.Section("").MapTo(p)

	if err != nil {
		fmt.Println("map err:",err)
	}
	fmt.Println("name:",p.Name)
	fmt.Println("age:",p.Age)
	fmt.Println("Male:",p.Male)
	
	cfg_map.Section("").NewKey("Name","666")

	cfg_map.SaveTo("mapconfig.ini")
}
//struct to ini
func map_struct_to_ini() {
		//从结构反射例子
		a := &Person{
			"zly",
			22,
			true,
			time.Now(),
			Note{
				"666",
				[]string{"13","23","ttt"},
			},
			time.Now(),
			}
		cfg := ini.Empty()
		err := ini.ReflectFrom(cfg, a)
		if err != nil {
			fmt.Println("err",err)
		}
		cfg.SaveTo("struct_map_v1.ini")
}

func main() {
	//加载
	cfg,err := ini.Load("config.ini")
	//全部转化为小写,如果有重复性的项目，保留最新的
	//cfg,err := ini.InsensitiveLoad("config.ini")
	
	if err != nil {
		fmt.Println("err")
	}
	// 操作默认分区的key，（新增和修改都是NewKey）
	cfg.Section("").NewKey("name","123")
	// 新增分区，并在新的分区添加key
	cfg.NewSection("demo1")
	cfg.Section("demo1").NewKey("name","666")
	key1,err := cfg.Section("demo1").GetKey("name")
	fmt.Println("key1:",key1)

	// demo2 常用值类型
	cfg.NewSection("demo2")
	v := cfg.Section("demo2").Key("BOOL").MustBool(true)
	v1 := cfg.Section("demo2").Key("BOOL1").MustBool(false)
	v2 := cfg.Section("demo2").Key("int_test").MustInt(5555)
	v3 := cfg.Section("demo2").Key("time_test").MustTimeFormat(time.RFC3339, time.Now())

	fmt.Println("is_ok:",v)
	fmt.Println("is_ok:",v1)
	fmt.Println("is_ok:",v2)
	fmt.Println("is_ok:",v3)


	// demo3 存取多行值（存取）
	cfg.NewSection("demo3")
	
	address := `123,
456,
123`
	v4 :=cfg.Section("demo3").Key("address").MustString(address)
	fmt.Println("v4:",v4)

	// demo4 ini 和struct的相互映射
	map_ini_to_struct()
	map_struct_to_ini()
	//保存
	cfg.SaveTo("config.ini")
}
