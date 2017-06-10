package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/iftekhersunny/api_testing_with_goconvey/http"
	"github.com/iftekhersunny/api_testing_with_goconvey/models"
	"github.com/iftekhersunny/api_testing_with_goconvey/router"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUsersApiRoutes(t *testing.T) {
	Convey("Given a HTTP request for /v1/users", t, func() {
		jsonByte, _ := json.Marshal(models.User{"iftekhersunny", 27})

		req := httptest.NewRequest("POST", "/v1/users", bytes.NewReader(jsonByte))
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			r := router.NewRouter()
			r = http.RegisterRoutes(r)

			r.ServeHTTP(resp, req)

			Convey("Then the response should be 200 with message 'User has been created'", func() {
				data := map[string]string{}
				json.Unmarshal(resp.Body.Bytes(), &data)

				So(resp.Code, ShouldEqual, 200)
				So(data["message"], ShouldEqual, "User has been created")
			})
		})
	})

	Convey("Given a HTTP request for /v1/users", t, func() {

		user := new(models.User)
		user.Clear()
		user.Create("iftekhersunny", 27)

		req := httptest.NewRequest("GET", "/v1/users", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			r := router.NewRouter()
			r = http.RegisterRoutes(r)

			r.ServeHTTP(resp, req)

			Convey("Then the response must return users list", func() {
				data := map[string]models.Users{}

				json.Unmarshal(resp.Body.Bytes(), &data)

				So(resp.Code, ShouldEqual, 200)

				So(len(data["data"]), ShouldEqual, 1)
				So(data["data"][0].Name, ShouldEqual, "iftekhersunny")
				So(data["data"][0].Age, ShouldEqual, 27)
			})
		})
	})
}
