# Sub Tests (Testing)

Today we will be learning about subtests and how to write them in golang.

Subtests are nested test that are written inside of a TestFunc() and they have multiple tests inside them.

With sub testing we can have different test that are well structured and organized.

## syntax
To write a subtest you write a function with the tag Test and then within the scope of the test function you then write a method `t.Run()` which takes a string(msg) and a function (the main definition of the test)
```Go
func TestDemo(t *testing.T) {
    t.Run("NAME_OF_TEST", FUNC(){...})
    t.Run("NAME_OF_TEST", FUNC(){...})
    t.Run("NAME_OF_TEST", FUNC(){...})
}
```
Multiple sub tests can be written inside of a single test `func()`
Here is a simple sub test in golang:
```Go
func TestPrint(t *testing.T) {
    t.Run("testing return value", func(t testing.T){
        got := printString(32)
        want := "32"

        if got != want {
            t.Errorf("expected: %v, got: %v, given: %v", want, got, 32)
        }
    })

    t.Run("testing a number greater than 100", func(t testing.T){
        got := printString(328)
        want := "328"

        if got != want {
            t.Errorf("expected: %v, got: %v, given: %v", want, got, 32)
        }
    })
}
```

And that is a quick-way to write sub tests in golang. 
