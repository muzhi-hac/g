package main

import (
	"fmt"
	"gee"
	"github.com/ipfs/go-ipfs-api"
	"html/template"
	"io"
	"net/http"
)

var uploadForm = `<html>
<body>
  <h2>Upload File to IPFS</h2>
  <form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="file" />
    <input type="submit" value="Upload" />
  </form>
</body>
</html>`

type catFile struct {
	filename string
	content  []byte
	offset   int64
}

func main() {
	router := gee.New()

	router.GET("/", func(c *gee.Context) {
		tmpl, err := template.New("index").Parse(uploadForm)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		tmpl.Execute(c.Writer, nil)
	})

	router.POST("/upload", uploadHandler)

	port := 8080
	fmt.Printf("Server started on :%d...\n", port)
	router.Run(":8081")
	//http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func uploadHandler(c *gee.Context) {
	r := c.Req
	w := c.Writer
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 将文件上传到IPFS
	cid, err := uploadToIPFS(file)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	msg := fmt.Sprintf("File Upload to IPFS Completed\nIPFS CID: %s\n", cid)
	c.String(http.StatusOK, msg)

	// 提示文件上传成功，并显示IPFS CID
	//c.String(http.StatusOK, "File Upload to IPFS Completed\nIPFS CID: %s\n" cid)
}

func uploadToIPFS(file io.Reader) (string, error) {
	// 连接到本地IPFS节点
	shell := shell.NewShell("localhost:5001")

	// 将文件上传到IPFS
	result, err := shell.Add(file)
	if err != nil {
		return "", err
	}

	return result, nil
}
func CatIpfs(filehash string, filename string) (http.File, error) {
	sh := shell.NewShell("localhost:5001")
	cat, err := sh.Cat(filehash)
	if err != nil {

		return nil, err
	}
	content, err := io.ReadAll(cat)
	if err != nil {
		return nil, err
	}
	file := &catFile{content: content,
		filename: filename,
	}

}
