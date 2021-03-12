package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"books"
)

var _ = Describe("Books", func() { // Describe描述一种行为或者一个方法
	var (
		longBook  books.Book
		shortBook books.Book
		book      *books.Book
		err       error
	)

	BeforeEach(func() { // BeforeEach会在每个小于等于beforeEach嵌套层的it函数之前运行，来设置公用的数据变量
		longBook = books.Book{
			Title:  "les miserables",
			Author: "victor hugo",
			Pages:  1488,
		}
		shortBook = books.Book{
			Title:  "Fox in socks",
			Author: "dr. seuss",
			Pages:  24,
		}
	})
	// JustBeforeEach 会在所有BeforeEach执行之后运行，在每个小于等于JustBeforeEach嵌套层的It函数之前运行，可以有效避免重复创建
	JustAfterEach(func() {
		book, err = books.NewBookFromJSON(`
			"title": "Le",
			"author": "victor",
			"pages": 1234
		`)
	})

	Describe("Categorizing book length", func() { // Describe描述一种行为或者一个方法
		Context("With more than 300 pages", func() { // Context 丰富Describe所描述的行为或方法，增加条件语句，尽可能全面覆盖各种condition
			It("should be a novel", func() { // It 申明用例，期望这个用例得到的结果
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL")) // Expect 这个是gomega提供的方法，用来断言结果
				Expect(longBook.Author).To(Equal("victor hugo"))
			})

		})
		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
	// 当嵌套Describe和Context块时，It执行时，围绕It的所有容器节点的BeforeEach块，从最外层到最内层运行.
	// 注意：每个It块都运行BeforeEach和AfterEach块，这确保每个spec的原始状态
	Describe("loading from JSON", func() {
		Context("when the json fails to parse", func() {
			BeforeEach(func() {
				book, err = books.NewBookFromJSON(`
					"title": "Le",
					"author": "victor",
					"pages": 1234aaa
				`)
			})
			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
