package handlers

import (
	"net/http"
	"net/url"

	"github.com/xray-web/web-check-api/checks"
)

func HandleGetHeaders(h *checks.Headers) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawURL := r.URL.Query().Get("url")
		if rawURL == "" {
			JSONError(w, ErrMissingURLParameter, http.StatusBadRequest)
			return
		}

		// Parse and validate the URL
		parsedURL, err := url.ParseRequestURI(rawURL)
		if err != nil || !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
			JSONError(w, ErrInvalidURL, http.StatusBadRequest)
			return
		}

		resp, err := http.Get(parsedURL.String())
		if err != nil {
			JSONError(w, err, http.StatusInternalServerError)
		}

		JSON(w, headers, http.StatusOK)
	})
}
