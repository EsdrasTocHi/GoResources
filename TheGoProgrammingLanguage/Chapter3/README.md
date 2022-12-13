# Other Notes

### Categories of data types
- Basic types
    include numbers, strings, and booleans
- Aggregate types
    arrays and structs
- Reference types
    pointers, slices, maps, functions and channels
- Interface types

### Precedence of binary operations


| | | | | | | |
|------|----|----|----|----|----|----|
| *    | /  | %  | << (left shift) | >> (right shift) | & (bitwise and)  | &^ (bit clear)|
| +    | -  | \| (bitwise or) | ^ (bitwise xor)  |    |    |    |
| ==   | != | <  | <= | >  | >= |    |
| &&   |    |    |    |    |    |    |
| \|\| |    |    |    |    |    |    |

Each operator in the first two lines of the table above, for instance +, has a corresponding assignment operator like += that may be used to abbreviate an assignment statement.

The operator ^ is bitwise exclusive OR (XOR) when used as a binary operator, but when used as a unary prefix operator it is bitwise negation or complement; that is, it returns a value with each bit in its operand inverted. The &^ operator is bit clear (AND NOT): in the expression z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; otherwise it equals the cor- responding bit of x.
