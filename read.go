package confs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Read 读取配置
func (t *Tags) ReadConf() {
	if len(cfg) == 0 {
		cfg = make(map[string]interface{})
		for i, s := range DB_EXT_TYPE {
			path := GLOAB_URL + "/" + t.NameSpace + "/" + s
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				panic(fmt.Sprintf("读取配置文件错误:%+v", err))
			}
			var data, _ = ioutil.ReadFile(path)
			var temp map[string]interface{}
			json.Unmarshal(data, &temp)
			cfg[i] = temp
		}
	} 
	//else {
	//	fmt.Println("已经初始化")
	//}

}
