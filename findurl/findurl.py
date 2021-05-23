log = """ 
hello world
this is text in python language
www.google.com
https://digikala.com
http://www.google.com
www.saminejad.ir
"""

for i in log.split("\n"):
    if i.find("www")!=-1 or i.find("http")!=-1:
        print(i)
