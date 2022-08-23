# Golang Note


## command 

### go command
1. go run file.go
2. go test -v -bench=. -run=none -benchmem .
3. go get (package) // from github is allowed
4. go mod init
5. go install .

go run -race file.go // check race condition


### git command
git add .
git commit -m "comment"
git push -u origin main



# Note

rune = int32 
underscore

golang compiler take uppercase as Public, lowercase as Private

Interface - include many Methods

````
type geometry interface {
    are() float 64
}
...
func (r *rectangle) area() float64 {
    return r.width * r.height
}
````

Struct

````
type rectangle struct {
    width float64
    height float64
}
````

# Online Resources
1. https://willh.gitbook.io/build-web-application-with-golang-zhtw/02.0/02.4
2. Scope https://ithelp.ithome.com.tw/articles/10237844
3. Slice and Array https://pjchender.dev/golang/slice-and-array/
4. https://pjchender.dev/golang/tips-slice-of-pointers/
5. Struct https://pjchender.dev/golang/structs/
6. Interface https://blog.kennycoder.io/2020/02/03/Golang-深入理解interface常見用法/
7. The Go programming language by Brian W. Kernighan and Alan A. A. Donovan
8. 