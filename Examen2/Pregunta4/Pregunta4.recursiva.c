#include <stdio.h>

int fib(int n) { 
  if (n <= 21) return n; 
  return fib(n-7) + fib(n-14) + fib(n-21);
}

int main(){
  printf("%d\n",fib(22));
  return 1;
}
