package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"math"
	"strconv"
	"time"
)

type Catalog struct {
	Id            int64
	ProjectNumber string    //项目编号
	ProjectName   string    //项目名称
	DesignStage   string    //阶段
	Section       string    //专业
	Tnumber       string    //成果编号
	Name          string    //成果名称
	Category      string    //成果类型
	Page          string    //成果计量单位
	Count         float64   //成果数量
	Drawn         string    //编制、绘制
	Designd       string    //设计
	Checked       string    //校核
	Examined      string    //审查
	Verified      string    //核定
	Approved      string    //批准
	Complex       float64   //难度系数
	Drawnratio    float64   //编制、绘制占比系数
	Designdratio  float64   //设计系数
	Checkedratio  float64   //校核系数
	Examinedratio float64   //审查系数
	Datestring    string    //保存字符型日期
	Date          time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created       time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"index;auto_now_add;type(datetime)"`
	Author        string    //上传者
	State         int       //1编写状态，未提交；2编写者提交，等待校核确认;3,校核确认，等待审查确认;4，审查确认
}

//任何人只能在线填写自己是编制和设计的，填写自己是校核和审查的不允许

//员工的编制、设计……分值——全部改成float浮点型小数
type Employeeachievement struct {
	Id       string  //员工Id
	Name     string  //员工姓名
	Drawn    float64 //编制、绘制
	Designd  float64 //设计
	Checked  float64 //校核
	Examined float64 //审查
	Verified float64 //核定
	Approved float64 //批准
	Sigma    float64 //合计
}

//分院里各个科室人员结构
type Secofficeachievement struct {
	Id       int64  //科室Id
	Name     string //科室
	Employee []Employeeachievement
}

func init() {
	orm.RegisterModel(new(Catalog))
}

//在线添加，批量导入
func SaveCatalog(catalog Catalog) (cid int64, err error, news string) {
	// orm := orm.NewOrm()
	// fmt.Println(user)
	//判断重复性
	o := orm.NewOrm()
	var catalog1 Catalog
	//保证成果的唯一性
	//出差必须在成果名称中写入自己的名字以示区别
	//Filter("Drawn", catalog.Drawn).Filter("Designd", catalog.Designd).Filter("Checked", catalog.Checked).
	err = o.QueryTable("catalog").Filter("Tnumber", catalog.Tnumber).Filter("Name", catalog.Name).Filter("Category", catalog.Category).One(&catalog1)
	if err == orm.ErrNoRows {
		cid, err = o.Insert(&catalog) //_, err = o.Insert(reply)
		if err != nil {
			return 0, err, "insert err"
		} else {
			return cid, nil, "insert ok"
		}
		// fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		return 0, err, "找不到主键"
		//     fmt.Println("找不到主键")
	} else {
		return 0, nil, "数据已存在"
	}
}

func GetAllCatalogs(cid string) (catalogs []*Catalog, err error) {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return nil, err
	}
	catalogs = make([]*Catalog, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	_, err = qs.Filter("parentid", cidNum).All(&catalogs)
	return catalogs, err
}

//用savecatalog，下面这个没用？
func AddCatalog(name, tnumber string) (id int64, err error) {
	// cid, err := strconv.ParseInt(categoryid, 10, 64)
	o := orm.NewOrm()
	catalog := &Catalog{
		Name:    name,
		Tnumber: tnumber,
		// Category:   category,
		// CategoryId: cid,
		// Content:    content,
		// Attachment: attachment,
		// Author:     uname,
		// Created:    time.Now(),
		// Updated:    time.Now(),
		// ReplyTime:  time.Now(),
	}
	//	qs := o.QueryTable("category") //不知道主键就用这个过滤操作
	//	err := qs.Filter("title", name).One(cate)
	//	if err == nil {
	//		return err
	//	}
	id, err = o.Insert(catalog)
	if err != nil {
		return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
	}
	if id == 0 {
		var catalog Catalog
		err = o.QueryTable("catalog").Filter("tnumber", tnumber).One(&catalog, "Id")
		id = catalog.Id
	}
	return id, err
}

//用户修改一条成果的某个字段
func ModifyCatalog(cid int64, fieldname, value string) error {
	o := orm.NewOrm()
	var catalog Catalog
	// catalog := &Catalog{Id: cid}
	err := o.QueryTable("catalog").Filter("Id", cid).One(&catalog)
	// err:=o.Read(catalog).One()
	if err == nil {
		type Duration int64
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
		// hours := 8

		const lll = "2006-01-02"
		catalog.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
		switch fieldname {
		case "ProjectNumber":
			catalog.ProjectNumber = value
			_, err := o.Update(&catalog, "ProjectNumber", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "ProjectName":
			catalog.ProjectName = value
			_, err := o.Update(&catalog, "ProjectName", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "DesignStage":
			catalog.DesignStage = value
			_, err := o.Update(&catalog, "DesignStage", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Section":
			catalog.Section = value
			_, err := o.Update(&catalog, "Section", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Tnumber":
			catalog.Tnumber = value
			_, err := o.Update(&catalog, "Tnumber", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Name":
			catalog.Name = value
			_, err := o.Update(&catalog, "Name", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Category":
			catalog.Category = value
			_, err := o.Update(&catalog, "Category", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Count":
			//转成float64
			catalog.Count, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Count", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Drawn":
			catalog.Drawn = value
			_, err := o.Update(&catalog, "Drawn", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Designd":
			catalog.Designd = value
			_, err := o.Update(&catalog, "Designd", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Checked":
			catalog.Checked = value
			_, err := o.Update(&catalog, "Checked", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Examined":
			catalog.Examined = value
			_, err := o.Update(&catalog, "Examined", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Drawnratio":
			catalog.Drawnratio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Drawnratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Designdratio":
			catalog.Designdratio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Designdratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Checkedratio":
			catalog.Checkedratio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Checkedratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Examinedratio":
			catalog.Examinedratio, err = strconv.ParseFloat(value, 64)
			_, err := o.Update(&catalog, "Examinedratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Complex":
			catalog.Complex, err = strconv.ParseFloat(value, 64)
			_, err := o.Update(&catalog, "Complex", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Datestring":
			const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
			catalog.Date, err = time.Parse(lll, value)
			if err != nil {
				return err
			}
			catalog.Datestring = value
			_, err := o.Update(&catalog, "Datestring", "Date", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
	} else {
		return o.Read(&catalog)
	}
	return nil
}

//用户修改一条成果状态
func ModifyCatalogState(cid int64, state int) error {
	o := orm.NewOrm()
	var catalog Catalog
	err := o.QueryTable("catalog").Filter("Id", cid).One(&catalog)
	if err == nil {
		catalog.State = state
		_, err = o.Update(&catalog, "State") //这里不能用&catalog
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

func DeletCatalog(cid string) error { //应该在controllers中显示警告
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	catalog := Catalog{Id: cidNum}
	if o.Read(&catalog) == nil {
		_, err = o.Delete(&catalog)
		if err != nil {
			return err
		}
	}
	return err
}

func GetCatalog(tid string) (*Catalog, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	catalog := new(Catalog)
	qs := o.QueryTable("catalog")
	err = qs.Filter("id", tidNum).One(catalog)
	if err != nil {
		return nil, err
	}
	// catalog.Views++
	// _, err = o.Update(topic)

	// attachments := make([]*Attachment, 0)
	// attachment := new(Attachment)
	// qs = o.QueryTable("attachment")
	// _, err = qs.Filter("topicId", tidNum).OrderBy("FileName").All(&attachments)
	// if err != nil {
	// 	return nil, err
	// }
	return catalog, err
}

// func GetPids(pid int64) ([]*Category, error) {
// 	o := orm.NewOrm()
// 	cates := make([]*Category, 0)
// 	qs := o.QueryTable("category")
// 	var err error
// 	//这里进行过滤
// 	_, err = qs.Filter("ParentId", pid).All(&cates)
// 	// _, err = qs.OrderBy("-created").All(&cates)
// 	// _, err := qs.All(&cates)
// 	return cates, err
// }

//由用户姓名取得所有编制、设计、校核分值
func Getemployeevalue(uname string, t1, t2 time.Time) (employeevalue []Employeeachievement, err error) {
	catalogs := make([]*Catalog, 0)
	cond := orm.NewCondition()
	cond1 := cond.And("date__gt", t1).And("Date__lte", t2)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)
	//1、查制图工作量
	_, err = qs.Filter("Drawn", uname).Filter("State", "5").All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	// slice1 := make([]Person, 0)
	var Drawnvalue float64
	var drawn float64
	var Designdvalue float64
	var designd float64
	var Checkedvalue float64
	var checked float64
	var Examinedvalue float64
	var examined float64
	aa := make([]Employeeachievement, 1)
	// var aa *Employeeachievement
	for _, v := range catalogs {
		//根据catalogs的category，再查出ratio中的系数
		ratio, err := GetRationumbycategory(v.Category)
		if err == orm.ErrNoRows {
			ratio = 0
		} else if err != nil {
			return nil, err
		}
		Drawnvalue = v.Count * ratio * v.Complex * v.Drawnratio
		// Category      string    //成果类型
		// Page          string    //成果计量单位
		// Count         int       //成果数量
		// Drawn         string    //编制、绘制
		// Designd       string    //设计
		// Checked       string    //校核
		// Examined      string    //审查
		// Verified      string    //核定
		// Approved      string    //批准
		// Complex       int       //难度系数
		// Drawnratio    int       //编制、绘制占比系数
		// Designdratio  int       //设计系数
		// Checkedratio  int       //校核系数
		// Examinedratio int       //审查系数
		// mark, err := strconv.Atoi(v.Count)
		// if err != nil {
		// 	return nil, err
		// }

		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		// switch v.Category {
		// case "图纸":
		// 	Drawnvalue = v.Count * v.Complex * v.Drawnratio
		// case "报告":
		// 	Drawnvalue = v.Count / 5 * v.Complex * v.Drawnratio
		// case "大纲":
		// 	Drawnvalue = v.Count / 5 * v.Complex * v.Drawnratio
		// case "计算书":
		// 	Drawnvalue = v.Count / 5 * v.Complex * v.Drawnratio
		// case "修改单":
		// 	Drawnvalue = v.Count * v.Complex * v.Drawnratio
		// default:
		// 	Drawnvalue = v.Count * v.Complex * v.Drawnratio
		// }
		drawn = drawn + Drawnvalue
	}
	aa[0].Drawn = Round(drawn, 1)

	//2、查设计工作量
	_, err = qs.Filter("Designd", uname).Filter("State", "5").All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		ratio, err := GetRationumbycategory(v.Category)
		if err == orm.ErrNoRows {
			ratio = 0
		} else if err != nil {
			return nil, err
		}
		Designdvalue = v.Count * ratio * v.Complex * v.Designdratio
		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		// switch v.Category {
		// case "图纸":
		// 	Designdvalue = v.Count * v.Complex * v.Designdratio
		// case "报告":
		// 	Designdvalue = v.Count / 5 * v.Complex * v.Designdratio
		// case "大纲":
		// 	Designdvalue = v.Count / 5 * v.Complex * v.Designdratio
		// case "计算书":
		// 	Designdvalue = v.Count / 5 * v.Complex * v.Designdratio
		// case "修改单":
		// 	Designdvalue = v.Count * v.Complex * v.Designdratio
		// default:
		// 	Designdvalue = v.Count * v.Complex * v.Designdratio
		// }
		designd = designd + Designdvalue
	}
	aa[0].Designd = Round(designd, 1)

	//3、查校核工作量
	_, err = qs.Filter("Checked", uname).Filter("State", "5").All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		ratio, err := GetRationumbycategory(v.Category)
		if err == orm.ErrNoRows {
			ratio = 0
		} else if err != nil {
			return nil, err
		}
		Checkedvalue = v.Count * ratio * v.Complex * v.Checkedratio
		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		// switch v.Category {
		// case "图纸":
		// 	Checkedvalue = v.Count * v.Complex * v.Checkedratio
		// case "报告":
		// 	Checkedvalue = v.Count / 5 * v.Complex * v.Checkedratio
		// case "大纲":
		// 	Checkedvalue = v.Count / 5 * v.Complex * v.Checkedratio
		// case "计算书":
		// 	Checkedvalue = v.Count / 5 * v.Complex * v.Checkedratio
		// case "修改单":
		// 	Checkedvalue = v.Count * v.Complex * v.Checkedratio
		// default:
		// 	Checkedvalue = v.Count * v.Complex * v.Checkedratio
		// }
		checked = checked + Checkedvalue
	}
	aa[0].Checked = Round(checked, 1)
	//4、查审查工作量
	_, err = qs.Filter("Examined", uname).Filter("State", "5").All(&catalogs) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	for _, v := range catalogs {
		ratio, err := GetRationumbycategory(v.Category)
		if err == orm.ErrNoRows {
			ratio = 0
		} else if err != nil {
			return nil, err
		}
		Examinedvalue = v.Count * ratio * v.Complex * v.Examinedratio
		//成果数量*难度系数*绘制占比
		//成果数量*难度系数*设计占比
		//如果是图纸
		// switch v.Category {
		// case "图纸":
		// 	Examinedvalue = v.Count * v.Complex * v.Examinedratio
		// case "报告":
		// 	Examinedvalue = v.Count / 5 * v.Complex * v.Examinedratio
		// case "大纲":
		// 	Examinedvalue = v.Count / 5 * v.Complex * v.Examinedratio
		// case "计算书":
		// 	Examinedvalue = v.Count / 5 * v.Complex * v.Examinedratio
		// case "修改单":
		// 	Examinedvalue = v.Count * v.Complex * v.Examinedratio
		// default:
		// 	Examinedvalue = v.Count * v.Complex * v.Examinedratio
		// }
		examined = examined + Examinedvalue
	}
	aa[0].Examined = Round(examined, 1)

	aa[0].Name = uname //这个是nickname，千万注意
	user1 := GetUserByNickname(uname)
	id := strconv.FormatInt(user1.Id, 10)
	aa[0].Id = id
	aa[0].Sigma = Round(drawn+designd+checked+examined, 1)
	employeevalue = aa
	// employeevalue = append(employeevalue, aa...)
	return employeevalue, err
}

//由用户Id取得所有编制、设计、校核详细catalog，按成果类型排列
func Getcatalogbyuserid(id, category string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	aa := make([]*Catalog, 0)
	bb := make([]*Catalog, 0)
	cc := make([]*Catalog, 0)
	dd := make([]*Catalog, 0)
	cond := orm.NewCondition()
	cond1 := cond.And("date__gt", t1).And("Date__lte", t2)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)

	//1、查图纸类型的工作
	_, err = qs.Filter("Drawn", user.Nickname).Filter("Category", category).All(&aa) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, aa...)
	_, err = qs.Filter("Designd", user.Nickname).Filter("Category", category).All(&bb) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, bb...)
	_, err = qs.Filter("Checked", user.Nickname).Filter("Category", category).All(&cc) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, cc...)
	_, err = qs.Filter("Examined", user.Nickname).Filter("Category", category).All(&dd) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, dd...)
	return catalogs, err
}

//由用户Id取得所有成果按时间顺序排列
//不返回重复的值
func Getcatalog2byuserid(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	aa := make([]*Catalog, 0)
	// bb := make([]*Catalog, 0)
	// cc := make([]*Catalog, 0)
	// dd := make([]*Catalog, 0)

	cond := orm.NewCondition()
	cond1 := cond.And("Date__gt", t1).And("Date__lte", t2).Or("Drawn", user.Nickname).Or("Designd", user.Nickname).Or("Checked", user.Nickname).Or("Examined", user.Nickname)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)

	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...

	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()
	//1、查图纸类型的工作
	_, err = qs.Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	// catalogs = append(catalogs, aa...)
	// _, err = qs.Filter("Designd", user.Nickname).All(&bb)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, bb...)
	// _, err = qs.Filter("Checked", user.Nickname).All(&cc)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, cc...)
	// _, err = qs.Filter("Examined", user.Nickname).All(&dd)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, dd...)
	return aa, err
}

//我发起，待提交
//由用户名（不是nickname）和时间段取得自己发起的成果
//不返回重复的值
//author=登录的人名，登录名所处制图-状态为1；设计-状态为2；校核-状态为3；审查-无
func GetcatalogMyself(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	catalogs = make([]*Catalog, 0) //没有这句，model返回null，导致json输出给table为null，而不是[]，导致错误
	aa := make([]*Catalog, 0)
	bb := make([]*Catalog, 0)
	cc := make([]*Catalog, 0)
	// dd := make([]*Catalog, 0)
	cond := orm.NewCondition()
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).Or("Drawn", user.Nickname).And("State", "1").Or("Designd", user.Nickname).And("State", "2").Or("Checked", user.Nickname).And("State", "3")
	cond1 := cond.And("Updated__gt", t1).And("Updated__lte", t2).And("Drawn", user.Nickname).And("State", "1")
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)

	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...

	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()
	_, err = qs.Filter("Author", user.Username).Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, aa...)

	cond2 := cond.And("Updated__gt", t1).And("Updated__lte", t2).And("Designd", user.Nickname).And("State", "2")
	qs = qs.SetCond(cond2)
	_, err = qs.Filter("Author", user.Username).Distinct().All(&bb) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, bb...)

	cond3 := cond.And("Updated__gt", t1).And("Updated__lte", t2).And("Checked", user.Nickname).And("State", "3")
	qs = qs.SetCond(cond3)
	_, err = qs.Filter("Author", user.Username).Distinct().All(&cc) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	catalogs = append(catalogs, cc...)
	// _, err = qs.Filter("Designd", user.Nickname).All(&bb)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, bb...)
	// _, err = qs.Filter("Checked", user.Nickname).All(&cc)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, cc...)
	// _, err = qs.Filter("Examined", user.Nickname).All(&dd)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, dd...)
	return catalogs, err
}

//自己发起，已经提交
//author=登录人名，状态>登录名字所处位置，且状态<5
//设计和制图同一人会重复
func GetcatalogRunning(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	catalogs = make([]*Catalog, 0) //没有这句，model返回null，导致json输出给table为null，而不是[]，导致错误
	aa := make([]*Catalog, 0)
	// bb := make([]*Catalog, 0)
	// cc := make([]*Catalog, 0)
	// dd := make([]*Catalog, 0)

	cond := orm.NewCondition()
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).Or("Drawn", user.Nickname).And("State", "1").Or("Designd", user.Nickname).And("State", "2").Or("Checked", user.Nickname).And("State", "3")
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).And("Drawn", user.Nickname).And("State__gt", 1).And("State__lt", 5)
	cond1 := cond.And("Updated__gt", t1).And("Updated__lte", t2)
	cond2 := cond.And("Drawn", user.Nickname).And("State__gt", 1).And("State__lt", 5)
	cond3 := cond.And("Designd", user.Nickname).And("State__gt", 2).And("State__lt", 5)
	cond4 := cond.And("Checked", user.Nickname).And("State__gt", 3).And("State__lt", 5)
	cond5 := cond.AndCond(cond1).AndCond(cond.OrCond(cond2).OrCond(cond3).OrCond(cond4))
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond5)

	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...

	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()
	_, err = qs.Filter("Author", user.Username).Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	// catalogs = append(catalogs, aa...)

	// cond2 := cond.And("Date__gte", t1).And("Date__lt", t2).And("Designd", user.Nickname).And("State__gt", 2).And("State__lt", 5)
	// qs = qs.SetCond(cond2)
	// _, err = qs.Filter("Author", user.Username).Distinct().All(&bb) //qs.Filter("Drawn", user.Nickname).All(&aa)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, bb...)

	// cond3 := cond.And("Date__gte", t1).And("Date__lt", t2).And("Checked", user.Nickname).And("State__gt", 3).And("State__lt", 5)
	// qs = qs.SetCond(cond3)
	// _, err = qs.Filter("Author", user.Username).Distinct().All(&cc) //qs.Filter("Drawn", user.Nickname).All(&aa)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, cc...)
	// _, err = qs.Filter("Designd", user.Nickname).All(&bb)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, bb...)
	// _, err = qs.Filter("Checked", user.Nickname).All(&cc)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, cc...)
	// _, err = qs.Filter("Examined", user.Nickname).All(&dd)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, dd...)
	return aa, err
}

//已经完成
//制图、设计、校核、审查中含有登录名字，状态为5
func GetcatalogCompleted(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	aa := make([]*Catalog, 0)
	// bb := make([]*Catalog, 0)
	// cc := make([]*Catalog, 0)
	// dd := make([]*Catalog, 0)

	cond := orm.NewCondition()
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).Or("Drawn", user.Nickname).And("State", "1").Or("Designd", user.Nickname).And("State", "2").Or("Checked", user.Nickname).And("State", "3")
	cond1 := cond.And("Date__gt", t1).And("Date__lte", t2).And("State", "5")
	cond2 := cond.AndCond(cond1).AndCond(cond.Or("Drawn", user.Nickname).Or("Designd", user.Nickname).Or("Checked", user.Nickname).Or("Examined", user.Nickname))
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond2)

	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)

	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...

	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()

	_, err = qs.Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	// catalogs = append(catalogs, aa...)

	// cond2 := cond.And("Date__gte", t1).And("Date__lt", t2).And("Designd", user.Nickname).And("State", "5")
	// qs = qs.SetCond(cond2)
	// _, err = qs.Distinct().All(&bb) //qs.Filter("Drawn", user.Nickname).All(&aa)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, bb...)

	// cond3 := cond.And("Date__gte", t1).And("Date__lt", t2).And("Checked", user.Nickname).And("State", "5")
	// qs = qs.SetCond(cond3)
	// _, err = qs.Distinct().All(&cc) //qs.Filter("Drawn", user.Nickname).All(&aa)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, cc...)

	// cond4 := cond.And("Date__gte", t1).And("Date__lt", t2).And("Examined", user.Nickname).And("State", "5")
	// qs = qs.SetCond(cond4)
	// _, err = qs.Distinct().All(&dd) //qs.Filter("Drawn", user.Nickname).All(&aa)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, dd...)
	// _, err = qs.Filter("Designd", user.Nickname).All(&bb)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, bb...)
	// _, err = qs.Filter("Checked", user.Nickname).All(&cc)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, cc...)
	// _, err = qs.Filter("Examined", user.Nickname).All(&dd)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogs = append(catalogs, dd...)
	return aa, err
}

//别人传来，自己处于设计位置的展示
//author！=登录名,状态为2，设计名=登录名
func GetcatalogDesignd(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	aa := make([]*Catalog, 0)
	cond := orm.NewCondition()
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).Or("Drawn", user.Nickname).And("State", "1").Or("Designd", user.Nickname).And("State", "2").Or("Checked", user.Nickname).And("State", "3")
	cond1 := cond.And("Updated__gt", t1).And("Updated__lte", t2).And("Designd", user.Nickname).And("State", "2").AndNot("Author", user.Username)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)
	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...
	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()
	_, err = qs.Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	return aa, err
}

//别人传来，自己处于校核位置的展示
//author！=登录名,状态为3，校核名=登录名
func GetcatalogChecked(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	aa := make([]*Catalog, 0)
	cond := orm.NewCondition()
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).Or("Drawn", user.Nickname).And("State", "1").Or("Designd", user.Nickname).And("State", "2").Or("Checked", user.Nickname).And("State", "3")
	cond1 := cond.And("Updated__gt", t1).And("Updated__lte", t2).And("Checked", user.Nickname).And("State", "3").AndNot("Author", user.Username)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)
	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...
	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()
	_, err = qs.Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	return aa, err
}

//别人传来，自己处于审查位置的展示
//author！=登录名,状态为4，审查名=登录名
func GetcatalogExamined(id string, t1, t2 time.Time) (catalogs []*Catalog, err error) {
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	user := GetUserByUserId(Id)
	aa := make([]*Catalog, 0)
	cond := orm.NewCondition()
	// cond1 := cond.And("Date__gte", t1).And("Date__lt", t2).Or("Drawn", user.Nickname).And("State", "1").Or("Designd", user.Nickname).And("State", "2").Or("Checked", user.Nickname).And("State", "3")
	cond1 := cond.And("Updated__gt", t1).And("Updated__lte", t2).And("Examined", user.Nickname).And("State", "4").AndNot("Author", user.Username)
	o := orm.NewOrm()
	qs := o.QueryTable("catalog")
	qs = qs.SetCond(cond1)
	// 	cond := NewCondition()
	// 	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	// 	qs := orm.QueryTable("user")
	// 	qs = qs.SetCond(cond1)
	// 	// WHERE ... AND ... AND NOT ... OR ...
	// 	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
	// 	qs = qs.SetCond(cond2).Count()
	// 	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
	// qs.Distinct()
	_, err = qs.Distinct().All(&aa) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return nil, err
	}
	return aa, err
}

//四舍五入
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
