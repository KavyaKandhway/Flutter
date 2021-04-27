package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"

	"strings"

	"github.com/furrble/go-cloudinary"
	"github.com/outofpluto/goconfig/config"
)

type Config struct {
	// Url to the CLoudinary service.
	CloudinaryURI *url.URL
	// Url to a MongoDB instance, used to track files and upload
	// only changed. Optional.
	MongoURI *url.URL
	// Regexp pattern to prevent remote file deletion.
	KeepFilesPattern string
	// An optional remote prepend path, used to generate a unique
	// data path to a remote resource. This can be useful if public
	// ids are not random (i.e provided as request arguments) to solve
	// any caching issue: a different prepend path generates a new path
	// to the remote resource.
	PrependPath string
	// ProdTag is an alias to PrependPath. If PrependPath is empty but
	// ProdTag is set (with at prodtag= line in the [global] section of
	// the config file), PrependPath is set to ProdTag. For example, it
	// can be used with a DVCS commit tag to force new remote data paths
	// to remote resources.
	ProdTag string
}

var service *cloudinary.Service

// Parses all structure fields values, looks for any
// variable defined as ${VARNAME} and substitute it by
// calling os.Getenv().
//
// The reflect package is not used here since we cannot
// set a private field (not exported) within a struct using
// reflection.
func replaceEnvVars(src string) (string, error) {
	r, err := regexp.Compile(`\${([A-Z_]+)}`)
	if err != nil {
		return "", err
	}
	envs := r.FindAllString(src, -1)
	for _, varname := range envs {
		evar := os.Getenv(varname[2 : len(varname)-1])
		if evar == "" {
			return "", errors.New(fmt.Sprintf("error: env var %s not defined", varname))
		}
		src = strings.Replace(src, varname, evar, -1)
	}
	return src, nil
}

func handleQuery(uri *url.URL) (*url.URL, error) {
	qs, err := url.QueryUnescape(uri.String())
	if err != nil {
		return nil, err
	}
	r, err := replaceEnvVars(qs)
	if err != nil {
		return nil, err
	}
	wuri, err := url.Parse(r)
	if err != nil {
		return nil, err
	}
	return wuri, nil
}
func (c *Config) handleEnvVars() error {
	// [cloudinary]
	if c.CloudinaryURI != nil {
		curi, err := handleQuery(c.CloudinaryURI)
		if err != nil {
			return err
		}
		c.CloudinaryURI = curi
	}
	if len(c.PrependPath) == 0 {
		// [global]
		if len(c.ProdTag) > 0 {
			ptag, err := replaceEnvVars(c.ProdTag)
			if err != nil {
				return err
			}
			c.PrependPath = cloudinary.EnsureTrailingSlash(ptag)
		}
	}

	// [database]
	if c.MongoURI != nil {
		muri, err := handleQuery(c.MongoURI)
		if err != nil {
			return err
		}
		c.MongoURI = muri
	}
	return nil
}

// LoadConfig parses a config file and sets global settings
// variables to be used at runtime. Note that returning an error
// will cause the application to exit with code error 1.
func LoadConfig(path string) (*Config, error) {
	settings := &Config{}

	c, err := config.ReadDefault(path)
	if err != nil {
		return nil, err
	}

	// Cloudinary settings
	var cURI *url.URL
	var uri string

	if uri, err = c.String("cloudinary", "uri"); err != nil {
		return nil, err
	}
	if cURI, err = url.Parse(uri); err != nil {
		return nil, errors.New(fmt.Sprint("cloudinary URI: ", err.Error()))
	}
	settings.CloudinaryURI = cURI

	// An optional remote prepend path
	if prepend, err := c.String("cloudinary", "prepend"); err == nil {
		settings.PrependPath = cloudinary.EnsureTrailingSlash(prepend)
	}
	settings.ProdTag, _ = c.String("global", "prodtag")

	// Keep files regexp? (optional)
	var pattern string
	pattern, _ = c.String("cloudinary", "keepfiles")
	if pattern != "" {
		settings.KeepFilesPattern = pattern
	}

	// mongodb section (optional)
	uri, _ = c.String("database", "uri")
	if uri != "" {
		var mURI *url.URL
		if mURI, err = url.Parse(uri); err != nil {
			return nil, errors.New(fmt.Sprint("mongoDB URI: ", err.Error()))
		}
		settings.MongoURI = mURI
	} else {
		fmt.Fprintf(os.Stderr, "Warning: database not set (upload sync disabled)\n")
	}

	// Looks for env variables, perform substitutions if needed
	if err := settings.handleEnvVars(); err != nil {
		return nil, err
	}
	return settings, nil
}

func fail(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
	os.Exit(1)
}

func printResources(res []*cloudinary.Resource, err error) {
	if err != nil {
		fail(err.Error())
	}
	if len(res) == 0 {
		fmt.Println("No resource found.")
		return
	}
	fmt.Printf("%-30s %-10s %-5s %s\n", "public_id", "Version", "Type", "Size")
	fmt.Println(strings.Repeat("-", 70))
	for _, r := range res {
		fmt.Printf("%-30s %d %s %10d\n", r.PublicId, r.Version, r.ResourceType, r.Size)
	}
}

func printResourceDetails(res *cloudinary.ResourceDetails, err error) {
	if err != nil {
		fail(err.Error())
	}
	if res == nil || len(res.PublicId) == 0 {
		fmt.Println("No resource details found.")
		return
	}
	fmt.Printf("%-30s %-6s %-10s %-5s %-8s %-6s %-6s %-s\n", "public_id", "Format", "Version", "Type", "Size", "Width", "Height", "Url")
	fmt.Printf("%-30s %-6s %-10d %-5s %-8d %-6d %-6d %-s\n", res.PublicId, res.Format, res.Version, res.ResourceType, res.Size, res.Width, res.Height, res.Url)

	fmt.Println()

	for i, d := range res.Derived {
		if i == 0 {
			fmt.Printf("%-25s %-8s %-s\n", "transformation", "Size", "Url")
		}
		fmt.Printf("%-25s %-8d %-s\n", d.Transformation, d.Size, d.Url)
	}
}

func perror(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	os.Exit(1)
}

func step(caption string) {
	fmt.Printf("==> %s\n", caption)
}

func main() {

	optRaw := flag.String("r", "", "raw filename or public id")
	optImg := flag.String("i", "", "image filename or public id")
	optVerbose := flag.Bool("v", false, "verbose output")
	optSimulate := flag.Bool("s", false, "simulate, do nothing (dry run)")
	optAll := flag.Bool("a", false, "applies to all resource files")
	flag.Parse()

	action := "up"
	*optRaw = ""
	*optImg = "C:/Users/User/Desktop/test.jpg"

	*optVerbose = false
	*optSimulate = false
	*optAll = false
	log.Print(optRaw)
	supportedAction := func(act string) bool {
		switch act {
		case "ls", "rm", "up", "url":
			return true
		}
		return false
	}(action)
	if !supportedAction {
		fmt.Fprintf(os.Stderr, "Unknown action '%s'\n", action)
		flag.Usage()
	}

	var err error
	settings, err := LoadConfig("C:/Users/User/Desktop/settings.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", flag.Arg(1), err.Error())
		os.Exit(1)
	}

	service, err = cloudinary.Dial(settings.CloudinaryURI.String())
	service.Verbose(*optVerbose)
	service.Simulate(*optSimulate)
	service.KeepFiles(settings.KeepFilesPattern)
	if settings.MongoURI != nil {
		if err := service.UseDatabase(settings.MongoURI.String()); err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting to mongoDB: %s\n", err.Error())
			os.Exit(1)
		}
	}

	if err != nil {
		fail(err.Error())
	}

	if *optSimulate {
		fmt.Println("*** DRY RUN MODE ***")
	}

	if len(settings.PrependPath) > 0 {
		fmt.Println("/!\\ Remote prepend path set to: ", settings.PrependPath)
	} else {
		fmt.Println("/!\\ No remote prepend path set")
	}

	switch action {
	case "up":
		if *optRaw == "" && *optImg == "" {
			fail("Missing -i or -r option.")
		}
		if *optRaw != "" {
			step("Uploading as raw data")
			if _, err := service.UploadStaticRaw(*optRaw, nil, settings.PrependPath); err != nil {
				perror(err)
			}
		} else {
			step("Uploading as images")
			if _, err := service.UploadStaticImage(*optImg, nil, settings.PrependPath); err != nil {
				perror(err)
			}
		}
		break

	case "rm":
		if *optAll {
			step(fmt.Sprintf("Deleting all resources..."))
			if err := service.DropAll(os.Stdout); err != nil {
				perror(err)
			}
		} else {
			if *optRaw == "" && *optImg == "" {
				fail("Missing -i or -r option.")
			}
			if *optRaw != "" {
				step(fmt.Sprintf("Deleting raw file %s", *optRaw))
				if err := service.Delete(*optRaw, settings.PrependPath, cloudinary.RawType); err != nil {
					perror(err)
				}
			} else {
				step(fmt.Sprintf("Deleting image %s", *optImg))
				if err := service.Delete(*optImg, settings.PrependPath, cloudinary.ImageType); err != nil {
					perror(err)
				}
			}
		}

	case "ls":
		if *optImg != "" {
			fmt.Println("==> Image Details:")
			printResourceDetails(service.ResourceDetails(*optImg))
		} else {
			fmt.Println("==> Raw resources:")
			printResources(service.Resources(cloudinary.RawType))
			fmt.Println("==> Images:")
			printResources(service.Resources(cloudinary.ImageType))
		}

	case "url":
		if *optRaw == "" && *optImg == "" {
			fail("Missing -i or -r option.")
		}
		if *optRaw != "" {
			fmt.Println(service.Url(*optRaw, cloudinary.RawType))
		} else {
			log.Print("yayaya")
			fmt.Println(service.Url(*optImg, cloudinary.ImageType))
		}
	}

	fmt.Println("")
	if err != nil {
		fail(err.Error())
	}
}
