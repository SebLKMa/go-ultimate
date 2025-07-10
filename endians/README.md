# Endianness

Endianness is about how computers store multi-byte data (like integers) in memory.

## Two Types:
Big Endian: Stores the most significant byte first (big end first).  

Little Endian: Stores the least significant byte first (little end first).  

Example (number = 0x12345678):  
This number has 4 bytes: 12, 34, 56, 78.  

**Big Endian**  
``` 
Address:   0   1   2   3  
Data:     12  34  56  78
```

**Little Endian**  
```
Address:   0   1   2   3  
Data:     78  56  34  12
```

That’s it! It’s just the ***order of bytes*** when storing multi-byte data.  
Most Intel CPUs use little endian.  

## Example integer 65534

What is 65534 in hexadecimal?  
```
65534 = 0xFFFE
```

That’s 2 bytes:  
```
FF (most significant byte)
FE (least significant byte)
```

### Memory representation
Big Endian (most significant byte first):  
```
Address:   0    1  
Data:     FF   FE
```
Little Endian (least significant byte first):  
```
Address:   0    1  
Data:     FE   FF
```

So:  
Big Endian    = FF FE  
Little Endian = FE FF  
