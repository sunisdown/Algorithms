package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack(t *testing.T) {
	Convey("创建一个栈", t, func() {
		s := NewArrayStack()
		println(s.size(), "xxx")
		So(s.size(), ShouldEqual, 0)
		Convey("添加数据", func() {
			s.Push(3)
			println(s.list, "xxx")
			So(s.size(), ShouldEqual, 1)
			svalue, err := s.Pop()
			So(err, ShouldEqual, nil)
			So(svalue, ShouldEqual, 3)
			Convey("添加数据", func() {
				s.Push(0)
				s.Push(1)
				s.Push(2)
				s.Push(3)
				s.Push(4)
				Convey("测试满了之后添加数据", func() {
					err := s.Push(5)
					So(err, ShouldNotEqual, nil)
				})
				Convey("读取数据", func() {
					svalue, err := s.Pop()
					So(err, ShouldEqual, nil)
					So(svalue, ShouldEqual, 4)
					svalue, err = s.Pop()
					So(err, ShouldEqual, nil)
					So(svalue, ShouldEqual, 3)
					svalue, err = s.Pop()
					So(err, ShouldEqual, nil)
					So(svalue, ShouldEqual, 2)
					svalue, err = s.Pop()
					So(err, ShouldEqual, nil)
					So(svalue, ShouldEqual, 1)
					svalue, err = s.Pop()
					So(err, ShouldEqual, nil)
					So(svalue, ShouldEqual, 0)
					Convey("读取一个empty stack", func() {
						svalue, err = s.Pop()
						So(err, ShouldNotEqual, nil)
						So(svalue, ShouldEqual, 0)
					})

				})
			})
		})
	})
}
