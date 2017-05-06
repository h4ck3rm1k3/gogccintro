Create optimal representation of static data.
Generate Go Language code to represent data.
Create structure types for each unique type.
       Create one go types per unique set of fields/node types.
       Each pointer can be typed to exactly the right one.
       Do a toplogical sort of unique types and create a graph of types and how they are used.
       Create specializations of types for each path in the graph.

       

Create unique naming system that contains the field types.
Create type hierarchy of sub/super types that contain more or less of the same fields. Lattice.

Create tables of records. Create tables of columns.
Create enums with probability for each variant type.
Is a field a containment or a pointer. 1-m, 1-1, m-1? Determine cardinality.
Create pairs of common fields in structs where they coccur.
Create tuples of common fields in structs where they coccur.
identify disjoint fields that never occur with each other.
Basic types : flag, integer, string, filename, line number.
List of Tables : list of columns. List of rows. List of pointers between rows, types, tables.
Determine all the forms of a give table.
Determine all the types referenced of a given column.




Reflection that does not suck.
       Create inline reflection data
       Create static reflection data for each type.
       Fast name lookup : dfa/parser or map tables or static switch.
       Generate parser for each type to match instances(json)
       support existing reflection apis.

Find reoccuring structures.
find patterns that occur again and again. Compare two objects and look for differences.
replaces common data with code. Re-Encode.
Create clusters of the date.

Look for differences in values :
     int
	+/- constants. 
	*/% scale
     string
	identifier :
		   camel case
		   underscore, different naming conventions
		   prefix, suffix, middle
		   shift string
		   duplicate strings
		   

Support common interfaces :
	upper sets of all fields that occur togeher for a give type
	common sets of fields togher.
	common fields of two types.
	Collect type trees, collect value trees.
	Walk graph with only a subset of fields.
	fields common for an operation.
	generate interfaces for fields needed, and match those against instances.
	set of all fields :
	is this field set :
	populate reflection data :

What functions use what fields.
