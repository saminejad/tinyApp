x = [10,24,56,78,99,124,256,345,390,440,450,600,660,770,799,801,920]

top ,bottom , counter ,search = len(x), 0,0,770

while 1:
    middle = (top + bottom) //2

    if x[middle -1]== search :
        print(f'index is {middle-1}')
        break
    elif x[middle-1] > search :
        top = middle -1
    else:
        bottom = middle +1
    
    counter +=1

    if counter >= len(x):
        print("not found")
        break