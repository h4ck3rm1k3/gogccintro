my $buf="";
my $c = 0;
while (<>)
{
    if (/ <\- /) {
	$c++;
	if (length($buf) < 2) {
	    next;
	}

	print ">test_".$c.".peg\n";
	open OUT,">test_".$c.".peg";
	print OUT q!
package main
type GccNode Peg {}
!;
	print OUT $buf;

	$buf = $_;      
    }
    else {
	$buf .= $_;      
    }
}
