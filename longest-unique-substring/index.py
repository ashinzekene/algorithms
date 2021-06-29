def lengthOfLongestSubstring(self, s: str) -> int:
        valueMap = dict()
        pointer = 0
        maxLen = 0
        for i in range(len(s)):
            char = s[i]
            if char in valueMap and pointer <= valueMap[char]:
                pointer = valueMap[char] + 1
            valueMap[char] = i
            diff = i - pointer + 1
            maxLen = max(diff, maxLen)
        return maxLen