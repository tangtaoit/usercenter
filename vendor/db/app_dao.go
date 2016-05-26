package db

import (
	"comm"
	"time"
)

type APP struct  {
	Id uint64
	//应用ID
	AppId string
	//应用KEY
	AppKey string
	//应用名称
	AppName string
	//应用描述
	AppDesc string
	//应用状态 0.待审核 1.已审核
	Status int
	//openID
	OpenId string
	//创建时间
	CreateTime time.Time
	//修改时间
	UpdateTime time.Time
}

func NewAPP() *APP  {

	return &APP{}
}

func (self *APP)  Insert() bool{

	self.CreateTime=time.Now()
	self.UpdateTime=time.Now()

	stmt,err :=db.Prepare("insert into app(app_id,app_key,app_name,app_desc,status,create_time,update_time) values(?,?,?,?,?,?,?)")
	comm.CheckErr(err)
	_,err =stmt.Exec(self.AppId,self.AppKey,self.AppName,self.AppDesc,self.Status,self.CreateTime,self.UpdateTime)
	comm.CheckErr(err)
	return true
}

//查询可用的APP
func (self *APP) QueryCanUseApp(appId string) *APP {

	stmt,err := db.Prepare("select id,app_id,app_key,app_name,app_desc,status from app where app_id=? and status=1")
	comm.CheckErr(err)

	rows,err := stmt.Query(appId);

	defer rows.Close()
	comm.CheckErr(err)
	if rows.Next()  {
		app :=NewAPP()
		rows.Scan(&app.Id,&app.AppId,&app.AppKey,&app.AppName,&app.AppDesc,&app.Status)

		return app
	}

	return nil;
}
