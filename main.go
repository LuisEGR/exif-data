package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

func main() {

	var args struct {
		Dir             string `arg:"positional" help:"Directory that contains the images (this tool will search recursively), this can be a relative/absolute path"`
		Output          string `default:"data" help:"Output filename to write the extracted data"`
		Format          string `default:"csv" help:"Format to write the output, this can be 'csv' or 'html', this will be appended to the output name as well"`
		DefaultLat      string `arg:"--default-lat" default:"" help:"default latitude to fill in the csv for those images which don't have valid GPS data"`
		DefaultLon      string `arg:"--default-lon" default:"" help:"default longitude to fill in the csv for those images which don't have valid GPS data"`
		HTMLMapLink     bool   `arg:"--html-map-link" default:"true" help:"Adds a link to see the GPS coordinates in Google Maps on each image (only HTML)"`
		HTMLImageInline bool   `arg:"--html-inline-img" default:"true" help:"Displays image inline (only HTML)"`
	}
	arg.MustParse(&args)
	dir := args.Dir

	if dir == "" {
		fmt.Println("No directory provided, using current one")
		dir = "."
	}

	fileOut, err := os.Create(args.Output + "." + args.Format)
	if err != nil {
		fmt.Println("err :", err)
		panic("Error creating " + args.Format + " file")
	}
	defer fileOut.Close()

	w := bufio.NewWriter(fileOut)

	if args.Format == "html" {
		w.WriteString(
			GetStartTableHTML(
				args.HTMLMapLink,
				args.HTMLImageInline))
	}

	e := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {

			// Check for content-type with magic bytes
			isImg, _ := IsImage(path)

			gpsLat := 0.0
			gpsLon := 0.0
			if isImg {
				gps, err := getGPSInfo(path)
				// Errors regarding exif reading can be ignored and
				// use the defaults for this, in order to not interrupt the
				// processing of the next images
				if err != nil {
					fmt.Printf("[error] %s: %s\n", path, err)
				} else {
					gpsLat = gps.Latitude.Decimal()
					gpsLon = gps.Longitude.Decimal()
				}
				if args.Format == "csv" {
					data := path + ","
					if gps != nil {
						data += fmt.Sprintf("%f", gpsLat)
						data += ","
						data += fmt.Sprintf("%f", gpsLon)
					} else {
						data += args.DefaultLat
						data += ","
						data += args.DefaultLon
					}
					w.WriteString(data + "\n")
				} else if args.Format == "html" {
					// <tr> <td>1</td><td>1</td><td>1</td></tr>
					w.WriteString("<tr>")

					w.WriteString(Td(GetALink(path, path)))

					if gps != nil {
						w.WriteString(Td(gpsLat))
						w.WriteString(Td(gpsLon))
						if args.HTMLMapLink {
							w.WriteString(Td(GetMapLink(gpsLat, gpsLon)))
						}

						if args.HTMLImageInline {
							w.WriteString(Td(GetImg(path)))
						}
					} else {
						w.WriteString(Td(args.DefaultLat))
						w.WriteString(Td(args.DefaultLon))
						w.WriteString(Td("")) // No Map link

						if args.HTMLImageInline {
							w.WriteString(Td(GetImg(path)))
						}
					}
					w.WriteString("</tr>")

				}

			}

		}
		return nil
	})

	if args.Format == "html" {
		w.WriteString(GetEndTableHTML())
	}

	w.Flush()
	if e != nil {
		panic(e)
	}

	fmt.Printf("\nResults saved to %s.%s\n", args.Output, args.Format)

}

func getGPSInfo(path string) (*exif.GpsInfo, error) {
	imageData, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	rawExif, err := exif.SearchAndExtractExif(imageData)
	if len(rawExif) == 0 { // No EXIF data
		return nil, fmt.Errorf("no EXIF data found (%s)", err)
	}

	if err == nil { // No errors, proceed with the reading

		im := exifcommon.NewIfdMapping()

		err = exifcommon.LoadStandardIfds(im)
		if err != nil {
			return nil, err
		}

		ti := exif.NewTagIndex()

		_, index, err := exif.Collect(im, ti, rawExif)
		if err != nil {
			return nil, fmt.Errorf("error extracting EXIF data (%s)", err)
			// no panic, we can handle with defaults
		} else {
			ifd, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
			if err != nil {
				return nil, err
			}

			gi, err := ifd.GpsInfo()
			if err != nil {
				return nil, err
			}

			fmt.Println("[info] Extracted data for: " + path)

			return gi, nil
		}
	}
	return nil, nil
}
