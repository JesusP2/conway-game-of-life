Arrays:
Arrays are data structures that have 2 things:
a length and a pointer its first element

A slice is a view of an array, it consists of 3 things:
length, capacity and a pointer to the array. The pointer to the array itself is a pointer to the first element of the array I believe

When you create a slice from an array, let's say s[:n], this slice is going to be created with the
next properties:
length: n;
capacity: len(array);
pointer: the same as the array

Now let's explain the reason for each one:
*length*: This basically is the difference between the left and right side of the slice,
    n - 0 = n in this case.
*capacity*: If length is the number of elements the slice has, capacity is the max number of elements
    a slice can have. Why is this important? well I have a reason, but I don't know if this is the only
    reason or if it's actually a reason.
    Arrays length are immutable, so to create a mutable ordered data structure you use a slice,
    you can either create more slices from this slice or append values to it.
    When you append a value to a slice you'll experience multiple behaviours:
    if len(append(slice, n1, n2..., nx)) is less than capacity of slice, then everything is good,
    append will just assign the values to the next elements in the array. Lets see an example:
    ```golang
    array := [5]int{1, 2, 3}
    s1 := array[:1]
    s2 := append(s1, 1, 2)
    fmt.Println(array) // [1, 1, 2]
    printSlice(s1) // [1]
    printSlice(s2) // [1, 1, 2]
    ```
    Now the second behaviour appears when your append exceeds the capacity...
    If you try to do that you'll end up copying the array, this new array is going to have len * 2
    so this means the slice is also going to have a capacity of len * 2, length of capacity will just be
    previous capacity + 1.
    ```golang
    array := [5]int{1, 2, 3}
    s1 := array[:1]
    s2 := append(s1, 1, 2, 3, 4, 5)
    fmt.Println(array) // [1, 2, 3]
    printSlice(s1) // [1]
    printSlice(s2) // [1, 1, 2]
    ```
    All this is just for s[:n], s[n:] has also cool things:
    s[n:] reduces the capacity of the slice by n and also changes the pointer. Weird? not really.
    This is expected if you want this slice to behave accordinly. This new slice first element
    is the n element of the array, so that's why the pointer changes, and the reason for the capacity
    to shrink is to make up for the 2 cases mentioned above. Once you reach the capacity then
    you need to create another array because you don't have space left in this one.
