# Exercises

## 1.1 Simple Conversion

Convert 9, 136, 247 into hexadecimal

```
dec     -   9
binary  -   1001
hex     -   0x9

dec     -   136
binary  -   1000 1000
hex     -   0x88

dec     -   247
binary  -   1111 0111
hex     -   0xF7
```

## 1.2 CSS colors

Guess the color?

```
hex     -   #B76D47
binary  -   1011 0111 0110 1101 0100 0111
            10110111  01101101  01000111

decimal -   183         109         71

color   -   Tufted Leather
```

Each of the 3 bytes in `rgb()` is represented by 8 bits.

```
8 * 3 = 24

2^24 = 16,777,216 colors

```

### Say hello to dex

Convert first 5 bytes of hellodex to binary

```
hex     -   68656c6c6f

binary  -
0110 1000 - 0110 0101 - 0110 1100 - 0110 1100 - 0110 1111

```

## 2.1 Simple Conversion

Convert the following decimals to binary

```
4       -   0000 0101

65      -   0100 0001

105     -   0110 1001

255     -   1111 1111
```

Convert the following unsigned ints to decimal

```
10      -   2

11      -   3

1101100 -  108

1010101 -  85

```

## 2.2 Unsigned Binary Addition

When adding binary numbers, we have to make sure we carry.

We are restricted to 1 and 0.

2 -> binary `10` - we would place the 0 and carry a 1 \
3 -> binary `11` - we would palce the 1 and carry a 1

```
11111111 + 00001101 = ??

carry  111111110
        11111111 -  255
        00001101 -  13
----------------
        00001100 -  12 ?

carry out -> 1
```

Expected return value = `268`

The actual returned value from this addition would be: `00001100` || `12`.\
This is referred to as an overflow. The max calue of 8bit register is 255.
Anything past that overflow and starts back from 0.
Much like an odometer in a car.

## 2.3 | 2s Complement Conversion

Give the 2s complement representations of the following integers

```
Decimal      Binary         1s Complement         2s Complement
127         0111 1111           1000 0000             1000 0001

-128        1000 0000           0111 1111             1000 0000

-1               0001                1110                  1111

1                0001                1110                  0001

-14        0000  1110           1111 0001             1111 0010
```

Convert the following 8-bit two’s complement numbers to decimal:

```
2s Complement                           Decimal

1000 0011            -128 + 3              -125


1100 0100           -128 + 68               -60
```

## 2.4 Addition of 2s complement Signed integers

```
0111 1111    (127)
1000 0000 + (-128)
------------------
1111 1111     (-1)
```

To negate a number in twos complement, you must first invert the bits
to its 1s complement then add 1.

The values of the most significat bit in 8bit twos complement is `-128`
The avlue of the most signifcant bit in 32bit twos complement is `-2,147,483,648`

## 2.5 Advanced: Integer overflow detection

There will be overflow if the `MSB` of each bit pattern is the same.

Example:

```
010000000 (carry row)
01000000 (64)
01000000 (64)

100000000 (carry row)
10000000 (-128)
10000000 (-128)
```

### _Answer_

`XOR(carry_in, carry_out).`
When there is a carry in but no carry out, overflows to `negative`
Where there is a carry out but no carry in, overflow to `positive`

## 3.1 Its over 9000

The number `9001` uses `big endian`

## 3.2 TCP

Sequence number: `01000100 00011110 01110011 01101000` -> `1142846312` \
Acknowledgement number: `11101111 11110010 10100000 00000010` -> `4025655298` \
Source port: `10101111 00000000` -> `44800` \
Desination port: `10111100 00000110` -> `48134`

## 4.1: IEEE Floating Point: Deconstruction

Identify the three components of 32 bit floating point number
`0 10000100 01010100000000000000000`

```
sign bit               exponent                fractional bits
0                      10000100              01010100000000000000000
                         132                        .328125
                      132 - 127                   1 + .328125
                          5                         1.328125

1.32815 * (2**5) = 42.5
```

What is the smallest (magnitude) incremental change that can be made to a number?

`0 11111110 0000000000000000000000` \
To make the smallest incremental change, you flip the least significant bit.\
`0 11111110 0000000000000000000001`

For the smallest (most negative) fixed exponent, what is the smallest (magnitude) incremental change that can be made to a number?

`0 00000000 0000000000000000000000` \
To make the smallest incremental changes to the number above, you flip the LSB \
`0 00000000 0000000000000000000001` \

What does this imply about the precision of IEEE Floating Point values?

Since we allocate 23 bits for the `mantissa` it means IEEE is high precision. \
If there were more bits for the exponent the `range` of numbers that could be represented would
increase.
