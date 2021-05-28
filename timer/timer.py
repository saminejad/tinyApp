import time , os

m=0
s=0
while True:
    time.sleep(1)
    os.system("clear")
    s+=1
    print(f'{m} m : {s} s')
    if s==59:
        m+=1
        s=0
    