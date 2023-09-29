# exif-data

## Description

`exif-data` is a command-line tool written in Go for extracting and exporting Exif metadata from images in a specified directory. It provides flexibility in choosing the output format (CSV or HTML) and allows customization of default GPS coordinates for images without valid GPS data. Additionally, it offers the option to generate HTML files with Google Maps links for easy visualization of GPS coordinates.

## Installation

You can install `exif-data` via Go using the following steps:

1. Get the source code.

2. Change to the project directory:

   ```bash
   cd exif-data
   ```

3. Build the tool using the `go build` command:

   ```bash
   go build -o exif-data *.go
   ```

4. Optionally, move the compiled binary to a directory in your PATH for easy access:

   ```bash
   mv exif-data /usr/local/bin/
   ```

## Usage

```bash
exif-data [--output OUTPUT] [--format FORMAT] [--default-lat DEFAULT-LAT] [--default-lon DEFAULT-LON] [--html-map-link] [--html-inline-img] [DIR]
```

### Positional arguments:

- `DIR` (optional): Directory that contains the images (this tool will search recursively), which can be a relative or absolute path.

### Options:

- `--output OUTPUT` (optional): Output filename to write the extracted data (default: `data`).
- `--format FORMAT` (optional): Format to write the output, which can be 'csv' or 'html'; this will be appended to the output name as well (default: `csv`).
- `--default-lat DEFAULT-LAT` (optional): Default latitude to fill in the CSV for those images that don't have valid GPS data.
- `--default-lon DEFAULT-LON` (optional): Default longitude to fill in the CSV for those images that don't have valid GPS data.
- `--html-map-link` (optional): Adds a link to see the GPS coordinates in Google Maps on each image (only applicable for HTML format) (default: true).
- `--html-inline-img` (optional): Displays the image inline (only applicable for HTML format) (default: true).
- `--help, -h`: Display this help and exit.

## Examples

1. Extract Exif data from images in the current directory, and subdirectories, using default settings, and save it as `data.csv`:

```bash
exif-data
```

2. Extract Exif data from images in a specific directory and customize the output filename and format:

```bash
exif-data my_images/ --output my_exif_data --format html
```

3. Specify default latitude and longitude for images without valid GPS data:

```bash
exif-data photos/ --default-lat 0 --default-lon 0
```

4. Generate HTML files without either Google Maps links or images inline:

```bash
exif-data vacation_photos/ --format html --html-map-link=false --html-inline-img=false
```

## License

This tool is licensed under the [MIT License](LICENSE.md).

## Author

[@LuisEGR](https://github.com/LuisEGR/)
