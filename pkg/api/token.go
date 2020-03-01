package api

import (
	"damingerdai/address/internal/models"
	service "damingerdai/address/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {
	value := reflect.ValueOf(obj)
	knd := value.Kind()
	if knd == reflect.Ptr {
		vPtr := value
		if value.IsNil() {
			vPtr = reflect.New(value.Type().Elem())
		}
		value.Set(vPtr)
		log.Printf("vPtr: %v", vPtr)
	}
	typ := reflect.TypeOf(obj)
	headers := req.Header
	log.Printf("typ.NumField: %d", typ.NumField)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		log.Printf("%v", field)
		tag := field.Tag
		log.Printf("%v", tag)
		tagName := tag.Get("header")
		log.Printf("%v", tagName)
		if len(tagName) != 0 {
			v := value.Field(i)
			// v = v.Elem()
			if v.CanSet() {
				switch v.Kind() {
				// TODO support other type
				case reflect.String: v.SetString(headers.Get(tagName))
				}
			}
		}
	}
	log.Printf("obj: %v", obj)
	return nil
}


func CreateToken(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")

	if len(username) == 0 {
		c.AbortWithStatusJSON(400, "username is required")
		return
	}

	if len(password) == 0 {
		c.AbortWithStatusJSON(400, "password is required")
		return
	}

	user := models.User{}
	user.Username = username
	user.Password = password

	service.CreateToken(c, &user)
	//secretKey := []byte("damingeridai")
	//expiresAt := time.Now().Add(time.Second * 5).Unix()
	//token, err := utils.CreateToken(&user, secretKey, expiresAt)
	//if err != nil {
	//	c.AbortWithStatusJSON(500, err.Error())
	//} else {
	//	c.AbortWithStatusJSON(200, token)
	//}
}
