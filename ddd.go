package main

import "fmt"

func main() {
	type QueryProjectTagModel struct {
		TagData string
		TagId   string
		Created string
	}

	list2 := []*QueryProjectTagModel{}

	list2 = append(list2,&QueryProjectTagModel{ 
		TagData: "123123",
		TagId:   "12321",
		Created: "12312312",
	})
	list2 = append(list2,&QueryProjectTagModel{
		TagData: "fdgdf",
		TagId:   "12dfgdf321",
		Created: "asda",
	})
	list2 = append(list2,&QueryProjectTagModel{
		TagData: "qwewq",
		TagId:   "12qweqw321",
		Created: "rtetre",
	})

	list1 := []*QueryProjectTagModel{}

	list1 = append(list1,&QueryProjectTagModel{
		TagData: "fghfg",
		TagId:   "as",
		Created: "hhj",
	})
	list1 = append(list1,&QueryProjectTagModel{
		TagData: "kjk",
		TagId:   "kjjk",
		Created: "jkjk",
	})
	list1 = append(list1,&QueryProjectTagModel{
		TagData: "wqw",
		TagId:   "wqqw",
		Created: "rtetrewqqw",
	})


	list1 = append(list1, list2...)


	fmt.Println(len(list1))
	fmt.Println(len(list2))

	for _, value := range list1 {
		fmt.Println(value)
	}
}





