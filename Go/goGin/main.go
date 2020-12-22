package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Login login request binding
type Login struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	// skip validate
	Email string `form:"email" json:"email" xml:"email" binding:"-"`
	// custom validator
	CheckIn time.Time `form:"check_in" json:"check_in" xml:"check_in" binding:"required,checkInValidator" time_foramt:"2006-01-02 15:04:05"`
}

// validator
var checkInValidator validator.Func = func(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(value) {
			return false
		}
	}
	return true
}

// Index request query binding
type Index struct {
	Page int32 `form:"page" json:"page" binding:"-"`
}

// Logger custom middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// 不带中间件的空白路由
	// router := gin.New()
	// 使用中间件
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())

	// log file
	logFile, _ := os.Create("log_file.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout) // 同时写入文件和控制台

	// with default log and recovery middleware
	router := gin.Default()

	// 注册validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("checkInValidator", checkInValidator)
	}
	// 自定义日志格式
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))

	// static files server
	router.Static("/templates", "./templates")
	router.StaticFS("/staticFs", http.Dir("./"))
	router.StaticFile("/ginmd", "./gin.md")

	// file data server
	router.GET("/local/file/ginmd", func(c *gin.Context) {
		c.File("./gin.md")
	})
	var fs http.FileSystem = http.Dir("./")
	router.GET("/fs/file/log", func(c *gin.Context) {
		c.FileFromFS("./log_file.log", fs)
	})

	// json response
	router.GET("/", func(c *gin.Context) {
		var index Index
		// 等价于 c.BindQuery(&index)
		if c.ShouldBindWith(&index, binding.Query) == nil {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		}
	})

	// model binding and validation
	router.POST("/login", func(c *gin.Context) {
		var login Login
		// Bind* If there is a binding error, the request is aborted with c.AbortWithError(400, err).SetType(ErrorTypeBind)
		// if err:=c.BindJSON(&login);err!=nil{}
		// ShouldBind*  If there is a binding error, the error is returned and the developer can handle the error appropriately.
		// ShouldBind(&login) 不指定特定类型的Bind时，会根据tag进行解析绑定
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.Username != "张三" || login.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": c.Request.URL.User, "status": "login success."})
	})

	// 分组路由
	user := router.Group("/user")
	{

		// path parameter :name 必须参数 *action可选参数，该参数存在时包含反斜杠/
		user.GET("/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			fmt.Printf("full path: %s\n", c.FullPath())
			c.String(http.StatusOK, "the name is %s, action is %s", name, action)
		})

		// query parameter
		user.GET("/", func(c *gin.Context) {
			// 不存在时取默认值
			firstName := c.DefaultQuery("firstname", "张三")
			// 是 c.Request.URL.Query().Get("lastName")的语法糖
			lastName := c.Query("lastname")
			c.JSON(http.StatusOK, gin.H{
				"firstname": firstName,
				"lastname":  lastName,
			})
		})

		// Multipart/Urlencoded Form
		user.POST("/", func(c *gin.Context) {
			message := c.PostForm("message")
			nick := c.DefaultPostForm("nick", "anonymous")
			c.JSON(200, gin.H{
				"status":  "posted",
				"message": message,
				"nick":    nick,
			})
		})
	}

	// PostFormMap and QueryMap
	// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	// Content-Type: application/x-www-form-urlencoded
	// names[first]=thinkerou&names[second]=tianou
	router.POST("/post", func(c *gin.Context) {
		queryMap := c.QueryMap("ids")
		postMap := c.PostFormMap("names")
		fmt.Println("queryMap:", queryMap, "postMap:", postMap)
	})

	// upload file
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file") // "file":form key
		fmt.Println(file.Filename)
		c.SaveUploadedFile(file, "./files/")
		// multi files
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			fmt.Println(file.Filename)
			c.SaveUploadedFile(file, "./files/")
		}
		c.String(http.StatusOK, fmt.Sprintf("'%d' files uploaded!", len(files)))
	})

	// secure json 防止json劫持
	router.GET("/secureJson", func(c *gin.Context) {
		colors := []string{"red", "blue", "green"}
		// 返回结果：["red","blue","green"]
		c.JSON(http.StatusOK, colors)
		// 返回结果： while(1);["red","blue","green"]
		// c.SecureJSON(http.StatusOK, colors)
	})

	router.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{"html": "<b>Hello World.</b>"})
		// c.JSON(200, gin.H{"html": "<b>Hello World.</b>"})
	})

	// render template
	// router.LoadHTMLGlob("templates/*") // 一级相对目录
	// router.LoadHTMLGlob("templates/**/*") //两级相对目录
	router.LoadHTMLFiles("./templates/index.html")
	router.GET("/tmpl", func(c *gin.Context) {
		// redirect
		c.Redirect(http.StatusFound, "/index")
		// redirect use handleCtx
		// c.Request.URL.Path = "/redirect"
		// router.HandleContext(c)
	})
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": " Gin-Go web framework", "colors": []string{"red", "blue", "green"}})
	})

	// middleware
	router.GET("/middleware", Logger(), func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// it would print: "12345"
		log.Println(example)
	})

	// basic-auth middleware
	secrets := gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	//Goroutines inside a middleware
	//inside a middleware or handler, SHOULD NOT use the original context instead use a read-only copy.
	router.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	

	// default port 8080
	router.Run(":8080")
}
