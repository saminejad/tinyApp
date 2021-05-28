const crpyto = require('crypto')

let hash = crpyto.createHash("md5")

hash.update(Math.random().toString())

console.log(hash.digest("base64").slice(0,10))