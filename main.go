package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"strconv"
	"time"
)

type List struct {
	gorm.Model
	Name    string `json:"name" gorm:"column:name"`
	State   string `json:"state" gorm:"column:state"`
	Phone   string `json:"phone" gorm:"column:phone"`
	Email   string `json:"email" gorm:"column:email"`
	Address string `json:"address" gorm:"column:address"`
}

/*
	业务码：
	正确：200,
	错误：400
*/

func get(r *gin.Engine) {
	// get请求
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "请求成功",
		})
	})
	r.Run(":3301")
}

func create(r *gin.Engine, db *gorm.DB) {
	// 增
	r.POST("/user/add", func(c *gin.Context) {
		var data List
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{
				"msg":  "添加失败",
				"data": gin.H{},
				"code": 400,
			})
		} else {
			//数据库操作
			db.Create(&data)
			c.JSON(200, gin.H{
				"msg":  "添加成功",
				"data": data,
				"code": 200,
			})
		}
	})
	r.Run(":3301")
}

func delete(r *gin.Engine, db *gorm.DB) {
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		var data []List
		id := c.Param("id")
		db.Where("id = ?", id).Find(&data)
		if len(data) == 0 {
			c.JSON(400, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": "400",
			})
		} else {
			db.Delete(&data)
			c.JSON(200, gin.H{
				"msg":  "删除成功",
				"code": "200",
			})
		}
	})
	r.Run(":3301")
}

func update(r *gin.Engine, db *gorm.DB) {
	r.PUT("/user/update/:id", func(c *gin.Context) {
		var find []List
		id := c.Param("id")
		db.Where("id = ?", id).Find(&find)
		if len(find) == 0 {
			c.JSON(400, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": "400",
			})
		} else {
			var data List
			err := c.ShouldBindJSON(&data)
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{
					"msg":  "修改失败",
					"code": 400,
				})
			} else {
				db.Where("id = ?", id).Updates(data)
				c.JSON(200, gin.H{
					"msg":  "修改成功",
					"code": 200,
				})
			}
		}
	})
	r.Run(":3301")
}

func read(r *gin.Engine, db *gorm.DB) {
	r.GET("/user/list/:name", func(c *gin.Context) {
		name := c.Param("name")
		var data []List
		db.Where("name = ?", name).Find(&data)
		if len(data) == 0 {
			c.JSON(400, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "查询到数据",
				"code": 200,
				"data": data,
			})
		}

	})
	r.Run(":3301")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
		//c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		fmt.Println(c.GetHeader("Origin"))
		// 必须，设置服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
		// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
		c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
		// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func crud(r *gin.Engine, db *gorm.DB) {
	r.POST("/user/add", func(c *gin.Context) {
		var data List
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(400, gin.H{
				"msg":  "添加失败",
				"data": gin.H{},
				"code": 400,
			})
		} else {
			//数据库操作
			db.Create(&data)
			c.JSON(200, gin.H{
				"msg":  "添加成功",
				"data": data,
				"code": 200,
			})
		}
	})
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		var data []List
		id := c.Param("id")
		db.Where("id = ?", id).Find(&data)
		if len(data) == 0 {
			c.JSON(400, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": "400",
			})
		} else {
			db.Delete(&data)
			c.JSON(200, gin.H{
				"msg":  "删除成功",
				"code": "200",
			})
		}
	})
	r.PUT("/user/update/:id", func(c *gin.Context) {
		var find []List
		id := c.Param("id")
		db.Where("id = ?", id).Find(&find)
		if len(find) == 0 {
			c.JSON(400, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": "400",
			})
		} else {
			var data List
			err := c.ShouldBindJSON(&data)
			if err != nil {
				fmt.Println(err)
				c.JSON(200, gin.H{
					"msg":  "修改失败",
					"code": 400,
				})
			} else {
				db.Where("id = ?", id).Updates(data)
				c.JSON(200, gin.H{
					"msg":  "修改成功",
					"code": 200,
				})
			}
		}
	})
	r.GET("/user/list/:name", func(c *gin.Context) {
		name := c.Param("name")
		var data []List
		db.Where("name = ?", name).Find(&data)
		if len(data) == 0 {
			c.JSON(400, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "查询到数据",
				"code": 200,
				"data": data,
			})
		}

	})
	r.GET("/user/list", func(c *gin.Context) {
		var datalist []List
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		pageNum, _ := strconv.Atoi(c.Query("pageNum"))
		var total int64
		db.Model(datalist).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&datalist)

		if len(datalist) == 0 {
			c.JSON(400, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "查询成功",
				"code": 200,
				"data": gin.H{
					"list":     datalist,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
			})
		}
	})
	r.Run(":3301")

}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println("gorm.Open err", err)
	}
	fmt.Println(db)

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	db.AutoMigrate(&List{})
	r := gin.Default()
	r.Use(Cors())

	// get(r)
	// create(r, db)
	// delete(r, db)
	// update(r, db)
	// read(r, db)
	crud(r, db)
}
