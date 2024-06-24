package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

func main() {
	// 创建一个新的Collector
	c := colly.NewCollector(colly.AllowedDomains("www.xbiquge.bz"))

	// 使用Colly框架中的OnHTML方法来指定一个CSS选择器和一个回调函数。
	// 当访问的HTML页面中出现了匹配该CSS选择器的元素时，就会触发这个回调函数。
	c.OnHTML("#hotcontent .l .item", func(e *colly.HTMLElement) {
		// 从当前元素的子元素中提取书名，子元素是通过CSS类名定位的。
		bookName := e.ChildText("div.image a") // 提取书名
		// 提取作者名，并使用strings.TrimSpace去除其前后的空格。
		authorName := strings.TrimSpace(e.ChildText("dt span")) // 提取并清理作者名
		// 提取图片URL，这里使用的是ChildAttr方法，它会提取元素属性的内容。
		picUrl := e.ChildAttr("div.image a img", "src") // 提取图片URL
		// 提取书籍描述，再次使用strings.TrimSpace去除其前后的空格。
		bookDesc := strings.TrimSpace(e.ChildText("dd")) // 提取并清理书籍描述
		// 提取最新章节名，使用strings.TrimSpace去除前后空格。
		lastChapterName := strings.TrimSpace(e.ChildText("dd.update a")) // 提取并清理最新章节名
		// 提取最新章节的URL。
		lastChapterURL := e.ChildAttr("dd.update a", "href") // 提取最新章节的URL

		// 打印提取的数据，这只是为了调试目的。
		fmt.Printf("书名: %s\n作者: %s\n封面URL: %s\n描述: %s\n最新章节: %s\n链接: %s\n\n",
			bookName, authorName, picUrl, bookDesc, lastChapterName, lastChapterURL)
		// 在这里可以调用一个函数将数据保存到数据库，但这个函数需要你自己实现。
	})

	// 设置Colly的配置，例如限速等
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*xbiquge.*",
		RandomDelay: 1 * time.Second, // 随机延迟
	})

	// 开始访问页面
	c.Visit("https://www.xbiquge.bz")
}
