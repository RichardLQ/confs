package confs

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var Default =NewContainer()

func NewContainer()  *Container {
	return &Container{}
}

type Tags struct {
	DataBaseName string `json:"data_base_name"` //数据库名称
	NameSpace    string `json:"name_space"`     //命令空间
	ECode        string `json:"e_code"`         //编码格式
	FileType     string `json:"file_type"`      //初始化类型
}

type FileStruct struct {
	SName string `json:"s_name"`
	SValue *reflect.Value `json:"s_value"`
	STag  *Tags 	`json:"s_tag"`
}

type Container struct {
	OldStruct *FileStruct `json:"old_struct"`
	Datas map[string]interface{} `json:"datas"`
}

func (c *Container)Start() *Container {
	GLOAB_URL = os.Getenv("DATABASEPATH")
	return nil
}

func (c *Container)StartConf(u ...interface{}) *Container {
	for _, v := range u {
		vl := reflect.ValueOf(v).Elem()
		tl := reflect.TypeOf(v).Elem()
		for i := 0; i < vl.NumField(); i++ {
			il := vl.Field(i)
			til := tl.Field(i)
			info :=&FileStruct{
				SName: til.Name,
			}
			tag := Tags{}
			UnmarshalTag(key, &tag, til.Tag)
			info.STag = &tag
			c.OldStruct = info
			if il.Kind() == reflect.Struct { //填补结构体
				temp :=Default.FIllStruct(&il)
				vl.FieldByName(til.Name).Set(temp)
			}
		}
	}
	return c
}

func (c *Container) FIllStruct(il *reflect.Value) reflect.Value {
	items := il.Type()
	data :=cfg[c.OldStruct.STag.FileType].(map[string]interface{})[c.OldStruct.SName].(map[string]interface{})
	c.Datas = data
	if c.OldStruct.SName == c.OldStruct.STag.DataBaseName {
		b,err :=c.InitServe(il)
		if err!= nil {
			panic(fmt.Sprintf("数据库启动失败:%+v", err))
		}
		return reflect.ValueOf(b)
	}
	newType:=reflect.New(items).Elem()
	for i := 0; i < items.NumField(); i++ {
		newType.FieldByName(items.Field(i).Name).Set(reflect.ValueOf(data[items.Field(i).Name]))
	}
	return newType
}

// UnmarshalTag ...
func UnmarshalTag(key string, info *Tags, tag reflect.StructTag) error {
	tagString := tag.Get(key)
	tags := strings.Split(tagString, ";")
	info.FileType = FILEEXT
	for _, s := range tags {
		tagList := strings.Split(s, "=")
		switch tagList[0] {
		case "name":
			info.DataBaseName = tagList[1]
			info.FileType = DATABASEEXT
			break
		case "read":
			info.NameSpace = tagList[1]
			break
		case "format":
			info.ECode = tagList[1]
			break
		}
	}
	info.ReadConf()
	return nil
}