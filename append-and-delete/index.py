def appendAndDelete(s, t, k):
    sameLength = 0
    toRemoveLength = 0
    toAddLength = 0
    if k >= (len(s) + len(t)):
        return "Yes"
    for i in range(min(len(s), len(t))):
        if s[i] == t[i]:
            sameLength+=1
        else:
            break
    toRemoveLength = len(s) - sameLength
    toAddLength = len(t) - sameLength
    whatsLeft = k - (toRemoveLength + toAddLength)
    if whatsLeft % 2 == 0:
        return "Yes"
    return "No"
