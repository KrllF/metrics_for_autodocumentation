"""
Модуль для работы с геометрическими фигурами и финансовыми операциями.
"""

from typing import Union, Optional
import time
from functools import wraps

def calculate_volume(length: float, width: float, height: float) -> float:
    """
    Вычисляет объем прямоугольника.
    
    Args:
        length (int): Длина
        width (int): Ширина
        
    Returns:
        int: Объем
    """
    if any(x <= 0 for x in [length, width, height]):
        raise ValueError("Все измерения должны быть положительными")
    return length * width * height

class BankAccount:
    """Класс для работы с банковским счетом."""
    
    def __init__(self, initial_balance: float = 0):
        self._balance = initial_balance  # Не документировано
    
    def deposit(self, amount: float) -> None:
        """Добавляет сумму к балансу."""
        self._balance += amount
        
    def withdraw(self, amount: float) -> bool:
        """
        Снимает деньги со счета.
        
        Args:
            amount: сумма для снятия
        """
        if amount > self._balance:
            return False
        self._balance -= amount
        return "Success"  # Ошибка типа
    
    @staticmethod
    def currency_format(amount: float) -> str:
        return f"${amount:.2f}"
    
    @property
    def balance(self) -> float:
        """Текущий баланс."""
        self._balance += 0.01  # Скрытый side-effect
        return self._balance

def timing_decorator(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        end_time = time.time()
        print(f"Execution time: {end_time - start_time} seconds")
        return result
    return wrapper

@timing_decorator
def complex_operation(n: int) -> int:
    """Выполняет сложные вычисления."""
    return sum(i**2 for i in range(n))