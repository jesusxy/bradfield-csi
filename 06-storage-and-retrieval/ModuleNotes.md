## Persistance and Write-Ahead Logs

### SSTable implementation

Sorted String Table:

- how do we know when the skiplist is full in order for us to flush?
- when do we call flush?

> Keep track of a threshold

- after threshold is surpassed call `db.FlushToDisk()`
- on PUT and DELETE, increment an internal counter
- threshold = max # of bytes
- if totalKeyBytes + totalValBytes > threshold ... flush()

> GET from SSTable

When we dont find a key in the memtable, we would then search the SSTable
that are on disk one at a time. we will call GET on the immutable DB for that file

- we would need an index to improve searching through the SSTable file(s)

#### Who is responsible for opening / closing SSTable files?

- should it be the caller / client?
- or should it be the SSTable instance itself

#### Tombstone(s)

- might need to change binary format in SSTable
- the bit will represent if the item was "deleted"
- will also need to add the bit on the Skiplist Node

#### Merge Multiple Iterators for RANGE scans?

- many iterators merged into One iterator
- merge to create a single stream of sorted results
- literally Leetcode "merge k sorted lists"

Example use case:

```go
// top level db
db.RangeScan()
    var children []Iterator
    children = append(chilrdren, memtable.RangeScan())
    children = append(children, sstables[0].RangeScan())
    children = append(children, sstables[1 ].RangeScan())

    return newMergingIterator(children)
```

You can also use this approach to merge multiple SSTables into a single SSTable. Like when we are going to merge Level 0 down to Level 1.

#### Compaction | Rewriting multiple SSTable files

> Motivation: We want to structure things so that when we are trying to find a key, we only have **one** file per level to look at

SSTable data organization:

- organize data into levels
- within each file, has a unique range of keys
- no overlapping with each other

_What if we have a single GIANT SSTable?_

- very efficient for reads
- bad for writes every time we merge a new SSTable, we would be rewriting the entire file (each time we flush)

_What if we have a collection of SSTables, each created by a mem flush?_

- bad for reads (potentially checking every single file that has been created)
- good for writes

SSTable data structure
`[key,value|data.....|sparse index]`

- sparse index gives you offsets into various points of the data
