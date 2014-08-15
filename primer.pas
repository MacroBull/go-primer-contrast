uses sysutils, linux, unixtype;

const LIM =  9999999;

var 
	i, n, m: longint;
	index: longint;
	primes: array[0..LIM] of longint;
// 	primes: array of longint; // dynamic
	tStart, tEnd: timespec;

begin
	clock_gettime(0, @tStart);

	primes[0] := 2;
	primes[1] := 3;
	index := 2;

	for n:=5 to LIM do
		for i:=0 to index do begin
			m := primes[i];
			if m*m > n then begin
				primes[index] := n;
				inc(index);
				break;
			end
			else if 0 = n mod m then break;
		end;
	
	clock_gettime(0, @tEnd);
	writeln(tEnd.tv_sec - tStart.tv_sec + 1e-9*(tEnd.tv_nsec - tStart.tv_nsec):0:9, ':	', index);

end. 
