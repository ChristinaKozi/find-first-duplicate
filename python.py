
def find_first_duplicate_1(arr):
    checked = []
    for i in range(len(arr)):
        if arr[i] in checked:
            return(arr[i])
        checked.append(arr[i])
    return -1

print(find_first_duplicate_1([2, 1, 3, 3, 2]))