package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const VERSIONSTART string = "v0.0.1\n"

func readPreviousVersion(f *os.File) (string, error) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), err
}

func incrementMajor(versionString string) (string, error) {
	s := strings.Split(versionString, ".")
	major := s[0]

	majorInt, err := strconv.Atoi(strings.TrimPrefix(major, "v"))
	if err != nil {
		return "", err
	}
	
	majorInt = majorInt + 1
	
	majorStr := strconv.Itoa(majorInt)
	s[0] = fmt.Sprintf("v%s", majorStr)
	
	return strings.Join(s, "."), nil
}

func decrementMajor(versionString string) (string, error) {
	s := strings.Split(versionString, ".")
	major := s[0]

	majorInt, err := strconv.Atoi(strings.TrimPrefix(major, "v"))
	if err != nil {
		return "", err
	}

	majorInt = majorInt - 1

	majorStr := strconv.Itoa(majorInt)
	s[0] = fmt.Sprintf("v%s", majorStr)

	return strings.Join(s, "."), nil
}

func incrementMinor(versionString string) (string, error) {
	s := strings.Split(versionString, ".")
	minor := s[1]

	minorInt, err := strconv.Atoi(minor)
	if err != nil {
		return "", err
	}

	minorInt = minorInt + 1

	minorStr := strconv.Itoa(minorInt)
	s[1] = minorStr

	return strings.Join(s, "."), nil
}

func decrementMinor(versionString string) (string, error) {
	s := strings.Split(versionString, ".")
	minor := s[1]

	minorInt, err := strconv.Atoi(minor)
	if err != nil {
		return "", err
	}

	minorInt = minorInt - 1

	minorStr := strconv.Itoa(minorInt)
	s[1] = minorStr

	return strings.Join(s, "."), nil
}

func incrementPatch(versionString string) (string, error) {
	s := strings.Split(versionString, ".")
	patch := s[2]

	patchInt, err := strconv.Atoi(strings.TrimSpace(patch))
	if err != nil {
		return "", err
	}

	patchInt = patchInt + 1

	patchStr := strconv.Itoa(patchInt)
	s[2] = patchStr

	return strings.Join(s, "."), nil
}

func decrementPatch(versionString string) (string, error) {
	s := strings.Split(versionString, ".")
	patch := s[2]

	patchInt, err := strconv.Atoi(strings.TrimSpace(patch))
	if err != nil {
		return "", err
	}

	patchInt = patchInt - 1

	patchStr := strconv.Itoa(patchInt)
	s[2] = patchStr

	return strings.Join(s, "."), nil
}

func main() {
	var (
		versionFile string
		MainCMD = &cobra.Command{
			Use: "version",
			Version: "0.0.1",
			Args: cobra.NoArgs,
			Short: "A simple method of automated Go module versioning",
			Run: func(ccmd *cobra.Command, args []string){
				_, err := os.Stat(versionFile)
				if os.IsNotExist(err) {
					f, err := os.OpenFile(versionFile, os.O_CREATE|os.O_WRONLY, 0660)
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()
					if _, err = f.WriteString(VERSIONSTART); err != nil {
						log.Fatal(err)
					}
				} else {
					fb, err := ioutil.ReadFile(versionFile)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("Version file exists with content: %s", fb)
				}
			},
		}

		MajorPlusPlusCMD = &cobra.Command{
			Use: "major++",
			Args: cobra.NoArgs,
			Short: "Increment Major Version",
			Run: func(ccmd *cobra.Command, args []string){
				reader, err := os.OpenFile(versionFile, os.O_RDONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				s, err := readPreviousVersion(reader)
				if err != nil {
					log.Fatal(err)
				}
				s, err = incrementMajor(s)
				if err != nil {
					log.Fatal(err)
				}
				writer, err := os.OpenFile(versionFile, os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer writer.Close()
				if _, err := writer.WriteString(s); err != nil {
					log.Fatal(err)
				} else {
					log.Print(s)
				}
			},
		}

		MinorPlusPlusCMD = &cobra.Command{
			Use: "minor++",
			Args: cobra.NoArgs,
			Short: "Increment Minor Version",
			Run: func(ccmd *cobra.Command, args []string){
				reader, err := os.OpenFile(versionFile, os.O_RDONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				s, err := readPreviousVersion(reader)
				if err != nil {
					log.Fatal(err)
				}
				s, err = incrementMinor(s)
				if err != nil {
					log.Fatal(err)
				}
				writer, err := os.OpenFile(versionFile, os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer writer.Close()
				if _, err := writer.WriteString(s); err != nil {
					log.Fatal(err)
				} else {
					log.Print(s)
				}
			},
		}

		PatchPlusPlusCMD = &cobra.Command{
			Use: "patch++",
			Args: cobra.NoArgs,
			Short: "Increment Patch Version",
			Run: func(ccmd *cobra.Command, args []string){
				reader, err := os.OpenFile(versionFile, os.O_RDONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				s, err := readPreviousVersion(reader)
				if err != nil {
					log.Fatal(err)
				}
				s, err = incrementPatch(s)
				if err != nil {
					log.Fatal(err)
				}
				writer, err := os.OpenFile(versionFile, os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer writer.Close()
				if _, err := writer.WriteString(s); err != nil {
					log.Fatal(err)
				} else {
					log.Print(s)
				}
			},
		}
		MajorMinusMinusCMD = &cobra.Command{
			Use: "major--",
			Args: cobra.NoArgs,
			Short: "Decrement Major Version",
			Run: func(ccmd *cobra.Command, args []string){
				reader, err := os.OpenFile(versionFile, os.O_RDONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				s, err := readPreviousVersion(reader)
				if err != nil {
					log.Fatal(err)
				}
				s, err = decrementMajor(s)
				if err != nil {
					log.Fatal(err)
				}
				writer, err := os.OpenFile(versionFile, os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer writer.Close()
				if _, err := writer.WriteString(s); err != nil {
					log.Fatal(err)
				} else {
					log.Print(s)
				}
			},
		}
		MinorMinusMinusCMD = &cobra.Command{
			Use: "minor--",
			Args: cobra.NoArgs,
			Short: "Decrement Minor Version",
			Run: func(ccmd *cobra.Command, args []string){
				reader, err := os.OpenFile(versionFile, os.O_RDONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				s, err := readPreviousVersion(reader)
				if err != nil {
					log.Fatal(err)
				}
				s, err = decrementMinor(s)
				if err != nil {
					log.Fatal(err)
				}
				writer, err := os.OpenFile(versionFile, os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer writer.Close()
				if _, err := writer.WriteString(s); err != nil {
					log.Fatal(err)
				} else {
					log.Print(s)
				}
			},
		}
		PatchMinusMinusCMD = &cobra.Command{
			Use: "patch--",
			Args: cobra.NoArgs,
			Short: "Decrement Patch Version",
			Run: func(ccmd *cobra.Command, args []string){
				reader, err := os.OpenFile(versionFile, os.O_RDONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				s, err := readPreviousVersion(reader)
				if err != nil {
					log.Fatal(err)
				}
				s, err = decrementPatch(s)
				if err != nil {
					log.Fatal(err)
				}
				writer, err := os.OpenFile(versionFile, os.O_TRUNC|os.O_WRONLY, 0660)
				if err != nil {
					log.Fatal(err)
				}
				defer writer.Close()
				if _, err := writer.WriteString(s); err != nil {
					log.Fatal(err)
				} else {
					log.Print(s)
				}
			},
		}
	)

	MainCMD.AddCommand(MajorPlusPlusCMD, MajorMinusMinusCMD, MinorPlusPlusCMD, MinorMinusMinusCMD, PatchPlusPlusCMD, PatchMinusMinusCMD)
	MainCMD.PersistentFlags().StringVarP(&versionFile, "version-file", "f", "./VERSION", "The file in which to store the versioning information")
	
	if err := MainCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}