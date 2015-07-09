package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type CloseableResonseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (w gzipResponseWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

func (w gzipResponseWriter) Close() {
	w.Writer.Close()
}

func (w gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (w closeableResponseWriter) Close() {}

func GetResponseWriter(w http.ResponseWriter, req *http.Request) CloseableResonseWriter {
	var rw CloseableResonseWriter

	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")

		rw = gzipResponseWriter{ResponseWriter: w, Writer: gzip.NewWriter(w)}
	} else {
		rw = closeableResponseWriter{ResponseWriter: w}
	}
	return rw
}
