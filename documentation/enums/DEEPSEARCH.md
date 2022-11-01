![Banner](../../images/gostack_SmallerTransparent.png)

<h2>DEEPSEARCH</h2>

 > True, False

Many functions have "deep search functionality".  This means they A) have a DEEPSEARCH parameter called `deepSearchType` and B) have an int/[]int/Stack{ints} parameter called `depth` which will act as a guide for which layers to consider in the deep search.  Sample:

 > enumerator type's name
 >> how the function will behave in its search

Enumerators:
 > DEEPSEARCH_True
 >> the function will listen to the `depth` parameter input in considering how deep it searches

 > DEEPSEARCH_False
 >> the function will set `depth` to 1, only considering the immediate children of this stack

 ---

 [> Return to enumerators](../enumsAPI.md)