go-primer-contrast
==================

I am new to Go, and I want to try the concurrent programming with Go and test its performance.

The code in [go-lang homepage](golang.org)
using daisy chain is pretty typical for illustrating concurrent, 
but it has a bad performance to this simple task.

Here gives the code to filter primes directly, the amount of process is constant, 
however without sorting the prime list, the result can rarely be correct.
I regard it only as a performance test, so won't fix it.

** BUT, why using array cost much more RAM and is slower than slice!!? **

----------------------------------
初学Go, 想学习一下它的并发, 看下性能.

[go主页](golang.org)上的的菊花链channel筛法确实很能体现并发, 但是实际上性能很差.

这儿还是直接开多个线程筛素数, 但是多线程下测试结果返回不是按顺序的, 所以得排序, 排序就不好比较了, 所以索性不排序,
只会随机多得到几个, 可以接受((
结果如下, pascal的性能略悲催啊, 虽然代码风格和Go有丁点像XD

** 但是为什么用数组比切片还慢, 还爆那么大内存? 搞不懂啊 **

result
---------------------------

    File		      environment		  result	      		  RAM(+shared)
      
    primer.py 	    pypy	    7.3335249424:   664579	19.452M + 11.316M
    primer.py	      python  	55.0426769257:  664579	23.548M + 1.852M
    primer.c  	    clang	    1.294029:       664579	1.632M  + 0.304M
    primer.c	      cl -O3  	0.908038:       664579	1.104M  + 0.304M
    primer.pas	    fpc264  	3.720380035:    664579	1.976M  + 0.136M
    primer_sin_sli	go      	3.156230691s:   664579	20.240M + 0.540M
    primer_sin_arr	go      	3.337368608s:	664579	  156.924M + 78.54M
    primer_s_a_iter	go       	....................	  235.216M + 78.66M
    primer_concur_4	go      	3.906765893s:   664616	31.516M + 0.584M
    primer_concur_1	go	      4.751173223s:   664579	21.588M + 0.584M
