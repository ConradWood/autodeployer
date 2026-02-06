package common

import (
	"flag"
	"fmt"
	"strings"

	dm "golang.conradwood.net/apis/deploymonkey"
)

var (
	use_artefact_server_url = flag.Bool("use_artefact_server_url", false, "use artefact:// instead of http:// urls")
)

// the url with variables resolved
func ResolvedDownloadURL(app *dm.ApplicationDefinition) string {
	use_artefact := *use_artefact_server_url
	idx := strings.Index(app.DownloadURL, "dist/")
	if idx == -1 {
		use_artefact = false
	}
	if !strings.HasPrefix(app.DownloadURL, "http://") {
		use_artefact = false
	}
	if !use_artefact {
		s := app.DownloadURL
		s = strings.ReplaceAll(s, "${BUILDID}", fmt.Sprintf("%d", app.BuildID))
		return s
	}
	path := app.DownloadURL[idx:]
	url := fmt.Sprintf("artefact://%d/%d/master/%s", app.ArtefactID, app.BuildID, path)
	return url

}
