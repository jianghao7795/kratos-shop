package data_test

import (
	"time"
	"user/internal/biz"
	"user/internal/data"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User
	BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil)
		birthday := time.Unix(693629981, 0)
		uD = &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: &birthday,
		}
	})

	It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		// 组装的数据 mobile 为 13803881388
		Ω(u.Mobile).Should(Equal("13803881388")) // 手机号应该为创建的时候写入的手机号
	})
})
