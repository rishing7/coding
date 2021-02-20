myList = [1, 3, 6, 8, 9, 3, 5]
print("My list is: ", myList)
print("My sorted list is: ", sorted(myList))
print("Size of list is: ", len(myList))
print("Iterating list: ", [myList[i] for i in range(0, len(myList))])
print("Search an item: ", 3 in myList)
print("Index of an item: ", myList.index(3))  # Left most index
