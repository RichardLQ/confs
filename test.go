package confs




//var GlobalFile = struct {
//	XinyueAMSConf    AMSConfig `confs:"read=rainbow;format=json"`
//	SuperCoreAMSConf AMSConfig `confs:"read=rainbow;format=json"`
//	SparkAMSConf     AMSConfig `confs:"read=rainbow;format=json"`
//}{}
//
//var Global = struct{
//	Work  gorm.DB `confs:"name=Work;read=rainbow"`
//	Test  redis.Conn `confs:"name=Test;read=rainbow"`
//	Worka  gorm.DB `confs:"name=Worka;read=rainbow"`
//}{}
//
////老师信息
//type Teacher struct {
//	Id         string `json:"id"`         //老师id
//	Nickname   string `json:"nickname"`   //昵称
//	Name       string `json:"name"`       //名字
//	Address    string `json:"address"`    //头像地址
//	Desc       string `json:"desc"`       //简介
//	Phone      int64  `json:"phone"`      //电话号码
//	Qrcode     string `json:"qrcode"`     //二维码
//	Type       int64  `json:"type"`       //舞蹈类型
//	Updateime  int64  `json:"updateime"`  //更新时间
//	Createtime int64  `json:"createtime"` //创建时间
//}
//
////AMSConfig ...
//type AMSConfig struct {
//	AppID      string
//	ActivityID string
//	Token      string
//}
//
//func main()  {
//	confs.NewStart().BinComb(&Global,&GlobalFile)
//	fmt.Println(GlobalFile.XinyueAMSConf.AppID)
//	list := &Teacher{}
//	err:=Global.Work.Table("teacher").Where("id = ?", 1).Find(list).Error
//	fmt.Println(err)
//	fmt.Println(list)
//}
