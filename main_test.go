package main

import (
	"fmt"
	"testing"
)

func TestDupe(t *testing.T) {
	var names1 []string = []string{"tom", "bill", "bob", "tom"}
	var names2 []string = []string{"tom", "bill", "bob", "tom", "bill", "bob"}
	var names3 []string = []string{"bill", "bob"}

	if dupe("tom", names1) == false {
		fmt.Println("Not noticing a dupe tom in" + fmt.Sprintf("%s", names1))
		t.Fail()
	}
	if dupe("tom", names2) == false {
		fmt.Println("Not noticing a dupe tom in" + fmt.Sprintf("%s", names2))
		t.Fail()
	}
	if dupe("tom", names3) == true {
		fmt.Println("shouldn't be a dupe tom" + fmt.Sprintf("%s", names3))
		t.Fail()
	}
}

func TestFirst(t *testing.T) {
	var names1 []string = []string{"tom", "bill", "bob", "tom"}
	var names2 []string = []string{"tom", "bill", "bob", "tom", "bill", "bob"}
	var names3 []string = []string{"bill", "bob"}

	if first(names1) != "tom" {
		fmt.Println("first of " + fmt.Sprintf("%s", names1) + " should be tom")
		t.Fail()
	}

	if first(names2) != "tom" {
		fmt.Println("first of " + fmt.Sprintf("%s", names1) + " should be tom")
		t.Fail()
	}
	if first(names3) != "bill" {
		fmt.Println("first of " + fmt.Sprintf("%s", names1) + " should be tom")
		t.Fail()
	}
}

func TestRest(t *testing.T) {
	var names1 []string = []string{"tom", "bill", "bob", "tom"}
	var names2 []string = []string{"tom", "bill", "bob", "tom", "bill", "bob"}
	var names3 []string = []string{"bill", "bob"}

	if len(rest(names1)) != 3 {
		fmt.Println("len of rest should be 3, rest of " + fmt.Sprintf("%s", rest(names1)) + " should be bill bob tom")
		t.Fail()
	}
	if len(rest(names2)) != 5 {
		fmt.Println("rest of " + fmt.Sprintf("%s", names2) + " should be tom")
		t.Fail()
	}
	if len(rest(names3)) != 1 {
		fmt.Println("rest of " + fmt.Sprintf("%s", names3) + " should be tom")
		t.Fail()
	}

}

func TestFilter(t *testing.T) {

	var names1 []string = []string{"tom", "bill", "bob", "tom"}
	var names2 []string = []string{"tom", "bill", "bob", "tom", "bill", "bob"}
	var names3 []string = []string{"bill", "jerry", "tom", "tom", "bob", "bob"}

	teststr := "tom"
	testfilter := func(v string) bool {
		if v == teststr {
			return false
		} else {
			return true
		}
	}

	filter(testfilter, names1)
	filter(testfilter, names2)
	filter(testfilter, names3)
}

func TestDedup(t *testing.T) {
	var names1 []string = []string{"tom", "bill", "bob", "tom"}
	var names2 []string = []string{"tom", "bill", "bob", "tom", "bill", "bob"}
	var names3 []string = []string{"tom", "bill", "bob", "tom", "bob"}
	var names4 []string = []string{"tom", "bill", "bob", "tom", "bill"}
	var names5 []string = []string{"tom", "bill", "jerry", "bob", "tom", "bill", "bob", "tom", "tom"}
	if "[bill bob]" != fmt.Sprintf("%s", dedup(names1)) {
		fmt.Println("names 1: [bill bob] != ", fmt.Sprintf("%s", dedup(names1)))
		//		t.Fail()
	}
	if "[]" != fmt.Sprintf("%s", dedup(names2)) {
		fmt.Println("names 2: [] != ", fmt.Sprintf("%s", dedup(names2)))
		t.Fail()
	}
	if "[bill]" != fmt.Sprintf("%s", dedup(names3)) {
		fmt.Println("names 3: [bill] != ", fmt.Sprintf("%s", dedup(names3)))
		t.Fail()
	}
	if "[bob]" != fmt.Sprintf("%s", dedup(names4)) {
		fmt.Println("names 4: [bob] !=", fmt.Sprintf("%s", dedup(names4)))
		t.Fail()
	}
	if "[jerry]" != fmt.Sprintf("%s", dedup(names5)) {
		fmt.Println("names 5: [jerry] != ", fmt.Sprintf("%s", dedup(names5)))
		t.Fail()
	}
}
