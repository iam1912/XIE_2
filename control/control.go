package control

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iam1912/XIE_2/model"
)

var (
	students = model.NewStuSlice()
)

func LoginerHandler(c *gin.Context) {
	var user model.Accout
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	if user.Name == "xjh" && user.Password == "1900" {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func PostHandler(c *gin.Context) {
	change := c.PostForm("change")
	switch change {
	case "显示列表":
		ShowHandle(c)
	case "排序":
		SortHandle(c)
	case "查询":
		SearchHandle(c)
	case "添加":
		AddHandle(c)
	case "删除":
		DeleteHandle(c)
	case "更新":
		ModifyHandle(c)
	}
}

func ShowHandle(c *gin.Context) {
	stu := students.List()
	c.JSON(http.StatusOK, stu)
}

func SortHandle(c *gin.Context) {
	stu := students.Sort()
	c.JSON(http.StatusOK, stu)
}

func SearchHandle(c *gin.Context) {
	id := c.PostForm("id")
	getint, err := mathvaild(id)
	if err != nil {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := students.FindIndex(getint[0]); err != nil {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
	} else {
		stu := students.Search(getint[0])
		c.JSON(http.StatusOK, stu)
	}
}

func DeleteHandle(c *gin.Context) {
	id := c.PostForm("id")
	getint, err := mathvaild(id)
	if err != nil {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := students.FindIndex(getint[0]); err != nil {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
	} else {
		err := students.Delete(getint[0])
		if err != nil {
			checkError(err)
			c.JSON(http.StatusBadRequest, nil)
		} else {
			c.JSON(http.StatusOK, "true")
		}
	}
}

func ModifyHandle(c *gin.Context) {
	var stu model.Student
	if err := c.ShouldBind(&stu); err != nil {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := students.FindIndex(stu.ID); err == nil {
		id := fmt.Sprintf("%d", stu.ID)
		sex := fmt.Sprintf("%d", stu.Sex)
		socre := fmt.Sprintf("%d", stu.Socre)

		_, err := mathvaild(id, sex, socre)
		_, errs := zhvaild(stu.Name, stu.Major, stu.Birthday)

		if err == nil && errs == nil {
			students.Modify(stu.ID, stu.Name, stu.Major, stu.Sex,
				stu.Birthday, stu.Socre, "")
			c.JSON(http.StatusOK, "true")
		} else {
			checkError(err)
			c.JSON(http.StatusBadRequest, nil)
		}
	} else {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
	}
}

func AddHandle(c *gin.Context) {
	var stu model.Student
	if err := c.ShouldBind(&stu); err != nil {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err := bindFormstr(stu.Name, stu.Major, stu.Birthday)
	errs := bindFormint(stu.ID, stu.Sex, stu.Socre)
	if err == nil && errs == nil {
		stu1 := model.NewStu(stu.ID, stu.Name, stu.Major,
			stu.Sex, stu.Birthday, stu.Socre, "")
		err = students.Add(stu1)
		if err != nil {
			checkError(err)
			c.JSON(http.StatusBadRequest, nil)
		} else {
			c.JSON(http.StatusOK, "true")
		}
	} else {
		checkError(err)
		c.JSON(http.StatusBadRequest, nil)
	}
}

func bindFormstr(args ...string) error {
	for _, val := range args {
		if val != "" {
			_, err := zhvaild(val)
			if err != nil {
				log.Printf("Fatal err %v\n", err)
				return err
			}
		} else {
			return errors.New("添加的数据有错误!")
		}
	}
	return nil
}

func bindFormint(args ...int) error {
	for index, val := range args {
		if val != 0 || index == 1 {
			valstr := fmt.Sprintf("%d", val)
			_, err := mathvaild(valstr)
			if err != nil {
				log.Printf("Fatal Error %v\n", err)
				return err
			}
		} else {
			return errors.New("添加的数据有误")
		}
	}
	return nil
}

func mathvaild(args ...string) ([]int, error) {
	var intslice []int

	for _, val := range args {
		getint, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		} else {
			intslice = append(intslice, getint)
		}
	}
	return intslice, nil
}

func zhvaild(args ...string) ([]string, error) {
	var strslice []string

	for _, val := range args {
		m, err := regexp.MatchString("^\\p{Han}+$", val)
		if !m {
			return nil, err
		} else {
			strslice = append(strslice, val)
		}
	}
	return strslice, nil
}

func checkError(err error) {
	log.Printf("Fatal Error: %v\n", err)
}
