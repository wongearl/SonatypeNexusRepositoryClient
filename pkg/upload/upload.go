package upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	klog "k8s.io/klog"
)

func Upload(username, password, filePath, repositoryName, nexusURL string) (err error) {
	client := http.Client{}
	// 读取文件
	file, err := os.Open(filePath)
	if err != nil {
		klog.Errorf("无法打开文件: %s\n", err)
		return
	}
	defer file.Close()

	// 创建一个带有文件内容的缓冲区
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		klog.Errorf("无法创建表单文件: %s\n", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		klog.Errorf("无法复制文件内容到表单文件: %s\n", err)
		return
	}
	writer.Close()

	// 准备上传请求
	uploadURL := nexusURL + "/repository/" + repositoryName + "/" + filePath
	req, err := http.NewRequest("PUT", uploadURL, body)
	if err != nil {
		klog.Errorf("无法创建上传请求: %s\n", err)
		return
	}

	// 设置请求头部
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(username, password)

	// 发送上传请求
	resp, err := client.Do(req)
	if err != nil {
		klog.Errorf("上传请求失败: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// 检查上传响应
	if resp.StatusCode != http.StatusCreated {
		klog.Errorf("上传失败: %s\n", resp.Status)
		return fmt.Errorf("上传失败: %s", resp.Status)
	}

	fmt.Println("成功推送文件到Sonatype Nexus Repository")

	return nil
}
