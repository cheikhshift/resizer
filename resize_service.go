package resizer

import (
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"strings"
)

var CacheFolder = "./cache"

func ResizeAndCache(imgpath string, cfolder string, width int, height int) (cacheref string, err error) {
	//untested
	cacheref = fmt.Sprintf("%s/%s_size_%v_%v.jpg", CacheFolder, strings.TrimPrefix(strings.Replace(imgpath, "/", ".", -1), "."), width, height)
	if _, err = os.Stat(cacheref); os.IsNotExist(err) {

		nimg, err := imaging.Open(imgpath)

		if err != nil {
			return cacheref, err
		}

		newimage := imaging.Resize(nimg, width, height, imaging.Lanczos)
		err = imaging.Save(newimage, cacheref)
		if err != nil {
			os.Remove(cacheref)

		}

	}
	return
}
