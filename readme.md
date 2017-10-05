# Resizer

Golangserver image resizer.

## Requirements
- [golangserver.com](http://golangserver.com) aka GOS

## How to install
Add this import tag to your `.gxml` file :

		<import src="github.com/cheikhshift/resizer/gos.gxml" />


## Configure autoresize
This procedure will resize any image being served.

### Properties
Set the following variables within your GOS `<main>` tag. You may also update the variables within methods and web services.

1. MaxWidth `int` : Trigger resize with image over (in pixels).
2. MaxHeight `int` : Maximum image height (in pixels).
3. ShrinkPercent `int` : 0 decimal percent value to reduce image by.
4. CacheFolder `string` : Path of folder to save resized images.
 
### API
You may also resize any image within your project's web root to your liking.

#### Format

	${protocol}://${hostname}/${resize}/${width}/${height}/${path}

Notes : 
- The path variable is relative to your GOS project web root. For example to get a resized image in path `web/img/set/img.png` with dimensions 40x40 the URI would be : http://localhost/resize/40/40/img/set/img.png

#### Golang package

##### Import

	import "github.com/cheikhshift/resizer"

Use the following function to resize and place in cache any image in your file system. Use the same function again to get the location of your file.

	func ResizeAndCache(imgpath string, cache_folder string,width int , height int) (cache_ref string,err error)

