package main

import "fmt"

const (
	january = 1 + iota
	february
	march
	april
	may
)

func main() {
	fmt.Println(january, february, march, april, may)

}
