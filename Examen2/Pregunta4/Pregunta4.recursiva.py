def funcionRecursiva(n):
    if n <= 21:
        return n
    return funcionRecursiva(n - 7) + funcionRecursiva(n - 14) + funcionRecursiva(n - 21)
