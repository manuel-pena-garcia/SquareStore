# SquareStore
SquareStore is a Store System based on the distribution of numbers in a square in 2D. It should be applied to cases where a great number of elements will be added, needing randomly access, and where the elements should be persistent (almost never deleting an element).

25 24 23 22 21
16 15 14 13 20
9  8  7  12 19
4  3  6  11 18
1  2  5  10 17

The insertion will be always at the end of the structure. The insertion cost is O(log(N)), as it has to make special checks to retrieve if the positions belongs to any of the main axis or the main diagonal.

Reading will be done using the best possible case, as it can be performed in both main diagonal and main edges, so the complexity will be reduced to O(sqrt(N)).

The deletion will be performed logically so the structure will not be rearranged.
