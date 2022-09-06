# [26. Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/)



## Solution

```

Array example:

[]int{1, 1, 1, 2, 3, 3, 4, 5, 5, 5, 5, 6, 8, 9, 10, 10, 11, 11, 12}
```

We start at the array's first element, backward to search the left numbers, if the next number equal the start number, skip it.
or the next number not equal the start number, move the search number behind start number.

### steps:
### todo animation
#### step 1
1, **1**, 1, 2, 3, 3

start key is 0,  value is 1

search key is 1,  value is 1

values is equal, skip, search key increase by 1 become to 2

1, **1**, 1, 2, 3, 3

search key become to 2

#### step 2
1, 1, **1**, 2, 3, 3

start key is 0,  value is 1

search key is 2,  value is 1

values is equal, skip, search key increase by 1 become to 3


search key become to 3

#### step 3

1, 2, 3, **2**, 3, 3

start key is 0,  value is 1

search key is 3,  value is 2

values is not equal, change the number behind start  to 2

1, 1, 1, 2, 3... => 1, 2, 1, ~~2~~, 3, 3...

search key become to 4

start key become to 1

#### step 4

1, 2, 3, ~~2~~, **3**, 3

start key is 1,  value is 2

search key is 4,  value is 3

array change:

1, 2, 1, ~~2~~, 3, 3... => 1, 2, 3, ~~2~~, ~~3~~, 3...

search key become to 5

start key become to 2

#### step 5

1, 2, 3, ~~2~~, ~~3~~,**3**

start key is 2,  value is 3

search key is 5,  value is 

1, 2, 1, ~~2~~, 3, 3... => 1, 2, 3, ~~2~~, ~~3~~, 3...


and so on
