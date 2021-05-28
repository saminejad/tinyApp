import hashlib

print(hashlib.md5(str("hello".__hash__()).encode('UTF-8')).hexdigest()[0:10])
