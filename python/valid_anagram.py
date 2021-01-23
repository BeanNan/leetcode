from typing import DefaultDict
from collections import defaultdict


class Solution:
    def isAnagram(self, s: str, t: str) -> bool:

        if len(s) != len(t):
            return False

        s_mapping = self.count_mapping(s)
        t_mapping = self.count_mapping(t)

        for char in s:
            if s_mapping[char] != t_mapping[char]:
                return False

        return True

    def count_mapping(self, s: str) -> DefaultDict[str, int]:
        mapping: DefaultDict[str, int] = defaultdict(int)

        for char in s:
            mapping[char] += 1

        return mapping

