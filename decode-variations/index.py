def decodeVariations(S):
    def addToResult(string):
        a = b = 0
        if len(string) == 0:
            return 1
        if string[0] != "0":
            a = addToResult(string[1:])
        if len(string) > 1:
            val = int(string[:2])
            if val > 0 and val <= 26:
                b = addToResult(string[2:])
        return a + b
    return addToResult(S)