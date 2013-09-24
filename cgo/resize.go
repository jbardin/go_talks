package main

import (
	"fmt"
	"os/exec"
)

/*
#cgo pkg-config: GraphicsMagick
#include <magick/api.h>
*/
import "C"

var (
	imageInfo     *C.ImageInfo
	exceptionInfo C.ExceptionInfo
)

func Read(path string) (*C.Image, error) {
	strCopy(&imageInfo.filename, []byte(path))

	img := C.ReadImage(imageInfo, &exceptionInfo)
	if img == nil {
		C.CatchException(&img.exception)
		return nil, fmt.Errorf(C.GoString(exceptionInfo.reason))
	}

	return img, nil
}

func Resize(img *C.Image, x, y int) (*C.Image, error) {
	newImg := C.ResizeImage(img, C.ulong(x), C.ulong(y), C.LanczosFilter, 1.0, &exceptionInfo)

	C.CatchException(&exceptionInfo)
	if newImg == nil {
		C.CatchException(&img.exception)
		return nil, fmt.Errorf(C.GoString(exceptionInfo.reason))
	}

	C.DestroyImage(img)
	return newImg, nil
}

func Save(img *C.Image, path string) error {
	strCopy(&img.filename, []byte(path))

	if C.WriteImage(imageInfo, img) == 0 {
		C.CatchException(&img.exception)
		return fmt.Errorf(C.GoString(img.exception.reason))
	}
	return nil
}

// strCopy OMIT
// fake strncpy so we don't need to allocate a new byte slice
// Careful, we're not checking the length of src!
func strCopy(dest *[C.MaxTextExtent]C.char, src []byte) {
	for i, c := range src {
		dest[i] = C.char(c)
	}
	// This is C, we need to terminate the string!
	dest[len(src)] = 0
}

func main() {
	// Some GraphicsMagick boilerplate
	C.InitializeMagick(nil)
	defer C.DestroyMagick()
	C.GetExceptionInfo(&exceptionInfo)
	imageInfo = C.CloneImageInfo(nil)
	defer C.DestroyExceptionInfo(&exceptionInfo)
	defer C.DestroyImageInfo(imageInfo)

	// main OMIT

	Show("./cgo/armed_gopher.jpg")

	img, err := Read("./cgo/armed_gopher.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	if img, err = Resize(img, 200, 800); err != nil {
		fmt.Println("resize error:", err)
		return
	}

	if err = Save(img, "./cgo/resized_gopher.jpg"); err != nil {
		fmt.Println("save error:", err)
		return
	}

	Show("./cgo/resized_gopher.jpg")
}

func Show(path string) {
	fmt.Println("opening", path)
	cmd := exec.Command("/usr/bin/open", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(output))
}
