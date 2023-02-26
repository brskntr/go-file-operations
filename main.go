package main

import "listfiles/copyfile"

func main() {
	copyfile.ListFile("/Users/bariskantar/playground/")
	copyfile.CopyFile("/Users/bariskantar/playground/test.rtf", "/Users/bariskantar/playground/test2.rtf")
	copyfile.MoveFile("/Users/bariskantar/playground/test/test2.rtf", "/Users/bariskantar/playground/test.rtf")
	copyfile.DeleteFile("/Users/bariskantar/playground/test2.rtf")

}
