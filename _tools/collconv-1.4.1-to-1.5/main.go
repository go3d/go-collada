package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"

	conv "github.com/go3d/go-collada/conv-1.4.1-to-1.5"
	xmlx "github.com/go-forks/xmlx"
	ugo "github.com/metaleap/go-util"
	uio "github.com/metaleap/go-util/io"
)

const ns = "http://www.collada.org/2005/11/COLLADASchema"

func convert(srcFilePath, dstFilePath string) {
	var (
		err     error
		data    []byte
		outFile *os.File
	)
	log.Printf("SRC: %s\n", srcFilePath)
	data = uio.ReadBinaryFile(srcFilePath, true)
	if outFile, err = os.Create(dstFilePath); err != nil {
		panic(err)
	}
	defer outFile.Close()
	if data, err = conv.ConvertBytes(data); err != nil {
		panic(err)
	}
	log.Printf("\tOUT: %s\n", dstFilePath)
	outFile.Write(data)
}

func main() {
	var (
		flagSrcFilePath = flag.String("src", "", "File path to the Collada 1.4.1 source document to be loaded.")
		flagDstFilePath = flag.String("dst", "", "File path to the Collada 1.5 destination document to be written.")
	)
	runtime.LockOSThread()
	xmlx.IndentPrefix = "\t"
	flag.Parse()
	if (len(*flagSrcFilePath) > 0) && (len(*flagDstFilePath) > 0) {
		convert(*flagSrcFilePath, *flagDstFilePath)
	} else {
		const dbp = "Dropbox/oga/collada"
		for _, baseDirPath := range []string{filepath.Join("Q:", dbp), filepath.Join(ugo.UserHomeDirPath(), dbp)} {
			if uio.DirExists(baseDirPath) {
				for _, subDirName := range []string{"cube-poly", "cube-tris", "duck-poly", "duck-tris", "mgmidget", "bikexsi", "diningroom", "berlin", "sponza"} {
					convert(filepath.Join(baseDirPath, subDirName, "obj.dae"), filepath.Join(baseDirPath, subDirName, "obj15.dae"))
				}
			}
		}
	}
}
