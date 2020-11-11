package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pgconfig/api/doc"
)

var (
	targetFolder string
)

func init() {
	flag.StringVar(&targetFolder, "target-folder", "/home/seba/projetos/github.com/sebastianwebber/pgconfig-api/pg_doc", "default target docs")
	flag.Parse()
}

func createFolder(param string) error {

	folder := fmt.Sprintf("%s/%s", targetFolder, param)

	_, err := os.Stat(folder)

	if os.IsNotExist(err) {
		return os.MkdirAll(folder, 0755)
	}

	return nil
}

func saveParam(param string, version float32, output doc.ParamDoc) error {

	log.Printf("Processing %s for %.1f", param, version)

	createFolder(param)

	f, err := os.Create(fmt.Sprintf("%s/%s/%s.txt", targetFolder, param, doc.FormatVer(version)))

	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("%s (%s)\n", output.Title, output.ParamType))
	f.WriteString(strings.Join(output.Text, "\n\n"))

	return nil
}

func main() {

	allVersions := []float32{
		9.1,
		9.2,
		9.3,
		9.4,
		9.5,
		9.6,
		10.0,
		11.0,
		12.0,
		13.0,
	}

	allParams := []string{
		"checkpoint_completion_target",
		"checkpoint_segments",
		"effective_cache_size",
		"effective_io_concurrency",
		"listen_addresses",
		"maintenance_work_mem",
		"max_connections",
		"max_parallel_workers",
		"max_parallel_workers_per_gather",
		"max_wal_size",
		"max_worker_processes",
		"min_wal_size",
		"random_page_cost",
		"shared_buffers",
		"wal_buffers",
		"work_mem",
	}

	for _, param := range allParams {
		for _, ver := range allVersions {

			a, err := doc.Get(param, ver)

			// 404 unsupported
			if err != nil {
				continue
			}

			saveParam(param, ver, a)
		}
	}
}
