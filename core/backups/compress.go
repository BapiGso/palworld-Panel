package backups

import (
	"archive/zip"
	"io"
	"os"
	"palworld-panel/core/conf"
	"path/filepath"
	"strings"
	"time"
)

func Compress() {
	src_dir := conf.Config.Get("serverPath").(string) + "\\Pal\\Saved"
	zipfile, _ := os.Create(time.Now().Format("20060102-150405") + ".zip")
	defer zipfile.Close()
	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	// 遍历路径信息
	err := filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == src_dir {
			return nil
		}
		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, src_dir+`/`)
		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}
		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	conf.Config.Set("lastBackupTime", time.Now().Unix())
}
