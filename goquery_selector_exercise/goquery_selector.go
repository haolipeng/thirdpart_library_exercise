package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func main() {
	//fmt.Println("begin call funtion pre_next_selector")
	//pre_next_selector()

	//fmt.Println("begin call funtion first_child_filter")
	//first_child_filter()

	//fmt.Println("begin call funtion first_of_type_filter")
	//first_of_type_filter()

	/*
		fmt.Println("begin call funtion element_selector")
		element_selector()

		fmt.Println("begin call funtion element_id_selector")
		element_id_selector()

		fmt.Println("begin call funtion class_selector")
		class_selector()

		fmt.Println("begin call funtion parent_child_selector")
		parent_child_selector()
	*/

	fmt.Println("begin call funtion element_class_selector")
	element_class_selector()
}

//封装goquery函数
func goguery_selector(htmlContent string, selectorContent string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent)) //io.Reader
	if err != nil {
		log.Fatal(err)
	}

	//selectorContent内容就是goquery的语法
	dom.Find(selectorContent).Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

//直接以element元素作为选择器，例子中直接查找div
func element_selector() {
	htmlContent := `
			<body>

				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>	
			`
	selectorContent := "div"
	goguery_selector(htmlContent, selectorContent)
}

//element id selector选择器
//语法规则：element#id
func element_id_selector() {
	htmlContent := `<body>

				<div id="div_one">you select is div element, and id is div_one</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`
	selectorContent := "div#div_one"
	goguery_selector(htmlContent, selectorContent)
}

//class 选择器 (单独查找class名称属性的)
//class 是很常用的html元素
//语法规则：.class_name
func class_selector() {
	htmlContent := `<body>
				<div>haolipeng</div>
				<div class="abc">codergeek</div>
				<div class="name">hello world</div>
				<span class="name">hello world span</span>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`

	//筛选出class为name的div元素
	selectorContent := ".name"
	goguery_selector(htmlContent, selectorContent)

	//筛选出class为abc的div元素
	selectorContent = "div[class=abc]"
	goguery_selector(htmlContent, selectorContent)
}

//element Class 选择器
func element_class_selector() {
	htmlContent := `<body>

				<div>DIV1</div>
				<div class="name name1 name2">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	//div[class]筛选出Element为div并且有class这个属性的
	//selectorContent := "div[class]"
	//goguery_selector(htmlContent,selectorContent)

	selectorContent := "div[class=name1]"
	goguery_selector(htmlContent, selectorContent)
}

func parent_child_selector() {
	//筛选出DIV1，DIV2，DIV3，但并没有筛选出DIV4，因为它不是子集，而是子集的子集
	htmlContent := `<body>

				<div lang="ZH">DIV1</div>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<div lang="japan">DIV4</div>
				<span>
					<div>DIV4</div>
				</span>

			</body>
			`
	//值筛选子集的
	selectorContent := "body>div"
	goguery_selector(htmlContent, selectorContent)

	//筛选父元素下所有符合条件的元素
	//selectorContent = "body~div"
	//goguery_selector(htmlContent,selectorContent)
}

//pre-next selector选择器
func pre_next_selector() {
	htmlContent := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div>DIV4</div>
				</span>
				<p>P2</p>

			</body>
			`
	//筛选当前元素的下一个元素
	selectorContent := "div[lang=zh]+p"
	goguery_selector(htmlContent, selectorContent)

	//只需要把+号换成~号,就可以把P2也筛选出来，因为P2、P1和DIV1都是兄弟。
	selectorContent = "div[lang=zh]~p"
	goguery_selector(htmlContent, selectorContent)
}

//first-child 过滤器
func first_child_filter() {
	htmlContent := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<div style="display:none;">DIV4</div>
					<div>DIV5</div>
				</span>
				<p>P2</p>
				<div></div>

			</body>
			`
	//筛选出的元素是他们的父元素的第一次子元素，如果不是，则不会被筛选出来
	selectorContent := "div:first-child"
	goguery_selector(htmlContent, selectorContent)
}

//first of type filter
func first_of_type_filter() {
	htmlContent := `<body>

				<div lang="zh">DIV1</div>
				<p>P1</p>
				<div lang="zh-cn">DIV2</div>
				<div lang="en">DIV3</div>
				<span>
					<p>P2</p>
					<div>DIV6</div>
					<div>DIV5</div>
				</span>
				<div></div>

			</body>
			`

	//如果还是用:first-child,<div>DIV5</div>是不会被筛选出来的，因为其不是父元素的第一个子元素
	//这时候我们使用:first-of-type就可以达到目的，因为它要求是同类型第一个就可以
	selectorContent := "div:first-of-type"
	goguery_selector(htmlContent, selectorContent)
}
