# The Go Type System

In this chapter, I'll discuss the Go type system. IMO the type system of a language defines its spirit. Go is no different. The Go designers started at usual with C and then deviated and extended it in interesting ways. I'll often compare how things are in C/C++ vs/ Go. First, I'll cover all the types, then I'll discuss some type declaration mechanics and smantics and for the grand finale reflection!

## Types	

### Numeric types

Go has integral, floating point and complex numeric types. You don't need much more than that most of the time. Here are all the types, with the range of values they represent:

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts
```

The integral types can be signed or unsigned and all the types have their bitwidth as part of the type. For example int16, is a signed integer that is occupies 16 bits and therefore can hold any integer value between -32768 and 32767. The internal rerpesentatin is 2's complement. Floating point numbers are always signed and come in two flavors of 32-bit and 64-bit. Finally, complex numbers are made of two floating point numbers that correpond to the real and imaginary parts.

The decision to explicitly specify the exact size/range of values of each type is great for writing platform-agnostic programs. It's a hard lesson from C where `int` would mean different things on different platforms.

That said, sometimes you're not worried about overflow and concatanation and you don't want to pontificate and choose between int8 and int16. Go accomodates you and provides 3 more liberal types whose size will depend on the specific implementation:

```
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

The interesting thing about these numeric types is that they are distinct types. You can't assign a variable of one type to another type without explicit conversion. These statements:

```
	var a int8 = 4
	var b int16 = i8
```

Will result in the following error message:


```
manuscript/code/chapter-3/main.go:10:6: cannot use i8 (type int8) as type int16 in assignment

Compilation finished with exit code 2
```

Fixing it will require an explict conversion as in:

```
	var a int8 = 4
	var b int16 = int16(a)
```

Go being helpful as it is allow you to even drop the type from the declaration since the explict conversion is enough to tell the type. This means the above code can be written as:

```
	var a int8 = 4
	var b int16 = int16(a)
```

Or even, more succintly as:

```
	var a int8 = 4
	b := int16(a)
```

You should be very careful when converting between numeric types. Go will silently overflow large values when using an explicit conversion. Let's see some examples. The largest valid value for int8 is 127. Let's try to assign it 300 in different ways and see what happens.

Directly assigning too large value resulta in a clear error (Good!):

```
var x int8 = 300

Output:

manuscript/code/chapter-3/main.go:15:6: constant 300 overflows int8

Compilation finished with exit code 2
```

Trying to convery it at the point of assignment also results in error (Good!)

```
var x int8 = int8(300)

Output:

manuscript/code/chapter-3/main.go:15:19: constant 300 overflows int8

Compilation finished with exit code 2
```

But the following code silently compiles and produces a wrong result:

```
package main

import (
	"fmt"
)

func main() {
	v := 300
	var x int8 = int8(v)
	fmt.Println(x)
}

Output:

44
```

What just happend? First we assigned 300 to v. The type of v in this case is the platform-specific integer 32-bit or 64-bit. Then in the next line we explicitly converted v to int8. This conversion overflows and the result is that x contains 300 % 256 = 44 (the modules operator).

It might occur in real-wprld code more often then you think when serializing data using across machines or between different languages or in general when trying to optimize and conserve space.


### Boolean

Go is strict with its booleans. There are two constatns `true` and `false` and they are the only valid values for a boolean variable.
Many other languages like C/C++, Python and Javascript automatically convert different expressions to bool. C/C++ treat 0 as false and everything else as true. Python is even crazier you can define for each type a special `__nozero__` (Python 2.x) or `__bool__` (Python 2.x) that returns True or False. In particular in Python None (the equivalent of Go's nil) and numeric values of 0 are considered as well empty strings and collections (sets, lists, dictionaries). Javascript has its own unintuitive definition of what's considered [truthy](https://developer.mozilla.org/en-US/docs/Glossary/Truthy)

In Go the value of an uninitialized boolean variable is false (its zero value). But you can't assign 0 or even nil to a boolean.

```
	var b bool   // good! b is false
	b := true    // good! b is true
	b := 0       // fail!
	b := nil     // fail!
```

Boolean expressions and variables are used a lot in conditionals that we will explore in depth later.

### Arrays and Slices

Arrays are a bedrock data strucuture for any programming language. Go arrays have contiguous blocks of memory with fixed capacity that store elemnts of the same type. Arrays are one of a few generic types in Go. This means array elements can be of any type. Here is a little program that illustrates how to declare arrays:

```
package main

import (
	"math"
	"fmt"
)

func main() {
	var ints [4]int
	var pies = [2]float64 {3.141592, math.Pi}
    bools := [6]bool{true, false, true}

	fmt.Println("ints:", ints)
	fmt.Println("pies:", pies)
	fmt.Println("bools:", bools)
}

Output:

ints: [0 0 0 0]
pies: [3.141592 3.141592653589793]
bools: [true false true false false false]
```

Note, that you have to specify the type of the element and the capacity of the array and it never changes. `[4]int32` is an arraty capable of of storing 4 int32 values. You can initialize an array when you declare as in [2]float64 {3.141592, math.Pi}. Elements that are not initialized explicitly will store the zero value for the element type (0 for numbers, false for boolean, etc).


#### Initializing Arrays


If you fully initialize an array you can let Go set the capacity for you, by using elipsis. Here is an example that also demonstrates the len() and cap() functions to determine the length anf capacity of an array.

```
x := [...]int8{1, 2, 3, 4, 5}
fmt.Println("x:", x, "length:", len(x), "capacity:", cap(x))


Output:

x: [1 2 3 4 5] length: 5 capacity: 5

```

#### Accessing Array Elements

You access array elements by indexing them. Go uses 0-based indexing of course in the tradition of C. You can declare multi-dimensional arrays too. The following snippet defines a 2x4 array, populates it and prints it to the screen

```
    multi := [2][4]int{}
    for i := 0; i < 2; i++ {
    	for j := 0; j < 4; j++ {
    		multi[i][j] = i + j
		}
	}

    fmt.Println(multi)

Output:

[[0 1 2 3] [1 2 3 4]]
```


#### Array Type

It's very important to understand that the type of an array includes its size. These are different types:

- [2]string
- [3]string


```
s2 := [2]string{}
var s3 [3]string = s2 // fail


Output:

manuscript/code/chapter-3/arrays_and_slices/main.go:36:6: cannot use s2 (type [2]string) as type [3]string in assignment

```


#### Slices

Arrays are useful and necessary. But, their most important role is as a backing store for slices. Slices looks like dynamic arrays. You declare them without a capacity and you can add or remove elements from slices. However, since slice elements are actually stored in arrays if you append an element to a slice whose backing array is at capacity then a new backing array will be allocated to the slice and all the slice elements will be copied.

```
s := []int{1,2,3}
fmt.Println("s:", s, "length:", len(s), "capacity:", cap(s))
s = append(s, 4)
fmt.Println("s:", s, "length:", len(s), "capacity:", cap(s))

Output:

s: [1 2 3] length: 3 capacity: 3
s: [1 2 3 4] length: 4 capacity: 6
```

Note that when I added another element to the slice the new backing array didn't have just enough capacity forthe new element (4), but actually doubled from 3 to 6 to accomodate appending a few more elements. Go uses heuristics to determine how to grow slices. But, you shouldn't rely on it because it's a pretty low-level implementation detail that can change. If you want predictability you can allocate a slice with predefined capacity using the `make()` function:

```
bytes := make([]byte, 100)
fmt.Println("length:", len(bytes), "capacity:", cap(bytes))

Output:

length: 100 capacity: 100
```

As you can see the slice has a length of 100. If you want an empty slice backed by an array of 100 elemnts use:

```
bytes := make([]byte, 100)[0:0]
fmt.Println("length:", len(bytes), "capacity:", cap(bytes))

Output:

length: 0 capacity: 100
```

#### Slice Type

The type of slice depends only on its element type. So a slice of strings is always the same type as another slice of strings even
if their backing arrays are of different types and thuse incompatible.

```
a2 := [2]string{"a", "b"}
a3 := [3]string{"c", "d", "e"}

s2 := a2[:]
s3 := a3[:]

fmt.Println(s2)
s2 = s3
fmt.Println(s2)

Output:

[a b]
[c d e]
```

#### Appending, Copying and Removing Elements

Slices are dynamic even if their backing arrays are immutable. The `append()` function can be used to append one or more elements to a slice as it is a variadic function (takes variable number of elements). The following snippet demnostrates different ways of appending elements (or slices) to a slice:

```
bytes := []byte{}
bytes = append(bytes, 1) // append one element
fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))
bytes = append(bytes, 2, 3, 4) // append multiple elements
fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))
bytes = append(bytes, []byte{5,6}...) // append another slice
fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))
bytes = append(bytes, bytes...) // append the slice to itself
fmt.Println("bytes", bytes, len(bytes), "capacity:", cap(bytes))

Output:

bytes [1] 1 capacity: 8
bytes [1 2 3 4] 4 capacity: 8
bytes [1 2 3 4 5 6] 6 capacity: 8
bytes [1 2 3 4 5 6 1 2 3 4 5 6] 12 capacity: 16
```

You can copy one slice into another using the copy() function. Unlike append, it will only copy as many elements from the source that will fit in the destination. The following snippet shows how to copy elements of the [Fibonacci sequence](https://www.mathsisfun.com/numbers/fibonacci-sequence.html) between slices.

```
fib := []byte{0, 1, 1, 2, 3, 5, 0, 0, 0, 0, 0}
fmt.Println("fib", fib, len(fib), "capacity:", cap(fib))
fib2 := []byte{8, 13, 21, 34, 55}
fmt.Println("fib", fib2, len(fib2), "capacity:", cap(fib2))
copy(fib[6:8], fib2) // copy only two elements
fmt.Println("fib", fib, len(fib), "capacity:", cap(fib))

copy(fib[6:], fib2) // copy all elements
fmt.Println("fib", fib, len(fib), "capacity:", cap(fib))

Output:

fib [0 1 1 2 3 5 0 0 0 0 0] 11 capacity: 11
fib [8 13 21 34 55] 5 capacity: 5
fib [0 1 1 2 3 5 8 13 0 0 0] 11 capacity: 11
fib [0 1 1 2 3 5 8 13 21 34 55] 11 capacity: 11
```

As you can see sub-slices like fib[6:8] share the same backing array and when they are modified the original fib slice is also modified because it's just pointing to the same array. But, appending to a sub-slice will create a new array for that sub-slice and the original slice will not be changed.

```
fib3 := fib[6:8]
fib3[0] = 77 // modifies the original backing array of fib
fmt.Println("fib:", fib)
fmt.Println("fib3:", fib3)

fib3 = append(fib3, 88) // fib3 now has a new array
fmt.Println("fib:", fib)
fmt.Println("fib3:", fib3)

Output:

fib: [0 1 1 2 3 5 77 13 21 34 55]
fib3: [77 13]
fib: [0 1 1 2 3 5 77 13 88 34 55]
fib3: [77 13 88]
```

That means you need to be careful when appending and modifying sub-slices. Depending on the capacity of the backing array you may modify the original slice (its backing array) or a fresh copy. Unless you want to modify shared storage I suggest in most cases to work on fresh copies.


Removing elements from a slice is really joining two sub-slices without the element or elements to remove. There is no `remove() function`.

```
a := [...]byte{1,2,3,4,5}
s := a[:]
fmt.Println("a", a, len(a), "capacity:", cap(a))
fmt.Println("s", s, len(s), "capacity:", cap(s))
s = append(a[:1], a[2:]...)
fmt.Println("a", a, len(a), "capacity:", cap(a))
fmt.Println("s", s, len(s), "capacity:", cap(s))

Output:

a [1 2 3 4 5] 5 capacity: 5
s [1 2 3 4 5] 5 capacity: 5
a [1 3 4 5 5] 5 capacity: 5
s [1 3 4 5] 4 capacity: 5
```

If you look carefully, you can see that the original backing array is still used, but all the elements after 2 got shifted to the left. That preserves the order, but for large slices can potentially be very expensive. If you don't care about the order you implement a quick remove that copies the last element of the slice onto the place of the removed element and resizes the slice:

```
a := [...]byte{1,2,3,4,5}
s := a[:]
fmt.Println("a", a, len(a), "capacity:", cap(a))
fmt.Println("s", s, len(s), "capacity:", cap(s))
s[1] = s[len(s)-1]
s = s[:len(s)-1]
fmt.Println("a", a, len(a), "capacity:", cap(a))
fmt.Println("s", s, len(s), "capacity:", cap(s))

Output:

a [1 2 3 4 5] 5 capacity: 5
s [1 2 3 4 5] 5 capacity: 5
a [1 5 3 4 5] 5 capacity: 5
s [1 5 3 4] 4 capacity: 5
```

As you can see backing arrays never shrink. As long as some slice is pointing to an array, the entire array will remain in memory even if the slice only contains one element of the array. That means you should watch out unused backing array lingering around and dispose of them.


### Bytes, Runes and Strings

Bytes, runes, and strings are the building blocks of text processing in Go.

#### Bytes

Bytes are 8-bit numbers. Each byte can represent one of a possible 256 values (2 to the power of 8). Each character in the ASCII character set can be represented by a single byte. But bytes are not characters. The reason is that Go as a modern language supports Unicode, where there are way more than 256 separate characters. Enter runes.

#### Runes

A rune in Go is another name for the int32 type. This means that each rune can represent more than four billion separate values (2 to the power of 32), which is good enough to cover the entire Unicode character set.

In the following code you can see that the rune '∆' (alt-J on the Mac) is just an int32. To print the character it represents to the screen, I have to convert it to a string.

```
package main

import (
    "fmt"
)


func main() {
    r := '∆'
    fmt.Println(r)
    fmt.Println(int32(r))
    fmt.Println(string(r))
}

Output:

8710
8710
∆
```

Unicode is complicated. A rune officially represents a Unicode code point. Unicode characters are usually represented by a single Unicode code point, but sometimes more than one. This means some Unicode characters will require more than one rune.

### Strings

Strings are officially just read-only slices of bytes. If you index a string, you get a byte back:

```
func main() {
    s := "abc"
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
    }
}

Output:

97
98
99
```

String literals are a sequence of UTF-8 characters enclosed in double quotes. They may contain escape sequences, which are a backslash followed by an ASCII character such as \n (newline) or \t (tab). Each escape sequence has a special meanings. Here is the full list:

```
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000b vertical tab
\\   U+005c backslash
\'   U+0027 single quote  (valid only within rune literals)
\"   U+0022 double quote  (valid only within string literals)
```

Sometimes you may want to store literal bytes directly in a string, regardless of escape sequences. You could escape each backslash, but that's tedious. A much better approach is to use raw strings that are enclosed in backticks.

Here is an example of a string with a \t (tab) escape sequence, which is represented once as is, then with the backslash escape, and then as a raw string:

```
func main() {
    s1 := "1\t2"
    s2 := "1\\t2"
    s3 := `1\t2`

    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3)
}

Output:

1    2
1\t2
1\t2
```
While strings are slices of bytes, when you iterate over a string with a for-range statement, you get a rune in each iteration. This means you may get one or more bytes. This is easy to see with the for-range index. Here is a crazy example. The Hebrew word "שלום" means "Hello" (and peace). Hebrew is also written right to left. I'll construct a string that mixes the Hebrew word with its English translation.

Then, I'll print it rune by rune, including the byte index of each rune within the string. As you'll see, each Hebrew rune takes two bytes, while the English characters take one byte, so the total length of this string is 16 bytes, even though it has four Hebrew characters, three symbols, and five English characters (12 characters). Also, the Hebrew characters will be displayed from right to left:

```
func main() {
    hello := "שלום = hello"
    fmt.Println("length:", len(hello))
    for i, r := range(hello) {
        fmt.Println(i, string(r))
    }
}

Output:

length: 16
0 ש
2 ל
4 ו
6 ם
8
9 =
10
11 h
12 e
13 l
14 l
15 o
```

When printing strings and byte slices, there are several format specifiers that work the same on both. The %s format prints the bytes as is, %x prints two lowercase hexadecimal characters per byte, %X prints two uppercase hexadecimal characters per byte, and %q prints a double quoted string escaped with go syntax.

To escape the % sign inside a format string specifier, just double it. To separate the bytes when using %x or %X, you can add a space, as in "% x" and "% X". Here is the demo:

```
func main() {
    s := "שלום"

    fmt.Printf("%%s format:  %s\n", s)
    fmt.Printf("%%x format:  %x\n", s)
    fmt.Printf("%%X format:  %X\n", s)
    fmt.Printf("%% x format:  % x\n", s)
    fmt.Printf("%% X format:  % X\n", s)
    fmt.Printf("%%q format:  %q\n", s)
}

Output:

%s format:  שלום
%x format:  d7a9d79cd795d79d
%X format:  D7A9D79CD795D79D
% x format:  d7 a9 d7 9c d7 95 d7 9d
% X format:  D7 A9 D7 9C D7 95 D7 9D
%q format:  "שלום"
```

### Maps



### Interfaces

### Structs

### Functions

### Pointers

Pointers in Go are not as prevalent as in C/C++. Their usage is sometimes masked becuase you don't need to invoke methods on object pointers with an arrow notation like in C/C++.


### Channels

## Type Declaration and Aliasing		

## Type Assertions

## Reflection
