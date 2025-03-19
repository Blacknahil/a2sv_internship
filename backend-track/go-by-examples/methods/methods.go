package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	fmt.Printf("areas arress %p inside area not a copy\n", r)
	r.height = 10
	r.width = 10
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	fmt.Printf("rectange address= %p inside perimeter copy\n", &r)
	r.height = 5
	r.width = 5
	return 2 * (r.width + r.height)
}

func main() {

	r := rectangle{width: 3, height: 4}
	fmt.Printf("r address= %p\n", &r)
	fmt.Println("perimeter=", r.perimeter())
	fmt.Println("check per change", "w=", r.width, " h=", r.height)
	fmt.Println("area=", r.area())
	fmt.Println("check area change", "w=", r.width, " h=", r.height)

	var new_r rectangle
	new_r = r
	fmt.Printf("r address = %p\n", &r)
	fmt.Printf("new r address = %p\n", &new_r)

}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdfjkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdjsdjhjksdjk
//kgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// kdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jdkfjjkfsdjjsfdkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjkshjjhfjhfsjjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjsdjhjksdjk
//jkdfbjkgkdfbjdf djbfjksdbfjsjsdjhjksdjk
//jkdfbjkgkdfbjdfjkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdjsdjhjksdjk
// /jsdjhjksdjk
// //jkdfbjkgjsdjksdjk
// jsdjhjksdjk
//jkdfbjkjhjfsdhjhjdsfsdjhjksdjk
// ksdjkfjksdjkfjkksjdjkkjjks
