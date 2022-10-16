package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	//blendFile := "test.blend"

	Blender.Custom = "C:/Users/zivno/Blender/stable/blender-3.2.0-windows-x64/blender.exe"

	if !Blender.Exists() {
		fmt.Println("install blender")
		return
	}

	//dir := renderFrames(blendFile)
	dir := "/mnt/C/Users/zivno/AppData/Local/Temp/test-randomStrinHere"
	convertPngsToGifs(dir)
}
func renderFrames(blendFile string) string {
	filename := strings.TrimSuffix(blendFile, path.Ext(blendFile))
	randString := "randomStrinHere"
	tempDir := path.Join(os.TempDir(), fmt.Sprintf("%s-%s", path.Base(filename), randString))
	//fmt.Println(tempDir)

	_ = os.Mkdir(tempDir, os.ModeDir)
	err := Blender.Run("-b", blendFile, "-o", path.Join(tempDir, "frame_####"), "-F", "PNG", "-a") //,"&>/dev/null")
	if err != nil {
		fmt.Println(err)
	}

	return tempDir
}

func convertPngsToGifs(directory string) {
	gifPath := path.Join(directory, "gifs")
	imgs, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = os.Mkdir(gifPath, os.ModeDir)

	imagick.Initialize()
	defer imagick.Terminate()

	ret, err := imagick.ConvertImageCommand([]string{
		"mogrify", "-format", "gif", path.Join(directory, "*.png"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Metadata:\n%s\n", ret.Meta)

	//for _, image := range imgs {
	//	if image.IsDir() {
	//		continue
	//	}
	//
	//f, err := os.Open(path.Join(directory, image.Name()))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//img, err := png.Decode(bufio.NewReader(f))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//f.Close()
	//buf := new(bytes.Buffer)
	//err = gif.Encode(buf, img, nil)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//name := strings.TrimSuffix(image.Name(), path.Ext(image.Name()))
	//err = os.WriteFile(path.Join(gifPath, name+".gif"), buf.Bytes(), image.Mode())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//}
}
