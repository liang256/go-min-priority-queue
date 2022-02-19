# go-min-priority-queue
Implement min-priority queue in Go

# Usage
```
go run main.go
```

Should print out this in Terminal:
```
&{[] 0}
&{[10] 0}
&{[10 2] 1}
&{[10 2 1] 2}
&{[10 2 1 3] 2}
&{[10 2 1 3 0] 4}
pop: 0 q: &{[10 2 1 3] 2}
pop: 1 q: &{[10 2 3] 1}
pop: 2 q: &{[10 3] 1}
pop: 3 q: &{[10] 0}
pop: 10 q: &{[] 0}
```
