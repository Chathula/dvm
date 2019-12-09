package util

import (
	"fmt"

	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

// Download file from URL to the filepath
func DownloadFile(filepath string, url string) error {
	tmpl := fmt.Sprintf(`{{string . "prefix"}}{{ green "%s" }} {{counters . }} {{ bar . "[" "=" ">" "-" "]"}} {{percent . }} {{speed . }}{{string . "suffix"}}`, filepath)

	// Get the data
	response, err := http.Get(url)

	if err != nil {
		return errors.Wrapf(err, "Download `%s` fail", url)
	}

	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		return errors.New(fmt.Sprintf("download file with status code %d", response.StatusCode))
	}

	// Create the file
	writer, err := os.Create(filepath)

	if err != nil {
		return errors.Wrapf(err, "Create `%s` fail", filepath)
	}

	defer func() {
		err = writer.Close()

		if err != nil {
			_ = os.Remove(filepath)
		}
	}()

	bar := pb.ProgressBarTemplate(tmpl).Start64(response.ContentLength)

	bar.SetWriter(os.Stdout)

	barReader := bar.NewProxyReader(response.Body)

	_, err = io.Copy(writer, barReader)

	bar.Finish()

	if err != nil {
		err = errors.Wrap(err, "copy fail")
	}

	return err
}
