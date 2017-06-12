package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/mlctrez/javaclassparser"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func failErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetOutput(os.Stdout)

	start := time.Now()
	if false {
		archive := "/some/path/to/somejar.ear"
		// archives within archives are supported
		rc, err := zip.OpenReader(archive)
		failErr(err)
		read(archive, rc)
	}

	if true {

		oc, err := exec.Command("javac", "-d", "java", "java/example/Sample.java").CombinedOutput()
		if err != nil {
			fmt.Println(string(oc))

			panic(err)
		}

		jcp := &javaclassparser.ClassParser{}
		r, err := os.Open("java/example/Sample.class")
		failErr(err)
		jcp.Parse(r)
		jcp.DebugOut()
	}
	log.Println(time.Since(start).Seconds())

}

func read(path string, rc *zip.ReadCloser) {
	for _, f := range rc.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rp, err := f.Open()
		failErr(err)

		if strings.HasSuffix(f.Name, ".jar") {

			tf, err := ioutil.TempFile(os.TempDir(), "jcp")
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(tf, rp)
			if err != nil {
				panic(err)
			}
			jarReader, err := zip.OpenReader(tf.Name())
			if err != nil {
				panic(err)
			}
			read(path+"!"+f.Name, jarReader)
			jarReader.Close()
			os.Remove(tf.Name())
		}

		if strings.HasSuffix(f.Name, ".class") {

			bb := &bytes.Buffer{}
			bc, err := io.Copy(bb, rp)
			if err != nil {
				log.Fatal(err)
			}

			if bc != int64(f.UncompressedSize64) {
				fmt.Println(bc, f.UncompressedSize64, f.Name, err)
				log.Fatal("unable to read entire file")
			}

			jcp := &javaclassparser.ClassParser{}
			jcp.Parse(bytes.NewReader(bb.Bytes()))

		}
		rp.Close()

	}

}

// (Lcom/erac/admin/presentation/user/CreateDemographicInfoForm;Lcom/erac/custmgmt/user/ProfiledWebUser;Lcom/erac/custmgmt/user/ProfiledWebUser;ZLcom/erac/arch/inf/common/exception/ArchExceptionContainer;)Lcom/erac/custmgmt/user/ProfiledWebUser;
// (Lcom/erac/admin/presentation/user/CreateDemographicInfoForm;Lcom/erac/custmgmt/user/ProfiledWebUser;Lcom/erac/custmgmt/user/ProfiledWebUser;ZLcom/erac/arch/inf/common/exception/ArchExceptionContainer;)LcomXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
