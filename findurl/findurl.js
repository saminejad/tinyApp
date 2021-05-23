let s = `
hello world
this is text in python language
www.google.com
https://digikala.com
http://www.google.com
www.saminejad.ir`

for (let i of s.split("\n")){
    if (i.includes('www') || i.includes('http'))
        console.log(i)
}