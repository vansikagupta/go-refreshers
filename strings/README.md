# String in Go 

### Literal Representation
* **"something"** supports escape characters
* **\`somethingelse\`** doesn't support escape characters. Can contain any character other than `\
See example

### Ranging through string
* String in go is immutable sequence of characters. 
* Characters in go are represented in UTF8 format. The character value in UTF8 format is called a code point.
* ASCII characters in UTF8 format occupy 1 byte or 8 bits.
* Non-ASCII characters can range upto 4 bytes or 32 bits.
* String is a read-only slice of bytes. \
    Number of bytes in string, i.e, len(string) != Number of characters. Reason: see above point.
* Rune is a datatype of 4 bytes, int32. It can represents any character's code point.
* Range loop on string returns index of byte and code point of the character as a rune type(int32)\
See example

*byte is an alias for unit8 and rune is an alias for int32; can be used interchangeably*

### Character comparision
A character is represented as rune, which is comparable
* character in single qoutes is a rune and it's type is int32.


## Package string
Homework: https://www.restapiexample.com/golang-tutorial/common-string-problems-solution-go/