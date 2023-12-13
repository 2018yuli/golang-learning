package main

import (
	"bufio"
	"fmt"
	"learning/channel/c28sort-pipeline-outer/pipeline"
	"os"
)

func createPipeline(filename string, fileSize, chunkCount int) (<-chan int, []*os.File) {
	chunkSize := fileSize / chunkCount

	sortResults := []<-chan int{}
	files := []*os.File{}
	for i := 0; i < chunkCount; i++ {
		// 此处需要 return []*File 等待归并结束后再 close
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		files = append(files, file)
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemorySort(source))
	}
	return pipeline.MergeN(sortResults...), files
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	// defer 栈中，flush 比 file.Close 先执行
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}

func main() {
	p, files := createPipeline("small.in", 400, 4)
	// TODO files 需要再 pipeline 结束后统一 close
	writeToFile(p, "small.out")
	printFile("small.out")
	for _, file := range files {
		defer file.Close()
	}
}
