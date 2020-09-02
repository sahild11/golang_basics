package main

import (
	"context"
	"flag"
	"fmt"
	"regexp"

	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/go/pkg/beam/runners/direct"
	"github.com/apache/beam/sdks/go/pkg/beam/transforms/stats"

	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/local"
)

var (
	input  = flag.String("input", "data/*", "File(s) to read.")
	output = flag.String("output", "outputs/wordcounts.txt", "Output filename.")
)

var wordRE = regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)

func main() {
	flag.Parse()
	fmt.Printf("Reading files from %v\n", *input)
	fmt.Printf("Writing files to %v\n", *output)


	beam.Init()

	pipeline := beam.NewPipeline()
	root := pipeline.Root()

	// Read lines from a text file.
	// lines := textio.Read(root, *input)
	lines := textio.Read(root, *input)

	// Use a regular expression to iterate over all words in the line.
	words := beam.ParDo(root, func(line string, emit func(string)) {
		for _, word := range wordRE.FindAllString(line, -1) {
			emit(word)
		}
	}, lines)

	// Count each unique word.
	counted := stats.Count(root, words)

	// Format the results into a string so we can write them to a file.
	formatted := beam.ParDo(root, func(word string, count int) string {
		return fmt.Sprintf("%s: %v", word, count)
	}, counted)

	// Finally, write the results to a file.
	textio.Write(root, *output, formatted)

	// We have to explicitly run the pipeline, otherwise it's only a definition.
	direct.Execute(context.Background(), pipeline)
}
