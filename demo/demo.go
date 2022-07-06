package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		log.Println(name, "starts reading waiting")
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
	log.Println(name, "starts finish")
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func doCond() {

	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	// write("writer", cond)

	time.Sleep(time.Second * 3)

}

func main2() {
	// imagine this is a C array `int arr[] = {0, 1, 2, 3, 4, 5}`
	arr := []int{0, 1, 2, 3, 4, 5}

	// get arr's reflect value.
	// note that we need to get address of first element in array.
	// just like what we do in C `&arr[0]`.
	arrValue := reflect.ValueOf(&arr[0])

	// now we can have an unsafe pointer points to start of array.
	arrPtr := arrValue.Pointer()

	// get arr's element type. we'll need it later.
	arrElemType := reflect.TypeOf(arr).Elem()

	// we can do arbitrary pointer calculation just like C.
	// suppose we want to write following C code.
	//     int *ptr = &arr[0];
	//     size_t offset = 2;
	//     int *valuePtr = ptr + offset;
	//     int value = *valuePtr; // same as arr[offset]
	// here is the equivalent go code.
	offset := uintptr(2)
	ptr := unsafe.Pointer(arrPtr + arrElemType.Size()*offset)
	valuePtr := reflect.NewAt(arrElemType, ptr)
	value := valuePtr.Elem()

	// see what we get. we should get 2, which is the value of arr[2].
	fmt.Println("New value is:", value.Int())

	// happy hacking!
}

func Clone[T any](data *T) T {
	temp := *data
	return temp

}

type S struct {
	Name string
}

func change(s *S) {
	s.Name = "yes"
}

func main() {

	s := &S{"name"}
	change(s)
	fmt.Println(s.Name)

	d := &S{"name"}
	e := Clone(d)
	change(&e)
	fmt.Println(d.Name)
}

func main_ptr_offset() {
	a := []int{2, 5}

	var b reflect.Value = reflect.ValueOf(&a)

	b = b.Elem()

	fmt.Println("Slice:", a)

	// use of Append method

	b = reflect.Append(b, reflect.ValueOf(80))

	fmt.Println("Slice after appending data:", b)

}

func b() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
}

func DoAny[E any](v E) {

	fmt.Printf("%T", v)
}
