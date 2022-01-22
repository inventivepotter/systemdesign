package inheritence

import "fmt"

type sharable interface {
	getLink(medium string) string
}

type actionable interface {
	getAllowedReactions() []string
}

// Inheriting sharable, actionable interfaces
type solvable interface {
	sharable
	actionable
	isResolved() bool
}

type Post struct {
	id string
}

type Question struct {
	Post
	resolved bool
}

type Answer struct {
	Post
}

type Reply struct {
	Post
}

func (q Question) getLink(medium string) string {
	return "https://" + medium + "/" + q.id
}

func (q Question) getAllowedReactions() []string {
	return []string{"love", "gotcha"}
}

func (q Question) isResolved() bool {
	return q.resolved
}

func (a Answer) getLink(medium string) string {
	return "https://" + medium + "/" + a.id
}

func (a Answer) getAllowedReactions() []string {
	return []string{"love", "gotcha"}
}

func main() {
	// q is implementing solvable interface which demostrates interface inheritence
	var q solvable
	q = Question{
		Post: Post{
			id: "xxx",
		},
		resolved: false,
	}
	fmt.Println(q.getLink("linkedin"))   // returns https://linkedin/xxx
	fmt.Println(q.getAllowedReactions()) // returns [love gotcha]
	fmt.Println(q.isResolved())          // returns false

	// Following doesn't work because Answer doesn't have method isResolved() thus can't implement solvable
	// var a solvable
	// a = Answer{
	// 	Post: Post{
	// 		id: "yyy",
	// 	},
	// }
	var a1 sharable
	a1 = Answer{
		Post: Post{
			id: "yyy1",
		},
	}
	fmt.Println(a1.getLink("linkedin"))
	// Following gives an error because sharabled doesn't have method getAllowedReactions althou answer has it.
	// fmt.Println(a1.getAllowedReactions())		// retruns Errors
	var a2 actionable
	a2 = Answer{
		Post: Post{
			id: "yyy2",
		},
	}
	// Following gives an error because sharabled doesn't have method getLink althou answer has it.
	// fmt.Println(a2.getLink("linkedin"))			// retruns Errors
	fmt.Println(a2.getAllowedReactions())
}
