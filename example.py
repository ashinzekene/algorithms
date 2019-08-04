# Python3 program to print all possible 
# substrings of a given string 
  
  
# Function to print all sub strings 
def subString(Str,n): 
      
    # Pick starting point 
    for Len in range(1,n + 1): 
          
        # Pick ending point 
        for i in range(n - Len + 1): 
              
            # Print characters from current 
            # starting point to current ending 
            # point.  
            j = i + Len - 1
  
            for k in range(i,j + 1): 
                print(Str[k],end="") 
            print() 
              
# Driver program to test above function 
Str = "abc"
subString(Str,len(Str)) 
  