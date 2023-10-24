#!/usr/bin/env python3
# -*- coding:utf-8 -*-
###
# File: qs.py
# Created Date: 2023-09-13 09:25:30
# Author: ysj
# Description:
###

def quick_sort(arr, left, right):
    if left >= right:
        return
    base = arr[left]
    low = left
    high = right
    while left < right:
        while left < right and arr[right] >= base:
            right -= 1
        arr[left] = arr[right]
        while left < right and arr[left] <= base:
            left += 1
        arr[right] = arr[left]
    arr[left] = base
    quick_sort(arr, low, left - 1)
    quick_sort(arr, left + 1, high)


# quick_sort unit tests
def test_quick_sort():
    # test 1
    arr = [5, 7, 3, 2, 1, 9, 10, 4]
    quick_sort(arr, 0, len(arr) - 1)
    assert arr == [1, 2, 3, 4, 5, 7, 9, 10]
    # test 2
    arr = [5, 7, 3, 2, 1, 9, 10, 4]
    quick_sort(arr, 0, len(arr) - 1)
    assert arr == [1, 2, 3, 4, 5, 7, 9, 10]
    # test 3
    arr = [5, 7, 3, 2, 1, 9, 10, 4]
    quick_sort(arr, 0, len(arr) - 1)
    assert arr == [1, 2, 3, 4, 5, 7, 9, 10]


test_quick_sort()
