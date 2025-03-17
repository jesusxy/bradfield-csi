## The global `pcap-savefile` Header

Whats the magic number? What does it tell you about the byte ordering in the pcap-specific aspects of file?

- The magic number is `d4c3b2a1`
- This tells me that the host who wrote the file used little endian, since my machine uses big-endian
- I will have to swap the bytes after the magic number in order to read the contents in the native byte order

What are the major and minor versions?

- Major version: 2
- Minor version: 4

Are the values that ought to be zero in fact zero?
Yes, the values that ought to be zero are zero in the header

What is the snapshot length?

The snapshot length is `1514` bytes

What is the link layer header type?

`LINKTYPE` = 1 = Ethernet

### Per-packet Headers

---

What is the size of the first packet?

- 78 bytes

Was any data truncated?

No data was truncated. Captured packet length == Original packet length
