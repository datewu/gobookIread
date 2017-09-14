// “In queuing theory, there is a law that — with enough sampling — predicts the throughput of your pipeline. It’s called Little’s Law, and you only need to know a few things to understand and make use of it.
// Let’s first define Little’s Law algebraicly. It is commonly expressed as: L=λW, where:
// L = the average number of units in the system.
// λ = the average arrival rate of units.
// W = the average time a unit spends in the system.”
//
// 摘录来自: Katherine Cox-Buday. “Concurrency in Go: Tools and Techniques for Developers”。 iBooks.
package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkUnbufferedWriter(b *testing.B) {
	performWrite(b, temFileOrFatal())
}

func BenchmarkBufferedWriter(b *testing.B) {
	bufferedFile := bufio.NewWriter(temFileOrFatal())
	performWrite(b, bufferedFile)
}

func temFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()

	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}

}

// “Let’s try and determine how many requests per second our pipeline can handle. Let’s assume we enable sampling on our pipeline and find that 1 request (r) takes about 1 second to make it through the pipeline. Let’s plug in those numbers!
// 3r = λr/s * 1s
// 3r/s = λr/s
// λr/s = 3r/s
// We set L to 3 because each stage in our pipeline is processing a request. We then set W to 1 second, do a little algebra, and voilà! In this pipeline, we can handle three requests per second.
// What about determining how large our queue needs to be to handle a desired number of requests. Can Little’s Law help us answer that?
// Let’s say our sampling indicates that a request takes 1 ms to process. What size would our queue have to be to handle 100,000 requests per second? Again, let’s plug in the numbers!
// Lr-3r = 100,000r/s * 0.0001s
// Lr-3r = 10r
// Lr = 7r
// Again, our pipeline has three stages, so we’ll decrement L by 3. We set λ to 100,000 r/s, and find that if we want to field that many requests, our queue should have a capacity of 7. Remember that as you increase the queue size, it takes your work longer to make it through the system! You’re effectively trading system utilization for lag.
// Something that Little’s Law can’t provide insight on is handling failure. Keep in mind that if for some reason your pipeline panics, you’ll lose all the requests in your queue. This might be something to guard against if re-creating the requests is difficult or won’t happen. To mitigate this, you can either stick to a queue size of zero, or you can move to a persistent queue, which is simply a queue that is persisted somewhere that can be later read from should the need arise.”

// “Queuing can be useful in your system, but because of its complexity, it’s usually one of the last optimizations I would suggest implementing.”
//
// 摘录来自: Katherine Cox-Buday. “Concurrency in Go: Tools and Techniques for Developers”。 iBooks.
