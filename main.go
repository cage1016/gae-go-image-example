package main

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
)

const (
	Bucket           = "<your-gae-default-bucket>"
	NotFoundImageURL = "img-api-example/img404.jpg"
)

/**
 * index page
 */
func indexGET(c echo.Context) error {
	c.Render(200, "index.tpl", map[string]interface{}{
		"title": "go gae image API example",
		"images": []string{
			"img?f=img-api-example/1.png",
			"img?f=img-api-example/1.png&s=200",
			"img?f=img-api-example/1.png&s=200-c",
			"img?f=img-api-example/2.jpg",
		},
	})
	return nil
}

/**
 * get not foound image default image, img404.jpg
 */
func getNotFoundImage(ctx context.Context) (urlString string, err error) {
	blobPath := fmt.Sprintf("/gs/%s/%s", Bucket, NotFoundImageURL)
	blobKey, err := blobstore.BlobKeyForFile(ctx, blobPath)
	if err != nil {
		return
	}
	opts := image.ServingURLOptions{Secure: false, Crop: true}
	url, err := image.ServingURL(ctx, blobKey, &opts)
	if err != nil {
		return
	}
	urlString = url.String()
	return
}

/**
 * image serve handler
 */
func imgServe(c echo.Context) error {
	ctx := appengine.NewContext(c.Request().(*standard.Request).Request)

	filePath := c.QueryParam("f")
	size := c.QueryParam("s")

	blobPath := fmt.Sprintf("/gs/%s/%s", Bucket, filePath)
	blobKey, err := blobstore.BlobKeyForFile(ctx, blobPath)
	if err != nil {
		c.String(http.StatusExpectationFailed, err.Error())
	}

	opts := image.ServingURLOptions{Secure: false, Crop: true}
	if url, err := image.ServingURL(ctx, blobKey, &opts); err != nil {
		if n, err := getNotFoundImage(ctx); err != nil {
			return err
		} else {
			c.Redirect(http.StatusTemporaryRedirect, n)
		}
	} else {
		if size != "" {
			c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s=s%s", url.String(), size))
		} else {
			c.Redirect(http.StatusTemporaryRedirect, url.String())
		}
	}

	return nil
}

func init() {
	e.GET("/", indexGET)
	e.GET("/img", imgServe)
}
