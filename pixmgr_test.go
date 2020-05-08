package main

import (
	"fmt"
	"testing"
)

func TestT1(t *testing.T) {
	want := make(map[string]int)
	want["p01.jpeg_9411"] = 2
	want["p02.jpeg_8024"] = 2
	want["p03.jpeg_7899"] = 3
	want["p04.jpeg_7479"] = 3
	want["p05.jpeg_7073"] = 3
	want["p06.jpeg_5524"] = 3
	want["p07.jpeg_4337"] = 2
	want["p08.jpeg_3331"] = 2
	want["p09.jpeg_2466"] = 2
	want["p10.jpeg_3461"] = 2
	want["p11.jpeg_9792"] = 2
	want["p12.jpeg_6264"] = 3
	want["p13.jpeg_4337"] = 3
	want["p14.jpeg_6743"] = 3
	want["p15.jpeg_5358"] = 2

	//dirPath := "testdata"
	dirPath := "/home/nurali/code/src"
	Start(dirPath)

	for key, val := range want {
		if val != FileDitto[key] {
			t.Errorf(fmt.Sprintf("for key=%s, want=%d, got=%d", key, val, FileDitto[key]))
			t.FailNow()
		}
	}
}