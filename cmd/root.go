package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/levinion/hulu/parse"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:  "hulu",
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		//解析文件内容到result
		filename := args[0]
		result, err := parse.Parse(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		//若传入第二个参数，忽略所有标志，拷贝解析的内容
		if len(args) == 2 {
			dir := args[1]
			basename := filepath.Base(filename)
			filename := filepath.Join(dir, basename)
			err := os.WriteFile(filename, []byte(result), os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}

		if overwrite {
			err := os.WriteFile(filename, []byte(result), os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else if copy {
			ext := filepath.Ext(filename)
			noext := strings.TrimRight(filename, ext)
			fmt.Println()
			var newFilename = ""
			i := 1
			for {
				newFilename = fmt.Sprintf(noext+"%d"+ext, i)
				_, err := os.Stat(newFilename)
				if err == nil {
					i++
				} else {
					break
				}
			}
			err := os.WriteFile(newFilename, []byte(result), os.ModePerm)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(result)
		}
	},
}

var overwrite = false
var copy = false

func Execute() {
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite", "w", false, "overwrite document")
	rootCmd.Flags().BoolVarP(&copy, "copy", "c", false, "just copy it")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
