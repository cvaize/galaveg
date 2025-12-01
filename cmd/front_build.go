package cmd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

// frontBuildCmd represents the frontBuild command
var frontBuildCmd = &cobra.Command{
	Use:   "front:build",
	Short: "Build, minify, gzip: css, js, svg",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new minifier and add rules for CSS and JS
		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		m.AddFunc("application/javascript", js.Minify)

		// Run the processing
		processStyles(m)
		processScripts(m)
		processSVG()

		fmt.Println("\n--- Build complete! ---")
	},
}

func init() {
	rootCmd.AddCommand(frontBuildCmd)
}

// findFiles recursively finds all files with a given extension in the root directory.
func findFiles(root, ext string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ext) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// minifyAndGzip minifies a file and creates its gzipped version.
func minifyAndGzip(m *minify.M, mimeType, inputPath, outputPath string) {
	// Read the source file
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", inputPath, err)
	}

	// Minify the data
	minifiedData, err := m.Bytes(mimeType, data)
	if err != nil {
		log.Fatalf("Error minifying %s: %v", inputPath, err)
	}

	// Write the minified version
	err = os.WriteFile(outputPath, minifiedData, 0644)
	if err != nil {
		log.Fatalf("Error writing minified file %s: %v", outputPath, err)
	}
	fmt.Printf("Minified: %s\n", outputPath)

	// Create the gzipped version
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(minifiedData)
	if err != nil {
		log.Fatalf("Error writing to gzip buffer for %s: %v", outputPath, err)
	}
	err = w.Close()
	if err != nil {
		log.Fatalf("Error closing gzip writer for %s: %v", outputPath, err)
	}

	// Write the .gz file
	gzPath := outputPath + ".gz"
	err = os.WriteFile(gzPath, b.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Error writing gzipped file %s: %v", gzPath, err)
	}
	fmt.Printf("Created gzipped file: %s\n", gzPath)
}

// processStyles handles CSS files.
func processStyles(m *minify.M) {
	fmt.Println("\n--- Processing Styles ---")
	const publicDir = "public/css"
	const outputFile = "app.css"
	const outputMinFile = "app.min.css"

	// Create directory if it doesn't exist
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		log.Fatalf("Could not create directory %s: %v", publicDir, err)
	}

	// Find all CSS files
	libFiles, err := findFiles("./resources/css/libraries", ".css")
	if err != nil {
		log.Fatalf("Error finding files in libraries: %v", err)
	}

	compFiles, err := findFiles("./resources/css/components", ".css")
	if err != nil {
		log.Fatalf("Error finding files in components: %v", err)
	}

	// Add layout styles before components
	compFiles = lo.Filter(compFiles, func(s string, _ int) bool {
		return s != "resources/css/components/layout.css"
	})

	// Combine lists, similar to the original
	styles := append(libFiles, "resources/css/components/layout.css")
	styles = append(styles, compFiles...)

	// Concatenate all files into one
	var builder strings.Builder
	for _, style := range styles {
		content, err := os.ReadFile(style)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", style, err)
		}
		builder.Write(content)
	}

	// Write the concatenated file
	concatenatedPath := filepath.Join(publicDir, outputFile)
	err = os.WriteFile(concatenatedPath, []byte(builder.String()), 0644)
	if err != nil {
		log.Fatalf("Error writing concatenated CSS file: %v", err)
	}
	fmt.Printf("Created concatenated file: %s\n", concatenatedPath)

	// Minify and create a .gz version
	minifiedPath := filepath.Join(publicDir, outputMinFile)
	minifyAndGzip(m, "text/css", concatenatedPath, minifiedPath)
}

// processScripts handles JS files.
func processScripts(m *minify.M) {
	fmt.Println("\n--- Processing Scripts ---")
	const publicDir = "public/js"
	const outputFile = "app.js"
	const outputMinFile = "app.min.js"

	// Create directory
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		log.Fatalf("Could not create directory %s: %v", publicDir, err)
	}

	// Find all JS files
	libFiles, err := findFiles("./resources/js/libraries", ".js")
	if err != nil {
		log.Fatalf("Error finding files in libraries: %v", err)
	}
	compFiles, err := findFiles("./resources/js/components", ".js")
	if err != nil {
		log.Fatalf("Error finding files in components: %v", err)
	}

	scripts := append(libFiles, compFiles...)

	// Concatenate all files
	var builder strings.Builder
	for _, script := range scripts {
		content, err := os.ReadFile(script)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", script, err)
		}
		builder.Write(content)
	}

	// Write the concatenated file
	concatenatedPath := filepath.Join(publicDir, outputFile)
	err = os.WriteFile(concatenatedPath, []byte(builder.String()), 0644)
	if err != nil {
		log.Fatalf("Error writing concatenated JS file: %v", err)
	}
	fmt.Printf("Created concatenated file: %s\n", concatenatedPath)

	// Minify and create a .gz version
	minifiedPath := filepath.Join(publicDir, outputMinFile)
	minifyAndGzip(m, "application/javascript", concatenatedPath, minifiedPath)
}

// processSVG handles SVG files.
func processSVG() {
	fmt.Println("\n--- Processing SVG ---")
	const publicDir = "public/svg"

	// Create directory
	if err := os.MkdirAll(publicDir, 0755); err != nil {
		log.Fatalf("Could not create directory %s: %v", publicDir, err)
	}

	// Find all SVG files
	svgFiles, err := findFiles("./resources/svg", ".svg")
	if err != nil {
		log.Fatalf("Error finding SVG files: %v", err)
	}

	for _, svgElement := range svgFiles {
		// Read the source file
		content, err := os.ReadFile(svgElement)
		if err != nil {
			log.Fatalf("Error reading SVG file %s: %v", svgElement, err)
		}

		// Determine the path for saving
		fileName := filepath.Base(svgElement)
		publicPath := filepath.Join(publicDir, fileName)

		// Copy the file to public/svg
		err = os.WriteFile(publicPath, content, 0644)
		if err != nil {
			log.Fatalf("Error copying SVG file %s: %v", publicPath, err)
		}
		fmt.Printf("Copied: %s\n", publicPath)

		// Create a gzipped version
		var b bytes.Buffer
		w, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(content)
		if err != nil {
			log.Fatalf("Error writing to gzip buffer for %s: %v", publicPath, err)
		}
		w.Close()

		gzPath := publicPath + ".gz"
		err = os.WriteFile(gzPath, b.Bytes(), 0644)
		if err != nil {
			log.Fatalf("Error writing gzipped SVG file %s: %v", gzPath, err)
		}
		fmt.Printf("Created gzipped file: %s\n", gzPath)
	}
}

func strSliceFilter(v []string) []string {
	return lo.Filter(lo.Map(v, func(s string, _ int) string {
		return strings.TrimSpace(s)
	}), func(s string, _ int) bool {
		return s != ""
	})
}
