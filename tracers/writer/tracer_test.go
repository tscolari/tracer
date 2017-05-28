package writer_test

import (
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/tscolari/tracer/tracers/writer"
)

var _ = Describe("Writer", func() {
	var (
		tracer *writer.Tracer
		buffer *gbytes.Buffer
		id     string
	)

	BeforeEach(func() {
		buffer = gbytes.NewBuffer()
		id = "id-1"
		tracer = writer.New(id, buffer)
	})

	It("writes the transaction time", func() {
		spanTracer := tracer.StartSpan("hello-world")
		time.Sleep(2 * time.Millisecond)
		Expect(spanTracer.End()).To(Succeed())

		Expect(buffer).To(gbytes.Say("id-1.hello-world:"))

		parts := strings.Split(string(buffer.Contents()), " ")
		Expect(parts).To(HaveLen(2))

		duration, err := strconv.Atoi(strings.Trim(parts[1], "\n"))
		Expect(err).NotTo(HaveOccurred())

		Expect(duration).To(BeNumerically("~", 2*time.Millisecond, time.Millisecond))
	})

	It("traces a sequence of events", func() {
		spanTracer1 := tracer.StartSpan("event1")
		spanTracer2 := spanTracer1.StartSpan("event2")
		spanTracer3 := spanTracer2.StartSpan("event3")

		spanTracer2.End()
		spanTracer3.End()
		spanTracer1.End()

		Expect(buffer).To(gbytes.Say("id-1.event1.event2:"))
		Expect(buffer).To(gbytes.Say("id-1.event1.event2.event3:"))
		Expect(buffer).To(gbytes.Say("id-1.event1:"))
	})
})
