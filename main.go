package main

import "listfiles/operations"

func main() {
	operations.ListFile("/Users/bariskantar/playground/")
	operations.CopyFile("/Users/bariskantar/playground/test.rtf", "/Users/bariskantar/playground/test2.rtf")
	operations.MoveFile("/Users/bariskantar/playground/test/test2.rtf", "/Users/bariskantar/playground/test.rtf")
	operations.DeleteFile("/Users/bariskantar/playground/test2.rtf")

}
