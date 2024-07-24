# Coralogix
Let’s say we want to implement a library of operations on big csv files (meaning tables that won’t necessarily fit in their entirety into memory). You can assume you have a library (encoding/csv) which implement the next functionality for reading csv files-



And also the next functionality for writing - 


You are asked to implement a library with the next functionality -
We want to be able to read an entire table and do some operations over it, then write the results back to another csv file. The inputs and output of every operation are tables, but we can control the flow of data between operations.
For instance - 




In this example, we filter all the rows by some dynamic logic, then get all the 10th columns (which will contain numbers), then operate an average operation and then ceiling the results and write it to a file.

Another example -



In this example we apply some operation for every cell in a row (double its value), then picking columns 3-5, then rows 7-20, summing every row and writing it to another file.

The main concerns and assumptions - 

1. As mentioned, the entire table won’t necessarily fit into memory (But you can assume every line will fit into memory).
2. You can assume that the operations (which you need to support) are those that running on a row produces a row (and not a column for example).
3.The implementation should be flexible so others can implement/extend the library.
4.You can assume the input is valid.

As a first step, write down all the structures/interfaces involved. Take into consideration the elements we discussed in the first part of the interview (SOLID principles).
As a second step, implement two of the functions you want.

