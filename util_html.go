package main

import "fmt"

func GetStartTableHTML(withMap bool, withImage bool) string {
	extra := ""
	if withMap {
		extra += `<td>Map</td>`
	}
	if withImage {
		extra += `<td>Image</td>`
	}
	return `<html> <head> <style>.exif-table{border-collapse: collapse;margin: 25px 0;font-size: 0.9em;font-family: sans-serif;min-width: 400px;box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);}.exif-table thead tr{background-color: #009879;color: #ffffff;text-align: left;}.exif-table th, .exif-table td{padding: 12px 15px;}.exif-table tbody tr{border-bottom: 1px solid #dddddd;}.exif-table tbody tr:nth-of-type(even){background-color: #f3f3f3;}.exif-table tbody tr:last-of-type{border-bottom: 2px solid #009879;}.exif-table tbody tr:hover{background-color: #00987910;}</style> <title>EXIF data for ...</title> </head> <body> <table class="exif-table"> <thead> <tr> <td>Path</td><td>Latitude</td><td>Longitude</td>` + extra + `</tr></thead> <tbody>`
}

func GetEndTableHTML() string {
	return `</tbody> </table> </body></html>`
}

func GetIconMap() string {
	return `<svg width="25px" height="25px" fill="#000000" version="1.1" viewBox="0 0 395.71 395.71" xml:space="preserve">
	<path d="m197.85 0c-75.718 0-137.32 61.609-137.32 137.33 0 72.887 124.59 243.18 129.9 250.39l4.951 6.738c0.579 0.792 1.501 1.255 2.471 1.255 0.985 0 1.901-0.463 2.486-1.255l4.948-6.738c5.308-7.211 129.9-177.5 129.9-250.39 0-75.72-61.61-137.33-137.33-137.33zm0 88.138c27.13 0 49.191 22.062 49.191 49.191 0 27.115-22.062 49.191-49.191 49.191-27.114 0-49.191-22.076-49.191-49.191 0-27.129 22.076-49.191 49.191-49.191z"/></svg>`
}

func GetALink(path string, content string) string {
	return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, path, content)
}

func GetImg(path string) string {
	return fmt.Sprintf(`<img height="50" src="%s" alt="%s"></img>`, path, path)
}

func Td(x interface{}) string {
	switch x.(type) {
	case string:
		return fmt.Sprintf("<td>%s</td>", x)
	case float32:
		return fmt.Sprintf("<td>%f</td>", x)
	case float64:
		return fmt.Sprintf("<td>%f</td>", x)
	}
	return ""
}

func GetMapLink(lat float64, lon float64) string {
	mapsLink := fmt.Sprintf(`http://www.google.com/maps/place/%f,%f`, lat, lon)
	return GetALink(mapsLink, GetIconMap())
}
