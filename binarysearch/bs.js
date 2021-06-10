
let x = [10,24,56,78,99,124,256,345,390,440,450,600,660,770,799,801,920]

let top = x.length , bottom = 0 , counter = 0 , search = 799

while (1){

    let middle = Math.floor((top + bottom)/2)
   
    if (x[middle-1]==search) 
    {
        console.log(`index is ${middle-1}`)
        break
    }
    else if (x[middle-1] > search){
        top = middle -1
    }
    else {
        bottom = middle + 1
    }

    counter ++ // if not found break the loop
    if (counter >= x.length){
    console.log('not found')
    break
    }
}