package main

import "fmt"

type NewIDSubscriber struct {
	newId string
	//TODO : maybe new
	//counter int
}

func (s NewIDSubscriber) EmailAuthor(msg string) {
	fmt.Println(msg)
	fmt.Println(" Emailing the author of post id: " + s.newId + " that someone commented with :\n")
}

func (s NewIDSubscriber) EmailOtherCommentators(msg string) {
	fmt.Println(msg)
	fmt.Println(" Emailing all other comment authors who commented on " + s.newId + " that someone commented with :\n")
}

func (s NewIDSubscriber) IncrementCommentCount(msg string) {
	fmt.Println(msg)
	fmt.Println(" Updating comment count to + 1 for blog post id: " + s.newId + "\n")
}
