# Go Notes
---------
>>>> @kaustubh-walokar
---------

## Misc Stuff
- `import(fmt)`  - does formatted io for prinfing or scanfing.

## Imports & Exports
- all structs and funcs begining with [A-Z] exported (public) ; [a-z] are not (private).

## Variables

- dealare by `[var/const] <name> <type>` 
```go
	var x int 
	var x int = 1 

	//implicit type assigned here, type cannot be changed
	x := 1
	a,b := 1,2

	var a [2]string
	b := [5]int{1, 2, 3, 4, 5}
	b := make()
```

## Testing

- Name file as xxx_test.go
- Run with `go test xxx_test.go`
- Name Test functions as `func TestXxxx(t *testing.T)`

