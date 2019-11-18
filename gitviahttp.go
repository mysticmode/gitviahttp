package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

type gitHandler struct {
	w    http.ResponseWriter
	r    *http.Request
	rpc  string
	dir  string
	file string
}

func getServiceType(r *http.Request) string {
	vars := r.URL.Query()
	serviceType := vars["service"][0]
	if !strings.HasPrefix(serviceType, "git-") {
		return ""
	}
	return strings.TrimPrefix(serviceType, "git-")
}

func gitCommand(dir string, args ...string) []byte {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v", err)
	}

	return out
}

func updateServerInfo(dir string) []byte {
	return gitCommand(dir, "update-server-info")
}

func (ghrs *gitHandler) sendFile(contentType string) {
	reqFile := path.Join(ghrs.dir, ghrs.file)
	fi, err := os.Stat(reqFile)
	if os.IsNotExist(err) {
		ghrs.w.WriteHeader(http.StatusNotFound)
		return
	}

	ghrs.w.Header().Set("Content-Type", contentType)
	ghrs.w.Header().Set("Content-Length", fmt.Sprintf("%d", fi.Size()))
	ghrs.w.Header().Set("Last-Modified", fi.ModTime().Format(http.TimeFormat))
	http.ServeFile(ghrs.w, ghrs.r, reqFile)
}

func packetWrite(str string) []byte {
	s := strconv.FormatInt(int64(len(str)+4), 16)
	if len(s)%4 != 0 {
		s = strings.Repeat("0", 4-len(s)%4) + s
	}
	return []byte(s + str)
}

func packetFlush() []byte {
	return []byte("0000")
}

func hdrNocache(w http.ResponseWriter) {
	w.Header().Set("Expires", "Fri, 01 Jan 1980 00:00:00 GMT")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-cache, max-age=0, must-revalidate")
}

func hdrCacheForever(w http.ResponseWriter) {
	now := time.Now().Unix()
	expires := now + 31536000
	w.Header().Set("Date", fmt.Sprintf("%d", now))
	w.Header().Set("Expires", fmt.Sprintf("%d", expires))
	w.Header().Set("Cache-Control", "public, max-age=31536000")
}

func serviceRPC(gh gitHandler, rpc string) {
	// vars := r.URL.Query()

	// rpcKey, ok := vars["rpc"]
	// if !ok || len(rpcKey[0]) < 1 {
	// 	log.Println("rpc key is missing or not valid")
	// 	return
	// }

	// rpc := rpcKey[0]

	// if rpc != "upload-pack" && rpc != "receive-pack" {
	// 	ghrs := gitHandler{}
	// 	updateServerInfo(ghrs.dir)
	// 	ghrs.sendFile("text/plain; charset=utf-8")
	// 	return
	// }

	if gh.r.Header.Get("Content-Type") != fmt.Sprintf("application/x-git-%s-request", rpc) {
		gh.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	gh.w.Header().Set("Content-Type", fmt.Sprintf("application/x-git-%s-result", rpc))

	var err error
	reqBody := gh.r.Body

	// Handle GZIP
	if gh.r.Header.Get("Content-Encoding") == "gzip" {
		reqBody, err = gzip.NewReader(reqBody)
		if err != nil {
			fmt.Printf("Fail to create gzip reader: %v", err)
			gh.w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	cmd := exec.Command("git", rpc, "--stateless-rpc", gh.dir)

	var stderr bytes.Buffer

	cmd.Dir = gh.dir
	cmd.Stdin = reqBody
	cmd.Stdout = gh.w
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Fail to serve RPC(%s): %v - %s", rpc, err, stderr.String()))
		return
	}
}

func main() {
	fmt.Println("Hello, World!")
}
