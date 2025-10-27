def funcionRecursiva(n):
    if n <= 21:
        return n
    else:
        return funcionRecursiva(n - 7) + funcionRecursiva(n - 14) + funcionRecursiva(n - 21)

print(f"El término F({10}) es: {funcionRecursiva(22)}")
