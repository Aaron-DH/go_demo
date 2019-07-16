package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

type I interface {
	notify()
}

func inotify(i I) {
	i.notify()
}

func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	inotify(bill)
	inotify(lisa)

	bill.changeEmail("bill@newmail.com")
	bill.notify()

	lisa.changeEmail("lisa@newmail.com")
	lisa.notify()
}
