package program

import (
	"github.com/spf13/cobra"
	"github.com/xrash/gofgen/pkg/gen"
	"strings"
)

type RootCommand struct {
	program *Program

	options struct {
		inputDir       string
		outputFile     string
		packageName    string
		compress       string
	}
}

func (rc *RootCommand) Run(cmd *cobra.Command, args []string) {

	inputDirname, err := handleInputDir(rc.options.inputDir)
	if err != nil {
		PrintlnErr("Error handling option input-dir", err)
		rc.program.Exit(1)
	}

	outputFilename, err := handleOutputFile(rc.options.outputFile)
	if err != nil {
		PrintlnErr("Error handling option output-file", err)
		rc.program.Exit(1)
	}

	packageName := handlePackageName(rc.options.packageName, inputDirname)

	compress, err := handleCompress(rc.options.compress)
	if err != nil {
		PrintlnErr("Error handling option compress", err)
		rc.program.Exit(1)
	}

	// Define useful variables.
	filesChan := make(chan *gen.File, 1024)

	// Initiate main structs.
	filesReader := &gen.FilesReader{
		RootDirname:       inputDirname,
		FilenamesToIgnore: make([]string, 0),
	}

	initWriter := &gen.InitWriter{
		PackageName:          packageName,
		CompressionAlgorithm: compress,
		OutputFilename:       outputFilename,
	}

	// If outfile is inside readdir, add outfile to ignorelist.
	if strings.HasPrefix(outputFilename, inputDirname) {
		relativePath := outputFilename[len(inputDirname):]
		filesReader.FilenamesToIgnore = append(filesReader.FilenamesToIgnore, relativePath)
	}

	// Run the reader/writer.
	errReader := filesReader.ReadFuture(filesChan)
	errWriter, writerDone := initWriter.WriteFuture(filesChan)

	// Wait for the result.
	select {
	case err := <-errReader:
		PrintlnErr("Reader error", inputDirname, err)
		rc.program.Exit(1)

	case err := <-errWriter:
		PrintlnErr("Writer error", outputFilename, err)
		rc.program.Exit(1)

	case <-writerDone:
		break
	}

	// Exit successfully.
	rc.program.Exit(0)
}

func (p *Program) createRootCmd() *cobra.Command {

	rc := &RootCommand{
		program: p,
	}

	cmd := &cobra.Command{
		Use:   "gofgen",
		Short: "gofgen generates go code loading your local files into memory at compile time",
		Long:  `gofgen generates go code loading your local files into memory at compile time`,
		Run:   rc.Run,
	}

	cmd.Flags().StringVarP(&rc.options.inputDir, "input-dir", "", ".",
		`dirname of the directory to read from`)

	cmd.Flags().StringVarP(&rc.options.outputFile, "output-file", "", "./init_gofgen.go",
		`filename of the output file`)

	cmd.Flags().StringVarP(&rc.options.packageName, "package-name", "", "",
		"package name to use, default value is basename of input-dirname")

	/*
	cmd.Flags().StringVarP(&rc.options.compress, "compress", "", "",
		"compression algorithm to use")
	 */

	return cmd
}
