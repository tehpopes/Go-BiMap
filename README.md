This repository contains a Bidirectional Map Implementation in Go

Details
 - A bidirectional map maintains unique keys and values
 - Uses two maps, one mapping keys to values and another mapping values to keys
 - insert key value pair
    - inserting duplicate key value pair does nothing
    - inserting pair with key in BiMap but not value replaces old value with new value
    - inserting pair with value in BiMap but not key replaces old key with new key
    - returns new value if successful and nil if failure
 - find and remove can be done by key or by value
    - returns nil if nothing found