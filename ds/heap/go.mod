module github.com/seblkma/go-ultimate/ds/heap

go 1.22.5

replace (
	github.com/seblkma/go-ultimate/ds/heap/heapint => ./heapint
	github.com/seblkma/go-ultimate/ds/heap/priorityqueueitem => ./priorityqueueitem
)

require github.com/seblkma/go-ultimate/ds/heap/heapint v0.0.0-00010101000000-000000000000
