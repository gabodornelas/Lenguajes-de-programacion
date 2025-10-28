#include <stdio.h>

long long fun(long long n){ 
  if (n <= 21) return n; 
  return fun(n-7) + fun(n-14) + fun(n-21);
}

int main(){
  printf("%lld\n",fun(22));
  return 0;
}
