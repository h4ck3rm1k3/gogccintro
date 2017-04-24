
The usage of reflection is 10x slower than using a dfa to lookup a field.
Using a switch on the fieldname is just a bit slower than the dfa and simpler to code.

BenchmarkSwitch-8    	 1000000	      1452 ns/op
BenchmarkReflect-8   	  100000	     14267 ns/op
BenchmarkDFA-8       	 1000000	      1133 ns/op
BenchmarkHash-8      	 2000000	       752 ns/op

