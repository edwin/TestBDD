package controllers_test

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session"
	"net/http"
	"net/http/httptest"
	_ "testbdd/models"
	_ "testbdd/routers"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"strings"
)

/*
 *  @Author Muhammad Edwin (muhammad.edwin@ruma.co.id)
 *  @Date 19/12/2018 10:50
 *
 *
 *
 */

func TestLogin(t *testing.T) {
	r, _ := http.NewRequest("GET", "/login?username=edwin&password=password", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("GIVEN Edwin as a USER", t, func() {
		Convey("WHEN Edwin login with the correct username and password", func() {
			Convey("THEN HTTP Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("AND The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
			Convey("AND The Result Should Contains \"login success\"", func() {
				So(strings.Count(w.Body.String(), "login success"), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func TestLoginFailing(t *testing.T) {
	r, _ := http.NewRequest("GET", "/login?username=edwin&password=password123", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("GIVEN Edwin as a USER", t, func() {
		Convey("WHEN Edwin login with the WRONG username and password", func() {
			Convey("THEN HTTP Code Should Be 404", func() {
				So(w.Code, ShouldEqual, 404)
			})
			Convey("AND The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
			Convey("AND The Result Should Contains \"user not exist\"", func() {
				So(strings.Count(w.Body.String(), "user not exist"), ShouldBeGreaterThan, 0)
			})
		})
	})
}

func httpGetRequest(url string) (*http.Request, *httptest.ResponseRecorder) {
	r, _ := http.NewRequest("GET", url, nil)
	setHeader(r)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	return r, w
}

func setHeader(r *http.Request) {
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Accept", "*/*")
	r.Header.Set("Accept-Encoding", "gzip, deflate")
	r.Header.Set("Accept-Language", "en-US,en;q=0.9,id;q=0.8")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	r.Header.Set("X-Requested-With", "XMLHttpRequest")
	r.Header.Set("Connection", "keep-alive")
}
