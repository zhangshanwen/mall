package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
	"github.zhangshanwen.mall/config"
	"github.zhangshanwen.mall/model"
	"github.zhangshanwen.mall/route"
	"os"
	"strconv"
)

func init() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	if config.DBConfig.SQLLog {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)

	model.DB = db
}

func main() {
	app := iris.New()

	route.Route(app)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errNo": model.ErrorCode.NotFound,
			"msg":   "Not Found",
			"data":  iris.Map{},
		})

	})

	app.OnErrorCode(500, func(ctx iris.Context) {
		ctx.JSON(
			iris.Map{
				"errNo": model.ErrorCode.ERROR,
				"msg":   "error",
				"data":  iris.Map{},
			})
	})

	app.Listen(":" + strconv.Itoa(config.ServerConfig.Port))
}
