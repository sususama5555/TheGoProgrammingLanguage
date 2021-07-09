package makeThumbnails

import (
	"../thumbnail"
)

// makeThumbnails3 makes thumbnails of the specified files in parallel.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// NOTE: incorrect!
	//匿名函数的循环变量快照问题，
	//上面这个单独的变量f是被所有的匿名函数所共享，
	//且会被连续的循环迭代所更新
	//for _, f := range filenames {
	//	go func() {
	//		thumbnail.ImageFile(f)
	//		// ...
	//	}()
	//}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}
