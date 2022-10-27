![Banner](../images/gostack_SmallerTransparent.png)

<h1>Benchmarks</h1>

To be entirely honest, there is a significant performance sacrifice in using ***gostack*** over ***native Go***.  Thus, optimization be the main focus of updates in future releases.  This sacrifice, however, especially apparent when performing basic tasks, but the performance difference exponentially alleviates as tasks become more complicated.

> My sample task: 123 *(gostack is this many times slower than native Go at doing this)*

Benchmarks:

> Creating an array/stack: 1137

> Creating a map/stack: 69

> Getting an element in a search: 840

> Adding an element in a search: 58

---

 [> How many lines does gostack save compared to native Go?](race.md)

 [> Return to glossary](../README.md)