package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStuff(t *testing.T) {
	Convey("Truth", t, func() {
		Convey("is falsey", func() {
			So(true, ShouldBeFalse)
		})
	})

}
