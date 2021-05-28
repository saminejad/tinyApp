
let m = 0 ,s = 0

let p = new Promise(res=>{
    
})
 async function f(){

    while(true){
    await   new Promise(res=>{


        setTimeout(() => {
            s++
            if (s==59){
              m++
              s=0}
            res([m,s]) 
              
          }, 1000);



    }) .then(data=>{
        console.clear()
        console.log(`${data[0]} m : ${data[1]} s`)
        
    })
}
 }

 f()

    

    

