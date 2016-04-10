package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestArray(t *testing.T) {
	Convey("创建一个环形队列", t, func() { // 数据表驱动
		q := NewArrayQueueA()
		So(q.head, ShouldEqual, 0)
		Convey("添加数据", func() {
			q.Enqueue(3)
			devalue, err := q.Dequeue() // 命名
			So(err, ShouldEqual, nil)
			So(devalue, ShouldEqual, 3)
			Convey("添加数据", func() {
				q.Enqueue(0)
				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)
				q.Enqueue(4) // 下一个空行
				Convey("测试队列满了之后添加数据", func() {
					err := q.Enqueue(5)
					So(err, ShouldNotEqual, nil)
				})
				Convey("读取数据", func() {
					devalue, err := q.Dequeue()
					So(err, ShouldEqual, nil)
					So(devalue, ShouldEqual, 0)
					devalue, err = q.Dequeue()
					So(err, ShouldEqual, nil)
					So(devalue, ShouldEqual, 1)
					devalue, err = q.Dequeue()
					So(err, ShouldEqual, nil)
					So(devalue, ShouldEqual, 2)
					devalue, err = q.Dequeue()
					So(err, ShouldEqual, nil)
					So(devalue, ShouldEqual, 3)
					devalue, err = q.Dequeue()
					So(err, ShouldEqual, nil)
					So(devalue, ShouldEqual, 4)
					Convey("读取一个空队列", func() {
						devalue, err = q.Dequeue()
						So(err, ShouldNotEqual, nil)
						So(devalue, ShouldEqual, 0)
					})

				})
			})
		})
	}) // 深度
}
