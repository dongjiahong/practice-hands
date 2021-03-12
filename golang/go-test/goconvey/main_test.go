package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("Start test new", t, func() {
		stu, err := NewStudent(0, "")
		Convey("have error", func() {
			So(err, ShouldBeError)
		})
		Convey("stu is nil", func() {
			So(stu, ShouldBeNil)
		})
	})
}

func TestScore(t *testing.T) {
	stu, _ := NewStudent(1, "test")
	Convey("if error", t, func() {
		_, err := stu.GetAve()
		Convey("have error", func() {
			So(err, ShouldBeError)
		})
	})

	Convey("normal", t, func() {
		stu.Math = 60
		stu.Chinaese = 70
		stu.English = 80
		score, err := stu.GetAve()
		Convey("have error", func() {
			So(err, ShouldBeNil)
		})
		Convey("score > 60", func() {
			So(score, ShouldBeGreaterThan, 60)
		})
	})
}
