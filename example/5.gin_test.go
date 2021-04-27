package main

import (
	"log"
	"strconv"
	"time"

	"github.com/yezihack/e"

	"github.com/gin-gonic/gin"
)

// curl -X GET http://localhost:8080/home/2
// curl -X GET http://localhost:8080/home/2a

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "1.0.0")
	})
	r.GET("/home/:id", Home)
	r.Run(":8080")
}

func Home(c *gin.Context) {
	id := c.Param("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		Err(c, 400, e.New("非数字类型"), err)
		return
	}
	result, err := Info(_id)
	if err != nil {
		Err(c, 500, err)
		return
	}
	OK(c, result)
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "ok",
		"data":   data,
	})
}
func Err(ctx *gin.Context, status int, err error, extra ...error) {
	// 对 err 进行判断是否是自定义错误
	h := gin.H{
		"code": status,
	}
	sysCode := 200
	if err != nil {
		if e.Assert(err) {
			if e.Convert(err).Code() > 0 {
				h["code"] = e.Convert(err).Code()
			}
			h["msg"] = e.Convert(err).Msg()
			log.Println(e.Convert(err).Errs())
		} else {
			h["code"] = status
			h["msg"] = err.Error()
		}
	}
	if len(extra) > 0 {
		err = extra[0]
		if err != nil {
			if e.Assert(err) {
				log.Println(e.Convert(err).ToStr())
				log.Println(e.Convert(err).Errs())
			} else {
				log.Println(err)
			}
		}
	}
	ctx.AbortWithStatusJSON(sysCode, h)
}

type InfoResponse struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
}

func Info(id int) (result *InfoResponse, err error) {
	switch id {
	case 1:
		result = &InfoResponse{
			ID:     1,
			Name:   "github.com/yezihack/e",
			Author: "百里江山",
			Date:   time.Now(),
		}
	case 2:
		result = &InfoResponse{
			ID:     2,
			Name:   "github.com/yezihack/algo",
			Author: "空树之空",
			Date:   time.Now(),
		}
	default:
		err = e.NewCode(400, "ID不存在", e.ErrorF("id:%d", id))
		return
	}
	return result, nil
}
