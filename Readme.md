## Go Mapper Solution
Solution for the Go mapper problem

## gomapper
gomapper - This package provides function that capitalizes *only* every third alphanumeric character of a string.

<<<<<<< HEAD
| File | Functionality |
| ------------- | ------------- |
| main.go | Starting point of this project |
| gomapper.go | This file contains solution to problem 1 |
| gomapper_test.go | This file contains test cases for the problem 1
| gomapperinterface.go | This file contains interface implementation for problem 2|
| gomapperinterface_test.go | This file contains test cases for interface implementation for problem 2|

- This module also contains a **main.go** file which can be run using the command ** go run main.go**
- Run test cases by navigating to gomapper directory using **go test -v**
## TestCases Output
```
=== RUN   TestCapitalizeEveryThirdAlphanumericChar
=== RUN   TestCapitalizeEveryThirdAlphanumericChar/CapitalizeEveryThirdAlphanumericChar_-_success
--- PASS: TestCapitalizeEveryThirdAlphanumericChar (0.00s)
    --- PASS: TestCapitalizeEveryThirdAlphanumericChar/CapitalizeEveryThirdAlphanumericChar_-_success (0.00s)
=== RUN   Test_skipString_GetValueAsRuneSlice
=== RUN   Test_skipString_GetValueAsRuneSlice/success
--- PASS: Test_skipString_GetValueAsRuneSlice (0.00s)
    --- PASS: Test_skipString_GetValueAsRuneSlice/success (0.00s)
=== RUN   Test_getPositionMapper
=== RUN   Test_getPositionMapper/success
--- PASS: Test_getPositionMapper (0.00s)
    --- PASS: Test_getPositionMapper/success (0.00s)
=== RUN   Test_skipString_String
=== RUN   Test_skipString_String/success
--- PASS: Test_skipString_String (0.00s)
    --- PASS: Test_skipString_String/success (0.00s)
PASS
ok      gomappersolution/gomapper       0.233s
```

## Time Complexity and Space complexity analysis
Time Complexity: O(N), N = length of the given string
Space Complexity: O(N), N = length of the given string

## Output of main.go

| input position to be capitalized| input string | output |
| ------------- | ------------- | ------------|
=======
- This module also contains a **main.go** file which can be run using the command **go run main.go**
>>>>>>> e850dd2849a9ef9581031b73eb2166e05fa9079d
