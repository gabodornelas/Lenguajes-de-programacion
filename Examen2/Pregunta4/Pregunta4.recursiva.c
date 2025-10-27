#include <stdio.h>

long long fib(long long n){ 
  if (n <= 21) return n; 
  return fib(n-7) + fib(n-14) + fib(n-21);
}

int main(){
  printf("%lld\n",fib(22));
  return 0;
}
