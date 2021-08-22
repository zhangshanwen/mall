
* **修改configuration.json**  
进入wemall目录，将`configuration.dev.json`改名为`configuration.json`, 再找到`UploadImgDir`，填写图片上传目录为 `{wemall所在目录}/upload/img`, 举例：C:\dev\src\wemall\upload\img (Windows系统）或 /dev/src/wemall/upload/img (Mac OS X 或Linux系统)
  
```
{
  "go": {
    "UploadImgDir": "" /*图片上传的目录*/
  }
}
```  
* **修改config.js**  
进入`wemall/weixin/config`目录，将`config.dev.js`改名为`config.js`

$ go run main.go
```


### 后台
* web框架: iris
* 路由: httprouter
* 持久层框架: gorm
* 数据库: mysql 

### go依赖的第三方库

| 库 | 说明              |
|:---------|:-----------------------|
| github.com/kataras/iris/v12   | iris web框架   |
| github.com/jinzhu/gorm     | gorm 持久层框架 |
| github.com/satori/go.uuid  | uuid生成工具    |
